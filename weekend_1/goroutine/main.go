package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Print(wg)
		//go HttpRequest(wg)
	}

	wg.Wait()
}

func Print(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello")
}

func HttpRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
}