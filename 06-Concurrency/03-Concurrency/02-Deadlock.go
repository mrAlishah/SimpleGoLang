package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func longProcess(wg *sync.WaitGroup, w http.ResponseWriter) {
	fmt.Println("\t----------------------------")
	time.Sleep(5 * time.Second) //simulate work ...

	for {
		fmt.Println("\t infinit loop 1\n")
		time.Sleep(time.Second)
	}

	// for true {
	// 	fmt.Println("\t infinit loop 1\n")
	// 	time.Sleep(time.Second)
	// }

	w.Write([]byte("I am long process that finished... \n"))
	wg.Done()

	fmt.Println("\t----------------------------")
}

func main() {
	var wg sync.WaitGroup

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("==========================")

		go longProcess(&wg, w)

		wg.Add(1)
		wg.Wait()
		w.Write([]byte("All long processes are gone \n"))

		fmt.Println("==========================")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
