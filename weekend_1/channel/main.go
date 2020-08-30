package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 10
		c <- 20
		c <- 30
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	bc := make(chan int, 3)
	bc <- 10
	bc <- 20
	bc <- 30
	fmt.Println(<-bc)
	fmt.Println(<-bc)
	fmt.Println(<-bc)
}

func Print(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello")
}
