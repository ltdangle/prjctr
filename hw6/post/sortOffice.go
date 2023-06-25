package post

import "errors"

// Sorting office.
type sortOffice struct {
	senders map[string]func(PackageI)
}

// Sorting office constructor.
func NewSortOffice() *sortOffice {
	return &sortOffice{
		senders: make(map[string]func(PackageI)),
	}
}

// Add sender function for each package type.
func (s *sortOffice) AddSender(packageType string, fn func(p PackageI)) {
	s.senders[packageType] = fn
}

// Send package via configured sender function.
func (s *sortOffice) Send(p PackageI) error {
	sendFn, ok := s.senders[p.PackageType()]

	if !ok {
		return errors.New("Sender for package is not configured.")
	}

	sendFn(p)

	return nil
}
