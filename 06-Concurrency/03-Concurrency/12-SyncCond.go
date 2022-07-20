//15. Sync.Cond Example One
package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	Action *sync.Cond
}

func main() {
	obj := Object{Action: sync.NewCond(&sync.Mutex{})}

	//allows for recursive call inside variable functions
	//var attachListener func(cd *sync.Cond, fn func())

	//cd like in court is a condition discharge... :)
	attachListener := func(cd *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1) //make sure go routin is runningAlert

		//launch go routine alert wait group after...
		go func() {
			wg.Done() //ok its running alert wait group
			cd.L.Lock()
			defer cd.L.Unlock()

			cd.Wait() //wait for conditional discharge of
			fn()      //run function provided by go routine

			//re-attaches event listener after firewall
			//go attachListener(cd, fn)

		}()

		wg.Wait() //wait for interal go func to fire exit this func
	}

	attachListener(obj.Action, func() {
		fmt.Println("Now,I feel like a Javascript thing: Fire One")
	})

	attachListener(obj.Action, func() {
		fmt.Println("Now,I feel like a Javascript thing: Fire Two")
	})

	attachListener(obj.Action, func() {
		fmt.Println("Now,I feel like a Javascript thing: Fire Three")
	})

	for range time.Tick(time.Second * 2) {
		obj.Action.Broadcast()
		//obj.Action.Signal() //this will signal One Two Three each second in order
	}
}
