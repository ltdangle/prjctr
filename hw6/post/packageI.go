package post

const BoxPackageType = "box"
const EnvelopePackageType = "envelope"

// Package interface.
type PackageI interface {
	PackageType() string
	SenderAddress() string
	RecepientAddress() string
}

// Box implementation of PackageI
type Box struct {
}

func (b *Box) PackageType() string {
	return BoxPackageType
}
func (b *Box) SenderAddress() string {
	return "sender address"
}
func (b *Box) RecepientAddress() string {
	return "recepient address"
}

// Envelope implementation of PackageI
type Envelope struct {
}

func (e *Envelope) PackageType() string {
	return EnvelopePackageType
}
func (e *Envelope) SenderAddress() string {
	return "sender address"
}
func (e *Envelope) RecepientAddress() string {
	return "recepient address"
}
