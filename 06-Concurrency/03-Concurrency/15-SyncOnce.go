//17. Sync.Once
package main

import (
	"fmt"
	"sync"
)

var doOnce sync.Once

func runMe() {

	doOnce.Do(func() {
		fmt.Println("Run Once !!!")
	})

	fmt.Println("I have been run.")
}

func main() {
	runMe()
	runMe()
}
