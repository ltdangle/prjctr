package main

type DirContents struct {
	contents map[string]int64
}

func NewDirContents() *DirContents {
	return &DirContents{contents: make(map[string]int64)}
}
