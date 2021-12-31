package helpers

import "fmt"

/** This function uses constants defined in 'config.go' */
func PrintBookingInfo(remainingTickets uint) {
	fmt.Printf("Welcome to the %s booking application\n", CONFERENCE_NAME)
	fmt.Printf(
		"You can book tickets here. %d of %v tickets remaining\n",
		remainingTickets, AVAILABLE_TICKETS,
	)
}

/** You can also print the Type of a variable
  fmt.Printf(
    "The type of conferenceName is %T, and remainingTickets is %T\n",
    conferenceName, remainingTickets,
  )
*/
