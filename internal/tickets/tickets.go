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

	//reading the csv and spliting it with a coma
	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")

		if len(line) < 6 {
			continue
		}

		//country is the country from each line on csv file, destination is from the string we pass to the function
		country := line[3]
		if country == destination {
			totalTickets++
		}
	}

	//in case arent tickets purchased
	if totalTickets == 0 {
		// fmt.Printf("No tickets have been purchased for that destination")
		return 0, fmt.Errorf("destination not found")
	}

	// fmt.Printf("Total tickets for %s: %d\n", destination, totalTickets)
	return totalTickets, nil
}

// Second requeriment, get all tickets for people traveling in specific period of time
func GetCountByPeriod(time string) (int, error) {
	data, err := ReadTickets("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	periodMin := 0
	periodMax := 0
	totalPeriod := 0

	//depending on the time string we get from the function, have 4 cases of periods
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
		return 0, fmt.Errorf("invalid period: %s", time) // Devolver un error apropiado
	}

	//spliting the csv against a coma for have the data cleared
	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")

		if len(line) < 6 {
			continue
		}

		//this is the hours and minutes we get from csv example "17:34"
		hoursAndMinutes := line[4]
		//here we split it against ":" for get in this case "17"
		hours := strings.Split(hoursAndMinutes, ":")
		//we parse it as int for use numeric expresions
		hourInt, err := strconv.Atoi(hours[0])
		if err != nil {
			fmt.Println("Error converting hours to int:", err)

		}

		//depending on the case, we compare the min and max hour from that period, against the one we insert on the function and count if checks
		if hourInt >= periodMin && hourInt <= periodMax {
			totalPeriod++
		}

	}

	if totalPeriod == 0 {
		return 0, fmt.Errorf("no tickets have been purchased for that destination")
	}
	// fmt.Printf("Total people traveling in %s: %d\n", time, totalPeriod)

	return totalPeriod, nil

}

// Third requeriment, percentage of people traveling on a specific country on one day
func PercentageDestination(destination string, total int) (float64, error) {

	// calling first requeriment function and asignt the return on a variable
	peopleTraveling, err := GetTotalTickets(destination)
	if err != nil {
		return 0, fmt.Errorf("getting total tickets for destination: %v", err)
	}

	// check if tickets isnt 0 for avoid exception
	if total == 0 {
		return 0, fmt.Errorf("total tickets cannot be zero")
	}

	// percentage calculation
	percentage := (float64(peopleTraveling) / float64(total)) * 100
	// fmt.Printf("Percentage of people traveling to %s: %.2f%%\n", destination, percentage)

	return percentage, nil
}
