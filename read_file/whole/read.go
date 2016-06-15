package main

import (
	"io/ioutil"
)

// reads the whole file
func read() {
	content, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		panic("file can not open")
	}
	println(string(content))
}

func main() {
	read()
}
