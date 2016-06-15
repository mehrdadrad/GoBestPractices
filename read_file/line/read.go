package main

import (
	"bufio"
	"os"
)

// read newline-delimited lines of text
func read() {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic("file can not open")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		println(scanner.Text())
	}
	file.Close()
}

func main() {
	read()
}
