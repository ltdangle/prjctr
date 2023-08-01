package main

import (
	"flag"
	"os"
)

func main() {

	path := flag.String("dir", "", "Directory to watch")
	flag.Parse()
	if *path == "" {
		*path, _ = os.Getwd()
	}

	pubsubService := NewPubSub()
	pubsubService.AddSubscriber(LoggingSubscriber{})
	pubsubService.AddSubscriber(UserNotifierSubscriber{})

	monitor := NewDirMonitor(*path, pubsubService)

	monitor.Watch()
}
