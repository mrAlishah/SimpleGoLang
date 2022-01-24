package main

import "fmt"

//? this program panics because there is no goroutine
//out side of `main` interacting with the `ch` channel:
//!fatal error: all qoroutines are asleep - deadlock!
func main() {
	var ch chan int 
	ch = make (chan int)
	//two line above can replaced with
	//ch := make(chan int)

	ch <- 10

	v := <- ch

	fmt.Println("recieved", v)


}