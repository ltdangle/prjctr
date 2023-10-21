package transport

type route struct {
	transport []Transport
}

func NewRoute() *route {
	return &route{}
}

func (r *route) AddTransport(t Transport) {
	r.transport = append(r.transport, t)
}

func (r *route) ShowTransport() string {
	var str string
	for _, t := range r.transport {
		str += t.Name() + "\n"
	}
	return str
}
