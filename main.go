package main

import (
	"fmt"
	vechicles "pio/vechicles"
)

func main() {
	fleet := vechicles.NewFleet()
	var vechicle vechicles.Vechicle
	x := 4

	if x == 2 {
		vechicle = vechicles.NewCar("Skoda")
	} else {
		if x%3 == 1 {
			vechicle = vechicles.NewVan("Renault")
			if van, isVan := vechicle.(*vechicles.Van); isVan {
				van.SetDlugosc(300)
			}
		} else {
			vechicle = vechicles.NewVanDlugosc("Renault", 400)
		}
	}

	vechicle.SetId("SY8X7NUL")
	vechicle.SetLiczbaMiejsc(5 - x)

	car := vechicles.NewCar("Skoda")
	car.SetId("SO998434")
	car.SetLiczbaMiejsc(5)
	fleet.AddVechicle(car)
	fleet.AddVechicle(vechicle)
	fleet.AddVechicle(car)
	fleet.RemoveVechicle(vechicle.GetId())
	fleet.RemoveVechicle(vechicle.GetId())

	vechicleInfo := vechicle.GetInfo()

	fmt.Println(vechicleInfo)
}
