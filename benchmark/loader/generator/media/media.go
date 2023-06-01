package media

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"microless/loader/generator"
	"microless/loader/utils"
	"net/http"
	"os"
)

type mediaGenerator struct {
	addr   string
	users  []user
	movies []movie
}

type user struct {
	UserId    string `json:"user_id"`
	NumReview int    `json:"num_review"`
}

type movie struct {
	MovieId   string `json:"movie_id"`
	NumReview int    `json:"num_review"`
}

func NewMediaGenerator(config *generator.Config) generator.Generator {
	var users []user
	var movies []movie
	// get user ids in dataset
	data, err := os.ReadFile(config.UserIdPath)
	utils.Check(err)
	err = json.Unmarshal(data, &users)
	utils.Check(err)
	// get movie ids in dataset
	data, err = os.ReadFile(config.MovieIdPath)
	utils.Check(err)
	err = json.Unmarshal(data, &movies)
	utils.Check(err)

	return &mediaGenerator{
		addr:   config.Address,
		users:  users,
		movies: movies,
	}
}

func (g *mediaGenerator) InitPrewarm(nThread int) {
	// do nothing
}

func (g *mediaGenerator) InitOpenLoop(ratio float64, rate int) {
	// do nothing
}

func (g *mediaGenerator) InitCloseLoop(rThread int, wThread int) {
	// do nothing
}

func (g *mediaGenerator) GenPrewarm(ctx context.Context, threadId int) *http.Request {
	return nil
}

func (g *mediaGenerator) GetPrewarmStatus() (int, int) {
	return 0, 0
}

func (g *mediaGenerator) GenRead(ctx context.Context) *http.Request {
	var url string
	p := rand.Float64()
	if p < 0.5 {
		// randomly select a user
		user := rand.Intn(len(g.users))
		userId := g.users[user].UserId
		n := g.users[user].NumReview
		// randomly select some reviews if user have more than 10 reviews
		var start, stop int
		if n <= 10 {
			start = 0
			stop = n
		} else {
			start = rand.Intn(n - 10)
			stop = start + 10
		}
		url = fmt.Sprintf("http://%s/api/v1/userreview/%v?start=%d&stop=%d", g.addr, userId, start, stop)
	} else {
		// randomly select a movie
		movie := rand.Intn(len(g.movies))
		movieId := g.movies[movie].MovieId
		n := g.movies[movie].NumReview
		// randomly select some reviews if movie have more than 10 reviews
		var start, stop int
		if n <= 10 {
			start = 0
			stop = n
		} else {
			start = rand.Intn(n - 10)
			stop = start + 10
		}
		url = fmt.Sprintf("http://%s/api/v1/page/%v?review_start=%d&review_stop=%d", g.addr, movieId, start, stop)
	}

	// generate request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	utils.Check(err)

	return req
}

func (g *mediaGenerator) GenWrite(ctx context.Context) *http.Request {
	url := "http://" + g.addr + "/api/v1/composereview"
	val := g.randComposeReview()

	// serialize value in JSON format
	data, err := json.Marshal(val)
	utils.Check(err)
	// generate request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	utils.Check(err)

	return req
}

type ComposeReviewRequest struct {
	MovieId string `json:"movie_id"`
	UserId  string `json:"user_id"`
	Text    string `json:"text"`
	Rating  int32  `json:"rating"`
}

func (g *mediaGenerator) randComposeReview() *ComposeReviewRequest {
	// randomly select a user
	user := rand.Intn(len(g.users))
	userId := g.users[user].UserId
	// randomly select a movie
	movie := rand.Intn(len(g.movies))
	movieId := g.movies[movie].MovieId
	// randomly generate a review text
	text := utils.RandString(rand.Intn(100))
	// randomly select a rating in range [1, 10]
	rating := rand.Int31n(10) + 1

	return &ComposeReviewRequest{
		MovieId: movieId,
		UserId:  userId,
		Text:    text,
		Rating:  rating,
	}
}
