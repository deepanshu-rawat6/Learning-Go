// Assignment

// We often will need to manipulate strings in our messaging app. For example, adding some personalization by using a customer's name within a template. The concat function should take two strings and smash them together.

//     hello + world = helloworld

// Fix the function signature of concat to reflect its behavior.


package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// func concat(s1, s2 string) string {
// 	return s1 + s2
// }

//  The above function also works beacuse of go syntaxing, of not writing datatypes of variables if they are same and in sequence.

// don't touch below this line

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
