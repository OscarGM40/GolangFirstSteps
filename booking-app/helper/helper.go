package helper

import (
	"strings"
)

const ConferenceTickets uint8 = 50

func ValidateUserInput(firstName string, lastname string, email string, userTickets int, remainingTickets int) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastname) >= 2

	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")

	isValidTickets := userTickets > 0

	stillRemaining := (userTickets) <= remainingTickets

	return isValidName, isValidEmail, isValidTickets, stillRemaining
}
