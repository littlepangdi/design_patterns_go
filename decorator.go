package main

import "fmt"

/*
	Allows adding new functionality to an existing object without changing its structure.
	This type of design pattern is a structural pattern, which acts as a wrapper around an existing class.
	imagine a person put on different kinds of clothes which can be taken off or put on.
	contains:
	1. basic component
	2. several layers share same behaviours
*/

type pizza interface {
	getPrice() int
}
type veggeMania struct{}

func (v *veggeMania) getPrice() int {
	return 10
}

type tomatoTopping struct {
	pizza
}

func (t *tomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 3
}

type cheeseTopping struct {
	pizza
}

func (c *cheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 5
}

func RunDecorator() {
	pizzaWithTomatoAndCheese := &cheeseTopping{
		&tomatoTopping{&veggeMania{}},
	}
	fmt.Printf("pizza with tomato and cheese costs $%d \n", pizzaWithTomatoAndCheese.getPrice())
}
