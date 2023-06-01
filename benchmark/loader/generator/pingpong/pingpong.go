package pingpong

import (
	"context"
	"fmt"
	"microless/loader/generator"
	"microless/loader/utils"
	"net/http"
)

type pingpongGenerator struct {
	addr string
}

func NewPingpongGenerator(config *generator.Config) generator.Generator {
	return &pingpongGenerator{
		addr: config.Address,
	}
}

func (g *pingpongGenerator) InitPrewarm(nThread int) {
	// do nothing
}

func (g *pingpongGenerator) InitOpenLoop(ratio float64, rate int) {
	// do nothing
}

func (g *pingpongGenerator) InitCloseLoop(rThread int, wThread int) {
	// do nothing
}

func (g *pingpongGenerator) GenPrewarm(ctx context.Context, threadId int) *http.Request {
	return nil
}

func (g *pingpongGenerator) GetPrewarmStatus() (int, int) {
	return 0, 0
}

func (g *pingpongGenerator) GenRead(ctx context.Context) *http.Request {
	url := fmt.Sprintf("http://%s/api/v1/ping", g.addr)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	utils.Check(err)
	return req
}

func (g *pingpongGenerator) GenWrite(ctx context.Context) *http.Request {
	url := fmt.Sprintf("http://%s/api/v1/ping", g.addr)
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	utils.Check(err)
	return req
}
