package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func longProcess(wg *sync.WaitGroup, w http.ResponseWriter) {
	fmt.Println("\t----------------------------")

	fmt.Println("go:before time.Sleep")
	time.Sleep(10 * time.Second) //simulate work ...

	w.Write([]byte("I am long process that finished... \n"))
	fmt.Println("go:w.Write")
	wg.Done()
	fmt.Println("go:wg.Done()")
	fmt.Println("\t----------------------------")
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("before http Handle")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("==========================")
		fmt.Println("before concurrency")
		go longProcess(&wg, w)
		fmt.Println("after concurrency")

		wg.Add(1)
		fmt.Println("wg.Add(1)")
		wg.Wait()
		fmt.Println("wg.Wait()")
		w.Write([]byte("All long processes are gone \n"))
		fmt.Println("w.Write")
		fmt.Println("==========================")
	})

	fmt.Println("http.ListenAndServe")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
