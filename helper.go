package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, location string) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets
	isValidLocation := location == "Houston" || location == "Miami"
	return isValidName, isValidEmail, isValidUserTicket, isValidLocation
}
