package utils

import (
	"strings"
)

// In Go, the first letter of a function name determines its visibility.
// If the first letter is uppercase, the function is public and can be imported by other packages.
func ValidateUserData(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}

// Variables can also be exported in a similar way. Just capitalize the first letter of the variable name.
