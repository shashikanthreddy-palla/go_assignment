// Assignment 2: Producer-Consumer with Channel Problem: Implement the producer-consumer problem
// using goroutines and channels. The producer should generate numbers from 1 to 100 and send them to a
// channel, and the consumer should print those numbers.
package main

import (
	"fmt"
	"sync"
)

var c = make(chan int)
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go producer()
	go consumer()
	wg.Wait()
}

func producer() {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	wg.Done()
	close(c)
}

func consumer() {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}
