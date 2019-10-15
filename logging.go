package main

import (
	"io"
	"log"
	"os"
)

func main() {
	logging("test.log")
}

func logging(file string) {
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiLogFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("end log")
}
