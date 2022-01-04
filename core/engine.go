package core

import "fmt"

var e *engine

type engine struct {
	isRunning bool
}

func (e *engine) Init() bool {
	e.isRunning = true
	return false
}

func (e *engine) Clean() bool {
	return false
}

func (e *engine) Quit() {
}

func (e *engine) Update() {
	fmt.Println("updating...")
}

func (e *engine) Render() {
}

func (e *engine) Events() {
}

func (e *engine) IsRunning() bool {
	return e.isRunning
}

func GetInstance() *engine {
	if e == nil {
		e = &engine{}
	}

	return e
}
