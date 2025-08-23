package retry

import (
	"fmt"
	"github.com/appellative-ai/common/iox"
	"net/http"
	"time"
)

func ExampleNewAgent() {
	a := NewAgent(0)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/operative/retry]

}

func ExampleExchange() {
	a := newAgent(time.Millisecond * 2000)
	a.primary.Store("localhost:8080")
	a.secondary.Store("google.com")

	req, _ := http.NewRequest("GET", "/search?q=golang", nil)
	req.Header.Set(iox.AcceptEncoding, iox.GzipEncoding)
	resp, err := a.Exchange(req)
	fmt.Printf("test: Exchange() -> [%v] [ce:%v] [err:%v]\n", resp.StatusCode, resp.Header.Get(iox.ContentEncoding), err)

	//Output:
	//test: Exchange() -> [200] [ce:] [err:<nil>]

}
