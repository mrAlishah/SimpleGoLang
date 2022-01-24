package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")

			ch <- i

			fmt.Println(time.Now(), i, "sent")

		}

		//TODO: there could be cases where this message is not completed,
		// this is solved in future example
		fmt.Println(time.Now(), "all completed")

	}()

	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "waiting for messages")

	fmt.Println(time.Now(), "recieved", <-ch)
	fmt.Println(time.Now(), "recieved", <-ch)

	//channel capacity is full and closed
	fmt.Println(time.Now(), "recieved", <-ch)
	fmt.Println(time.Now(), "recieved", <-ch)

	fmt.Println(time.Now(), "exiting")
}
