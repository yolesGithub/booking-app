package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string
	var userTickets = 2

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
