package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Startup")
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Errorf("dings")
	}
	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("dings")
	}
	fmt.Println(string(body))
}
