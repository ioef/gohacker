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
	}
}
