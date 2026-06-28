package challenges

import "gohacker/models"

// GetIntermediateChallenges returns intermediate-level challenges
func GetIntermediateChallenges() []models.Challenge {
	return []models.Challenge{
		{
			ID:            "011",
			Title:         "Structs: Define a User",
			Description:   "Create a User struct with Name and Age fields, create an instance, and print it",
			Story:         "🏗️ Structs let you create custom data types. Build your first one!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      50,
			RequiredLevel: 6,
			Hints: []string{
				"Define struct: type User struct { Name string; Age int }",
				"Create instance: user := User{Name: \"Alice\", Age: 25}",
				"Print with fmt.Println(user)",
			},
			StarterCode: `package main

import "fmt"

// TODO: Define User struct with Name (string) and Age (int)

func main() {
	// TODO: Create a User with Name="Alice" and Age=25
	// Print the user
}`,
			Solution: `package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Alice", Age: 25}
	fmt.Println(user)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "{Alice 25}\n",
					Description:    "Should print the user struct",
					Hidden:         false,
				},
			},
			Order: 11,
		},
		{
			ID:            "012",
			Title:         "Methods: Add Behavior to Structs",
			Description:   "Create a Rectangle struct and add an Area() method",
			Story:         "⚙️ Methods add behavior to your structs. Make them come alive!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      60,
			RequiredLevel: 7,
			Hints: []string{
				"Define struct: type Rectangle struct { Width, Height int }",
				"Add method: func (r Rectangle) Area() int { return r.Width * r.Height }",
				"Call method: rect.Area()",
			},
			StarterCode: `package main

import "fmt"

// TODO: Define Rectangle struct with Width and Height
// TODO: Add Area() method

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println(rect.Area())
}`,
			Solution: `package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println(rect.Area())
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "50\n",
					Description:    "Should calculate area",
					Hidden:         false,
				},
			},
			Order: 12,
		},
		{
			ID:            "013",
			Title:         "Pointers: Reference Types",
			Description:   "Create a function that modifies a value using a pointer",
			Story:         "👉 Pointers let you modify values directly. Master memory manipulation!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 8,
			Hints: []string{
				"Function signature: func double(n *int)",
				"Dereference with *: *n = *n * 2",
				"Pass address with &: double(&x)",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create double function that takes *int and doubles the value

func main() {
	x := 5
	double(&x)
	fmt.Println(x)
}`,
			Solution: `package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func main() {
	x := 5
	double(&x)
	fmt.Println(x)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "10\n",
					Description:    "Should double the value",
					Hidden:         false,
				},
			},
			Order: 13,
		},
		{
			ID:            "014",
			Title:         "Interfaces: Polymorphism",
			Description:   "Create a Shape interface with Area() method and implement it for Circle",
			Story:         "🎭 Interfaces enable polymorphism. Write flexible code!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      80,
			RequiredLevel: 9,
			Hints: []string{
				"Define interface: type Shape interface { Area() float64 }",
				"Implement for Circle with radius",
				"Use math.Pi for circle area calculation",
			},
			StarterCode: `package main

import (
	"fmt"
	"math"
)

// TODO: Define Shape interface with Area() float64
// TODO: Define Circle struct with Radius float64
// TODO: Implement Area() for Circle

func main() {
	var s Shape = Circle{Radius: 5}
	fmt.Printf("%.2f\n", s.Area())
}`,
			Solution: `package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	var s Shape = Circle{Radius: 5}
	fmt.Printf("%.2f\n", s.Area())
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "78.54\n",
					Description:    "Should calculate circle area",
					Hidden:         false,
				},
			},
			Order: 14,
		},
		{
			ID:            "015",
			Title:         "Error Handling: Return Errors",
			Description:   "Create a divide function that returns an error for division by zero",
			Story:         "⚠️ Proper error handling prevents crashes. Be defensive!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 10,
			Hints: []string{
				"Return (int, error) from function",
				"Use errors.New() or fmt.Errorf()",
				"Check if b == 0 and return error",
			},
			StarterCode: `package main

import (
	"errors"
	"fmt"
)

// TODO: Create divide function that returns (int, error)
// Return error if b is 0

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}`,
			Solution: `package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Error: division by zero\n",
					Description:    "Should handle division by zero",
					Hidden:         false,
				},
			},
			Order: 15,
		},
		{
			ID:            "016",
			Title:         "JSON Encoding",
			Description:   "Create a Person struct and encode it to JSON",
			Story:         "📡 JSON is the language of APIs. Learn to encode data!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 8,
			Hints: []string{
				"Import encoding/json package",
				"Use json.Marshal() to encode",
				"Print the JSON string",
			},
			StarterCode: `package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	// TODO: Encode to JSON and print
	
}`,
			Solution: `package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "{\"Name\":\"Alice\",\"Age\":30}\n",
					Description:    "Should print JSON",
					Hidden:         false,
				},
			},
			Order: 16,
		},
		{
			ID:            "017",
			Title:         "Defer Statement",
			Description:   "Use defer to print \"End\" after \"Start\"",
			Story:         "⏰ Defer delays execution until function returns. Very useful!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      60,
			RequiredLevel: 8,
			Hints: []string{
				"defer executes after function completes",
				"Use defer fmt.Println(\"End\")",
				"Print \"Start\" normally",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Use defer to print "End" after "Start"
	fmt.Println("Start")
}`,
			Solution: `package main

import "fmt"

func main() {
	defer fmt.Println("End")
	fmt.Println("Start")
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Start\nEnd\n",
					Description:    "Should print Start then End",
					Hidden:         false,
				},
			},
			Order: 17,
		},
		{
			ID:            "018",
			Title:         "Closures",
			Description:   "Create a closure that increments a counter",
			Story:         "🔒 Closures capture variables from outer scope. Powerful pattern!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 9,
			Hints: []string{
				"Return a function from a function",
				"The inner function can access outer variables",
				"Each call increments the counter",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create counter function that returns a closure

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}`,
			Solution: `package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\n3\n",
					Description:    "Should increment counter",
					Hidden:         false,
				},
			},
			Order: 18,
		},
		{
			ID:            "019",
			Title:         "Slice Operations",
			Description:   "Create a slice and use slicing to get elements from index 1 to 3",
			Story:         "✂️ Slice operations let you extract portions of data!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      55,
			RequiredLevel: 9,
			Hints: []string{
				"Use slice[start:end] syntax",
				"End index is exclusive",
				"slice[1:4] gets elements at index 1, 2, 3",
			},
			StarterCode: `package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	// TODO: Get elements from index 1 to 3 (inclusive) and print
	
}`,
			Solution: `package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	subset := numbers[1:4]
	fmt.Println(subset)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "[20 30 40]\n",
					Description:    "Should print subset",
					Hidden:         false,
				},
			},
			Order: 19,
		},
		{
			ID:            "020",
			Title:         "Map Iteration",
			Description:   "Iterate over a map and print all key-value pairs",
			Story:         "🗺️ Learn to traverse maps efficiently!",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      60,
			RequiredLevel: 10,
			Hints: []string{
				"Use range to iterate over maps",
				"for key, value := range myMap",
				"Print each key and value",
			},
			StarterCode: `package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	// TODO: Iterate and print each key-value pair as "key: value"
	
}`,
			Solution: `package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	for key, value := range scores {
		fmt.Printf("%s: %d\n", key, value)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Alice: 90\nBob: 85\n",
					Description:    "Should print all pairs",
					Hidden:         false,
				},
			},
			Order: 20,
		},
	}
}
