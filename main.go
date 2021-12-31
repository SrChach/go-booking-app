package main

import (
	"fmt"
	"go-booking-app/helpers"
)

func main() {
	/** Constants */
	totalTickets := helpers.AVAILABLE_TICKETS

	/** Variables */
	var conferenceName string = "My Conference"

	/** Self typed variables */
	remainingTickets := 50

	var bookings = []string{}

	/** You can also print the Type of a variable
	  fmt.Printf(
	    "The type of conferenceName is %T, and remainingTickets is %T\n",
	    conferenceName, remainingTickets,
	  )
	*/

	for {
		fmt.Printf("Welcome to the %s booking application\n", conferenceName)
		fmt.Printf(
			"You can book tickets here. %d of %v tickets remaining\n",
			remainingTickets, totalTickets,
		)

		clientName, clientLastName, ticketsNumber := helpers.GetUserBookingData()

		fmt.Printf(
			"The client %s %s has reserved %d tickets\n",
			clientName, clientLastName, ticketsNumber,
		)

		bookings = append(bookings, clientName)
		fmt.Printf("Total bookings %v\n", bookings)
	}

}
