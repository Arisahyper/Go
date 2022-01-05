package main

import (
	"log"
	"os"
)

func main() {
	log.Println("hello world")
	log.Printf("%T %v", "hello", "hello")

	_, err := os.Open("/tmp/test.txt")
	if err != nil {
		log.Fatalln("エラーです >> ",err)
	}
}
