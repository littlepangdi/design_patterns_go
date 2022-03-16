package main

import (
	"fmt"
	"sync"
)

//Gives two examples both of which are goroutine safe.
//--------1. use lock ----------
var (
	lock           sync.Mutex //also can be defined as  `var lock=&sync.Mutex{}` but it's not necessary
	singleInstance *single
)

type single struct {
	item string
}

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now:")
			singleInstance = &single{
				item: "ITEM1",
			}
		} else {
			fmt.Println("Single instance already exists.")
		}
	} else {
		fmt.Println("Single instance already exists.")
	}
	return singleInstance
}

func RunSingletonThreadSafe() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

////--------2. use sync.Once ----------
//also, we can create instance in init() func, since init() will only be called once
var once sync.Once

var singleInstance2 *single

func getInstance2() *single {
	if singleInstance2 == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now:")
			singleInstance = &single{
				item: "ITEM2",
			}
		})
	} else {
		fmt.Println("Single instance already exists.")
	}
	return singleInstance2
}

func RunSingletonThreadSafe2() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
