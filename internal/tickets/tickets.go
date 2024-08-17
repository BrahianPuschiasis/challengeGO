package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          int
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      int
}

func MakeList(filePath string) ([]Ticket, error) {
	res, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	data := strings.Split(string(res), "\n")
	var tickets []Ticket

	for i := 0; i < len(data); i++ {
		line := strings.Split(data[i], ",")
		if len(line) < 6 {
			continue
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("error converting ID: %v", err)
		}

		precio, err := strconv.Atoi(line[5])
		if err != nil {
			return nil, fmt.Errorf("error converting Precio: %v", err)
		}

		ticket := Ticket{
			ID:          id,
			Nombre:      line[1],
			Email:       line[2],
			PaisDestino: line[3],
			HoraVuelo:   line[4],
			Precio:      precio,
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// First requeriment, get all tickets from a specific country on "tickets.csv" file
func GetTotalTickets(destination string, tickets []Ticket) (int, error) {
	total := 0

	for _, ticket := range tickets {
		if ticket.PaisDestino == destination {
			total++
		}
	}

	return total, nil
}

// Second requeriment, get all tickets for people traveling in specific period of time
func GetCountByPeriod(time string, tickets []Ticket) (int, error) {

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
		return 0, fmt.Errorf("invalid period: %s", time)
	}

	//reading the list
	for _, ticket := range tickets {

		//here we split it against ":" for get in this case "17" instead "17:30"
		hours := strings.Split(ticket.HoraVuelo, ":")
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

	return totalPeriod, nil

}

// Third requeriment, percentage of people traveling on a specific country on one day
func PercentageDestination(destination string, total int, tickets []Ticket) (float64, error) {
	// check if tickets isnt 0 for avoid exception

	if total == 0 {
		return 0, fmt.Errorf("total tickets cannot be zero")
	}

	// finding how much tickets where selled today in this destination
	peopleTraveling := 0

	for _, ticket := range tickets {
		if ticket.PaisDestino == destination {
			peopleTraveling++
		}
	}

	// percentage calculation
	percentage := (float64(peopleTraveling) / float64(total)) * 100

	return percentage, nil
}
