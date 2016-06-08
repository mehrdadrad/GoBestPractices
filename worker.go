package main

import (
	"fmt"
	"time"
)

type Work struct {
	request   int
	err       error
	timestamp time.Duration
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.request++
		out <- w
	}
}

func publisher(work chan *Work) {
	for i := 0; i < 1000; i++ {
		work <- &Work{request: 1}
	}
}

func consumer(result chan *Work) {
	var (
		ok   int  = 0
		loop bool = true
	)
	for loop {
		select {
		case r := <-result:
			if r.request == 2 {
				ok++
			}
		case <-time.After(1 * time.Second):
			loop = false
		}
	}
	fmt.Printf("result %d done successfully\n", ok)
}

func main() {
	input, output := make(chan *Work), make(chan *Work)
	for i := 0; i < 10; i++ {
		go worker(input, output)
	}

	go publisher(input)
	consumer(output)
}
