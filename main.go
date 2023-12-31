package main

import "fmt"

func main() {
	ConferenceName := "go conference" //Synatic Sugar   :variable creation
	const ConferenceTickets int = 50  //declaring constants
	var remainingTickets int = 50
	var bookings [50]string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T conferenceName is %T\n", ConferenceTickets, remainingTickets, ConferenceName)

	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaiable.\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) //scan function assign value using pointer

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceName)
}
