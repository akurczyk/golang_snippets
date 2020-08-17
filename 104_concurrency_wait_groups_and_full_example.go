package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	a, b int
	operation string
	result int
}

func main() {
	var workersWg sync.WaitGroup

	tasks := make(chan task)
	results := make(chan task)
	done := make(chan bool)

	for i := 0; i < 5; i++ { go worker(tasks, results, &workersWg) }
	go scheduler(tasks, results, &workersWg)
	go receiver(results, done)

	<-done
}

func worker(tasks <-chan task, results chan<- task, workersWg *sync.WaitGroup) {
	for task := range tasks {
		time.Sleep(1 * time.Second)

		switch task.operation {
		case "+":
			task.result = task.a + task.b
		case "-":
			task.result = task.a - task.b
		case "*":
			task.result = task.a * task.b
		case "/":
			task.result = task.a / task.b
		}

		results <- task
	}

	workersWg.Done()
}

func scheduler(tasks chan<- task, results chan<- task, workersWg *sync.WaitGroup) {
	for i := 1; i < 101; i++ {
		tasks <- task{i * 1, i * 3, "+", 0}
		tasks <- task{i * 2, i * 4, "-", 0}
		tasks <- task{i * 2, i * 4, "*", 0}
		tasks <- task{i * 2, i * 4, "/", 0}
	}

	close(tasks)
	workersWg.Wait()
	close(results)
}

func receiver(results <-chan task, done chan<- bool) {
	for result := range results {
		fmt.Printf("%d %s %d = %d\r\n", result.a, result.operation, result.b, result.result)
	}

	done <- true
}
