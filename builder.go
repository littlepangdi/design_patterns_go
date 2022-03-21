package main

import "fmt"

/*
	Creating complex objects in steps.
	This pattern allows you to generate objects of different types and forms using the same creation code.
	1. avoid telescopic constructor
	2. in steps
	3. each step can have several concrete patterns
*/
var (
	iceType    = "ice"
	woodenType = "wood"
)

//Our product
type house struct {
	windowType string
	doorType   string
	floor      string
}

//Builder builds house
type builder interface {
	setWindowType()
	setDoorType()
	setFloor()
	getHouse() house
}

type iglooBuilder struct {
	house
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = iceType
}
func (i *iglooBuilder) setDoorType() {
	i.doorType = iceType
}
func (i *iglooBuilder) setFloor() {
	i.floor = iceType
}
func (i *iglooBuilder) getHouse() house {
	return i.house
}

type woodBuilder struct {
	house
}

func (w *woodBuilder) setWindowType() {
	w.windowType = woodenType
}
func (w *woodBuilder) setDoorType() {
	w.doorType = woodenType
}
func (w *woodBuilder) setFloor() {
	w.floor = woodenType
}
func (w *woodBuilder) getHouse() house {
	return w.house
}

//Director uses builder build house and controls steps
type director struct {
	builder builder
}

func newDirector(material string) *director {
	switch material {
	case iceType:
		return &director{
			builder: &iglooBuilder{},
		}
	case woodenType:
		return &director{
			builder: &woodBuilder{},
		}
	default:
		return nil
	}
}

func (d *director) setBuilder(b builder) *director {
	return &director{builder: b}
}
func (d *director) buildHouse() house {
	d.builder.setFloor()
	d.builder.setDoorType()
	d.builder.setWindowType()
	return d.builder.getHouse()
}

func RunBuilder() {
	d := newDirector(iceType)
	iceHouse := d.buildHouse()
	fmt.Printf("Igloo House Door Type: %s\n", iceHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iceHouse.windowType)
	fmt.Printf("Igloo House Floor Type: %s\n", iceHouse.floor)
	d.setBuilder(&woodBuilder{})
	woodHouse := d.buildHouse()
	fmt.Printf("Wooden House Door Type: %s\n", woodHouse.doorType)
	fmt.Printf("Wooden House Window Type: %s\n", woodHouse.windowType)
	fmt.Printf("Wooden House Floor Type: %s\n", woodHouse.floor)
}
