//12-simulate work lab setup
package main

import (
	//Native Go
	"fmt"
	"runtime"
	"time"
	//3rd Party
	//Our Packages
)

func init() {
	fmt.Println("Initializing Go Application")
}

func longProcess() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(5 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
}

func main() {
	fmt.Println("Go Program Running")

	go func() {
		for range time.Tick(time.Second * 2) {
			fmt.Println("\t Engine #2 is working:")
		}

	}()

	for i := 0; i < 2; i++ {
		go longProcess()
	}

	for range time.Tick(time.Second * 2) {
		fmt.Println("Engine #1 is working:", runtime.NumGoroutine(), " tasks(Go Routines) running")
	}
}
