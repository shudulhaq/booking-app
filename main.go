package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Print("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email : ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets : ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remining for %v\n", remainingTickets, conferenceName)
	}
}
