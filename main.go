package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("The Rate Limiter")

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.NewTicker(time.Second)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}
}
