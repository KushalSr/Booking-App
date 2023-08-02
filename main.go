package main

import (
	"booking-app/validation"
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Tech Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validation.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email address")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}

}

// Greeting User

func greetUsers() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend the conference")
}

// Printing First Names of Users

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

//Getting User Inputs

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("\nEnter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("\nEnter your email address: ")
	fmt.Scan(&email)

	fmt.Print("\nHow many tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("\nList of bookings is %v", bookings)

	fmt.Printf("\nThank you %v %v for booking %v tickets. You'll receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("\n%v Tickets remaining for %v", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(4 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n##################################")
	fmt.Printf("\nSending ticket:\n %v To Email address: %v", ticket, email)
	fmt.Println("\n##################################")
}
