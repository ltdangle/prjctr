package main

type FileName string
type FileSize int64

type DirContents struct {
	contents map[FileName]FileSize
}

func NewDirContents() *DirContents {
	return &DirContents{contents: make(map[FileName]FileSize)}
}
