package main

import (
	"fmt"
	"go-booking-app/booking"
)

func main() {

	var remainingTickets uint = 50
	var bookings = make([]booking.BookingData, 0)

	for {
		booking.PrintBookingInfo(remainingTickets)

		bookingInfo := booking.GetBookingStruct()

		fmt.Printf(
			"The client %s %s has reserved %d tickets\n",
			bookingInfo.Name, bookingInfo.LastName, bookingInfo.Quantity,
		)

		bookings = append(bookings, bookingInfo)
		remainingTickets -= bookingInfo.Quantity
		fmt.Printf("Total bookings %v\n", bookings)
	}

}
