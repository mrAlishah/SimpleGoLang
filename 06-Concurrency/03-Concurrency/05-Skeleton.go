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

func main() {
	fmt.Println("Go Program Running")

	go func() {
		for range time.Tick(time.Second * 2) {
			fmt.Println("\t Engine #2 is working:")
		}

	}()
	for range time.Tick(time.Second * 5) {
		fmt.Println("Engine #1 is working:", runtime.NumGoroutine(), " tasks(Go Routines) running")
	}
}
