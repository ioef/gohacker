package challenges

import "gohacker/models"

// GetAdvancedChallenges returns advanced-level challenges (concurrency)
func GetAdvancedChallenges() []models.Challenge {
	return []models.Challenge{
		{
			ID:            "016",
			Title:         "Goroutines: Concurrent Execution",
			Description:   "Start sayHello in a goroutine and keep main alive long enough for it to print \"Hello from goroutine\".",
			Story:         "🚀 Run sayHello in a goroutine and ensure the mission log appears before main exits.",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      100,
			RequiredLevel: 10,
			Hints: []string{
				"Start the goroutine with go sayHello()",
				"Use time.Sleep after launching the goroutine",
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
			Description:   "Create a channel, send a string from a goroutine, and intercept the message in main.",
			Story:         "📡 Channels enable goroutines to communicate safely!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      120,
			RequiredLevel: 10,
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
			Description:   "Use sync.WaitGroup to wait for multiple goroutines to finish.",
			Story:         "⏳ WaitGroups help coordinate multiple goroutines!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      130,
			RequiredLevel: 10,
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
			Description:   "Create a buffered channel, send two values, and receive them without blocking.",
			Story:         "📦 Buffered channels allow asynchronous communication!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      120,
			RequiredLevel: 10,
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
			Description:   "Use select to receive from whichever channel is ready first.",
			Story:         "🎛️ Select lets you handle multiple channel operations!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      130,
			RequiredLevel: 10,
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
			Description:   "Use a mutex to safely increment a shared counter from multiple goroutines.",
			Story:         "🔒 Mutexes prevent race conditions in concurrent code!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      140,
			RequiredLevel: 10,
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
			Description:   "Create a worker pool with 3 workers processing jobs from a channel.",
			Story:         "👷 Worker pools efficiently distribute work across goroutines!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      150,
			RequiredLevel: 10,
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
			Description:   "Use context cancellation to stop a long-running goroutine.",
			Story:         "⏱️ Context enables cancellation and timeouts in Go!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      160,
			RequiredLevel: 10,
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
		{
			ID:            "026",
			Title:         "Channel Directions",
			Description:   "Create functions with send-only and receive-only channel parameters.",
			Story:         "➡️ Channel directions enforce correct usage patterns!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      130,
			RequiredLevel: 10,
			Hints: []string{
				"Send-only: chan<- type",
				"Receive-only: <-chan type",
				"Helps prevent misuse",
			},
			StarterCode: `package main

import "fmt"

func send(ch chan<- string) {
	ch <- "message"
}

func receive(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go send(ch)
	receive(ch)
}`,
			Solution: `package main

import "fmt"

func send(ch chan<- string) {
	ch <- "message"
}

func receive(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go send(ch)
	receive(ch)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "message\n",
					Description:    "Should use directional channels",
					Hidden:         false,
				},
			},
			Order: 26,
		},
		{
			ID:            "027",
			Title:         "Closing Channels",
			Description:   "Close a channel and detect its closure in the receiver.",
			Story:         "🚪 Closing channels signals completion!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      120,
			RequiredLevel: 10,
			Hints: []string{
				"Use close(ch) to close",
				"val, ok := <-ch to check",
				"ok is false when closed",
			},
			StarterCode: `package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	
	// TODO: Receive all values and print "closed" when done
	
}`,
			Solution: `package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("closed")
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\nclosed\n",
					Description:    "Should detect channel closure",
					Hidden:         false,
				},
			},
			Order: 27,
		},
		{
			ID:            "028",
			Title:         "sync.Once",
			Description:   "Use sync.Once to run initialization exactly once.",
			Story:         "1️⃣ sync.Once guarantees single execution!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      140,
			RequiredLevel: 10,
			Hints: []string{
				"var once sync.Once",
				"once.Do(func() { ... })",
				"Safe for concurrent calls",
			},
			StarterCode: `package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once

func initialize() {
	count++
	fmt.Println("initialized")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}`,
			Solution: `package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once

func initialize() {
	count++
	fmt.Println("initialized")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "initialized\n1\n",
					Description:    "Should initialize only once",
					Hidden:         false,
				},
			},
			Order: 28,
		},
		{
			ID:            "029",
			Title:         "RWMutex",
			Description:   "Use sync.RWMutex for concurrent reads and exclusive writes.",
			Story:         "📖 RWMutex allows multiple readers or one writer!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      150,
			RequiredLevel: 10,
			Hints: []string{
				"RLock() for reading",
				"Lock() for writing",
				"Multiple readers can proceed",
			},
			StarterCode: `package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.RWMutex
	v  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.v
}

func main() {
	c := SafeCounter{}
	c.Inc()
	c.Inc()
	fmt.Println(c.Value())
}`,
			Solution: `package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.RWMutex
	v  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.v
}

