// In go,everything is a package, and here defualt convention of package name is main
package main

// importing a package named "fmt"
import "fmt"

func main() {
	// var is used for variable declaration
	var conferenceName = "Go Conference"

	// const is used for constant declaration
	const conferenceTickets = 50

	var remainingTickets = 50

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
	var userName string
	var userTickets int

	userName = "Deepanshu"
	userTickets = 2

	fmt.Printf("User: %v \nBooked: %v tickets\n", userName, userTickets)

	//Basic datatype of Go: Integers(whole numbers) and String(textual data)

	//Other data types of Go: boolean, maps, arrays...

	//Although, Go being statically typed, but still Go has type inference to automatically detect vars based on value assigned to them

}
