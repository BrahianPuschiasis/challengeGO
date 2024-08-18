package main

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	destination := "Poland"
	period := "noche"
	countryPercentage := "Brazil"
	totalTickets := 1000

	list, err := tickets.MakeList("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// making the channels
	totalTicketsCh := make(chan int)
	countByPeriodCh := make(chan int)
	percentageCh := make(chan float64)
	errorCh := make(chan error, 3) // channel for errors

	// First Requeriment
	go func(destination string, list []tickets.Ticket) {
		total, err := tickets.GetTotalTickets(destination, list)
		if err != nil {
			errorCh <- err
			return
		}
		totalTicketsCh <- total
	}(destination, list)

	// Second Requeriment
	go func(period string, list []tickets.Ticket) {
		count, err := tickets.GetCountByPeriod(period, list)
		if err != nil {
			errorCh <- err
			return
		}
		countByPeriodCh <- count
	}(period, list)

	// Third Requeriment
	go func(countryPercentage string, totalTickets int, list []tickets.Ticket) {
		percentage, err := tickets.PercentageDestination(countryPercentage, totalTickets, list)
		if err != nil {
			errorCh <- err
			return
		}
		percentageCh <- percentage
	}(countryPercentage, totalTickets, list)

	// Results prints in console
	for i := 0; i < 3; i++ {
		select {
		case total := <-totalTicketsCh:
			fmt.Printf("Total tickets for %s: %d\n", destination, total)
		case count := <-countByPeriodCh:
			fmt.Printf("Total people traveling in %s: %d\n", period, count)
		case percentage := <-percentageCh:
			fmt.Printf("Percentage of people traveling to %s: %.2f%%\n", countryPercentage, percentage)
		case err := <-errorCh:
			fmt.Println("Error:", err)
		}
	}

	// press a key to stop the exection
	fmt.Scanln()

}
