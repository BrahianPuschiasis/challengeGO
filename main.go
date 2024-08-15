package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	//First requeriment call
	destination := "Venezuela"
	totalTickets, err := tickets.GetTotalTickets(destination)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Total tickets for %s: %d\n", destination, totalTickets)
	}

	//Second requeriment call
	period := "noche"
	totalPeriod, err := tickets.GetCountByPeriod(period)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Total people traveling in %s: %d\n", period, totalPeriod)
	}

	//Third requeriment call
	country := "Brazil"
	percentage, err := tickets.PercentageDestination(country, 1000)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Percentage of people traveling to %s: %.2f%%\n", destination, percentage)
	}
}
