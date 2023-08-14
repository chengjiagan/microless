package socialnetwork

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
	"strconv"
	"sync/atomic"
)

const NumPostOnce = 1000

type socialnetworkGenerator struct {
	api   string // home, user, or mix
	addr  string
	users []user

	// prewarm
	cnt        atomic.Int32
	nThread    int
	curUserIdx []int
	curPostIdx []int
}

type user struct {
	UserId   string `json:"user_id"`
	NumPost  int    `json:"num_post"`
	HomePost int    `json:"home_post"`
}

func NewSocialnetworkGenerator(config *generator.Config) generator.Generator {
	var users []user
	// get user ids in dataset
	data, err := os.ReadFile(config.UserIdPath)
	utils.Check(err)
	err = json.Unmarshal(data, &users)
	utils.Check(err)

	return &socialnetworkGenerator{
		api:   config.Api,
		addr:  config.Address,
		users: users,
	}
}

func (g *socialnetworkGenerator) InitPrewarm(nThread int) {
	g.curUserIdx = make([]int, nThread)
	g.curPostIdx = make([]int, nThread)
	g.nThread = nThread

	for i := 0; i < nThread; i++ {
		g.curUserIdx[i] = i
	}
}

func (g *socialnetworkGenerator) InitOpenLoop(ratio float64, rate int) {
	// do nothing
}

func (g *socialnetworkGenerator) InitCloseLoop(rThread int, wThread int) {
	// do nothing
}

func (g *socialnetworkGenerator) GenPrewarm(ctx context.Context, threadId int) *http.Request {
	userIdx := g.curUserIdx[threadId]
	postIdxStart := g.curPostIdx[threadId]

	if userIdx >= len(g.users) {
		return nil
	}

	var postIdxEnd int
	var function string
	for {
		// home timeline
		function = "hometimeline"
		if postIdxStart < g.users[userIdx].HomePost {
			postIdxEnd = postIdxStart + NumPostOnce
			if postIdxEnd > g.users[userIdx].HomePost {
				postIdxEnd = g.users[userIdx].HomePost
			}
			g.curPostIdx[threadId] = postIdxEnd
			break
		}

		// user timeline
		function = "usertimeline"
		postIdxStart -= g.users[userIdx].HomePost
		if postIdxStart < g.users[userIdx].NumPost {
			postIdxEnd = postIdxStart + NumPostOnce
			if postIdxEnd > g.users[userIdx].NumPost {
				postIdxEnd = g.users[userIdx].NumPost
			}
			g.curPostIdx[threadId] = postIdxEnd + g.users[userIdx].HomePost
			break
		}

		// try next user
		userIdx += g.nThread
		g.curUserIdx[threadId] = userIdx
		postIdxStart = 0
		g.cnt.Add(1)

		// no more user
		if userIdx >= len(g.users) {
			return nil
		}
	}

	// generate request
	url := fmt.Sprintf("http://%s/api/v1/%s/%s?start=%d&stop=%d", g.addr, function, g.users[userIdx].UserId, postIdxStart, postIdxEnd)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	utils.Check(err)

	return req
}

func (g *socialnetworkGenerator) GetPrewarmStatus() (int, int) {
	return int(g.cnt.Load()), len(g.users)
}

func (g *socialnetworkGenerator) GenRead(ctx context.Context) *http.Request {
	// randomly select a user
	user := rand.Intn(len(g.users))
	userid := g.users[user].UserId

	function, n := g.getReadFunctionAndNum(user)

	// randomly select some posts if user have more than 10 posts
	var start, stop int
	if n <= 10 {
		start = 0
		stop = n
	} else {
		start = rand.Intn(n - 10)
		stop = start + 10
	}

	// generate request
	url := fmt.Sprintf("http://%s/api/v1/%s/%s?start=%d&stop=%d", g.addr, function, userid, start, stop)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	utils.Check(err)

	return req
}

func (g *socialnetworkGenerator) getReadFunctionAndNum(user int) (string, int) {
	var n int
	var function string

	// get function based on api mode
	switch g.api {
	case "home":
		function = "hometimeline"
	case "user":
		function = "usertimeline"
	case "mix":
		// select ramdonly
		if rand.Intn(2) == 0 {
			function = "hometimeline"
		} else {
			function = "usertimeline"
		}
	}

	switch function {
	case "hometimeline":
		n = g.users[user].HomePost
	case "usertimeline":
		n = g.users[user].NumPost
	}

	return function, n
}

func (g *socialnetworkGenerator) GenWrite(ctx context.Context) *http.Request {
	url := "http://" + g.addr + "/api/v1/composepost"
	val := g.randComposePost()

	// serialize value in JSON format
	data, err := json.Marshal(val)
	utils.Check(err)
	// generate request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	utils.Check(err)

	return req
}

type ComposePostRequest struct {
	Username   string   `json:"username"`
	UserId     string   `json:"user_id"`
	Text       string   `json:"text"`
	MediaIds   []int64  `json:"media_ids"`
	MediaTypes []string `json:"media_types"`
	PostType   string   `json:"post_type"`
}

func (g *socialnetworkGenerator) randComposePost() *ComposePostRequest {
	// randomly select a user
	user := rand.Intn(len(g.users))
	userId := g.users[user].UserId
	username := strconv.FormatInt(int64(user), 10)

	// randomly generate a post
	mention := rand.Intn(len(g.users))
	text := fmt.Sprintf(
		"%s\nhttp://url_%d.com\n@%d\n",
		utils.RandString(rand.Intn(10)),
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
