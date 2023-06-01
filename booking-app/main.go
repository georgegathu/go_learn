package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTickets = 69
	var remainingTickets = 69

	fmt.Printf("Welcome to our %v booking app\n", confrenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available!\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
	

	var userName string // Ask users for their name
	var userTickets int // Displays how many tickets the user bought

	userName = "Allan"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
