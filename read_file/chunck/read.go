package main

import (
	"os"
)

// reads file content chunck by chunck
func read() {
	buffSize := 64
	file, err := os.Open("../test.txt")
	if err != nil {
		panic("file can not open")
	}
	for {
		buff := make([]byte, buffSize)
		n, err := file.Read(buff)
		if err != nil || n == 0 {
			break
		} else {
			print(string(buff))
		}
	}
}

func main() {
	read()
}
