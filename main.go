package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	// conferenceName := "Go Conference"
	const conferenceTickets int = 50
	// uint only positive values
	var remainingTickets uint = 50
	// slice , not an array .
	// var bookings []string
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// %T - dataType
	// %v - value
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//asking user for there name
		// fmt.Scan used for input/ output. print msgs and collect data, write into a file
		// pointer & printing the address of memory where this variable is stored
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		// the short version := , use only for var
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// .contains will give us true or false
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isInvalidCity := city != "Singapore" && city != "London"
		// !isValidCity := city == "Singapore" && city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			// append - build in function , adds an elements at the end of the slice , grows and returned an update slice value
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking the %v tickets. You will receive a confirmation e-mail at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			printFirstNames(bookings)

			fmt.Printf("The first names of bookings are: %v.\n", firstNames)
			// fmt.Printf("The whole slice: %v.\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("The type of the slice: %T.\n", bookings[0])
			// fmt.Printf("The size of the slice: %v.\n", len(bookings))
			// firstName = "Tom"
			// userTickets = 2

			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out, come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid.")
			}
		}
	}
}

// else if userTickets == remainingTickets {
// 	// do something else
// }

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v!\n", confName)
	fmt.Printf("We have total of %v and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	// _ is a black identifier. we want to use index but we wont be using it
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
}
