package main

import "fmt"

/*
	The Mediator pattern is a behavioral design pattern that allows you to reduce cluttered dependencies between objects.
	This pattern restricts direct interaction between objects,
	forcing them to cooperate through a mediator object.

	Pilots communicate through the control tower
*/

var (
	cargoType = "cargo"
	civilType = "civil"
)

type aircraft interface {
	arrive()
	depart()
	//called by mediator
	permitLanding()
}
type airportMediator interface {
	//called by aircraft
	canArrive(aircraft) bool
	//called by aircraft
	notifyAboutDeparture()
}

type cargoAircraft struct {
	name     string
	mediator airportMediator
}

func newAircraft(name string, style string, m airportMediator) aircraft {
	switch style {
	case cargoType:
		return &cargoAircraft{
			name:     name,
			mediator: m,
		}
	case civilType:
		return &civilAircraft{
			name:     name,
			mediator: m,
		}
	default:
		return nil
	}
}
func (c *cargoAircraft) arrive() {
	if c.mediator.canArrive(c) {
		fmt.Printf("cargo %s : Arrived\n", c.name)
	} else {
		fmt.Printf("cargo %s : Arrival blocked, waiting\n", c.name)
	}
}

func (c *cargoAircraft) depart() {
	fmt.Printf("cargo %s : Leaving\n", c.name)
	c.mediator.notifyAboutDeparture()
}
func (c *cargoAircraft) permitLanding() {
	fmt.Printf("cargo %s : Arrival permitted\n", c.name)
	c.arrive()
}

type civilAircraft struct {
	name     string
	mediator airportMediator
}

func (c *civilAircraft) arrive() {
	if c.mediator.canArrive(c) {
		fmt.Printf("civil %s : Arrived\n", c.name)
	} else {
		fmt.Printf("civil %s : Arrival blocked, waiting\n", c.name)
	}
}

func (c *civilAircraft) depart() {
	fmt.Printf("civil %s : Leaving\n", c.name)
	c.mediator.notifyAboutDeparture()
}
func (c *civilAircraft) permitLanding() {
	fmt.Printf("civil %s : Arrival permitted\n", c.name)
	c.arrive()
}

type airportControlTower struct {
	space  int
	blocks []aircraft
}

func (a *airportControlTower) canArrive(ac aircraft) bool {
	if a.space >= 1 {
		a.space--
		return true
	} else {
		a.blocks = append(a.blocks, ac)
		return false
	}
}

func (a *airportControlTower) notifyAboutDeparture() {
	a.space++
	if len(a.blocks) > 1 {
		a.blocks[0].permitLanding()
		a.blocks = a.blocks[1:]
	}
}

func RunMediator() {
	ct := &airportControlTower{
		space:  2,
		blocks: make([]aircraft, 0),
	}
	airCraftList := []aircraft{
		newAircraft("no.1", cargoType, ct),
		newAircraft("no.2", cargoType, ct),
		newAircraft("no.1", civilType, ct),
		newAircraft("no.2", civilType, ct),
		newAircraft("no.3", civilType, ct),
	}
	for _, air := range airCraftList {
		air.arrive()
	}
	for _, air := range airCraftList {
		air.depart()
	}
}
