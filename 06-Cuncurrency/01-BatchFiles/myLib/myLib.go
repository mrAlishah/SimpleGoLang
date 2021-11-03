package myLib

import (
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Println("\n================================================")
	log.Printf("%s took %s", name, elapsed)
}
