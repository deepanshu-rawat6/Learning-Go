package helper

import "strings"

// ValidateUserInput For exporting a function we just need to capitalize first letter of the function name
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// For valid input values
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// For valid emial
	isValidEmail := strings.Contains(email, "@")

	// For valid user tickets
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // we need to import variable from main.go

	// In Go, we can return any number of return values from a function
	return isValidName, isValidEmail, isValidTicketNumber

}

/*
Scope of variables in Go

* Local : Declaration within a function (can be used only inside the function) -> same applies for block of code

* Package : Declared outside all functions

* Global : Declared outside all functions and uppercase first letter(can be used everywhere across all packages)
Variable Scope : Scope is the region of a program, where a defined variable can be accessed

*/
