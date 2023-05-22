// In go,everything is a package, and here defualt convention of package name is main
package main

// importing a package named "fmt"
import "fmt"

func main() {
	// var is used for variable declaration
	var conferenceName string = "Go Conference"

	// const is used for constant declaration
	const conferenceTickets int = 50

	var remainingTickets uint = 50

	//Using %T placeholder to identify the type of variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	// Normal way to formatting output
	// fmt.Println("Welcome to the", conferenceName, "booking applcation");

	//Formatting output by using Printf() function
	fmt.Printf("Welcome to the %v booking applcation\n", conferenceName)

	//Using %v placeholder to maps value in clean manner
	fmt.Printf("We have a total of %v tickets available for the %v conference: %v \n", conferenceTickets, conferenceName, remainingTickets)

	fmt.Println("Get your tickets here to attend")

	//asks for the user for their name

	//Basic datatype of Go: Integers(whole numbers) and String(textual data)

	//Other data types of Go: boolean, maps, arrays...

	//Although, Go being statically typed, but still Go has type inference to automatically detect vars based on value assigned to them

	// More on variables: Float point types

	//SyntacticSugar in programming

	//	We can also use this syntax for using variables in Go, but this only works on variables not constant or on defined datatypes like "strings" or "int"

	//  userName := <value>

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//	Taking std Input
	// Using pointers

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	// Remaining tickets logic

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, conferenceName)

}
