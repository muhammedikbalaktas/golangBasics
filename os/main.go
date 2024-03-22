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
func createFile() {
	file, err := os.OpenFile("muhammed.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	fmt.Println(os.O_WRONLY | os.O_APPEND | os.O_CREATE)

	if err != nil {
		panic(err)
	}
	file.WriteString("Muhammed ikbal\n")
	defer file.Close()
}
func copyImage() {
	fileStat, err := os.Stat("golang.png")
	if err != nil {
		panic(err)
	}
	bufferSize := fileStat.Size()
	fileToCopy, err := os.Open("golang.png")
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, bufferSize)

	n, err := fileToCopy.Read(buffer)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("copy.png", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	s, err := file.Write(buffer[:n])
	if err != nil {
		panic(err)
	}
	fmt.Printf("buffer size is %d\n", s)

}

func main() {
	copyImage()
}
