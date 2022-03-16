package main

import "fmt"

/*
Observer--->subscribe--->Subject
Subject--->notify--->Observer

*/

// Subject ------------Subject-------------
type Subject interface {
	register(observer Observer)
	unregister(observer Observer)
	notifyAll()
}

type item struct {
	observerList []Observer
	//self-attributes
	name    string
	inStock bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (t *item) register(observer Observer) {
	t.observerList = append(t.observerList, observer)
}

func (t *item) unregister(observer Observer) {
	t.observerList = removeFromSlice(t.observerList, observer)
}
func (t *item) notifyAll() {
	for _, observer := range t.observerList {
		observer.update(t.name)
	}
}
func (t *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", t.name)
	t.inStock = true
	t.notifyAll()
}

func removeFromSlice(observerList []Observer, observer Observer) []Observer {
	l := len(observerList)
	for i, o := range observerList {
		if o.getID() == observer.getID() {
			observerList[l-1], observerList[i] = observerList[i], observerList[l-1]
			return observerList[:l-1]
		}
	}
	return observerList
}

// Observer ------------Observer-------------
type Observer interface {
	update(subjectMessage string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(subjectMessage string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, subjectMessage)
}
func (c *customer) getID() string {
	return c.id
}

//RunObserverSubject gives simple example of how this works
func RunObserverSubject() {
	shirtItem := newItem("regular-size shirt")

	observerA := &customer{
		id: "fake_email_A",
	}
	observerB := &customer{
		id: "fake_email_B",
	}

	shirtItem.register(observerA)
	shirtItem.register(observerB)

	shirtItem.updateAvailability()
}
