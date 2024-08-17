package main

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	destination := "Brazil"
	period := "madrugada"
	countryPercentage := "Poland"
	totalTickets := 1000

	list, err := tickets.MakeList("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//First requeriment call
	go func() {
		total, err := tickets.GetTotalTickets(destination, list)
		if err != nil {
			fmt.Println("Error getting total tickets:", err)
			return
		}
		fmt.Printf("Total tickets for %s: %d\n", destination, total)
	}()

	//Second requeriment call
	go func() {
		count, err := tickets.GetCountByPeriod(period, list)
		if err != nil {
			fmt.Println("Error getting count by period:", err)
			return
		}
		fmt.Printf("Total people traveling in %s: %d\n", period, count)
	}()

	//Third requeriment call
	go func() {
		percentage, err := tickets.PercentageDestination(countryPercentage, totalTickets, list)
		if err != nil {
			fmt.Println("Error getting percentage:", err)
			return
		}
		fmt.Printf("Percentage of people traveling to %s: %.2f%%\n", countryPercentage, percentage)
	}()

	// //Wait until the user press a key for stop
	fmt.Scanln()

}
