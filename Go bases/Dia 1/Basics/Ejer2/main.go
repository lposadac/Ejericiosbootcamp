package main

import "fmt"

func main() {

	var (
		temperatura int
		humedad     float32
		presion     float64
	)

	temperatura = 24
	humedad = 88.9
	presion = 45

	fmt.Printf("Temperatura: %v \nHumedad: %v \nPresion: %v\n", temperatura, humedad, presion)

}