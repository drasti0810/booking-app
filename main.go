package main
import (
	"fmt"
	"booking-app/helper"
	"strconv"
)
var name = "Goa"
const ticketTotal = 50 //constant
var ticketCount uint = 50 //variable
var bookingsslice = []string{} //Initialising slice
var bookingsmap = make([]map[string]string, 0) //Intialising list of maps
func main(){
	greetUsers()
	for {	
		// ask user for input
		userName, userTickets := getUserInput()
		// validate user input
		isValidName, isValidTicket := helper.ValidateInput(userName, userTickets, ticketCount)
		if isValidName && isValidTicket {
			// book ticket
			bookTicket(userName, userTickets)
			// fmt.Printf("First user %v\n", bookings[0])
			// fmt.Printf("Type %T \n", bookings)
			// fmt.Printf("No. of bookings: %v\n", len(bookings))
			if ticketCount == 0 {
				fmt.Println("No tickets remaining, come back next year")
				break
			}
		}	else {
				if !isValidName {
					fmt.Println("Name is too short")
				}
				if !isValidTicket {
					fmt.Println("No is wrong")
				}
		}
		}
}
func greetUsers() {
	fmt.Println("Welcome to your Final Destination.....", name)
	fmt.Println("Buckle Up!!")
	fmt.Printf("There are %v total tickets and %v are still available\n", ticketTotal, ticketCount)
}	
func getUserInput() (string, uint) {
	var userName string
	var userTickets uint
	//ask user input
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets you want")
	fmt.Scan(&userTickets)
	return userName, userTickets
}
func bookTicket(userName string, userTickets uint) {
	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)
	ticketCount = ticketCount - userTickets
	//Create map for details
	var userData = make(map[string]string) //create a map with key and value as string data-type
	userData["userName"] = userName
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) //convert ticketCount to string from uint using builtin strconv module

	bookingsslice = append(bookingsslice, userName)
	bookingsmap = append(bookingsmap, userData)
	fmt.Printf("Remaining tickets %v\n", ticketCount)
	fmt.Printf("These are all the bookings stored in slice %v\n", bookingsslice)
	fmt.Printf("These are all the bookings stored in map %v\n", bookingsmap)
}