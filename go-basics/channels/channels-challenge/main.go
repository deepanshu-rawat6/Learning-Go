package main

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})

	done := make(chan struct{})

	go ponger(pings, pongs, done)
	go pinger(pings, pongs, numPings, done)

	<-done
}

// don't touch below this line

func pinger(pings, pongs chan struct{}, numPings int, done chan struct{}) {
	go func() {
		sleepTime := 50 * time.Millisecond
		for i := 0; i < numPings; i++ {
			fmt.Println("ping", i, "sent")
			pings <- struct{}{}
			time.Sleep(sleepTime)
			sleepTime *= 2
		}
		close(pings)
	}()
	i := 0
	for range pongs {
		fmt.Println("pong", i, "got")
		i++
	}
	fmt.Println("pongs done")
	done <- struct{}{}
}

func ponger(pings, pongs chan struct{}, done chan struct{}) {
	i := 0
	for range pings {
		fmt.Println("ping", i, "got", "pong", i, "sent")
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}

