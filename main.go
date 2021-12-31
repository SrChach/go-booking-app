package main

import (
	"fmt"
	"go-booking-app/helpers"
)

func main() {

	var remainingTickets uint = 50
	var bookings = make([]string, 0)

	for {
		helpers.PrintBookingInfo(remainingTickets)

		booking := helpers.GetBookingStruct()

		fmt.Printf(
			"The client %s %s has reserved %d tickets\n",
			booking.Name, booking.LastName, booking.Quantity,
		)

		bookings = append(bookings, booking.Name)
		fmt.Printf("Total bookings %v\n", bookings)
	}

}
