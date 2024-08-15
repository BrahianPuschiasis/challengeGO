package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
}

// Function for read the csv
func ReadTickets(filePath string) ([]string, error) {
	res, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	data := strings.Split(string(res), "\n")

	return data, nil
}

// First requeriment, get all tickets from a specific country on "tickets.csv" file
func GetTotalTickets(destination string) (int, error) {

	data, err := ReadTickets("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	totalTickets := 0

	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")

		if len(line) < 6 {
			continue
		}

		country := line[3]
		if country == destination {
			totalTickets++
		}
	}

	if totalTickets == 0 {
		fmt.Printf("No tickets have been purchased for that destination")
		return 0, fmt.Errorf("destination not found")
	}

	fmt.Printf("Total tickets for %s: %d\n", destination, totalTickets)
	return totalTickets, nil
}

// // Second asignement, get all tickets for people traveling in specific period of time
func GetCountByPeriod(time string) (int, error) {
	data, err := ReadTickets("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	periodMin := 0
	periodMax := 0
	totalPeriod := 0

	switch time {
	case "madrugada":
		periodMin = 0
		periodMax = 6
	case "mañana":
		periodMin = 7
		periodMax = 12
	case "tarde":
		periodMin = 13
		periodMax = 19
	case "noche":
		periodMin = 20
		periodMax = 23
	default:
		fmt.Println("Invalid period")
		return 1, nil
	}

	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")

		if len(line) < 6 {
			continue
		}

		hoursAndMinutes := line[4]
		hours := strings.Split(hoursAndMinutes, ":")
		hourInt, err := strconv.Atoi(hours[0])

		if err != nil {
			fmt.Println("Error converting hours to int:", err)

		}

		if hourInt >= periodMin && hourInt <= periodMax {
			totalPeriod++
		}

	}

	if totalPeriod == 0 {
		fmt.Printf("No tickets have been purchased for that destination")
		return 0, fmt.Errorf("destination not found")
	}
	fmt.Printf("Total people traveling in %s: %d\n", time, totalPeriod)

	return 1, nil

}

// // ejemplo 3
// func AverageDestination(destination string, total int) (int, error) {

// }
