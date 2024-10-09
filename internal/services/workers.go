package services

import (
	"sync"
)

func ConcurrentTask() {
	var wg sync.WaitGroup
	tasks := []string{"Hello", "Go Microservices", "Gadag"}
	results := make(chan string)

	for _, task := range tasks {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			results <- "Processed " + t
		}(task)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		println(result)
	}
}