func main() {
	c := SafeCounter{}
	c.Inc()
	c.Inc()
	fmt.Println(c.Value())
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "2\n",
					Description:    "Should safely increment",
					Hidden:         false,
				},
			},
			Order: 29,
		},
		{
			ID:            "030",
			Title:         "Atomic Operations",
			Description:   "Use sync/atomic to increment a shared counter without locks.",
			Story:         "⚛️ Atomic operations are faster than mutexes!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      160,
			RequiredLevel: 10,
			Hints: []string{
				"Use atomic.AddInt64",
				"Use atomic.LoadInt64",
				"No locks needed",
			},
			StarterCode: `package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&counter))
}`,
			Solution: `package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&counter))
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "10\n",
					Description:    "Should atomically increment",
					Hidden:         false,
				},
			},
			Order: 30,
		},
		{
			ID:            "031",
			Title:         "Fan-Out Pattern",
			Description:   "Distribute work across workers using fan-out goroutines.",
			Story:         "📤 Fan-out distributes work for parallel processing!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      170,
			RequiredLevel: 10,
			Hints: []string{
				"Send jobs to multiple workers",
				"Each worker processes independently",
				"Collect results from all",
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
	
	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	wg.Wait()
	close(results)
	
	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println(sum)
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
	
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	wg.Wait()
	close(results)
	
	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println(sum)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "30\n",
					Description:    "Sum of doubled values (2+4+6+8+10)",
					Hidden:         false,
				},
			},
			Order: 31,
		},
		{
			ID:            "032",
			Title:         "Pipeline Pattern",
			Description:   "Create a pipeline with generator, squarer, and printer stages.",
			Story:         "🔗 Pipelines chain operations for data processing!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      180,
			RequiredLevel: 10,
			Hints: []string{
				"Each stage is a function returning channel",
				"Connect stages: gen() -> square() -> print()",
				"Data flows through pipeline",
			},
			StarterCode: `package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range square(gen(1, 2, 3)) {
		fmt.Println(n)
	}
}`,
			Solution: `package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range square(gen(1, 2, 3)) {
		fmt.Println(n)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n4\n9\n",
					Description:    "Should square numbers",
					Hidden:         false,
				},
			},
			Order: 32,
		},
		{
			ID:            "033",
			Title:         "Rate Limiting",
			Description:   "Implement rate limiting with time.Ticker to control request flow.",
			Story:         "⏱️ Rate limiting controls request frequency!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      170,
			RequiredLevel: 10,
			Hints: []string{
				"Use time.NewTicker",
				"Receive from ticker.C",
				"Limits operations per second",
			},
			StarterCode: `package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()
	
	for req := range requests {
		<-limiter.C
		fmt.Println(req)
	}
}`,
			Solution: `package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()
	
	for req := range requests {
		<-limiter.C
		fmt.Println(req)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\n3\n4\n5\n",
					Description:    "Should rate limit requests",
					Hidden:         false,
				},
			},
			Order: 33,
		},
		{
			ID:            "034",
			Title:         "Timeout Pattern",
			Description:   "Implement timeout using select and time.After.",
			Story:         "⏰ Timeouts prevent indefinite waiting!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      160,
			RequiredLevel: 10,
			Hints: []string{
				"Use time.After in select",
				"Returns channel that fires after duration",
				"Print timeout if triggered",
			},
			StarterCode: `package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()
	
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}`,
			Solution: `package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()
	
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "timeout\n",
					Description:    "Should timeout after 1 second",
					Hidden:         false,
				},
			},
			Order: 34,
		},
		{
			ID:            "035",
			Title:         "HTTP Server Basics",
			Description:   "Stand up a basic HTTP server that responds with \"Hello World\".",
			Story:         "🌐 HTTP servers are the foundation of web services!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      190,
			RequiredLevel: 10,
			Hints: []string{
				"Use http.HandleFunc",
				"Use http.ListenAndServe",
				"Handler writes to ResponseWriter",
			},
			StarterCode: `package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server ready")
	// Server would run here in real app
}`,
			Solution: `package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server ready")
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Server ready\n",
					Description:    "Should setup server",
					Hidden:         false,
				},
			},
			Order: 35,
		},
		{
			ID:            "036",
			Title:         "HTTP Client",
			Description:   "Make an HTTP GET request and print the response body.",
			Story:         "🌐 HTTP clients consume APIs!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      200,
			RequiredLevel: 10,
			StarterCode: `package main
import ("fmt"; "net/http"; "io")
func main() {
	resp, _ := http.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(len(body) > 0)
}`,
			Solution: `package main
import ("fmt"; "net/http"; "io")
func main() {
	resp, _ := http.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(len(body) > 0)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "true\n", Description: "Should get response", Hidden: false}},
			Order:     36,
		},
		{
			ID:            "037",
			Title:         "JSON API Response",
			Description:   "Parse JSON from an HTTP response and print a field.",
			Story:         "📡 APIs speak JSON!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      210,
			RequiredLevel: 10,
			StarterCode: `package main
import "fmt"
type Response struct { Status string }
func main() {
	r := Response{Status: "ok"}
	fmt.Println(r.Status)
}`,
			Solution: `package main
import "fmt"
type Response struct { Status string }
func main() {
	r := Response{Status: "ok"}
	fmt.Println(r.Status)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "ok\n", Description: "Should parse response", Hidden: false}},
			Order:     37,
		},
		{
			ID:            "038",
			Title:         "Middleware Pattern",
			Description:   "Create HTTP middleware that logs each request.",
			Story:         "🔗 Middleware chains request processing!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      220,
			RequiredLevel: 10,
			StarterCode: `package main
import ("fmt"; "net/http")
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logged")
		next(w, r)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
func main() {
	http.HandleFunc("/", logger(handler))
	fmt.Println("middleware ready")
}`,
			Solution: `package main
import ("fmt"; "net/http")
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logged")
		next(w, r)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
func main() {
	http.HandleFunc("/", logger(handler))
	fmt.Println("middleware ready")
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "middleware ready\n", Description: "Should setup middleware", Hidden: false}},
			Order:     38,
		},
		{
			ID:            "039",
			Title:         "Table-Driven Tests",
			Description:   "Write a table-driven test and ensure it passes.",
			Story:         "✅ Table tests cover multiple cases efficiently!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      200,
			RequiredLevel: 10,
			StarterCode: `package main
import "fmt"
func add(a, b int) int { return a + b }
func main() {
	tests := []struct{ a, b, want int }{
		{1, 2, 3},
		{2, 3, 5},
	}
	for _, tt := range tests {
		if got := add(tt.a, tt.b); got == tt.want {
			fmt.Println("pass")
		}
	}
}`,
			Solution: `package main
import "fmt"
func add(a, b int) int { return a + b }
func main() {
	tests := []struct{ a, b, want int }{
		{1, 2, 3},
		{2, 3, 5},
	}
	for _, tt := range tests {
		if got := add(tt.a, tt.b); got == tt.want {
			fmt.Println("pass")
		}
	}
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "pass\npass\n", Description: "Should pass tests", Hidden: false}},
			Order:     39,
		},
		{
			ID:            "040",
			Title:         "Benchmarking",
			Description:   "Understand benchmark pattern by writing a simple benchmark.",
			Story:         "⚡ Benchmarks measure performance!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      210,
			RequiredLevel: 10,
			StarterCode: `package main
import "fmt"
func fibonacci(n int) int {
	if n < 2 { return n }
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	fmt.Println(fibonacci(10))
}`,
			Solution: `package main
import "fmt"
func fibonacci(n int) int {
	if n < 2 { return n }
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	fmt.Println(fibonacci(10))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "55\n", Description: "Should calculate fibonacci", Hidden: false}},
			Order:     40,
		},
		{
			ID:            "041",
			Title:         "Context with Timeout",
			Description:   "Use context.WithTimeout to limit an operation to a deadline.",
			Story:         "⏱️ Timeouts prevent hanging operations!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      220,
			RequiredLevel: 10,
			StarterCode: `package main
import ("context"; "fmt"; "time")
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}`,
			Solution: `package main
import ("context"; "fmt"; "time")
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "timeout\n", Description: "Should timeout", Hidden: false}},
			Order:     41,
		},
		{
			ID:            "042",
			Title:         "Graceful Shutdown",
			Description:   "Implement graceful server shutdown with context cancellation.",
			Story:         "🛑 Graceful shutdown ensures clean exit!",
			Difficulty:    models.DifficultyAdvanced,
			Category:      models.CategoryConcurrency,
			XPReward:      250,
			RequiredLevel: 10,
			StarterCode: `package main
import ("context"; "fmt"; "net/http"; "time")
func main() {
	srv := &http.Server{Addr: ":8080"}
	go func() {
		srv.ListenAndServe()
	}()
	time.Sleep(10 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("shutdown complete")
}`,
			Solution: `package main
import ("context"; "fmt"; "net/http"; "time")
func main() {
	srv := &http.Server{Addr: ":8080"}
	go func() {
		srv.ListenAndServe()
	}()
	time.Sleep(10 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("shutdown complete")
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "shutdown complete\n", Description: "Should shutdown gracefully", Hidden: false}},
			Order:     42,
		},
	}
}
