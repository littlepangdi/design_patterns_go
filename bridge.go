package main

import "fmt"

/*
	The bridge pattern separates the abstraction from its implementation so that they can both change independently.
	which is to say,delegate work to the implementation part layer.
*/
type computer interface {
	setPrinter(printer)
	Print()
}

type printer interface {
	printFiles()
}

type mac struct {
	printer printer
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}
func (m *mac) Print() {
	fmt.Println("Print request from mac")
	m.printer.printFiles()
}

type lenovo struct {
	printer printer
}

func (l *lenovo) setPrinter(p printer) {
	l.printer = p
}
func (l *lenovo) Print() {
	fmt.Println("Print request from lenovo")
	l.printer.printFiles()
}

type epson struct {
}

func (p *epson) printFiles() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct {
}

func (h *hp) printFiles() {
	fmt.Println("Printing by a HP Printer")
}
func Print(c computer, p printer) {
	c.setPrinter(p)
	c.Print()
}
func RunBridge() {
	Print(&mac{}, &hp{})
	Print(&lenovo{}, &hp{})
	Print(&mac{}, &epson{})
	Print(&lenovo{}, &epson{})

}
