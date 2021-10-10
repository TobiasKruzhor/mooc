package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello form goroutine %d\n", i)

				/*
					a[i]++
					runtime.Gosched()
				*/
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
