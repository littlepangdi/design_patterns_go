package main

import "fmt"

/*
	The Strategy pattern is a behavioral design pattern that allows you to define a series of algorithms
		and put each algorithm into a separate class so that the algorithm objects can be replaced with each other.
*/

type meal struct {
	material *material
	cookMethod
}

func (m *meal) setCookMethod(method cookMethod) {
	m.cookMethod = method
}
func makeMeal(ms []material) []*meal {
	result := make([]*meal, 0, len(ms))
	for _, m := range ms {
		meal := &meal{material: &m}
		if m.fragile() {
			meal.setCookMethod(&boil{})
		} else {
			meal.setCookMethod(&fry{})
		}
		meal.cook()
	}
	return result
}

type cookMethod interface {
	cook()
}

type fry struct{}

func (f *fry) cook() {
	fmt.Println("cook with fry method")
}

type boil struct{}

func (b *boil) cook() {
	fmt.Println("cook with boil method")
}

type material interface {
	fragile() bool
}
type vegetable struct{}

func (v *vegetable) fragile() bool {
	return true
}

type meat struct{}

func (v *meat) fragile() bool {
	return false
}

func RunStrategy() {
	materials := []material{
		&meat{},
		&vegetable{},
		&meat{},
	}
	makeMeal(materials)

}
