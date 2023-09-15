package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference" //same as {var conferenceName string = "Go conference"}
	conferenceTickets := 50
	var remainingTickets uint = 50
	var bookings []string // slice (expandable array without defined length)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // use { & } to point the variable

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter your tickets: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets) // convert variables to the same dataType
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thankyou %v %v for booking %v tickets. You will recive confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

			fmt.Printf("All the bookings first names are: %v\n", getFirstNames(bookings))

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out :(")
				break // end program
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}

			if !isValidEmail {
				fmt.Println("Your email address doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
		}

	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	// %T gives the dataType of the variable
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // forEach function // to ignore a variable we use { _ }
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidName, isValidEmail, isValidTicketNumber
}
