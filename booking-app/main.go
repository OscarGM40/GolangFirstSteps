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

	var userName string;
	var userTickets int;
	/* ask user for their name */
	userName = "Tom"
	userTickets =2;
	fmt.Printf("User %v has booked %v tickets\n",userName,userTickets);

}