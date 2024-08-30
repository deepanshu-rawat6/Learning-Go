package main

import (
	"fmt"
)

func tryingNewForLoop(n int) int {
	for i := range n {
		fmt.Println(n - i)
	}
	return n
}

func bulkSend(numMessages int) float64 {
	totalCost := 1.0
	for i := 0; i < numMessages; i++ {
		totalCost += 0.01
	}
	return totalCost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
	tryingNewForLoop(10)
}
