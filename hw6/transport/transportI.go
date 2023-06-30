package transport

// Transport interface.
type Transport interface {
	Name() string
	Move()
	Stop()
	ChangeSpeed()
	TakePassengers()
	DropPassengers()
}
