package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day4/errorHandle.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data) // errは上書きされてるよ
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	fmt.Println(string(data))

	if err = os.Chdir("/tmp"); err != nil {
		log.Fatal(err)
	}
}
