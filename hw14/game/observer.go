package main

import "fmt"

type Observer interface {
	GetNotified(subject any)
	GetID() string
}

// LoggerObserver.
type LoggerObserver struct {
}

func (l LoggerObserver) GetNotified(subject any) {
	fmt.Printf("\nLogger observer got notified about %v\n", subject)
}

func (l LoggerObserver) GetID() string {
	return "logger"
}

// ObserverRegistrar.
type ObserverRegistrar struct {
	observers map[string]Observer
}

func NewObserverRegistrar() *ObserverRegistrar {
	return &ObserverRegistrar{
		observers: make(map[string]Observer),
	}
}

func (s *ObserverRegistrar) NotifyAll(subject any) {
	for _, o := range s.observers {
		o.GetNotified(subject)
	}
}

func (s *ObserverRegistrar) Register(o Observer) {
	s.observers[o.GetID()] = o
}

func (s *ObserverRegistrar) Unregister(o Observer) {
	delete(s.observers, o.GetID())
}
