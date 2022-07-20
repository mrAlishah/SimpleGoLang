package main

import (
	"fmt"
	"sync"
	"time"
)

var resource1 bool
var resource2 bool
var resource3 bool

func serviceA(wg *sync.WaitGroup) {

	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service A")

		if resource1 == true {
			resource2 = true
			wg.Done()
		} //else {  resource1 = true }
	}
}

func serviceB(wg *sync.WaitGroup) {

	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service B")

		if resource2 == true {
			resource3 = true
			wg.Done()
		}
	}
}

func serviceC(wg *sync.WaitGroup) {

	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service C")

		if resource3 == true {
			resource1 = true
			wg.Done()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	resource1 = false
	resource2 = false
	resource3 = false

	wg.Add(3)
	go serviceA(&wg)
	go serviceB(&wg)
	go serviceC(&wg)

	wg.Wait()
}
