package main

import (
	"fmt"
	"log"
	"os"
)

func openFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	read := make([]byte, 50)
	count, err := file.Read(read)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, read[:count])
	fmt.Println(read[0])
}

func changePermission(filename string, mode os.FileMode) {
	if err := os.Chmod(filename, mode); err != nil {
		log.Fatal(err)
	}
}
func main() {
	mode := os.FileMode(0777)
	changePermission("test.txt", mode)
}
