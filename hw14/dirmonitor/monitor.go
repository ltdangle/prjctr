package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	FileCreatedEvent     = "file_created"
	FileSizeChangedEvent = "filesize_changed"
	FileRemovedEvent     = "file_removed"
)

// FsChangedEvent filesystem changed event.
type FsChangedEvent struct {
	fileName  FileName
	eventName string
}

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

		changed, event := m.compareDirContents(dir, newDir)
		if changed {
			fmt.Println(event)
			m.publish(event)
			dir = newDir
		}
	}
}

func (m *DirMonitor) publish(event *FsChangedEvent) {
	errChan := m.pub.Publish(event)
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
		dir.contents[FileName(file.Name())] = FileSize(fileInfo.Size())
	}

	return dir
}

func (m *DirMonitor) compareDirContents(original *DirContents, new *DirContents) (changed bool, event *FsChangedEvent) {

	for newFileName, newFileSize := range new.contents {
		originalFileSize, exists := original.contents[newFileName]
		if !exists {
			return true, &FsChangedEvent{fileName: newFileName, eventName: FileCreatedEvent}
		}
		if originalFileSize != newFileSize {
			return true, &FsChangedEvent{fileName: newFileName, eventName: FileSizeChangedEvent}
		}
	}

	for originalFileName := range original.contents {
		_, exists := new.contents[originalFileName]
		if !exists {
			return true, &FsChangedEvent{fileName: originalFileName, eventName: FileRemovedEvent}
		}
	}

	return false, nil
}
