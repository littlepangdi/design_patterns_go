package main

import "fmt"

/*
	The memento pattern is a behavioral design pattern
		that allows saving and restoring an object's previous state without exposing the object's implementation details.

	1. object creates memento proactively
	2. caretaker stores memento
    3. caretaker supplies historical mementos
	4. object restores from memento
*/

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{
		state: o.state,
	}
}
func (o *originator) restoreFromMemento(m *memento) {
	o.state = m.state
}
func (o *originator) setState(s string) {
	o.state = s
}
func (o *originator) getState() string {
	return o.state
}

type memento struct {
	state string
}
type caretaker struct {
	memos []*memento
}

func newCaretaker() *caretaker {
	return &caretaker{
		memos: make([]*memento, 0),
	}
}
func (c *caretaker) addMemento(m *memento) {
	c.memos = append(c.memos, m)
}
func (c *caretaker) getMemento(i int) (m *memento) {
	return c.memos[i]
}

func RunMemento() {
	object := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", object.getState())

	caretaker := newCaretaker()
	caretaker.addMemento(object.createMemento())
	object.setState("B")
	caretaker.addMemento(object.createMemento())
	fmt.Printf("Originator Current State: %s\n", object.getState())

	object.setState("C")
	fmt.Printf("Originator Current State: %s\n", object.getState())

	object.restoreFromMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", object.getState())
	object.restoreFromMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", object.getState())
}
