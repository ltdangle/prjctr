package post

import "errors"

// Sorting office.
type sortOffice struct {
	senders map[string]func(Package)
}

// Sorting office constructor.
func NewSortOffice() *sortOffice {
	return &sortOffice{
		senders: make(map[string]func(Package)),
	}
}

// Add sender function for each package type.
func (s *sortOffice) AddSender(packageType string, fn func(p Package)) {
	s.senders[packageType] = fn
}

// Send package via configured sender function.
func (s *sortOffice) Send(p Package) error {
	sendFn, ok := s.senders[p.PackageType()]

	if !ok {
		return errors.New("Sender for package is not configured.")
	}

	sendFn(p)

	return nil
}
