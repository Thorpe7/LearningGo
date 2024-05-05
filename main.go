package main

import "fmt"

func main() {
	var appName string = "My_Application_Name"
	const numTickets int = 100
	var remainingTickets int = 100

	fmt.Printf("Welcome to the our %v .\n", appName)
	fmt.Printf("We have %v tickets available for sale and we have %v tickets remaining.\n", numTickets, remainingTickets)

	// Defining var & types for users
	var userName string
	var userTickets int

	// Assigning values to user input
	userName = "Tom"
	userTickets = 5
	fmt.Printf("User %v has bought %v tickets.\n", userName, userTickets)

}
