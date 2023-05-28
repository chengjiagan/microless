package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"microless/loader/generator"
	"microless/loader/generator/media"
	"microless/loader/generator/pingpong"
	"microless/loader/generator/socialnetwork"
	"net/http"
	"os"
	"sync"
	"time"
)

// required by all modes
var addr = flag.String("addr", "", "address to the gateway service")
var service = flag.String("service", "", "kind of service to test: social-network, media, pingpong")
var mode = flag.String("mode", "", "load test mode: close for close-loop, open for open-loop, prewarm for pre-warming")

// required by social-network and media
var pathUserIds = flag.String("userid", "", "path to json file that contains user ids")
var pathMovieIds = flag.String("movieid", "", "path to json file that contains movie ids")

// required by pre-warming
var nThread = flag.Int("nthread", 1, "number of threads sending requests")

// required by close-loop and open-loop
var seconds = flag.Int("time", 0, "load duration in seconds")
var output = flag.String("output", "", "path to output file")

// close-loop load test
var rThread = flag.Int("rthread", 0, "number of threads sending read requests")
var wThread = flag.Int("wthread", 0, "number of threads sending write requests")

// open-loop load test
var ratio = flag.Float64("ratio", 0, "ratio of read requests")
var rate = flag.Int("rate", 0, "request rate in QPS, 0 if rate is unlimited")

var client *http.Client
var gen generator.Generator

// record a test sample
type sample struct {
	start time.Time
	end   time.Time
	t     string
	code  int
}

func main() {
	flag.Parse()

	// check params
	checkParams()

	// init http client
	client = &http.Client{
		Timeout: time.Minute,
	}

	// init generator
	gen = newGenerator(
		*service,
		&generator.Config{
			Address:     *addr,
			UserIdPath:  *pathUserIds,
			MovieIdPath: *pathMovieIds,
		},
	)
	if gen == nil {
		log.Fatal("unknown service")
	}

	// start test
	switch *mode {
	case "close":
		closeLoop()
	case "open":
		openLoop()
	case "prewarm":
		prewarm()
	default:
		log.Fatal("unknown mode")
	}
}

func checkParams() {
	// addr, pathUserIds, output are required
	if *addr == "" || *pathUserIds == "" || *mode == "" ||
		(*mode != "prewarm" && *output == "") {
		flag.Usage()
		os.Exit(1)
	}

	// exit if nothing to do
	if (*mode != "prewarm" && *seconds == 0) ||
		(*mode == "close" && (*rThread == 0 && *wThread == 0)) ||
		(*mode == "open" && *rate == 0) {
		os.Exit(0)
	}
}

func newGenerator(service string, config *generator.Config) generator.Generator {
	switch service {
	case "socialnetwork":
		return socialnetwork.NewSocialnetworkGenerator(config)
	case "media":
		return media.NewMediaGenerator(config)
	case "pingpong":
		return pingpong.NewPingpongGenerator(config)
	default:
		return nil
	}
}

func closeLoop() {
	gen.InitCloseLoop(*rThread, *wThread)

	// start load test
	ctx, cancel := context.WithCancel(context.Background())
	out := load(ctx)

	// wait and stop
	for i := 0; i < *seconds; i++ {
		fmt.Printf("\r%d/%d", i, *seconds)
		time.Sleep(time.Second)
	}
	fmt.Println()
	cancel()

	// get result
	ss := make([]sample, 0)
	for _, ch := range out {
		for s := range ch {
			ss = append(ss, s)
		}
	}

	// print metrics and save raw data
	print(ss)
	save(ss)
}

func openLoop() {
	gen.InitOpenLoop(*ratio, *rate)

	ctx, cancel := context.WithCancel(context.Background())
	out := loadWithRate(ctx)

	// wait and stop
	go func() {
		for i := 0; i < *seconds; i++ {
			fmt.Printf("\r%d/%d", i, *seconds)
			time.Sleep(time.Second)
		}
		fmt.Println()
		cancel()
	}()

	// get result
	ss := make([]sample, 0)
	for s := range out {
		ss = append(ss, s)
	}

	// print metrics and save raw data
	print(ss)
	save(ss)
}

func prewarm() {
	gen.InitPrewarm(*nThread)

	var wg sync.WaitGroup
	wg.Add(*nThread)
	for t := 0; t < *nThread; t++ {
		go func(t int) {
			ctx := context.Background()
			for {
				req := gen.GenPrewarm(t)
				if req == nil {
					break
				}
				code := sendRequest(ctx, req)
				if code != http.StatusOK {
					log.Fatalf("prewarm failed: %v", code)
				}
			}
			wg.Done()
		}(t)
	}

	go func() {
		for {
			cur, total := gen.GetPrewarmStatus()
			fmt.Printf("\r%d/%d", cur, total)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println()
}

func print(ss []sample) {
	// calculate estimated throughput
	n := len(ss)
	tp := float64(n) / float64(*seconds)
	// calculate average latency in ms
	var total int64
	for _, s := range ss {
		total += s.end.Sub(s.start).Milliseconds()
	}
	avg := float64(total) / float64(n)
	// output
	fmt.Printf("throughput: %v qps\naverage latency: %v ms\n", tp, avg)
}

func save(ss []sample) {
	// save samples to file
	// open file
	fp, err := os.Create(*output)
	check(err)
	defer fp.Close()
	// write
	_, err = fp.WriteString("start,end,type,code\n")
	check(err)
	for _, s := range ss {
		_, err = fp.WriteString(fmt.Sprintf("%v,%v,%v,%v\n", s.start.UnixMilli(), s.end.UnixMilli(), s.t, s.code))
		check(err)
	}
}

// close-loop load test
func load(ctx context.Context) []chan sample {
	out := make([]chan sample, *rThread+*wThread)

	for i := 0; i < *rThread+*wThread; i++ {
		ch := make(chan sample)
		out[i] = ch

		go func(i int) {
			ss := make([]sample, 0)
			for ctx.Err() == nil {
				var s sample
				if i < *rThread {
					s = send(ctx, gen.GenRead(), "read")
				} else {
					s = send(ctx, gen.GenWrite(), "write")
				}
				ss = append(ss, s)
			}

			for _, s := range ss {
				ch <- s
			}
			close(ch)
		}(i)
	}

	return out
}

// open-loop load test
func loadWithRate(ctx context.Context) chan sample {
	ch := make(chan sample)
	var wg sync.WaitGroup

	go func() {
		timer := time.Tick(time.Second / time.Duration(*rate))
		for ctx.Err() == nil {
			<-timer
			wg.Add(1)
			go func() {
				p := rand.Float64()
				var s sample
				if p < *ratio {
					s = send(ctx, gen.GenRead(), "read")
				} else {
					s = send(ctx, gen.GenWrite(), "write")
				}
				ch <- s
				wg.Done()
			}()
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func send(ctx context.Context, req *http.Request, t string) sample {
	start := time.Now()
	code := sendRequest(ctx, req)
	end := time.Now()
	return sample{
		start: start,
		end:   end,
		t:     t,
		code:  code,
	}
}

// send a http request
func sendRequest(ctx context.Context, req *http.Request) int {
	// send request
	resp, err := client.Do(req)
	check(err)

	// read respond and close
	_, err = io.ReadAll(resp.Body)
	check(err)
	err = resp.Body.Close()
	check(err)

	return resp.StatusCode
}

// panic if encounter error
func check(err error) {
	if err != nil {
		panic(err)
	}
}
