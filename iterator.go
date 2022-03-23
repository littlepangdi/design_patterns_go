package main

import "fmt"

/*
	The iterator pattern is a behavioral design pattern
	that allows you to iterate over all the elements in a collection
		without exposing the underlying representation of the collection (lists, stacks, trees, etc.).
	1. make a collection
	2. make an iterator
	3. connect
*/
type object interface {
	play()
}
type collection interface {
	iterator() iterator
}
type dollyCollection struct {
	dollies []object
}

func (d *dollyCollection) getIterator() iterator {
	return &dollyIterator{
		dollies: d.dollies,
	}
}

type iterator interface {
	getNext() object
	hasNext() bool
}

type dollyIterator struct {
	index   int
	dollies []object
}

func (d *dollyIterator) getNext() object {
	if d.hasNext() {
		temp := d.dollies[d.index]
		d.index++
		return temp
	}
	return nil
}
func (d *dollyIterator) hasNext() bool {
	if d.index < len(d.dollies) {
		return true
	}
	return false
}

type dolly struct {
	name string
}

func (d *dolly) play() {
	fmt.Printf("dolly %s is playing\n", d.name)
}

func RunIterator() {
	collection := &dollyCollection{
		dollies: []object{
			&dolly{
				name: "Duck duck",
			},
			&dolly{
				name: "Pikachu",
			},
		},
	}
	iter := collection.getIterator()
	for iter.hasNext() {
		iter.getNext().play()
	}
}
