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
		{
			ID:            "021",
			Title:         "Buffered Channels",
			Description:   "Create a buffered channel and send/receive values without blocking",
			Story:         "📦 Buffered channels allow asynchronous communication!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      120,
			RequiredLevel: 16,
			Hints: []string{
				"Create buffered channel: ch := make(chan int, 2)",
				"Send values without blocking: ch <- 1",
				"Receive with: val := <-ch",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create buffered channel with capacity 2
	// Send 1 and 2 to channel
	// Receive and print both values
}`,
			Solution: `package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\n",
					Description:    "Should print 1 and 2",
					Hidden:         false,
				},
			},
			Order: 21,
		},
		{
			ID:            "022",
			Title:         "Select Statement",
			Description:   "Use select to receive from multiple channels",
			Story:         "🎛️ Select lets you handle multiple channel operations!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      130,
			RequiredLevel: 17,
			Hints: []string{
				"Use select { case val := <-ch1: ... case val := <-ch2: ... }",
				"Select picks the first ready channel",
				"Can include default case for non-blocking",
			},
			StarterCode: `package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "from ch2"
	}()
	
	// TODO: Use select to receive from either channel and print
	time.Sleep(50 * time.Millisecond)
}`,
			Solution: `package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "from ch2"
	}()
	
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
	time.Sleep(50 * time.Millisecond)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "from ch1\n",
					Description:    "Should print from ch1 (arrives first)",
					Hidden:         false,
				},
			},
			Order: 22,
		},
		{
			ID:            "023",
			Title:         "Mutex: Safe Concurrent Access",
			Description:   "Use a mutex to safely increment a counter from multiple goroutines",
			Story:         "🔒 Mutexes prevent race conditions in concurrent code!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      140,
			RequiredLevel: 18,
			Hints: []string{
				"Import sync package",
				"Use sync.Mutex with Lock() and Unlock()",
				"Protect shared variable access",
			},
			StarterCode: `package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	// TODO: Add mutex here
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: Lock, increment counter, unlock
			counter++
		}()
	}
	
	wg.Wait()
	fmt.Println(counter)
}`,
			Solution: `package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	
	wg.Wait()
	fmt.Println(counter)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "10\n",
					Description:    "Should safely increment to 10",
					Hidden:         false,
				},
			},
			Order: 23,
		},
		{
			ID:            "024",
			Title:         "Worker Pool Pattern",
			Description:   "Create a worker pool with 3 workers processing jobs from a channel",
			Story:         "👷 Worker pools efficiently distribute work across goroutines!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      150,
			RequiredLevel: 19,
			Hints: []string{
				"Create jobs channel and results channel",
				"Start multiple worker goroutines",
				"Workers receive from jobs channel",
			},
			StarterCode: `package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup
	
	// TODO: Start 3 workers
	
	// Send 5 jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	
	wg.Wait()
	close(results)
	
	// Print results
	for result := range results {
		fmt.Println(result)
	}
}`,
			Solution: `package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup
	
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	
	wg.Wait()
	close(results)
	
	for result := range results {
		fmt.Println(result)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "2\n4\n6\n8\n10\n",
					Description:    "Should process all jobs (order may vary)",
					Hidden:         false,
				},
			},
			Order: 24,
		},
		{
			ID:            "025",
			Title:         "Context: Cancellation",
			Description:   "Use context to cancel a long-running goroutine",
			Story:         "⏱️ Context enables cancellation and timeouts in Go!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      160,
			RequiredLevel: 20,
			Hints: []string{
				"Import context package",
				"Use context.WithCancel()",
				"Check ctx.Done() in goroutine",
			},
			StarterCode: `package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancelled")
			return
		default:
			// Simulate work
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	// TODO: Create context with cancel
	// Start worker goroutine
	// Cancel after 50ms
	// Wait a bit for cleanup
}`,
			Solution: `package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancelled")
			return
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(20 * time.Millisecond)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "cancelled\n",
					Description:    "Should cancel worker",
					Hidden:         false,
				},
			},
			Order: 25,
		},
	}
}
