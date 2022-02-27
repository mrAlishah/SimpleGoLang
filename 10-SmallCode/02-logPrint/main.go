package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	// Package log implements simple struct, methods & functions to log program runtime information

	// Why not use package fmt?
	// 1) Package log is safe from concurrent goroutines while plain fmt isn't safe
	// 2) Log can attach information automatically, such as time, date, file path, etc (Setting output flags)

	//---------- segment 1
	// log.Println("Hi there.")
	// log.SetFlags(log.Ldate | log.Lshortfile)
	// log.Println("log with setting flags")

	//---------- segment 2
	// log.Panic("Panicking...") //panic&log message
	// log.Fatal("finished, successfully")

	//---------- segment 3: Output a file
	// file, _ := os.Create("file.log")
	// log.SetOutput(file)
	// log.Println("Hello Log file")
	// file.Close()

	// log.SetOutput(os.Stdout)
	// log.Println("Printing into standard out again")

	//---------- segment 4:Common loggers
	//https://pkg.go.dev/log#pkg-constants
	// flags := log.LstdFlags | log.Lshortfile
	// infoLogger := log.New(os.Stdout, "INFO: ", flags)
	// warnLogger := log.New(os.Stdout, "WARN: ", flags)
	// errorLogger := log.New(os.Stdout, "ERROR: ", flags)
	// infoLogger.Println("This is an info log")
	// warnLogger.Println("This is a warning log")
	// errorLogger.Println("This is an error log")

	//---------- segment 5: aggregate all tree into one
	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, color.HiCyanString("WARN: "), flags)
	errorLogger := log.New(os.Stdout, color.RedString("ERROR: "), flags)
	agLog := aggregateLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}

	agLog.info("This is an info log")
	agLog.warn("This is a warning log")
	agLog.error("This is an error log")

	//---------- segment 6: use package github.com/fatih/color to colorize log
	// go get github.com/fatih/color
	// Create SprintXxx functions to mix strings with other non-colorized strings:
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
	log.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Printf("This %s rocks!\n", info("package"))
	log.Printf("This %s rocks!\n", info("package"))

	// Use helper functions
	fmt.Println("This", color.RedString("warning"), "should be not neglected.")
	fmt.Printf("%v %v\n", color.GreenString("Info:"), "an important message.")

}

type aggregateLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *aggregateLogger) info(v ...interface{}) {
	l.infoLogger.Println(v...)
}
func (l *aggregateLogger) warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}
func (l *aggregateLogger) error(v ...interface{}) {
	l.errorLogger.Println(v...)
}
