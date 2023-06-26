package transport

type Route struct {
	transport []TransportI
}

func (r *Route) AddTransport(t TransportI) {
	r.transport = append(r.transport, t)
}

func (r *Route) ShowTransport() string {
	var str string
	for _, t := range r.transport {
		str += t.Name() + "\n"
	}
	return str
}
