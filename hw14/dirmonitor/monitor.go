package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type DirMonitor struct {
	path string
	pub  EventPublisher
}

type EventPublisher interface {
	Publish(event any) <-chan error
}

func NewDirMonitor(path string, pub EventPublisher) *DirMonitor {
	return &DirMonitor{
		path: path,
		pub:  pub,
	}
}
func (m *DirMonitor) Watch() {

	dir := m.collectDirData(m.path)

	// Periodically check filesystem for updates.
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Checking filesystem...")

		newDir := m.collectDirData(m.path)

		err := m.compareDirContents(dir, newDir)
		if err != nil {
			fmt.Println(err.Error())
			m.publish()
			dir = newDir
		}
	}
}

func (m *DirMonitor) publish() {
	errChan := m.pub.Publish("dircontents_changed")
	go func() {
		for err := range errChan {
			if err != nil {
				fmt.Println("Failed to publish order processed event")
			}
		}
	}()
}

func (m *DirMonitor) collectDirData(path string) *DirContents {

	dir := NewDirContents()

	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Collect directory information.
	for _, file := range files {
		fileInfo, err := file.Info()

		if err != nil {
			continue
		}
		dir.contents[file.Name()] = fileInfo.Size()
	}

	return dir
}

func (m *DirMonitor) compareDirContents(d1 *DirContents, d2 *DirContents) error {
	if len(d1.contents) != len(d2.contents) {
		return errors.New("Directory contents have changed.")
	}
	return nil
}
