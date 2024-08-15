package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.GetTotalTickets("Brazil")
	tickets.GetCountByPeriod("madrugada")
	tickets.PercentageDestination("Poland", 1000)
}
