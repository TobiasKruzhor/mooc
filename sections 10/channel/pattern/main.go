package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("serverce %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func fanIn(chs ...chan string) chan string { // chs: len:2,cap:2
	c := make(chan string)   // c: chan string
	for _, ch := range chs { // ch: chan string
		// chCopy := ch //chCopy: chan string
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func main() {
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m3 := msgGen("service3")
	m := fanIn(m1, m2, m3)
	for {
		fmt.Println(<-m)
	}

}
