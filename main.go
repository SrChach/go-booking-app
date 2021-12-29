package main

import "fmt"

func main() {
	/** Constants */
	const totalTickets = 50

	/** Variables */
	var conferenceName string = "My Conference"

	/** Self typed variables */
	remainingTickets := 50

	/** You can also print the Type of a variable
	  fmt.Printf(
	    "The type of conferenceName is %T, and remainingTickets is %T\n",
	    conferenceName, remainingTickets,
	  )
	*/

	fmt.Printf("Welcome to the %s booking application\n", conferenceName)
	fmt.Printf(
		"You can book tickets here. %d of %v tickets remaining\n",
		remainingTickets, totalTickets,
	)
}
