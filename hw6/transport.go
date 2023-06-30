package main

import (
	"fmt"
	"hw6/transport"
)

func main() {
	route := transport.NewRoute()

	car := transport.NewCar(4)
	_ = car.TakePassengers(&transport.Passenger{})
	_ = car.TakePassengers(&transport.Passenger{})

	train := transport.NewTrain()
	plane := transport.NewPlane()

	route.AddTransport(car)
	route.AddTransport(train)
	route.AddTransport(plane)

	fmt.Println("Transport on the route:")
	fmt.Println(route.ShowTransport())
}
