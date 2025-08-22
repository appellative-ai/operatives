package retry

import (
	"fmt"
	"net/http"
)

func ExampleNewURL() {
	req, err := http.NewRequest("GET", "/search?q=golang", nil)
	if err != nil {
		fmt.Printf("test: http.NewRequest() [err:%v]\n", err)
	}
	fmt.Printf("test: http.NewRequest() [host:%v] [url:%v]\n", req.Host, req.URL.String())

	err = newURL(req, "localhost:8080")
	if err != nil {
		fmt.Printf("test: newURL() [err:%v]\n", err)
	}
	fmt.Printf("test: http.NewRequest() [host:%v] [url:%v]\n", req.Host, req.URL.String())

	req, err = http.NewRequest("GET", "/search?q=golang", nil)
	err = newURL(req, "google.com")
	if err != nil {
		fmt.Printf("test: newURL() [err:%v]\n", err)
	}
	fmt.Printf("test: http.NewRequest() [host:%v] [url:%v]\n", req.Host, req.URL.String())

	//Output:
	//test: http.NewRequest() [host:] [url:/search?q=golang]
	//test: http.NewRequest() [host:localhost:8080] [url:http://localhost:8080/search?q=golang]
	//test: http.NewRequest() [host:google.com] [url:https://google.com/search?q=golang]

}
