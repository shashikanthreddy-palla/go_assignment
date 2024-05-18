// Assignment 1: Goroutine with Channel Problem: Write a Go program that calculates the
//  sum of numbers from 1 to N concurrently using goroutines and channels.
//   The program should take the value of N as input from the user.

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("not provided args")
		os.Exit(0)
	}
	var wg sync.WaitGroup
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("provided invalid number as input", os.Args[1])
		os.Exit(0)
	}
	c := make(chan int64)
	i := int64(0)
	sum := int64(0)
	wg.Add(2)
	go func(num int64) {
		for i = 1; i <= num; i++ {
			c <- i
		}
		wg.Done()
		close(c)
	}(n)
	go func() {
	loop:
		for {
			select {
			case d, ok := <-c:
				if !ok {
					wg.Done()
					break loop
				}
				sum += d
			}
		}

	}()
	wg.Wait()
	fmt.Println("sum of numbers ", n, "is ", sum)
}
