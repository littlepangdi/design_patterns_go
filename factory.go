package main

import (
	"errors"
	"fmt"
)

/*
	Factory ---> create ---> Product
	concrete Product:
		1. shares similar attributes
		2. extends same Product interface

	Product: candyBar
	concrete Product:chocolate,toffees
	Factory create candyBars orders
*/

const (
	Chocolate = "Chocolate"
	Toffees   = "Toffees"
)

// CandyBar : Product
type CandyBar interface {
	setName(string)
	getName() string
	setPrice(string)
	getPrice() string
	String() string
}

//candyBar: concrete Product
type candyBar struct {
	name  string
	price string
}

func (c *candyBar) setName(name string) {
	c.name = name
}
func (c *candyBar) getName() string {
	return c.name
}
func (c *candyBar) setPrice(p string) {
	c.price = p
}
func (c *candyBar) getPrice() string {
	return c.price
}

func (c *candyBar) String() string {
	return fmt.Sprintf("%s, sells for %s", c.getName(), c.getPrice())
}

type chocolate struct {
	candyBar
}

type toffees struct {
	candyBar
}

func newChocolate() CandyBar {
	return &chocolate{
		candyBar: candyBar{
			name:  "Chocolate",
			price: "¥40/kg",
		},
	}
}

func newToffees() CandyBar {
	return &toffees{
		candyBar: candyBar{
			name:  "Toffees",
			price: "¥50/kg",
		},
	}
}

type Factory struct{}

func (f *Factory) OrderFor(candyType string) (CandyBar, error) {
	switch candyType {
	case Chocolate:
		return newChocolate(), nil
	case Toffees:
		return newToffees(), nil
	default:
		return nil, errors.New("not supported candy type")
	}
}

func RunFactory() {
	f := &Factory{}
	if candy, err := f.OrderFor(Chocolate); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produces %s\n", candy.String())
	}
	if candy, err := f.OrderFor(Toffees); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produces %s\n", candy.String())
	}
}
