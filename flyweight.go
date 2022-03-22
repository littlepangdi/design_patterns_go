package main

import "fmt"

/*
	The Flyweight pattern is a structural design pattern
	that eliminates the need to store all data in each object,
	allowing you to load more objects into a limited amount of memory
	by sharing the same state shared by multiple objects.

	Flyweight refers to object only stores internal states, its construction function should only be called once

	imagine game CS,players with same identification share same clothes
*/

var (
	TerroristType        = "terrorist"
	DefenderType         = "defender"
	TerroristDressType   = "terrorist dress"
	DefenderDressType    = "defender dress"
	dressFactoryInstance = &dressFactory{
		dressMap: make(map[string]fDress),
	}
)

//flyweight interface
type fDress interface {
	getColor() string
}

//concrete flyweight object
type defenderDress struct {
	color string
}

func (d *defenderDress) getColor() string {
	return d.color
}
func newDefenderDress() *defenderDress {
	return &defenderDress{color: "green"}
}

//concrete flyweight object
type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type dressFactory struct {
	dressMap map[string]fDress
}

func (d *dressFactory) getDressByType(dressType string) (fDress, error) {
	if dress, ok := d.dressMap[dressType]; ok {
		return dress, nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == DefenderDressType {
		d.dressMap[dressType] = newDefenderDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

type fPlayer struct {
	dress     fDress
	pType     string
	latitude  int
	longitude int
}

func newFPlayer(playerType, dressType string) *fPlayer {
	if dress, err := dressFactoryInstance.getDressByType(dressType); err != nil {
		return nil
	} else {
		return &fPlayer{
			dress: dress,
			pType: playerType,
		}
	}

}
func (f *fPlayer) fPlayerOperations() {}

type fGame struct {
	terrorists        []*fPlayer
	counterTerrorists []*fPlayer
}

func newGame() *fGame {
	return &fGame{
		terrorists:        make([]*fPlayer, 1),
		counterTerrorists: make([]*fPlayer, 1),
	}
}

func (c *fGame) addTerrorist(dressType string) {
	player := newFPlayer(TerroristType, dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *fGame) addDefender(dressType string) {
	player := newFPlayer(DefenderType, dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

func RunFlyweight() {
	game := newGame()
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addDefender(DefenderDressType)
	game.addDefender(DefenderType)

	fmt.Printf("total number of dresses = %d \n", len(dressFactoryInstance.dressMap))
	for t, d := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", t, d.getColor())
	}

}
