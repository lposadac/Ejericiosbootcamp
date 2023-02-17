package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Ticket struct {
	id     int
	name   string
	email  string
	pais   string
	hora   string
	precio int
}

var tickets []*Ticket

func OpenTickets(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error casting id:", err)
			break
		}
		precio, err := strconv.Atoi(record[5])
		if err != nil {
			fmt.Println("Error casting precio:", err)
			break
		}
		ticket := &Ticket{
			id:     id,
			name:   record[1],
			email:  record[2],
			pais:   record[3],
			hora:   record[4],
			precio: precio,
		}
		tickets = append(tickets, ticket)
	}
	return
}

func CloseTickets() (err error) {
	tickets = nil
	return
}

func PrintTickets() (err error) {
	fmt.Println("Customers:")
	for _, c := range tickets {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Pais: %s, Hora: %s, Precio: %d\n", c.id, c.name, c.email, c.pais, c.hora, c.precio)
	}
	fmt.Println()
	return
}


func GetTotalTickets(destination string) (count int, err error) {
	if len(tickets) == 0 {
		err = errors.New("el archivo se encuentro vacio")
		return
	}

	count = 0
	for _, t := range tickets {
		if t.pais == destination {
			
			count++
		}
	}

	return
}


const (
	madrugada = "madrugada"
	mañana    = "mañana"
	tarde     = "tarde"
	noche     = "noche"
)

func GetCountByPeriod(time string) (count int, err error) {
	switch time {
	case madrugada:
		count, err = TravelersInTimeRange(0, 6)
	case mañana:
		count, err = TravelersInTimeRange(7, 12)
	case tarde:
		count, err = TravelersInTimeRange(13, 19)
	case noche:
		count, err = TravelersInTimeRange(20, 23)
	default:
		err = errors.New("el rango horario no es valido")
	}
	return
}

func TravelersInTimeRange(start, end int) (count int, err error) {

	if len(tickets) == 0 {
		err = errors.New("el archivo se encuentro vacio")
		return
	}

	count = 0
	for _, t := range tickets {
		if len(t.hora) == 5 {
			hour, _ := strconv.Atoi(t.hora[:2])
			if hour >= start && hour <= end {
				
				count++
			}

		} else if len(t.hora) == 4 {
			hour, _ := strconv.Atoi(t.hora[:1])
			if hour >= start && hour <= end {
				
				count++
			}
		} else {
			err = errors.New("formato de tiempo invalido")
			break
		}
	}
	return
}

func AverageDestination(destination string) (f float64, err error) {

	if len(tickets) == 0 {
		err = errors.New("el archivo se encuentro vacio")
		return
	}
	count, err := GetTotalTickets(destination)

	if err != nil {
		fmt.Println("Error rin GetTotalTickets():", err)
		return
	}

	f = float64(count) / float64(len(tickets))
	return
}