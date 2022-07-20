//https://stackoverflow.com/questions/36857167/how-to-correctly-use-sync-cond
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Mutex{}
	m.Lock() // main gouroutine is owner of lock
	c := sync.NewCond(&m)
	go func() {
		m.Lock() // obtain a lock
		defer m.Unlock()
		fmt.Println("3. goroutine is owner of lock")
		time.Sleep(2 * time.Second) // long computing - because you are the owner, you can change state variable(s)
		c.Broadcast()               // State has been changed, publish it to waiting goroutines
		fmt.Println("4. goroutine will release lock soon (deffered Unlock")
	}()
	fmt.Println("1. main goroutine is owner of lock")
	time.Sleep(1 * time.Second) // initialization
	fmt.Println("2. main goroutine is still lockek")
	c.Wait() // Wait temporarily release a mutex during waiting and give opportunity to other goroutines to change the state.
	// Because you don't know, whether this is state, that you are waiting for, is usually called in loop.
	m.Unlock()
	fmt.Println("Done")
}
