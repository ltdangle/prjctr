package library

// User struct.
type User struct {
	name string
}

// Customer embeds User.
type Customer struct {
	User
}

// NewCustomer constructor.
func NewCustomer(name string) *Customer {
	return &Customer{
		User: User{name: name},
	}
}

// Manager embeds User.
type Manager struct {
	User
}

// NewManager constructor.
func NewManager(name string) *Manager {
	return &Manager{
		User: User{name: name},
	}
}
