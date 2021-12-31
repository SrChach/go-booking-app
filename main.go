package main

import (
	"fmt"
	"go-booking-app/booking"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

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

		/** Add 1 thread to wait before the main thread closes */
		waitGroup.Add(1)

		/** The "go" keyword make a new thread */
		go SendMailTicket()

		/** Waits until all the open Threads finishes */
		waitGroup.Wait()

		bookings = append(bookings, bookingInfo)
		remainingTickets -= bookingInfo.Quantity
		fmt.Printf("Total bookings %v\n", bookings)
	}

}

func SendMailTicket() {
	time.Sleep(1 * time.Second)
	fmt.Println("###############")
	fmt.Println("Your ticket are being sent")
	fmt.Println("###############")

	/** Closes the thread */
	waitGroup.Done()
}
