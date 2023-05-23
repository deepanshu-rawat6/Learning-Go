// In go,everything is a package, and here default convention of package name is main
package main

// importing a package named "fmt"
import (
	// importing our own package using "<import-path>/<package name>"
	"booking-app/helper"
	"fmt"
	"strings"
)

// These variables are package level variables, accessible to all functions, rest variables inside functions are local variables
// These variables can be accessed by any file under the same package

// var is used for variable declaration
var conferenceName = "Go Conference"

// const is used for constant declaration
const conferenceTickets int = 50

var remainingTickets uint = 50

// Use arrays for bookings
//var bookings = [50]string{"Deepanshu", "Docker", "Kube"}

//Array implementation
//var bookings [50]string

// Using slice instead of arrays

var bookings []string

// Another way of implementing slice
// var bookings = []string{}

// Another way of implementing slice
// bookings := []string{}

func main() {

	greetUsers()

	//Using %T placeholder to identify the type of variable
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	// Normal way to formatting output
	// fmt.Println("Welcome to the", conferenceName, "booking application");

	//Formatting output by using Printf() function
	//fmt.Printf("Welcome to the %v booking application\n", conferenceName)

	//Using %v placeholder to maps value in clean manner
	//fmt.Printf("We have a total of %v tickets available for the %v conference: %v \n", conferenceTickets, conferenceName, remainingTickets)
	//
	//fmt.Println("Get your tickets here to attend")

	//asks for the user for their name

	//Basic datatype of Go: Integers(whole numbers) and String(textual data)

	//Other data types of Go: boolean, maps, arrays...

	//Although, Go being statically typed, but still Go has type inference to automatically detect vars based on value assigned to them

	// More on variables: Float point types

	//SyntacticSugar in programming

	//	We can also use this syntax for using variables in Go, but this only works on variables not constant or on defined datatype like "strings" or "int"

	//  userName := <value>

	/*
		Loops in Go [only one type of loop is present -> for loop]

		Inner categories:
			Infinite Loop
					for {
						<statements>
					}

					Either use for true { <statements>} or the above syntax for infinite loop

			For-eac Loop
					for <variable for index>, <variable for holding values> := range <structure or slice or arrays to be iterated>{
						<statements>
					}
	*/

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Remaining tickets logic

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstName()

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Check to see if remaining tickets are zero or less than zero

			// noTicketsRemaining := remainingTickets <= 0

			if remainingTickets <= 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Either the first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address, because it doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you've entered is invalid")
			}
		}
	}

	//	 Example of switch case in Go

	/*
		city := "London"

		switch city {
			case "New York":
				// execute code for booking New York tickets

			case "Singapore":
				// execute code for booking New York tickets

			case "London", "Berlin":  // if same logic for London and Berlin
				// execute code for booking New York tickets

			case "Mexico City":
				// execute code for booking New York tickets

			case "Hong Kong":
				// execute code for booking New York tickets

			default:
				// execute code for booking New York tickets
		}
	*/
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets available for the %v conference: %v \n", conferenceTickets, conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	var firstNames []string

	// Using blank identifier (_) to ignore a variable
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//	Taking std Input
	// Using pointers

	fmt.Println("Enter your first name:")
	_, err := fmt.Scan(&firstName)
	if err != nil {
		return "", "", "", 0
	}

	fmt.Println("Enter your last name:")
	_, err = fmt.Scan(&lastName)
	if err != nil {
		return "", "", "", 0
	}

	fmt.Println("Enter your email:")
	_, err = fmt.Scan(&email)
	if err != nil {
		return "", "", "", 0
	}

	fmt.Println("Enter number of tickets:")
	_, err = fmt.Scan(&userTickets)
	if err != nil {
		return "", "", "", 0
	}

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// Array way of adding values to indices
	//bookings[0] = firstName + " " + lastName

	// Slice way of adding values, but accessing values from slices is same as arrays
	bookings = append(bookings, firstName+" "+lastName)

	// For arrays
	//fmt.Printf("The whole array: %v\n", bookings)
	//fmt.Printf("The first element of the array: %v\n", bookings[0])
	//fmt.Printf("Array type: %T\n", bookings)
	//fmt.Printf("Array length: %v\n", len(bookings))

	// For slice
	//fmt.Printf("The whole slice: %v\n", bookings)
	//fmt.Printf("The first element of the slice: %v\n", bookings[0])
	//fmt.Printf("Slice type: %T\n", bookings)
	//fmt.Printf("Slice length: %v\n", len(bookings))

	// Problem with arrays(static sized -> need of dynamic sized structure) -> Slice(dynamic sized, more efficient)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, conferenceName)
}
