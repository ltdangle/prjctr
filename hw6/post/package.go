package post

const BoxPackageType = "box"
const EnvelopePackageType = "envelope"

// Package interface.
type Package interface {
	PackageType() string
	SenderAddress() string
	RecepientAddress() string
}

// box implementation of Package
type box struct {
}

func NewBox() *box {
	return &box{}
}

func (b *box) PackageType() string {
	return BoxPackageType
}
func (b *box) SenderAddress() string {
	return "sender address"
}
func (b *box) RecepientAddress() string {
	return "recepient address"
}

// Envelope implementation of Package
type Envelope struct {
}

func NewEnvelope() *Envelope {
	return &Envelope{}
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
