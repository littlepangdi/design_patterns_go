package main

import "fmt"

/*
	The Chain of Responsibility pattern is a behavioral design pattern that allows you to send requests down a chain of handlers.
		Once a request is received, each handler can process the request or pass it on to the next handler on the chain.

	For example, hospital takes over patient
*/

type department interface {
	process(*patient)
	setNext(department)
}
type reception struct {
	next department
}

func (d *reception) process(p *patient) {
	if p.registrationDone {
		fmt.Printf("%s already finished registration\n", p.name)
	} else {
		fmt.Printf("reception registers  %s\n", p.name)
		p.registrationDone = true
	}
	d.next.process(p)
}
func (d *reception) setNext(next department) {
	d.next = next
}

type doctor struct {
	next department
}

func (d *doctor) process(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Printf("%s already finished doctor check\n", p.name)
	} else {
		fmt.Printf("doctor checks  %s\n", p.name)
		p.registrationDone = true
	}
	d.next.process(p)
}
func (d *doctor) setNext(next department) {
	d.next = next
}

type cashier struct {
	next department
}

func (d *cashier) process(p *patient) {
	if p.paymentDone {
		fmt.Printf("%s already finished payment\n", p.name)
	} else {
		fmt.Printf("cashier charges  %s\n", p.name)
		p.paymentDone = true
	}
	if d.next != nil {
		d.next.process(p)
	}
}
func (d *cashier) setNext(next department) {
	d.next = next
}

type medicine struct {
	next department
}

func (d *medicine) process(p *patient) {
	if p.medicineDone {
		fmt.Printf("%s already gets medicine\n", p.name)
	} else {
		fmt.Printf("pharmacy provides medicine to %s\n", p.name)
		p.medicineDone = true
	}
	if d.next != nil {
		d.next.process(p)
	}
}
func (d *medicine) setNext(next department) {
	d.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	paymentDone       bool
	medicineDone      bool
}
type hospital struct {
	name string
	next department
}

func newHospital(name string) *hospital {
	hospital := &hospital{
		name: name,
	}
	doctor := &doctor{}
	reception := &reception{}
	cashier := &cashier{}
	medicine := &medicine{}

	switch name {
	case "xie he": //get to pay before get medicine
		hospital.setNext(reception)
		reception.setNext(doctor)
		doctor.setNext(cashier)
		cashier.setNext(medicine)
	default:
		return nil
	}
	return hospital
}

func (d *hospital) process(p *patient) {
	fmt.Printf("%s arrives \"%s\" hospital \n", p.name, d.name)
	d.next.process(p)
}
func (d *hospital) setNext(next department) {
	d.next = next
}

func RunChainOfResponsibility() {
	hospital := newHospital("xie he")
	xiaoming := &patient{
		name: "xiaoming",
	}
	hospital.process(xiaoming)

}
