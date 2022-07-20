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
		fmt.Println("I am service A", resource1)

		if resource1 == true {
			resource2 = true
			wg.Done()
		} //else {  resource1 = true }
	}
}

func serviceB(wg *sync.WaitGroup) {

	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service B", resource2)

		if resource2 == true {
			resource3 = true
			wg.Done()
		}
	}
}

func serviceC(wg *sync.WaitGroup) {

	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service C", resource3)

		if resource3 == true {
			resource1 = true
			wg.Done()
		}
	}
}

var round int64

func watchdog(wg *sync.WaitGroup) {
	//Kill the Go funcs reset the system clock
	//if we end up in the same situation in a livelock...
}

var everythingDone bool

func mainProcess(wg *sync.WaitGroup) {

	resource1 = false
	resource2 = false
	resource3 = false

	wg.Add(3)
	go serviceA(wg)
	go serviceB(wg)
	go serviceC(wg)
	fmt.Println("go:everythingDone ", everythingDone)

	wg.Wait()
	if resource1 && resource2 && resource3 {
		everythingDone = true
	}

}

func main() {
	round = 1
	var wg sync.WaitGroup

	fmt.Println("everythingDone ", everythingDone)
	go watchdog(&wg)
	go mainProcess(&wg)

	//* time for finished infinit goroutine
	if !everythingDone {
		time.Sleep(5 * time.Second)
		fmt.Println("waiting...")
	}

	fmt.Println("Finally")
}
