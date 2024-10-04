package routine

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send the result back
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) // Start 3 workers
	}

	for j := 1; j <= 5; j++ {
		jobs <- j // Send 5 jobs
	}
	close(jobs) // Close the jobs channel

	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results) // Collect all results
	}
}
