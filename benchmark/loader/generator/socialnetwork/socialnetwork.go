package socialnetwork

import (
	"bytes"
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

type socialnetworkGenerator struct {
	addr  string
	users []user

	// prewarm
	cnt        int32
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

func (g *socialnetworkGenerator) GenPrewarm(threadId int) *http.Request {
	userIdx := g.curUserIdx[threadId]
	postIdxStart := g.curPostIdx[threadId]

	if userIdx >= len(g.users) {
		return nil
	}

	if postIdxStart >= g.users[userIdx].NumPost {
		userIdx++
		postIdxStart = 0
		atomic.AddInt32(&g.cnt, 1)
	}

	if userIdx >= len(g.users) {
		return nil
	}

	postIdxEnd := postIdxStart + 100
	if postIdxEnd > g.users[userIdx].NumPost {
		postIdxEnd = g.users[userIdx].NumPost
	}

	// generate request
	url := fmt.Sprintf("http://%s/api/v1/hometimeline/%s?start=%d&stop=%d", g.addr, g.users[userIdx].UserId, postIdxStart, postIdxEnd)
	req, err := http.NewRequest("GET", url, nil)
	utils.Check(err)

	return req
}

func (g *socialnetworkGenerator) GetPrewarmStatus() (int, int) {
	return int(atomic.LoadInt32(&g.cnt)), len(g.users)
}

func (g *socialnetworkGenerator) GenRead() *http.Request {
	// randomly select a user
	user := rand.Intn(len(g.users))
	userid := g.users[user].UserId
	n := g.users[user].HomePost

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
	url := fmt.Sprintf("http://%s/api/v1/hometimeline/%s?start=%d&stop=%d", g.addr, userid, start, stop)
	req, err := http.NewRequest("GET", url, nil)
	utils.Check(err)

	return req
}

func (g *socialnetworkGenerator) GenWrite() *http.Request {
	url := "http://" + g.addr + "/api/v1/composepost"
	val := g.randComposePost()

	// serialize value in JSON format
	data, err := json.Marshal(val)
	utils.Check(err)
	// generate request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
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
