package main

import (
	"fmt"
	"time"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("woker %d received %d \n", id, n)
		done <- true
	}
}

func createWorker(id int) chan int {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return c
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
}

func main() {
	chanDemo()
}
