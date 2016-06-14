package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

// read newline-delimited lines of text
func readFile01() {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic("file can not open")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		println(scanner.Text())
	}
	file.Close()
}

// reads the whole file
func readFile02() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		panic("file can not open")
	}
	println(string(content))
}

func main() {
	readFile01()
	readFile02()
}
