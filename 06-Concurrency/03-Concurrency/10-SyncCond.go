//https://stackoverflow.com/questions/36857167/how-to-correctly-use-sync-cond
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := sync.Mutex{}
	lock.Lock()

	cond := sync.NewCond(&lock)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		fmt.Println("2) First go routine has started and waits for 1 second before broadcasting condition")

		time.Sleep(1 * time.Second)

		fmt.Println("4) First go routine broadcasts condition")

		cond.Broadcast()
	}()

	go func() {
		defer waitGroup.Done()

		fmt.Println("3) Second go routine has started and is waiting on condition")

		cond.Wait()

		fmt.Println("5) Second go routine unlocked by condition broadcast")
	}()

	fmt.Println("1) Main go routine starts waiting")

	waitGroup.Wait()

	fmt.Println("6) Main go routine ends")
}
