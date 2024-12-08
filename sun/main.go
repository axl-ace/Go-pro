package main

import "fmt"

func main() {

	var booking [50]string
	var conferenceName = "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Println("--------------------------------------------------------------")

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	fmt.Println("--------------------------------------------------------------")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please fill in the form")

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets: ")
	fmt.Scan(&userTickets)

	booking[0] = firstName + " " + lastName

	remainingTickets = remainingTickets - userTickets
	fmt.Println("--------------------------------------------------------------")

	fmt.Printf("The whole array %v\n", booking)
	fmt.Printf("The first value: %v\n", booking[0])
	fmt.Printf("The length of the array: %v\n", len(booking))
	fmt.Printf("The type of array: %T\n", booking)

	fmt.Println("--------------------------------------------------------------")

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v. \n", remainingTickets, conferenceName)
}
