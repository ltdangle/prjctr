package main

import (
	"fmt"
	"hw6/post"
)

func main() {
	// Task 1: Post office.
	office := *post.NewSortOffice()

	// Configure senders.
	boxSender := func(p post.Package) {
		fmt.Println("BoxSender is sending " + p.PackageType() + " from " + p.SenderAddress() + " to " + p.RecepientAddress())
	}
	office.AddSender(post.BoxPackageType, boxSender)

	envelopeSender := func(p post.Package) {
		fmt.Println("EnvelopeSender is sending " + p.PackageType() + " from " + p.SenderAddress() + " to " + p.RecepientAddress())
	}
	office.AddSender(post.EnvelopePackageType, envelopeSender)

	// Create packages.
	box := post.NewBox()
	envelope := post.NewEnvelope()

	// Send packages.
	office.Send(box)
	office.Send(envelope)
}
