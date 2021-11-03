package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fld := os.Args[1]

	//Read Folder files
	err := filepath.Walk(fld, func(file string, filInfo os.FileInfo, err error) error {
		if !filInfo.IsDir() {
			md5, err := md5Generate(file)
			fmt.Printf("MD5= %s , File= %s , Error=%v \n", md5, file, err)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

//Create MD5
func md5Generate(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hash.Sum(nil)
	md5 := string(hex.EncodeToString(checksum))
	return md5, nil

}
