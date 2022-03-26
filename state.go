package main

import (
	"fmt"
)

/*
	This pattern abstracts state-related behavior into separate state classes,
		letting the original object delegate work to instance of these classes instead of handling it itself.

	 e.g.  state machine
*/

//abstract state
type state interface {
	handler(ctx *context)
}

//main object
type context struct {
	state state
}

func NewContext() *context {
	return &context{state: &defaultState{}}
}

func (c *context) SetState(state state) {
	c.state = state
}
func (c *context) Handle() {
	c.state.handler(c)
}

//concrete state 1
type defaultState struct{}

func (d *defaultState) handler(ctx *context) {
	fmt.Println("current state: default")
	ctx.SetState(&concreteStateA{})
}

//concrete state 2
type concreteStateA struct{}

func (d *concreteStateA) handler(ctx *context) {
	fmt.Println("current state: concreteStateA")
	ctx.SetState(&concreteStateB{})
}

//concrete state 3
type concreteStateB struct{}

func (d *concreteStateB) handler(ctx *context) {
	fmt.Println("current state: concreteStateB")
	ctx.SetState(&defaultState{})
}
func RunState() {
	context := NewContext()
	context.Handle()
	context.Handle()
	context.Handle()
	context.Handle()
	context.Handle()
}
