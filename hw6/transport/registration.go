package transport

type registration struct {
	passengers []*passenger
}

func NewRegistration() *registration {
	return &registration{}
}

func (r *registration) AddPassengers(p *passenger) {
	r.passengers = append(r.passengers, p)
}
