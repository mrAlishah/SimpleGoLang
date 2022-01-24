package main

import "fmt"

func main(){
	//print nothing, because main thread is fastest and completed goroutine
	go hello()
}

func hello() {
	fmt.Println("It's most likely you will never see this.")
}