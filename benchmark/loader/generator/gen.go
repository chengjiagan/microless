package generator

import (
	"context"
	"net/http"
)

type Generator interface {
	InitPrewarm(nThread int)
	InitOpenLoop(ratio float64, rate int)
	InitCloseLoop(rThread int, wThread int)

	GenPrewarm(ctx context.Context, threadId int) *http.Request
	GetPrewarmStatus() (int, int)
	GenRead(ctx context.Context) *http.Request
	GenWrite(ctx context.Context) *http.Request
}

type Config struct {
	Api         string
	Address     string
	UserIdPath  string
	MovieIdPath string
}
