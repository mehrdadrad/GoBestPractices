package main

import (
	"fmt"
	"sync"
)

type Contact struct {
	Name string
	Age  uint
}

func main() {
	var contactPool = sync.Pool{
		New: func() interface{} {
			return &Contact{}
		},
	}
	func() {
		c := contactPool.Get().(*Contact)
		defer contactPool.Put(c)

		c.Name = "Mehrdad"
		c.Age = 40

		fmt.Println(c.Name)
	}()
}
