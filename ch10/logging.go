package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important Information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical Problem
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 06666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}
	// defer file.Close()

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	Trace.Println("I have something to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know")
	Error.Println("Something has failed")
	Error.Println("Something has failed")
	Error.Println("Something has failed")
}
