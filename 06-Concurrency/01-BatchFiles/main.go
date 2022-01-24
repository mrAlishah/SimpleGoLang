package main

import (
	"batchFils/myLib"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func main() {
	fld := os.Args[1]
	outfile := os.Args[2]
	defer myLib.TimeTrack(time.Now(), "main")

	//-----------------------------------------
	//fan in in order to avoid race condition
	librerian := make(chan []string)
	writeDone := make(chan struct{})

	go func() {
		cache := make([][]string, 0)
		for result := range librerian {
			cache = append(cache, result)

			if len(cache) >= 1000 {
				writeToFile(cache, outfile)
				cache = nil
			}
		}
		if len(cache) > 0 {
			writeToFile(cache, outfile)
			cache = nil
		}
		writeDone <- struct{}{}
	}()
	//-----------------------------------------

	wg := &sync.WaitGroup{}

	//fan out
	pipline := make(chan string)
	createWorkers(10, pipline, librerian, wg)

	go func() {
		wg.Wait() //wait wg emtpty and worker finished
		close(librerian)
	}()
	//-----------------------------------------

	//Read Folder files
	err := filepath.Walk(fld, func(file string, filInfo os.FileInfo, err error) error {
		if !filInfo.IsDir() {
			//ِdistributed task between workers
			fmt.Printf("- pipline <- file => %s \n", file)
			pipline <- file

		}
		return nil
	})

	//-----------------------------------------
	//1-finished all files and workers done => used close(pipeline)
	//2-close witer channel after complete worker read file finished => Used WaitGroup
	//3- wait till write channel will be free => used wg.Wait(); close(librarian)
	//4- keep alive writeDone

	close(pipline)

	if err != nil {
		panic(err)
	}

	<-writeDone
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
//func writeToFile(filename, outfilename, md5sum, workerId string) error {
func writeToFile(inputs [][]string, outfilename string) error {
	file, err := os.OpenFile(outfilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755) //ModePerm FileMode = 0777 // Unix permission bits
	if err != nil {
		return err
	}
	defer file.Close()

	//file.WriteString(fmt.Sprintf("%s ; %s \n", md5sum, filename))
	//fmt.Printf("-> Worker #%s Write=> %s \n", workerId, filename)
	for _,input := range inputs{
		file.WriteString(fmt.Sprintf("%s ; %s \n",input[0], input[1]))
	}


	//https://stackoverflow.com/questions/10862375/when-to-flush-a-file-in-go
	return file.Sync()
}

func createWorkers(count int, pipline <-chan string, fanIn chan<- []string, wg *sync.WaitGroup) {
	for i := 0; i <= count; i++ {
		wg.Add(1) //add to the waitgroup counter

		go func(workerId int) {
			fmt.Printf("Worker #%d is ready to ricieve job ... \n", workerId)
			for file := range pipline {
				md5, _ := md5Generate(file)
				fmt.Printf("<- Worker #%d getMD5=> %s \n", workerId, file)

				//worker deliver his work to librerian
				fanIn <- []string{md5, file, strconv.Itoa(workerId)}
			}

			wg.Done() //signal that the worker work is done

		}(i)

	}
}
