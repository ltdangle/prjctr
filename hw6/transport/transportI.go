package transport

// TransportI interface.
type TransportI interface {
	Name() string
	Move()
	Stop()
	ChangeSpeed()
	TakePassengers()
	DropPassengers()
}
