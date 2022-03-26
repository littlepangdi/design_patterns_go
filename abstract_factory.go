package main

import (
	"errors"
	"fmt"
)

/*
	based on Factory design pattern, introduce new attributes to Product,
	for example, different "brands" produce same kind of Product.
	in this case, we upgrade our candyFactory
*/
const (
	Dove   = "Dove"
	Nestle = "Nestle"
)

type AbstractFactory interface {
	OrderFor(string) (CandyBar, error)
	MakeChocolate() CandyBar
	MakeToffees() CandyBar
}

type DoveFactory struct{}

//it's not elegant, since Golang does not support default extends.

func (f *DoveFactory) OrderFor(candyType string) (CandyBar, error) {
	switch candyType {
	case Chocolate:
		return f.MakeChocolate(), nil
	case Toffees:
		return f.MakeToffees(), nil
	default:
		return nil, errors.New("not supported candy type")
	}
}
func (f *DoveFactory) MakeChocolate() CandyBar {
	return &chocolate{
		candyBar: candyBar{
			name:  "Dove Chocolate",
			price: "¥100/kg",
		},
	}
}

func (f *DoveFactory) MakeToffees() CandyBar {
	return &toffees{
		candyBar: candyBar{
			name:  "Dove Toffees",
			price: "¥100/kg",
		},
	}
}

type NestleFactory struct{}

func (n *NestleFactory) OrderFor(candyType string) (CandyBar, error) {
	switch candyType {
	case Chocolate:
		return n.MakeChocolate(), nil
	case Toffees:
		return n.MakeToffees(), nil
	default:
		return nil, errors.New("not supported candy type")
	}
}
func (n *NestleFactory) MakeChocolate() CandyBar {
	return &chocolate{
		candyBar: candyBar{
			name:  "Nestle Chocolate",
			price: "¥110/kg",
		},
	}
}

func (n *NestleFactory) MakeToffees() CandyBar {
	return &toffees{
		candyBar: candyBar{
			name:  "Nestle Toffees",
			price: "90/kg",
		},
	}
}

func ShopForFamous(candy string) AbstractFactory {
	if candy == Chocolate {
		return &DoveFactory{}
	} else if candy == Toffees {
		return &NestleFactory{}
	}
	return nil
}

func RunAbstractFactory() {

	if candy, err := ShopForFamous(Chocolate).OrderFor(Chocolate); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("buy %s\n", candy.String())
	}
	if candy, err := ShopForFamous(Toffees).OrderFor(Toffees); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("buy %s\n", candy.String())
	}
}
