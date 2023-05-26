package generator

import (
	"net/http"
)

type Generator interface {
	InitPrewarm(nThread int)
	InitOpenLoop(ratio float64, rate int)
	InitCloseLoop(rThread int, wThread int)

	GenPrewarm(threadId int) *http.Request
	GetPrewarmStatus() (int, int)
	GenRead() *http.Request
	GenWrite() *http.Request
}

type Config struct {
	Address     string
	UserIdPath  string
	MovieIdPath string
}
