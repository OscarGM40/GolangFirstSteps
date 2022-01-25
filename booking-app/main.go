package main

import (
	"fmt"
	"strings"
)
/* package level variables - no syntactic sugar */
var conferenceName = "Go Conference"
const conferenceTickets uint8 = 50
var remainingTickets uint8 = conferenceTickets

func main() {
	/* Syntactic Sugar <varName> := value */
  /* firstName := "Go" */
	/*  var array_name = [size]type{optional initial values}*/
	var bookings = [50]string{}

	/* otra forma,lo mismo pero sin el = ni los {} */
	// var users [50]string;

	/* un slice se declara igual pero sin el size */
	var bookingsSlice []string

	/* %T me dirá el data type de una variable/constante */
	fmt.Printf("conferenceTickets type is %T while conferenceName's data type is %T\n", conferenceTickets, conferenceName)

	/* %v formatea por defecto */
	greetUsers(conferenceName, int(conferenceTickets), int(remainingTickets))

	for {

		city := ""

		fmt.Println("Enter your city(New York | Paris | London | Singapore")
		fmt.Scan(&city)
		switch city {
		case "New York":
			fmt.Println("Welcome to New York")
		case "Paris", "London":
			fmt.Println("Welcome to Europe")
		case "Singapore":
			fmt.Println("Welcome to Singapore")
		default:
			fmt.Println("Welcome to the city of your dreams")
		}

		firstName, lastname, email, userTickets := getUserInput()

		/* validate user input */
		isValidName, isValidEmail, isValidTickets, stillRemaining := validateUserInput(firstName, lastname, email, userTickets, int(remainingTickets))

		// isValidCity := city == "singapore" || city == "new york" || city == "paris"

		if !stillRemaining {
			fmt.Printf("Sorry, we don't have enough tickets to complete your booking. Only %v tickets are available.\n", remainingTickets)
			fmt.Println("Please try again")
			continue
		} else if !isValidTickets {
			fmt.Println("You have to book at least one ticket")
			continue
		} else if !isValidEmail {
			fmt.Println("You have to enter a valid email")
			continue
		} else if !isValidName {
			fmt.Println("You have to enter a valid name")
			continue
		}

		remainingTickets = remainingTickets - uint8(userTickets)
		bookings[0] = firstName + " " + lastname

		bookingsSlice = append(bookingsSlice, firstName+" "+lastname)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The whole slice: %v\n", bookingsSlice)
		fmt.Printf("The type of the array is: %T\n", bookings)
		fmt.Printf("The type of the slice is: %T\n", bookingsSlice)
		fmt.Printf("The size of the array is: %v\n", len(bookings))
		fmt.Printf("The size of the slice is: %v\n", len(bookingsSlice))

		fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email shortly at %v.\n", firstName, lastname, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		firstNames := getFirstNames(bookingsSlice)
		fmt.Printf("These are all the firstnames of the bookings we have made so far: %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets <= 0

		if noTicketsRemaining {
			fmt.Println("Our conference is booked out.Come back next year.")
			break
		}

	}

}

func greetUsers(confName string, confTickets int, rTickets int) {
	fmt.Printf("Welcome to %v booking conference\n", confName)

	fmt.Printf("We have total of %v tickets available and %v are still available.\n", confTickets, rTickets)

	fmt.Println("Get your tickets here to attend")
	fmt.Println("--------------------------------------")
	fmt.Println("----WELCOME TO OUR BOOKING SYSTEM-----")
	fmt.Println("--------------------------------------")
}

func getFirstNames(bookingsSlice []string) []string {
	firstNames := []string{}

	for _, booking := range bookingsSlice {
		firstNames = append(firstNames, strings.Split(booking, " ")[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastname string, email string, userTickets int, remainingTickets int) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastname) >= 2

	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")

	isValidTickets := userTickets > 0

	stillRemaining := (userTickets) <= remainingTickets

	return isValidName, isValidEmail, isValidTickets, stillRemaining
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastname string
	var email string
	var userTickets int

	/* ask user for their name */
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastname, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string,firstName string, lastname string, email string,conferenceName string){

}