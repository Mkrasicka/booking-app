package main

import (
	// "booking-app/helper"
	"fmt"
	"sync"

	// "strconv"
	"strings"
	"time"
	// "strings"
)

var conferenceName string = "Go Conference"

// conferenceName := "Go Conference"
const conferenceTickets int = 50

// uint only positive values
var remainingTickets uint = 50

// slice , not an array .
// var bookings []string
// list of strings but now we are using list of maps
// var bookings = []string{}
// we are creating empty LIST of map and 0 is for initializing the size which will increase
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	// greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// %T - dataType
	// %v - value
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// for {

	firstName, lastName, email, userTickets := getUserInput()
	// helper. if it will be in the other package
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	// isInvalidCity := city != "Singapore" && city != "London"
	// !isValidCity := city == "Singapore" && city == "London"

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)

		go sendingTicket(userTickets, firstName, lastName, email)
		// because this function is returning firstNames and we are assigning it to the variable firstNames so we can print it in the line below
		firstNames := getFirstNames()
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
			// break
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
	wg.Wait()
}

// }

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	// the short version := , use only for var
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// .contains will give us true or false
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

//	else if userTickets == remainingTickets {
//		// do something else
//	}
//
// func greetUsers(confName string, confTickets int, remainingTickets uint)
func greetUsers() {
	fmt.Printf("Welcome to %v!\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	// _ is a black identifier. we want to use index but we wont be using it
	for _, booking := range bookings {
		// extract the first name from the booking
		// var names = strings.Fields(booking)
		// var firstName = names[0] - that was for booking when was a string
		// var firstName = bookings["firstName"] - extracting name from the map
		// firstNames = append(firstNames, booking["firstName"]) - that syntax for maps
		firstNames = append(firstNames, booking.firstName)
		// syntax above is from struct
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// create a struct
	// create an EMPTY MAP for a user
	// var userData = make(map[string]string)
	// // key value pair
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// // in map we can't mix data types. that's why we need to use build-in convert function from string convertion package. 10 stands for decimal number
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// user data object
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName
	// append - build in function , adds an elements at the end of the slice , grows and returned an update slice value
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking the %v tickets. You will receive a confirmation e-mail at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendingTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	// Sprintf saving text into a variable because Printf is only printing
	var ticket = fmt.Sprintf("%v tickets for %v %v.\n", userTicket, firstName, lastName)
	fmt.Println("###############################")
	fmt.Printf("Sending tickets: \n %v \n to an email address %v\n", ticket, email)
	fmt.Println("###############################")
	wg.Done()
}
