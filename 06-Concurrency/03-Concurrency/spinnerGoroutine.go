package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, x := range `-\|/` {
			fmt.Printf("\r%c", x) //\r = Carrage Return for replacement
			time.Sleep(delay)
		}
	}
}

func waitAndPrint(delay time.Duration) {
	fmt.Println("working")
	time.Sleep(delay)
	fmt.Println("\nFunction Finished")
}

func main() {
	go spinner(100 * time.Millisecond)
	waitAndPrint(10 * time.Second)

}
