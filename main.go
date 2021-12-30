package main

import "fmt"

func getClientName() string {
	var clientName string
	fmt.Println("Please enter your name")
	fmt.Scan(&clientName)

	return clientName
}

func getNumberOfTickets() int {
	var ticketsNumber int

	fmt.Println("Please enter the number of tickets you want")
	fmt.Scan(&ticketsNumber)

	return ticketsNumber
}

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

	var clientName string = getClientName()
	var ticketsNumber int = getNumberOfTickets()

	fmt.Printf("The client %s has reserved %d tickets\n", clientName, ticketsNumber)
}
