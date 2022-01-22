package main

import "fmt"

func main() {
	/* Syntactic Sugar <varName> := value */
	conferenceName := "Go Conference";
	const conferenceTickets uint8 = 50;
	var remainingTickets uint8 = conferenceTickets;

	/* %T me dirá el data type de una variable/constante */
	fmt.Printf("conferenceTickets type is %T while conferenceName's data type is %T\n", conferenceTickets, conferenceName);

	/* %v formatea por defecto */
	fmt.Printf("Welcome to our %v booking application\n",conferenceName);
	fmt.Printf("We have total of %v tickets available and %v are still available.\n",conferenceTickets,remainingTickets);

	fmt.Println("Get your tickets here to attend");
	fmt.Println("--------------------------------------");
	fmt.Println("----WELCOME TO OUR BOOKING SYSTEM-----");
	fmt.Println("--------------------------------------");

	/*  var array_name = [size]type{optional initial values}*/
	var bookings = [50]string{};
	
	/* otra forma,lo mismo pero sin el = ni los {} */
	// var users [50]string;

	/* un slice se declara igual pero sin el size */
	var bookingsSlice []string;
	
	var firstName string;
	var lastname string;
	var email string;
	var userTickets int;

	/* ask user for their name */
	fmt.Println("Enter your first name:");
	fmt.Scan(&firstName);
	
	fmt.Println("Enter your last name:");
	fmt.Scan(&lastname);
	
	fmt.Println("Enter your email:");
	fmt.Scan(&email);

	fmt.Println("Enter the number of tickets you want:");
	fmt.Scan(&userTickets);

	remainingTickets = remainingTickets - uint8(userTickets);
	bookings[0] = firstName +" "+ lastname;
	bookingsSlice = append(bookingsSlice, firstName +" "+ lastname);

	fmt.Printf("The whole array: %v\n",bookings);
	fmt.Printf("The whole slice: %v\n",bookingsSlice);
	fmt.Printf("The type of the array is: %T\n",bookings);
	fmt.Printf("The type of the slice is: %T\n",bookingsSlice);
	fmt.Printf("The size of the array is: %v\n",len(bookings));
	fmt.Printf("The size of the slice is: %v\n",len(bookingsSlice));

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email shortly at %v.\n",firstName,lastname,userTickets,email);
	fmt.Printf("%v tickets remaining for %v.\n",remainingTickets,conferenceName);
	fmt.Printf("These are all the bookings we have made so far: %v\n",bookingsSlice);

}