package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference" //this is a variable
const conferenceTickets int = 50     //this is a constant
var remainingTickets uint = 50
var bookings []string //Size= how many elements can the array hold

func main() {

	greetUser()

	for {
		//getUserinput() function
		firstName, lastName, email, userTickets := getUserinput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			//call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			fmt.Println()

			if remainingTickets == 0 {
				fmt.Println("[-]Conference is booked out comeback next year!!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name entered is too short..")
			}
			if !isValidEmail {
				fmt.Println("Email Address entered lacks @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("NUmber of tickets entered is invalid")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	//create a slice
	firstNames := []string{}

	//iterate over bookings to get the first name
	for _, booking := range bookings {
		//split the array elements using white space
		var names = strings.Fields(booking)

		//add it to our SLICE using append
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	//generate a bool for our expression
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	//we use strings library to search for the "@" sign
	isValidEmail := strings.Contains(email, "@") ///will be true or false
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserinput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	//get user input
	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email: ")
	fmt.Scan(&email)

	fmt.Println("Number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
	//fmt.Printf("These are all the bookings: %v \n", bookings)

}
