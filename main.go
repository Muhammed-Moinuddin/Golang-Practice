package main

import "fmt"

func main() {
	bookingName := "Golang Project" //syntactic sugar (does not work with constants and data types)
	const bookingTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to my First %v - booking app\n", bookingName)
	fmt.Println("We have total of", bookingTickets, "and this much are left", remainingTickets)
	fmt.Println("Fork and Clone it to use it")

	fmt.Println(bookingName)

	//Printing data types
	fmt.Printf("Booking name type is %T and booking tickets type is %T \n", bookingName, bookingTickets)

	fmt.Println(bookingName)
	fmt.Println(&bookingName) //Pointer

	//Below Data types
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email?")
	fmt.Scan(&email)

	fmt.Println("How much tickets do you want?")
	fmt.Scan(&userTickets)

	fmt.Printf("Hello %v %v. Thank you for buying %v tickets. You'll get confirmation email at: %v \n", firstName, lastName, userTickets, email)

}
