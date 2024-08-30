package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		// Simulate some work with the value
	}
}

func main() {
	var wg sync.WaitGroup

	// Unbuffered Channel
	unbufferedCh := make(chan int)
	wg.Add(2)
	start := time.Now()
	go producer(unbufferedCh, &wg)
	go consumer(unbufferedCh, &wg)
	wg.Wait()
	fmt.Println("Unbuffered channel time:", time.Since(start))

	// Buffered Channel
	bufferedCh := make(chan int, 1000)
	wg.Add(2)
	start = time.Now()
	go producer(bufferedCh, &wg)
	go consumer(bufferedCh, &wg)
	wg.Wait()
	fmt.Println("Buffered channel time:", time.Since(start))
}
