package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	in := make(chan string)
	out := make(chan int)

	for i := 0; i < 1000; i++ {
		go Worker(in, out)
	}
	go Print(out)

	for i := 0; i < 1000; i++ {
		in <- "https://example.com" + strconv.Itoa(i)
	}
	time.Sleep(2 * time.Second)
}

func Worker(in chan string, out chan int) {
	for val := range in {
		//res, err := http.Get(val)
		//if err != nil {
		//	panic(err)
		//}
		//
		//out <- res.StatusCode

		//time.Sleep(1 * time.Second)
		//fmt.Println(val)
		//out <- 200

		go Request(val, out)
	}
}

func Print(in chan int) {
	for val := range in {
		fmt.Println(val)
	}
}

func Request(val string, out chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println(val)
	out <- 200
}