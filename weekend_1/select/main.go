package main

import (
	"fmt"
	"strconv"
)

func main() {
	taskChanA := make(chan int)
	taskChanB := make(chan string)
	taskChanC := make(chan bool)

	go taskA(taskChanA, 10)
	go taskB(taskChanB, 10)
	go taskC(taskChanC, 10)

label1:
	for {
		select {
		case val := <-taskChanA:
			fmt.Println("A:", val)
		case val := <-taskChanB:
			fmt.Println("B:", val)
		case val := <-taskChanC:
			fmt.Println("C:", val)
		default:
			break label1
		}
	}
}

func taskA(c chan int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}
}

func taskB(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- strconv.Itoa(i)
	}
}

func taskC(c chan bool, n int) {
	for i := 0; i < n; i++ {
		c <- i%2 == 0
	}
}