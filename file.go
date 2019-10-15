package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	read("./file.go", 100)
}

func read(filename string, bytenum int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
	data := make([]byte, bytenum)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(count, string(data))
}
