package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fld := os.Args[1]

	//Read Folder files
	err := filepath.Walk(fld, func(filename string, filInfo os.FileInfo, err error) error {
		if !filInfo.IsDir() {
			fmt.Println(filename)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
