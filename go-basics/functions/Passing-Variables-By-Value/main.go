
// It's critical in Textio that we keep track of how many SMS messages we send on behalf of our clients. Fix the bug to accurately track the number of SMS messages sent.

//  Alter the incrementSends function so that it returns the result after incrementing the sendsSoFar.
//  Alter main() to capture the return value from incrementSends() and overwrite the previous sendsSoFar value.


package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
