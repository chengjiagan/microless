package loadbalancer

import (
	"context"
	"log"
	"microless/loadbalancer/internal/queue"
	"microless/loadbalancer/internal/utils"
	"sync"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const avgAcc = 1000
const resourceUnit = 1000

type ServerlessLB struct {
	// params from config
	totalResources     int
	maxCapacity        int
	methodRequirements map[string]int

	mu   sync.Mutex
	cond *sync.Cond

	// for concurrency controlling
	taskId           atomic.Int64
	q                *queue.TaskQueue
	concurrency      int
	currentResources int

	// for calculate time average queue length
	e       []event // protected by mu
	taskAvg atomic.Int32
}

type event struct {
	t time.Time
	l int
}

func NewServerlessLB(stats *Stats) *ServerlessLB {
	config := utils.GetServerlessConfig()
	if !config.Enable {
		return nil
	}

	e := make([]event, 0)
	e = append(e, event{
		t: time.Now(),
		l: 0,
	})

	req := make(map[string]int)
	for k, v := range config.MethodReqirements {
		req[k] = int(v * resourceUnit)
	}

	sl := &ServerlessLB{
		totalResources:     config.MaxConcurrency * resourceUnit,
		maxCapacity:        config.MaxCapacity,
		methodRequirements: req,
		q:                  queue.NewTaskQueue(config.MaxCapacity),
		e:                  e,
	}
	sl.cond = sync.NewCond(&sl.mu)
	go sl.updateLoop()

	stats.reg.MustRegister(
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: NameServerlessTaskTotal,
				Help: HelpServerlessTaskTotal,
			},
			func() float64 {
				return float64(sl.taskAvg.Load()) / avgAcc
			},
		),
	)

	return sl
}

func (lb *ServerlessLB) updateLoop() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		lb.updateConcurrency()
	}
}

func (lb *ServerlessLB) updateConcurrency() {
	newAvg := lb.averageConcurrency() * avgAcc
	oldAvg := float64(lb.taskAvg.Load())
	lb.taskAvg.Store(int32(0.5*oldAvg + 0.5*newAvg))
}

func (lb *ServerlessLB) averageConcurrency() float64 {
	now := time.Now()

	lb.mu.Lock()
	l := lb.concurrency + lb.q.Len()
	e := lb.e
	// add start event
	lb.e = make([]event, 0)
	lb.e = append(lb.e, event{
		t: now,
		l: l,
	})
	lb.mu.Unlock()

	// add stop event
	e = append(e, event{
		t: now,
		l: l,
	})

	avg := 0.0
	for i := 1; i < len(e); i++ {
		avg += float64(e[i].l) * e[i].t.Sub(e[i-1].t).Seconds()
	}
	avg /= e[len(e)-1].t.Sub(e[0].t).Seconds()
	return avg
}

func (lb *ServerlessLB) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if lb == nil {
			return handler(ctx, req)
		}

		taskid := lb.taskId.Add(1)
		_, method := utils.GetServiceAndMethod(info)
		methodRequirement := lb.getMethodRequirement(method)
		err = lb.requestResource(taskid, methodRequirement)
		if err != nil {
			return
		}

		resp, err = handler(ctx, req)
		lb.releaseResource(methodRequirement)
		return
	}
}

func (lb *ServerlessLB) getMethodRequirement(method string) int {
	if v, ok := lb.methodRequirements[method]; ok {
		return v
	}
	// default value
	return resourceUnit
}

func (lb *ServerlessLB) requestResource(taskid int64, req int) error {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.addEvent()

	// check if the request can be processed
	if lb.currentResources+req > lb.totalResources {
		// check if the queue is full
		if lb.q.Len() >= lb.maxCapacity {
			return status.Error(codes.ResourceExhausted, "Serverless queue is full")
		}

		// wait in the queue
		lb.q.Push(taskid)
		for lb.q.Front() != taskid || lb.currentResources+req > lb.totalResources {
			if lb.q.Front() == taskid {
				log.Printf("task %d at front but not enough resources", taskid)
			}
			lb.cond.Wait()
		}
		lb.q.Pop()
	}
	lb.currentResources += req
	lb.concurrency++

	return nil
}

func (lb *ServerlessLB) releaseResource(amount int) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.addEvent()

	lb.currentResources -= amount
	lb.concurrency--
	lb.cond.Broadcast()
}

func (lb *ServerlessLB) addEvent() {
	lb.e = append(lb.e, event{
		t: time.Now(),
		l: lb.concurrency + lb.q.Len(),
	})
}
