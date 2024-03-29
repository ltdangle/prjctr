package main

import (
	"errors"
	"fmt"
	"time"
)

type LoggingSubscriber struct {
}

func (l LoggingSubscriber) GetNotified(subject any) error {
	fmt.Printf("Logger subscriber got notified about %v\n", subject)
	return nil
}

func (l LoggingSubscriber) GetID() string {
	return "logger"
}

type UserNotifierSubscriber struct {
}

func (l UserNotifierSubscriber) GetNotified(subject any) error {
	time.Sleep(2 * time.Second)
	fmt.Printf("User notifier subscriber got notified about %v\n", subject)
	return errors.New("unable to process notification")
}

func (l UserNotifierSubscriber) GetID() string {
	return "user_notifier"
}
