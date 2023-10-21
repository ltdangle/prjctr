package main

import (
	"fmt"
	"hw6/transport"
)

func main() {
	route := transport.NewRoute()

	car := transport.NewCar(4)
	_ = car.TakePassengers(transport.NewPassenger())
	_ = car.TakePassengers(transport.NewPassenger())

	train := transport.NewTrain(200)

	planeRegistration := transport.NewRegistration()
	planeRegistration.AddPassengers(transport.NewPassenger())
	plane := transport.NewPlane(planeRegistration)

	route.AddTransport(car)
	route.AddTransport(train)
	route.AddTransport(plane)

	fmt.Println("Transport on the route:")
	fmt.Println(route.ShowTransport())
}
