package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	rw      sync.RWMutex
)

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	rw.RLock()
	fmt.Println("Counter:", counter)
	rw.RUnlock()
}

func write(wg *sync.WaitGroup) {
	defer wg.Done()
	rw.Lock()
	counter++
	rw.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go write(&wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go read(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
