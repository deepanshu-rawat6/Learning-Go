// Ignoring Return Values

// A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _

// For example:

// func getPoint() (x int, y int) {
//   return 3, 4
// }

// ignore y value
// x, _ := getPoint()




// Assignment

// In this code snippet, we only need our customer's first name. Ignore the last name so that the code compiles.

package main

import "fmt"

func main() {
	// firstName, lastName := getNames() // This will return a error because lastName is not used.
	firstName, _ := getNames() // This will not return a error because lastName is ignored.
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}

