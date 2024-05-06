package main

import "fmt"

func main() {
	appName := "My First Go Application" // Suagr syntax for defining variables
	const numTickets int = 100           // Defining constant w/ implicit type
	var remainingTickets uint = 100      // Defining variable w/ explicit type

	fmt.Printf("Welcome to the our %v .\n", appName)
	fmt.Printf("We have %v tickets available for sale and we have %v tickets remaining.\n", numTickets, remainingTickets)

	// Printing out the types of the variables and constants
	fmt.Printf("Type of appName: %T, Type of numTickets is %T, Type of remainingTickets is %T\n", appName, numTickets, remainingTickets)

	// Defining var & types for users
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask for user input information
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName) // Scan used for user I/O, need pointer to reference firstName variable where it is stored in memory

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you would like to buy: ")
	fmt.Scan(&userTickets)

	// Update the amount of remaining tickets
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Example of pointer object: %v\n", &firstName)

	// Assigning values to user input
	fmt.Printf("Purchase confirmed for %v %v. A confirmation email will be sent to %v for purchase of %v tickets.\n", firstName, lastName, email, userTickets)
	fmt.Println("Thank you for your purchase!")
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

}
