package main

import (
	"fmt"
	"strings"
)

func main() {
	//var is variable
	conName := "Go Conference"
	//const is constant and cannot be changed
	const conTickets = 50
	var remTickets uint = 50
	//array
	//var bookings [50]string
	//slice
	var bookings []string

	//printing datatypes with %T
	fmt.Printf("conName is %T ,conTickets is %T,remTickets is %T\n", conName, conTickets, remTickets)

	//format line print
	fmt.Printf("Welcome to %v booking application\n", conName)
	//new line print
	fmt.Println("We have total of", conTickets, "tickets and", remTickets, "are still available")
	//normal print
	fmt.Print("Get your tickets here to attend\n")

	for {
		//predefined variables
		var firstName string
		var lastName string
		var usertickets uint
		var email string

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter no. of tickets you bought:")
		fmt.Scan(&usertickets)

		remTickets = remTickets - usertickets

		//creating arrays
		//assigning values
		//var bookings = [50]string{"adjfab","kjgvvb"}
		//predefined array

		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		//fmt.Printf("the whole slice:%v\n", bookings)
		//fmt.Printf("the first value:%v\n", bookings[0])
		//fmt.Printf("slice type:%T\n", bookings)
		//fmt.Printf("slice length:%v\n", len(bookings))

		fmt.Printf("Thank You %v %v for booking %v tickets.\nYou'll recieve a confirmation in %v", firstName, lastName, usertickets, email)
		fmt.Printf("\n%v tickets remaining for %v\n", remTickets, conName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings:%v\n", firstNames)
	}

}
