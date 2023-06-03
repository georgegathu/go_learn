package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTickets = 69
	var remainingTickets = 69

	fmt.Printf("confrenceName is %T, confrenceTickets is %T, remainingTickets is %T \n", confrenceName, confrenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking app\n", confrenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available!\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
	

	var firstName string // Ask users for their fistname
	var lastName string // Ask users for their lastname
	var email string // Ask users for their email
	var userTickets int // Displays how many tickets the user bought

	// Ask User for Input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
