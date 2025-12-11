package main

import (
	"fmt"
	"strconv"
	"strings"
)

//learnt: if/else else if conditionals len(for checking lenght of strings/slice/)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// var bookings = [50]string{}
	// var bookings [50]string
	// var bookings []string
	bookings := []string{"Precious Morgrace", "Joy Morgrace", "Nana Smith"}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")

		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0

		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
	// Creating a map
	var userData = make(map[string]string)
	userData["firstName"] = "Precious"
	userData["lastName"] = "Morgrace"
	userData["number"] = strconv.FormatUint(uint64(40), 10)

	// Creating a Struct
	type UserData struct {
		firstName              string
		lastName               string
		email                  string
		numberOfTickets        uint
		isOptedForblahblahblah bool
	}
	var usingStructs = UserData{
		firstName: "Savior",
		lastName:  "Brown",
	}
	fmt.Printf("usingStructs: %v\n", usingStructs)
}

//  Program flow
// book ticket
// generate ticket
// send email
