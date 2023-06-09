package main

import "fmt"

type scene struct {
	// Scene name.
	name string
	// Scene action.
	action string
	// Scene description.
	description string
	// Previous scene.
	previous *scene
	// Next scene.
	next []*scene
}

func (s *scene) addNextScene(scene *scene) {
	scene.previous = s
	s.next = append(s.next, scene)
}

func (s *scene) hasNextScene() bool {
	if len(s.next) > 0 {
		return true
	}
	return false
}
func (s *scene) nextSceneCount() int {
	if len(s.next) == 0 {
		return 0
	}
	return len(s.next) - 1
}

func (s *scene) hasPreviousScene() bool {
	if s.previous == nil {
		return false
	}
	return true
}

func (s *scene) gotoNextScene(sceneIndex int) *scene {
	return s.next[sceneIndex]
}

func (s *scene) gotoPreviousScene() *scene {
	if s.hasPreviousScene() {
		return s.previous
	}
	return nil
}
func (s *scene) String() string {
	var str string
	str += "\n"
	str += "You are here: " + s.breadcrumbs()
	str += "\n\n"
	str += fmt.Sprintf("%s\n\n", s.description)

	// Print next available actions.
	for i, next := range s.next {
		str += fmt.Sprintf("%d: %s\n", i, next.action)
	}
	return str
}

func (s *scene) breadcrumbs() string {
	var breadcrumbs string
	for s.hasPreviousScene() {
		breadcrumbs = s.name + " > " + breadcrumbs
		s = s.gotoPreviousScene()
	}
	breadcrumbs = s.name + " > " + breadcrumbs
	return breadcrumbs
}
