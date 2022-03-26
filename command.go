package main

import "fmt"

/*
	The Command pattern is a behavioral design pattern that transforms a request into a separate object that contains all the information related to the request.
	This transformation allows you to parameterize methods,
		delay request execution or enqueue them, and implement undoable operations based on different requests.

	client initialization order:
	1. create receiver
	2. create command, relate it to receiver if necessary
	3. create sender, and bonds it to specific command
*/

//sender
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

//command
type command interface {
	execute()
}
type onCommand struct {
	tv device
}

func (c *onCommand) execute() {
	c.tv.on()
}

type offCommand struct {
	tv device
}

func (c *offCommand) execute() {
	c.tv.off()
}

//receiver
type device interface {
	on()
	off()
}
type tv struct{}

func (t *tv) on() {
	fmt.Printf("tv turn on\n")
}
func (t *tv) off() {
	fmt.Printf("tv turn off\n")
}

func RunCommand() {
	tv := &tv{}
	onCommand := &onCommand{tv: tv}
	offCommand := &offCommand{tv: tv}

	onButton := &button{
		command: onCommand,
	}
	offButton := &button{
		command: offCommand,
	}
	onButton.press()
	offButton.press()
}
