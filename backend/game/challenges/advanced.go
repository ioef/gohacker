package challenges

import "gohacker/models"

// GetAdvancedChallenges returns advanced-level challenges (concurrency)
func GetAdvancedChallenges() []models.Challenge {
	return []models.Challenge{
		{
			ID:            "016",
			Title:         "Goroutines: Concurrent Execution",
			Description:   "Create a goroutine that prints \"Hello from goroutine\" and wait for it",
			Story:         "🚀 Goroutines enable concurrency. Run code in parallel!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      100,
			RequiredLevel: 15,
			Hints: []string{
				"Use 'go' keyword to start goroutine",
				"Use time.Sleep to wait for goroutine",
				"Import time package",
			},
			StarterCode: `package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	// TODO: Start sayHello as a goroutine
	// Wait for it to complete
}`,
			Solution: `package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Hello from goroutine\n",
					Description:    "Should print from goroutine",
					Hidden:         false,
				},
			},
			Order: 16,
		},
		{
			ID:            "017",
			Title:         "Channels: Communication Between Goroutines",
			Description:   "Create a channel, send a value in a goroutine, and receive it in main",
			Story:         "📡 Channels enable goroutines to communicate safely!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      120,
			RequiredLevel: 16,
			Hints: []string{
				"Create channel: ch := make(chan string)",
				"Send: ch <- \"message\"",
				"Receive: msg := <-ch",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create a string channel
	// Start goroutine that sends "Hello" to channel
	// Receive and print the message
}`,
			Solution: `package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello"
	}()
	msg := <-ch
	fmt.Println(msg)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Hello\n",
					Description:    "Should receive message from channel",
					Hidden:         false,
				},
			},
			Order: 17,
		},
		{
			ID:            "018",
			Title:         "WaitGroups: Synchronization",
			Description:   "Use sync.WaitGroup to wait for multiple goroutines",
			Story:         "⏳ WaitGroups help coordinate multiple goroutines!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      130,
			RequiredLevel: 18,
			Hints: []string{
				"Create: var wg sync.WaitGroup",
				"Add: wg.Add(1) before goroutine",
				"Done: defer wg.Done() in goroutine",
				"Wait: wg.Wait() in main",
			},
			StarterCode: `package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	// TODO: Add defer wg.Done()
	fmt.Printf("Worker %d\n", id)
}

func main() {
	// TODO: Create WaitGroup
	// Start 3 workers
	// Wait for all to complete
}`,
			Solution: `package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Worker 1\nWorker 2\nWorker 3\n",
					Description:    "Should print all workers (order may vary)",
					Hidden:         false,
				},
			},
			Order: 18,
		},
	}
}
