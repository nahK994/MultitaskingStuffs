package main

import (
	"fmt"
	"sync"
	"time"
)

func routine1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Routine 1 finished")
}

func routine2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Routine 2 finished")
}

func routine3(wg *sync.WaitGroup, dependencyWG *sync.WaitGroup) {
	defer wg.Done()
	dependencyWG.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("Routine 3 finished")
}

func routine4(wg *sync.WaitGroup, dependencyWG *sync.WaitGroup) {
	defer wg.Done()
	dependencyWG.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("Routine 4 finished")
}

func routine5(wg *sync.WaitGroup, dependencyWG *sync.WaitGroup) {
	defer wg.Done()
	dependencyWG.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("Routine 5 finished")
}

func routine6(wg *sync.WaitGroup, dependencyWG *sync.WaitGroup) {
	defer wg.Done()
	dependencyWG.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("Routine 6 finished")
}

func main() {
	var wg1, wg2, wg3 sync.WaitGroup

	wg1.Add(2)
	wg2.Add(3)
	wg3.Add(1)

	go routine1(&wg1)
	go routine2(&wg1)
	go routine3(&wg2, &wg1)
	go routine4(&wg2, &wg1)
	go routine5(&wg2, &wg1)
	go routine6(&wg3, &wg2)

	wg3.Wait()
	fmt.Println("All routines finished")
}
