package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// required by all modes
var addr = flag.String("addr", "", "address to the gateway service")
var pathUserIds = flag.String("userid", "", "path to json file that contains user ids")
var mode = flag.String("mode", "", "load test mode: close for close-loop, open for open-loop, prewarm for pre-warming")

// required by close-loop and open-loop
var seconds = flag.Int("time", 0, "load duration in seconds")
var output = flag.String("output", "", "path to output file")

// close-loop load test
var rThread = flag.Int("rthread", 0, "number of threads sending read requests")
var wThread = flag.Int("wthread", 0, "number of threads sending write requests")

// open-loop load test
var ratio = flag.Float64("ratio", 0, "ratio of read requests")
var rate = flag.Int("rate", 0, "request rate in QPS, 0 if rate is unlimited")

var users []struct {
	UserId   string `json:"user_id"`
	NumPost  int    `json:"num_post"`
	HomePost int    `json:"home_post"`
}

var client *http.Client

// record a test sample
type sample struct {
	start time.Time
	end   time.Time
	t     string
}

func main() {
	flag.Parse()

	// check params
	checkParams()

	// init http client
	client = &http.Client{
		Timeout: time.Minute,
	}

	// get user ids in dataset
	data, err := os.ReadFile(*pathUserIds)
	check(err)
	err = json.Unmarshal(data, &users)
	check(err)

	// start test
	switch *mode {
	case "close":
		closeLoop()
	case "open":
		openLoop()
	case "prewarm":
		prewarm()
	default:
		fmt.Println("unknown mode")
		os.Exit(1)
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

func closeLoop() {
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
	ctx := context.Background()
	for i, user := range users {
		fmt.Printf("\r%d/%d", i, len(users))
		sendRead(ctx, user.UserId, 0, user.HomePost)
	}
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
	_, err = fp.WriteString("start,end,type\n")
	check(err)
	for _, s := range ss {
		_, err = fp.WriteString(fmt.Sprintf("%v,%v,%v\n", s.start.UnixMilli(), s.end.UnixMilli(), s.t))
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
					s = send(ctx, loadRead, "read")
				} else {
					s = send(ctx, sendWrite, "write")
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

func loadRead(ctx context.Context) {
	userid, start, stop := randHomeTimeline()
	sendRead(ctx, userid, start, stop)
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
					s = send(ctx, loadRead, "read")
				} else {
					s = send(ctx, sendWrite, "write")
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

func send(ctx context.Context, sendfunc func(context.Context), t string) sample {
	start := time.Now()
	sendfunc(ctx)
	end := time.Now()
	return sample{
		start: start,
		end:   end,
		t:     t,
	}
}

func sendWrite(ctx context.Context) {
	url := "http://" + *addr + "/api/v1/composepost"
	val := randComposePost()

	// serialize value in JSON format
	data, err := json.Marshal(val)
	check(err)
	// generate request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	check(err)

	sendRequest(ctx, url, req)
}

type ComposePostRequest struct {
	Username   string   `json:"username"`
	UserId     string   `json:"user_id"`
	Text       string   `json:"text"`
	MediaIds   []int64  `json:"media_ids"`
	MediaTypes []string `json:"media_types"`
	PostType   string   `json:"post_type"`
}

func randComposePost() *ComposePostRequest {
	// randomly select a user
	user := rand.Intn(len(users))
	userId := users[user].UserId
	username := strconv.FormatInt(int64(user), 10)

	// randomly generate a post
	mention := rand.Intn(len(users))
	text := fmt.Sprintf(
		"%s\nhttp://url_%d.com\n@%d\n",
		randString(rand.Intn(10)),
		user,
		mention,
	)

	return &ComposePostRequest{
		Username:   username,
		UserId:     userId,
		Text:       text,
		MediaIds:   []int64{1},
		MediaTypes: []string{"png"},
		PostType:   "POST",
	}
}

// send read home timeline request
func sendRead(ctx context.Context, userid string, start, stop int) {
	// generate request
	url := fmt.Sprintf("http://%s/api/v1/hometimeline/%s?start=%d&stop=%d", *addr, userid, start, stop)
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	// send
	sendRequest(ctx, url, req)
}

// randomly generate a home timeline request
// return userid, start, stop
func randHomeTimeline() (string, int, int) {
	// randomly select a user
	user := rand.Intn(len(users))
	userid := users[user].UserId
	n := users[user].HomePost

	// randomly select some posts if user have more than 10 posts
	var start, stop int
	if n <= 10 {
		start = 0
		stop = n
	} else {
		start = rand.Intn(n - 10)
		stop = start + 10
	}

	return userid, start, stop
}

// send a http request
func sendRequest(ctx context.Context, url string, req *http.Request) {
	// send request
	resp, err := client.Do(req)
	check(err)

	// read respond and close
	_, err = io.ReadAll(resp.Body)
	check(err)
	err = resp.Body.Close()
	check(err)
}

// panic if encounter error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

var alphanum = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(b)
}
