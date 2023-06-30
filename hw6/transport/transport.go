package transport

// Transport interface.
type Transport interface {
	Name() string
	Move()
	Stop()
	ChangeSpeed()
	TakePassengers(p *Passenger) error
	DropPassengers()
}
