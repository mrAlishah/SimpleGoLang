package main

import (
	"batchFils/myLib"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fld := os.Args[1]
	outfile := os.Args[2]
	defer myLib.TimeTrack(time.Now(), "main")

	//Read Folder files
	err := filepath.Walk(fld, func(file string, filInfo os.FileInfo, err error) error {
		if !filInfo.IsDir() {

			md5, err := md5Generate(file)
			fmt.Printf("MD5= %s , File= %s , Error=%v \n", md5, file, err)

			if err := writeToFile(file, outfile, md5); err != nil {
				panic(err)
			}
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

//Write info to args[2] filename
func writeToFile(filename, outfilename, md5sum string) error {
	file, err := os.OpenFile(outfilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755) //ModePerm FileMode = 0777 // Unix permission bits
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%s ; %s \n", md5sum, filename))

	//https://stackoverflow.com/questions/10862375/when-to-flush-a-file-in-go
	return file.Sync()
}
