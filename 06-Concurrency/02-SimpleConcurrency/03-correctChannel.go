package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		//2
		fmt.Println(time.Now(), "goroutine taking a nap")

		time.Sleep(2 * time.Second)

		ch <- "hello"
	}()

	//1
	fmt.Println(time.Now(), "waiting for message")

	v := <-ch
	//v := "hi"

	fmt.Println(time.Now(), "recieved", v)

}
