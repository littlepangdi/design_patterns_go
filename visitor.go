package main

import "fmt"

/*
	The visitor pattern allows you to add behavior to a struct without actually changing the struct.

*/

type visitor interface {
	visit(interface{})
}

type shape interface {
	getType() string
	accept(visitor)
}

//concrete interviewee 1
type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visit(s)
}

func (s *square) getType() string {
	return "Square"
}

//concrete interviewee 2
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visit(c)
}

func (c *circle) getType() string {
	return "Circle"
}

//concrete visitor 1
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visit(s interface{}) {
	switch s.(type) {
	case *circle:
		fmt.Println("Calculating area for circle")
	case *square:
		fmt.Println("Calculating area for square")
	default:
		fmt.Println("Calculating area for unknown")
	}
}

//concrete visitor 2
type coordinateCalculator struct {
	x, y int
}

func (a *coordinateCalculator) visit(s interface{}) {
	switch s.(type) {
	case *circle:
		fmt.Println("Calculating middle point coordinates for for circle")
	case *square:
		fmt.Println("Calculating middle point coordinates for for square")
	default:
		fmt.Println("Calculating middle point coordinates for for unknown")
	}
}

func RunVisitor() {
	square := &square{side: 2}
	circle := &circle{radius: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &coordinateCalculator{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
}
