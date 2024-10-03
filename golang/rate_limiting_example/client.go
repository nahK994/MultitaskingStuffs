package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func processConn(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte(fmt.Sprintf("task %d", taskId)))
	if err != nil {
		log.Fatal(err)
	}

	bt := make([]byte, 2048)
	_, err = conn.Read(bt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bt))
}

func main() {
	t := time.Now()
	batchSize := 3
	totalRequest := 100
	var wg sync.WaitGroup
	for i := 1; i <= totalRequest; i += batchSize {
		for j := i; j < i+batchSize && j <= totalRequest; j++ {
			wg.Add(1)
			go processConn(j, &wg)
		}
		wg.Wait()
		fmt.Println()
	}
	fmt.Println("time taken", time.Since(t).Seconds())
}
