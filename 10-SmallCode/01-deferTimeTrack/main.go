package main

import (
	"fmt"
	"log"
	"time"
)

func TimeTrack(start time.Time) string {
	elapsed := time.Since(start)
	log.Printf("%s", elapsed)
	return fmt.Sprint(elapsed)
}

func print(str string) {
	log.Println(str)
}
func TimeTrackPrint(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Println("\n================================================")
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	//defer log.Printf("%s took %s", "main", TimeTrack(time.Now()))
	//defer TimeTrackPrint(time.Now(), "main")

	defer log.Println("bye")           //2
	defer print(TimeTrack(time.Now())) //1

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
	}
	//fmt.Scanln()
}
