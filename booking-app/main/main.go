package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	// "strconv"
	"strings"
	"sync"
)

/* package level variables - no syntactic sugar */
var conferenceName = "Go Conference"
var remainingTickets uint8 = helper.ConferenceTickets
var bookingsMap = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastname        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

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
	fmt.Printf("conferenceTickets type is %T while conferenceName's data type is %T\n", helper.ConferenceTickets, conferenceName)

	/* %v formatea por defecto */
	greetUsers(conferenceName, int(helper.ConferenceTickets), int(remainingTickets))

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
		isValidName, isValidEmail, isValidTickets, stillRemaining := helper.ValidateUserInput(firstName, lastname, email, userTickets, int(remainingTickets))

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

		bookTicket(uint8(userTickets), firstName, lastname, email)
		fmt.Printf("List of bookings looks like this: %v\n", bookingsMap)

		wg.Add(1)
		go sendTicket(uint8(userTickets), firstName, lastname, email)
		wg.Wait()

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The whole slice: %v\n", bookingsSlice)
		fmt.Printf("The type of the array is: %T\n", bookings)
		fmt.Printf("The type of the slice is: %T\n", bookingsSlice)
		fmt.Printf("The size of the array is: %v\n", len(bookings))
		fmt.Printf("The size of the slice is: %v\n", len(bookingsSlice))

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

func bookTicket(userTickets uint8, firstName string, lastname string, email string) {
	remainingTickets = remainingTickets - uint8(userTickets)

	/* create a map for a user */
	var userData = UserData{
		firstName:       firstName,
		lastname:        lastname,
		email:           email,
		numberOfTickets: uint(userTickets),
	}
	/* 	var user = make(map[string]string)

	   	user["firstName"] = firstName
	   	user["lastName"] = lastname
	   	user["email"] = email
	   	user["numberOfTickets"] = strconv.FormatInt((int64(userTickets)), 10) */

	bookingsMap = append(bookingsMap, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email shortly at %v.\n", firstName, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint8, firstName string, lastName string, email string) {
	time.Sleep(time.Second * 10)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
