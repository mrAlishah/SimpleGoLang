//14. Wait Group Use Case
package main

import (
	"net/http"
	"sync"
	"time"
)

func longProcess(wg *sync.WaitGroup, w http.ResponseWriter) {

	time.Sleep(3 * time.Second) //simulate work ...
	w.Write([]byte("I am long process that finished... \n"))

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)

		go longProcess(&wg, w)

		wg.Wait()

		w.Write([]byte("All long processes are gone \n"))

	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
