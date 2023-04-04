// package helper

// // remember to run main.go helper.go - to execute this code OR go run .

// import "strings"

// // we need to export this function to use in the main package by capitalizing the first letter
// // we also added remainingTickets argument in this function and in the main.go file. We can also capitalize thsi variable in the main.go file
// // and skip adding it as an argument
// func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
// 	// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
// 	// the short version := , use only for var
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	// .contains will give us true or false
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

// 	return isValidName, isValidEmail, isValidTicketNumber
// }
