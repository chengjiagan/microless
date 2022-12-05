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
	"time"
)

var addr = flag.String("addr", "localhost:8080", "address to the gateway service")
var pathUserIds = flag.String("userid", "user_ids.json", "path to json file that contains user ids")
var pathMovieIds = flag.String("movieid", "movie_ids.json", "path to json file that contains movie ids")
var seconds = flag.Int("time", 1, "load duration in seconds")
var ratio = flag.Float64("ratio", 0.8, "ratio of read requests")
var nThread = flag.Int("thread", 1, "number of threads")
var rate = flag.Int("rate", 0, "request rate, 0 if rate is unlimited")
var output = flag.String("output", "", "path to output file")

var users []struct {
	UserId    string `json:"user_id"`
	NumReview int    `json:"num_review"`
}
var movies []struct {
	MovieId   string `json:"movie_id"`
	NumReview int    `json:"num_review"`
}

var client *http.Client

// record a test sample
type sample struct {
	start time.Time
	end   time.Time
}

func main() {
	flag.Parse()

	// init http client
	client = &http.Client{
		Timeout: time.Minute,
	}

	// get user ids in dataset
	data, err := os.ReadFile(*pathUserIds)
	check(err)
	err = json.Unmarshal(data, &users)
	check(err)
	// get movie ids in dataset
	data, err = os.ReadFile(*pathMovieIds)
	check(err)
	err = json.Unmarshal(data, &movies)
	check(err)

	// set random seed
	rand.Seed(time.Now().UnixNano())

	// start load test
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	var out []chan sample
	if *rate == 0 {
		out = load(ctx)
	} else {
		out = loadWithRate(ctx, *rate)
	}

	// wait and stop
	for i := 0; i < *seconds; i++ {
		fmt.Printf("\r%d/%d", i, *seconds)
		time.Sleep(time.Second)
	}
	fmt.Println()
	cancel()
	end := time.Now()
	duration := end.Sub(start).Seconds()

	// get result
	ss := make([]sample, 0)
	for _, ch := range out {
		for s := range ch {
			ss = append(ss, s)
		}
	}

	// calculate estimated throughput
	n := len(ss)
	tp := float64(n) / duration
	// calculate average latency in ms
	var total int64
	for _, s := range ss {
		total += s.end.Sub(s.start).Milliseconds()
	}
	avg := float64(total) / float64(n)
	// output
	fmt.Printf("throughput: %v qps\naverage latency: %v ms\n", tp, avg)

	// save samples to file
	// generate output filename if not given
	if *output == "" {
		*output = fmt.Sprintf("load_socialnetwork_%v_t%v.csv", start.Format("200601021504"), *nThread)
	}
	// open file
	fp, err := os.Create(*output)
	check(err)
	defer fp.Close()
	// write
	_, err = fp.WriteString("start,end\n")
	check(err)
	for _, s := range ss {
		_, err = fp.WriteString(fmt.Sprintf("%v,%v\n", s.start.UnixMilli(), s.end.UnixMilli()))
		check(err)
	}
}

// close-loop load test
func load(ctx context.Context) []chan sample {
	out := make([]chan sample, *nThread)

	for i := 0; i < *nThread; i++ {
		ch := make(chan sample)
		out[i] = ch

		go func() {
			ss := make([]sample, 0)
			for ctx.Err() == nil {
				s := send(ctx)
				ss = append(ss, s)
			}

			for _, s := range ss {
				ch <- s
			}
			close(ch)
		}()
	}

	return out
}

// open-loop load test
func loadWithRate(ctx context.Context, r int) []chan sample {
	// TODO
	return nil
}

func send(ctx context.Context) sample {
	start := time.Now()

	// randomly send read or write request
	p := rand.Float64()
	if p < *ratio {
		sendRead(ctx)
	} else {
		sendWrite(ctx)
	}

	end := time.Now()
	return sample{
		start: start,
		end:   end,
	}
}

func sendWrite(ctx context.Context) {
	url := "http://" + *addr + "/api/v1/composereview"
	val := randComposeReview()

	// serialize value in JSON format
	data, err := json.Marshal(val)
	check(err)
	// generate request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	check(err)

	sendRequest(ctx, url, req)
}

type ComposeReviewRequest struct {
	MovieId string `json:"movie_id"`
	UserId  string `json:"user_id"`
	Text    string `json:"text"`
	Rating  int32  `json:"rating"`
}

func randComposeReview() *ComposeReviewRequest {
	// randomly select a user
	user := rand.Intn(len(users))
	userId := users[user].UserId
	// randomly select a movie
	movie := rand.Intn(len(movies))
	movieId := movies[movie].MovieId
	// randomly generate a review text
	text := randString(rand.Intn(100))
	// randomly select a rating in range [1, 10]
	rating := rand.Int31n(10) + 1

	return &ComposeReviewRequest{
		MovieId: movieId,
		UserId:  userId,
		Text:    text,
		Rating:  rating,
	}
}

func sendRead(ctx context.Context) {
	var url string
	p := rand.Float64()
	if p < 0.5 {
		// randomly select a user
		user := rand.Intn(len(users))
		userId := users[user].UserId
		n := users[user].NumReview
		// randomly select some reviews if user have more than 10 reviews
		var start, stop int
		if n <= 10 {
			start = 0
			stop = n
		} else {
			start = rand.Intn(n - 10)
			stop = start + 10
		}
		url = fmt.Sprintf("http://%s/api/v1/userreview/%v?start=%d&stop=%d", *addr, userId, start, stop)
	} else {
		// randomly select a movie
		movie := rand.Intn(len(movies))
		movieId := movies[movie].MovieId
		n := movies[movie].NumReview
		// randomly select some reviews if movie have more than 10 reviews
		var start, stop int
		if n <= 10 {
			start = 0
			stop = n
		} else {
			start = rand.Intn(n - 10)
			stop = start + 10
		}
		url = fmt.Sprintf("http://%s/api/v1/page/%v?review_start=%d&review_stop=%d", *addr, movieId, start, stop)
	}

	// generate request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	check(err)

	sendRequest(ctx, url, req)
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
