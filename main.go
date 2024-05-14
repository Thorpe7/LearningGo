package main

import (
	"fmt"
	"intro/utils"
	"strconv"
	// For custom imports, have to specify path from 'go.mod' with package name
)

func main() {
	// Call the greetUsers function
	appName := "My First Go Application" // Suagr syntax for defining variables
	const numTickets uint = 100          // Defining constant w/ implicit type
	var remainingTickets uint = 100      // Defining variable w/ explicit type
	greetUsers(appName, numTickets, remainingTickets)

	// var bookings [50]string // Arrays of fixed sizes
	var bookings = make([]map[string]string, 0) // Slices of dynamic sizes
	// Initializing a slice w/ a map, need to specify that the starting size is 0.

	// Replacing slice w/ a map
	var userData = make(map[string]string) // Use make to initialize a map

	// Maps in go cannot have mixed data types! Keys and values must be of the same type

	for {

		// Get user data
		firstName, lastName, email, userTickets := getUserData()
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numTickets"] = strconv.FormatUint(uint64(userTickets), 10)
		// Uses string conversion package to convert uint to stringa and the '10' is specifying that is a base 10 number.

		isValidName, isValidEmail, isValidTickets := utils.ValidateUserData(firstName, lastName, email, userTickets, remainingTickets)

		// Check if the user is trying to buy more tickets than available
		if isValidTickets && isValidName && isValidEmail {

			remainingTickets, bookings := bookTickets(remainingTickets, userTickets, bookings, userData)

			fmt.Printf("Example of pointer object: %v\n", &firstName)

			// Assigning values to user input
			fmt.Printf("Purchase confirmed for %v %v. A confirmation email will be sent to %v for purchase of %v tickets.\n", firstName, lastName, email, userTickets)
			fmt.Println("Thank you for your purchase!")
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			firstNames := getFirstNames(bookings)

			fmt.Printf("There are %v bookings: %v\n", len(bookings), firstNames)

			if remainingTickets == 0 {
				// End the program
				fmt.Println("All the tickets have been purchased. Thank you for your interest in our event!")
				break
			}
		} else {
			fmt.Println("Sorry, your input data is invalid.")
			if !isValidName {
				fmt.Println("Please enter a valid first and last name.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email.")
			}
			if !isValidTickets {
				fmt.Println("Please enter a valid number of tickets.")
			}
		}
	}

	// Example for switch statement
	// var city := "new york"
	//     switch city {
	//        case "new york":
	//             execute code for new york specific booking
	//        case "london", "manchester":
	//             execute code for london specific booking
	//        default:
	//			   execute code for default booking
	// }

}

func greetUsers(confName string, numTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v.", confName)

	fmt.Printf("Welcome to the our %v .\n", confName)
	fmt.Printf("We have %v tickets available for sale and we have %v tickets remaining.\n", numTickets, remainingTickets)

	// Printing out the types of the variables and constants
	fmt.Printf("Type of appName: %T, Type of numTickets is %T, Type of remainingTickets is %T\n", confName, numTickets, remainingTickets)
}

func getFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	// Iterate through the bookings slice, and append only first names
	for _, booking := range bookings { // _ is a blank identifier, used to ignore unused variables
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserData() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, bookings []map[string]string, userData map[string]string) (uint, []map[string]string) {
	// Update the amount of remaining tickets & booked users
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName // Assignment of array element
	bookings = append(bookings, userData) // Append to the slice

	return remainingTickets, bookings
}
