package main

import (
	"fmt"
	vechicles "pio/vechicles"
)

func main() {
	fmt.Println("HelloThere")
	car := vechicles.NewCar("Mazda")
	car.SetId("SG35602").SetLiczbaMiejsc(7)
	carInfo := car.GetInfo()

	van := vechicles.NewVan("Volkswagen")
	van.SetId("SO984AH").SetLiczbaMiejsc(3)
	van.SetDlugosc(300)
	vanInfo := van.GetInfo()

	fmt.Println(carInfo)
	fmt.Println()
	fmt.Println(vanInfo)

}
