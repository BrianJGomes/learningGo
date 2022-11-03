package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 1)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	location    string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//Call function to greet users
	greetUsers()

	firstName, lastName, email, location, userTickets := getUserInput()

	isValidName, isValidEmail, isValidUserTicket, isValidLocation := validateUserInput(firstName, lastName, email, userTickets, location)

	if isValidEmail && isValidLocation && isValidName && isValidUserTicket && userTickets < remainingTickets {

		bookTicket(userTickets, firstName, lastName, email, location)

		wg.Add(1)

		go sendTicket(firstName, lastName, userTickets, email)

		//Call function print first names
		firstNames := FirstNames()
		fmt.Printf("These are all the bookings so far: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is now booked.  Please come again next year!!")

		}
	} else {
		if !isValidName {
			fmt.Println("Please ensure that your first and last name have at least 2 letters")
		}
		if !isValidEmail {
			fmt.Println("Please ensure that your email address is written correctly")
		}
		if !isValidUserTicket {
			fmt.Println("Please ensure that the number of tickets ordered is a positive number and that there are enough remaining tickets to cover your order")
		}
		if !isValidLocation {
			fmt.Println("Location must be either Houston or Miami")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Initial conference tickets total = %v.\nThere are currently %v tickets left to purchase.\nGet your tickets now!!\n", conferenceTickets, remainingTickets)
}

func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var location string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter the location you would like to attend (Houston, Miami): ")
	fmt.Scan(&location)
	fmt.Printf("Enter the number of tickets you would like to purchase: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, location, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, location string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		location:    location,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket(s).  You will be receiving a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// return remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName  (uint, uint, []string, string, string, string, string)
}

func sendTicket(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v %v purchased %v tickets for the %v", firstName, lastName, userTickets, conferenceName)
	fmt.Println("------------------")
	fmt.Printf("Sending ticket %v to %v\n", ticket, email)
	fmt.Println("------------------")
	wg.Done()
}
