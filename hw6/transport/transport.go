package transport

// Transport interface.
type Transport interface {
	Name() string
	Move()
	Stop()
	ChangeSpeed(speedChange int) error
	TakePassengers(p *passenger) error
	DropPassengers()
}
