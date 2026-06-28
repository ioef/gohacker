package challenges

import "gohacker/models"

// GetIntermediateChallenges returns intermediate-level challenges
func GetIntermediateChallenges() []models.Challenge {
	return []models.Challenge{
		{
			ID:            "011",
			Title:         "Structs: Define a User",
			Description:   "Forge a User struct with Name and Age, instantiate it, and print the struct.",
			Story:         "🧬 Build a custom data blueprint to represent a target profile in the hacker dossier.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      50,
			RequiredLevel: 5,
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
			Description:   "Define a Rectangle struct and add an Area() method to calculate its area.",
			Story:         "🔧 Enhance your struct with behavior — turn raw data into an active tool.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      60,
			RequiredLevel: 5,
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
			Description:   "Write a function that doubles a value through a pointer parameter.",
			Story:         "🧠 Reach into memory and modify the payload directly, like a true systems hacker.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 5,
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
			Description:   "Define a Shape interface with Area() and implement it for a Circle.",
			Story:         "🌀 Build a stealthy abstraction that can swap shapes without breaking the system.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      80,
			RequiredLevel: 5,
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
			Description:   "Create a divide function that returns an error on zero to protect the mission pipeline.",
			Story:         "🛡️ Handle bad input like a hacker avoids traps; return errors instead of crashing.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
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
			Description:   "Create a Person struct, encode it to JSON, and print the JSON payload.",
			Story:         "📦 Serialize your target profile into JSON and prepare it for API exfiltration.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 5,
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
			Description:   "Use defer so \"End\" prints after \"Start\" during the shutdown sequence.",
			Story:         "⏳ Defer schedules cleanup after the mission, ensuring the log closes properly.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      60,
			RequiredLevel: 5,
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
			Description:   "Create a closure that captures a counter, increments it, and returns the new value.",
			Story:         "🔐 Capture state inside a function and create a hidden counter that stays alive.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
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
			Description:   "Slice a list to extract elements from index 1 through 3.",
			Story:         "🗡️ Carve out the exact segment you need from the data stream.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      55,
			RequiredLevel: 5,
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
			Description:   "Iterate a map and print each key-value pair.",
			Story:         "🗺️ Traverse the map like a cyber scout reading every node and value.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      60,
			RequiredLevel: 5,
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
		{
			ID:            "021",
			Title:         "Empty Interface",
			Description:   "Store mixed typed values in a []interface{} payload and print it.",
			Story:         "🎭 Build a mixed payload container that can hold any data type.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 5,
			Hints: []string{
				"var items []interface{}",
				"Can store int, string, bool, anything",
				"Print the slice",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create slice of interface{}, add 42, "hello", true, print it
	
}`,
			Solution: `package main

import "fmt"

func main() {
	var items []interface{}
	items = append(items, 42)
	items = append(items, "hello")
	items = append(items, true)
	fmt.Println(items)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "[42 hello true]\n",
					Description:    "Should print mixed types",
					Hidden:         false,
				},
			},
			Order: 21,
		},
		{
			ID:            "022",
			Title:         "Type Assertion",
			Description:   "Use type assertion to extract a string from interface{} and print it.",
			Story:         "🔎 Inspect the unknown and pull out the exact string when you need it.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
			Hints: []string{
				"value, ok := i.(string)",
				"Check ok before using value",
				"Print the extracted string",
			},
			StarterCode: `package main

import "fmt"

func main() {
	var i interface{} = "hello"
	// TODO: Extract string using type assertion and print
	
}`,
			Solution: `package main

import "fmt"

func main() {
	var i interface{} = "hello"
	str, ok := i.(string)
	if ok {
		fmt.Println(str)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "hello\n",
					Description:    "Should extract and print string",
					Hidden:         false,
				},
			},
			Order: 22,
		},
		{
			ID:            "023",
			Title:         "Type Switch",
			Description:   "Use a type switch to handle int and string values and print the int case.",
			Story:         "🧩 Branch on runtime type like a savvy decoder handling multiple packet formats.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      80,
			RequiredLevel: 5,
			Hints: []string{
				"switch v := i.(type) { case int: ... case string: ... }",
				"Handle int and string cases",
				"Print type-specific message",
			},
			StarterCode: `package main

import "fmt"

func printType(i interface{}) {
	// TODO: Use type switch to print "int: X" or "string: X"
	
}

func main() {
	printType(42)
}`,
			Solution: `package main

import "fmt"

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	}
}

func main() {
	printType(42)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "int: 42\n",
					Description:    "Should identify int type",
					Hidden:         false,
				},
			},
			Order: 23,
		},
		{
			ID:            "024",
			Title:         "Custom Error Type",
			Description:   "Define a custom error type with extra context in its Error() method.",
			Story:         "🧨 Create an error that speaks with more detail when the breach fails.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      85,
			RequiredLevel: 5,
			Hints: []string{
				"Create struct with Error() string method",
				"Implement error interface",
				"Return your custom error",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create MyError struct with Code int and Message string
// TODO: Implement Error() string method

func main() {
	err := &MyError{Code: 404, Message: "Not Found"}
	fmt.Println(err.Error())
}`,
			Solution: `package main

import "fmt"

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {
	err := &MyError{Code: 404, Message: "Not Found"}
	fmt.Println(err.Error())
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Error 404: Not Found\n",
					Description:    "Should format custom error",
					Hidden:         false,
				},
			},
			Order: 24,
		},
		{
			ID:            "025",
			Title:         "Error Wrapping",
			Description:   "Wrap an error with fmt.Errorf and print the wrapped message.",
			Story:         "🧭 Preserve the error trail so every failure can be traced back.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      80,
			RequiredLevel: 5,
			Hints: []string{
				"Use fmt.Errorf with %w verb",
				"Wrap original error with context",
				"Print wrapped error",
			},
			StarterCode: `package main

import (
	"errors"
	"fmt"
)

func main() {
	original := errors.New("database error")
	// TODO: Wrap with "failed to connect: %w" and print
	
}`,
			Solution: `package main

import (
	"errors"
	"fmt"
)

func main() {
	original := errors.New("database error")
	wrapped := fmt.Errorf("failed to connect: %w", original)
	fmt.Println(wrapped)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "failed to connect: database error\n",
					Description:    "Should wrap error",
					Hidden:         false,
				},
			},
			Order: 25,
		},
		{
			ID:            "026",
			Title:         "Panic and Recover",
			Description:   "Catch a panic with recover and print a recovery message.",
			Story:         "🛑 Shield your program from crashes by recovering from a rogue panic.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      90,
			RequiredLevel: 5,
			Hints: []string{
				"Use defer with recover()",
				"Check if recover() != nil",
				"Print \"recovered\" if panic caught",
			},
			StarterCode: `package main

import "fmt"

func main() {
	defer func() {
		// TODO: Recover from panic and print "recovered"
		
	}()
	panic("oops")
}`,
			Solution: `package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}()
	panic("oops")
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "recovered\n",
					Description:    "Should recover from panic",
					Hidden:         false,
				},
			},
			Order: 26,
		},
		{
			ID:            "027",
			Title:         "Read File",
			Description:   "Read a file using os.ReadFile and print its contents.",
			Story:         "📂 Pull the file contents from the archive and reveal the stored secret.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
			Hints: []string{
				"Use os.WriteFile to create test file first",
				"Use os.ReadFile to read it",
				"Print the content as string",
			},
			StarterCode: `package main

import (
	"fmt"
	"os"
)

func main() {
	// Create test file
	os.WriteFile("/tmp/test.txt", []byte("Hello File"), 0644)
	// TODO: Read and print the file content
	
}`,
			Solution: `package main

import (
	"fmt"
	"os"
)

func main() {
	os.WriteFile("/tmp/test.txt", []byte("Hello File"), 0644)
	data, _ := os.ReadFile("/tmp/test.txt")
	fmt.Println(string(data))
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Hello File\n",
					Description:    "Should read file content",
					Hidden:         false,
				},
			},
			Order: 27,
		},
		{
			ID:            "028",
			Title:         "Time Duration",
			Description:   "Create a 2 second duration and print it.",
			Story:         "⏱️ Build the time delay before launching the next stage.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      65,
			RequiredLevel: 5,
			Hints: []string{
				"Use time.Second constant",
				"Multiply by 2",
				"Print the duration",
			},
			StarterCode: `package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO: Create 2 second duration and print
	
}`,
			Solution: `package main

import (
	"fmt"
	"time"
)

func main() {
	d := 2 * time.Second
	fmt.Println(d)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "2s\n",
					Description:    "Should print 2s",
					Hidden:         false,
				},
			},
			Order: 28,
		},
		{
			ID:            "029",
			Title:         "Regular Expression Match",
			Description:   "Use regexp to inspect a string and detect digit characters.",
			Story:         "🔍 Scan the input for numeric signatures using regex.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      85,
			RequiredLevel: 5,
			Hints: []string{
				"Use regexp.MatchString",
				"Pattern: \\d+ matches digits",
				"Print true or false",
			},
			StarterCode: `package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "hello123"
	// TODO: Check if text contains digits, print result
	
}`,
			Solution: `package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "hello123"
	matched, _ := regexp.MatchString("\\d+", text)
	fmt.Println(matched)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "true\n",
					Description:    "Should find digits",
					Hidden:         false,
				},
			},
			Order: 29,
		},
		{
			ID:            "030",
			Title:         "Sort Integers",
			Description:   "Sort a slice of integers using sort.Ints and print the sorted list.",
			Story:         "📊 Reorder the data payload into ascending sequence.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 5,
			Hints: []string{
				"Use sort.Ints()",
				"Sorts in place",
				"Print sorted slice",
			},
			StarterCode: `package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 8, 1, 9}
	// TODO: Sort and print
	
}`,
			Solution: `package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 8, 1, 9}
	sort.Ints(numbers)
	fmt.Println(numbers)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "[1 2 5 8 9]\n",
					Description:    "Should sort ascending",
					Hidden:         false,
				},
			},
			Order: 30,
		},
		{
			ID:            "031",
			Title:         "Interface Composition",
			Description:   "Compose smaller interfaces into a ReadWriter composite type.",
			Story:         "🧩 Combine simple contracts into a powerful composite interface.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      90,
			RequiredLevel: 5,
			StarterCode: `package main
import "fmt"
type Reader interface { Read() string }
type Writer interface { Write(string) }
type ReadWriter interface { Reader; Writer }
type File struct{}
func (f File) Read() string { return "data" }
func (f File) Write(s string) {}
func main() {
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
}`,
			Solution: `package main
import "fmt"
type Reader interface { Read() string }
type Writer interface { Write(string) }
type ReadWriter interface { Reader; Writer }
type File struct{}
func (f File) Read() string { return "data" }
func (f File) Write(s string) {}
func main() {
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "data\n", Description: "Should compose interfaces", Hidden: false}},
			Order:     31,
		},
		{
			ID:            "032",
			Title:         "Stringer Interface",
			Description:   "Implement fmt.Stringer so Point prints custom formatted output.",
			Story:         "📝 Define how your type prints itself for clearer hacker logs.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
			StarterCode: `package main
import "fmt"
type Point struct { X, Y int }
func (p Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }
func main() { fmt.Println(Point{3, 4}) }`,
			Solution: `package main
import "fmt"
type Point struct { X, Y int }
func (p Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }
func main() { fmt.Println(Point{3, 4}) }`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "(3,4)\n", Description: "Should format point", Hidden: false}},
			Order:     32,
		},
		{
			ID:            "033",
			Title:         "Error Interface",
			Description:   "Understand the error interface by printing an error message.",
			Story:         "⚠️ Treat errors like any other interface and implement them cleanly.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      80,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "errors")
func main() {
	err := errors.New("test error")
	fmt.Println(err.Error())
}`,
			Solution: `package main
import ("fmt"; "errors")
func main() {
	err := errors.New("test error")
	fmt.Println(err.Error())
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "test error\n", Description: "Should print error", Hidden: false}},
			Order:     33,
		},
		{
			ID:            "034",
			Title:         "Sentinel Errors",
			Description:   "Use sentinel errors for direct comparison.",
			Story:         "🚩 Flag common failures with sentinel error values.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      85,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "errors")
var ErrNotFound = errors.New("not found")
func find(id int) error {
	if id != 1 { return ErrNotFound }
	return nil
}
func main() {
	err := find(2)
	if err == ErrNotFound { fmt.Println("not found") }
}`,
			Solution: `package main
import ("fmt"; "errors")
var ErrNotFound = errors.New("not found")
func find(id int) error {
	if id != 1 { return ErrNotFound }
	return nil
}
func main() {
	err := find(2)
	if err == ErrNotFound { fmt.Println("not found") }
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "not found\n", Description: "Should detect sentinel", Hidden: false}},
			Order:     34,
		},
		{
			ID:            "035",
			Title:         "errors.Is",
			Description:   "Use errors.Is to check a wrapped error against its base error.",
			Story:         "🔍 Trace the root cause through wrapped error chains.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      90,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "errors")
var ErrBase = errors.New("base")
func main() {
	wrapped := fmt.Errorf("wrapped: %w", ErrBase)
	fmt.Println(errors.Is(wrapped, ErrBase))
}`,
			Solution: `package main
import ("fmt"; "errors")
var ErrBase = errors.New("base")
func main() {
	wrapped := fmt.Errorf("wrapped: %w", ErrBase)
	fmt.Println(errors.Is(wrapped, ErrBase))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "true\n", Description: "Should find base error", Hidden: false}},
			Order:     35,
		},
		{
			ID:            "036",
			Title:         "Write File",
			Description:   "Write data to a file and verify the save by reading it back.",
			Story:         "💾 Exfiltrate content to disk and confirm the write succeeded.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "os")
func main() {
	os.WriteFile("/tmp/out.txt", []byte("test"), 0644)
	data, _ := os.ReadFile("/tmp/out.txt")
	fmt.Println(string(data))
}`,
			Solution: `package main
import ("fmt"; "os")
func main() {
	os.WriteFile("/tmp/out.txt", []byte("test"), 0644)
	data, _ := os.ReadFile("/tmp/out.txt")
	fmt.Println(string(data))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "test\n", Description: "Should write and read", Hidden: false}},
			Order:     36,
		},
		{
			ID:            "037",
			Title:         "Buffered I/O",
			Description:   "Use bufio to scan and print a line from input efficiently.",
			Story:         "⚡ Stream data with buffering to keep the pipeline fast.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      85,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "strings"; "bufio")
func main() {
	r := strings.NewReader("line1\nline2")
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fmt.Println(scanner.Text())
}`,
			Solution: `package main
import ("fmt"; "strings"; "bufio")
func main() {
	r := strings.NewReader("line1\nline2")
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fmt.Println(scanner.Text())
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "line1\n", Description: "Should scan line", Hidden: false}},
			Order:     37,
		},
		{
			ID:            "038",
			Title:         "File Paths",
			Description:   "Use filepath to extract the filename from a path.",
			Story:         "📁 Normalize the path and pull out the target file name.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "path/filepath")
func main() {
	p := filepath.Join("dir", "file.txt")
	fmt.Println(filepath.Base(p))
}`,
			Solution: `package main
import ("fmt"; "path/filepath")
func main() {
	p := filepath.Join("dir", "file.txt")
	fmt.Println(filepath.Base(p))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "file.txt\n", Description: "Should get basename", Hidden: false}},
			Order:     38,
		},
		{
			ID:            "039",
			Title:         "Time Formatting",
			Description:   "Format a date string using Go time layout templates.",
			Story:         "🕐 Stamp the packet with the correct timestamp format.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "time")
func main() {
	t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t.Format("2006-01-02"))
}`,
			Solution: `package main
import ("fmt"; "time")
func main() {
	t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t.Format("2006-01-02"))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "2024-01-01\n", Description: "Should format date", Hidden: false}},
			Order:     39,
		},
		{
			ID:            "040",
			Title:         "Timers",
			Description:   "Use time.NewTimer for delayed execution.",
			Story:         "⏲️ Set a timed trigger for the next operation.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      80,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "time")
func main() {
	timer := time.NewTimer(10 * time.Millisecond)
	<-timer.C
	fmt.Println("done")
}`,
			Solution: `package main
import ("fmt"; "time")
func main() {
	timer := time.NewTimer(10 * time.Millisecond)
	<-timer.C
	fmt.Println("done")
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "done\n", Description: "Should wait for timer", Hidden: false}},
			Order:     40,
		},
		{
			ID:            "041",
			Title:         "Regex Find",
			Description:   "Find the first digit substring in text using regexp.FindString.",
			Story:         "🔎 Extract numeric markers from the payload.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      85,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "regexp")
func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindString("abc123def"))
}`,
			Solution: `package main
import ("fmt"; "regexp")
func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindString("abc123def"))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "123\n", Description: "Should find digits", Hidden: false}},
			Order:     41,
		},
		{
			ID:            "042",
			Title:         "Flag Package",
			Description:   "Parse command-line flags and print the default value.",
			Story:         "🚩 Read mission options from flags so the hacker tool is configurable.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      75,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "flag")
func main() {
	name := flag.String("name", "default", "a name")
	flag.Parse()
	fmt.Println(*name)
}`,
			Solution: `package main
import ("fmt"; "flag")
func main() {
	name := flag.String("name", "default", "a name")
	flag.Parse()
	fmt.Println(*name)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "default\n", Description: "Should use default", Hidden: false}},
			Order:     42,
		},
		{
			ID:            "043",
			Title:         "fmt Verbs",
			Description:   "Master fmt verbs by printing both type and value.",
			Story:         "📝 Print both the type and value with precise formatting.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      70,
			RequiredLevel: 5,
			StarterCode: `package main
import "fmt"
func main() {
	fmt.Printf("%T %v\n", 42, 42)
}`,
			Solution: `package main
import "fmt"
func main() {
	fmt.Printf("%T %v\n", 42, 42)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "int 42\n", Description: "Should print type and value", Hidden: false}},
			Order:     43,
		},
		{
			ID:            "044",
			Title:         "Sort Custom Types",
			Description:   "Implement sort.Interface so custom data can be sorted by length.",
			Story:         "🗃️ Teach the sorter how to rank your custom data structures.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      95,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "sort")
type ByLen []string
func (s ByLen) Len() int { return len(s) }
func (s ByLen) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByLen) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func main() {
	words := []string{"aaa", "b", "cc"}
	sort.Sort(ByLen(words))
	fmt.Println(words)
}`,
			Solution: `package main
import ("fmt"; "sort")
type ByLen []string
func (s ByLen) Len() int { return len(s) }
func (s ByLen) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByLen) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func main() {
	words := []string{"aaa", "b", "cc"}
	sort.Sort(ByLen(words))
	fmt.Println(words)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "[b cc aaa]\n", Description: "Should sort by length", Hidden: false}},
			Order:     44,
		},
		{
			ID:            "045",
			Title:         "JSON Decoding",
			Description:   "Decode JSON into a struct and print the decoded fields.",
			Story:         "📥 Parse the incoming JSON packet and extract the payload.",
			Difficulty:    models.DifficultyIntermediate,
			Category:      models.CategoryStructs,
			XPReward:      85,
			RequiredLevel: 5,
			StarterCode: `package main
import ("fmt"; "encoding/json")
type User struct { Name string; Age int }
func main() {
	data := []byte("{\"Name\":\"Bob\",\"Age\":25}")
	var u User
	json.Unmarshal(data, &u)
	fmt.Println(u.Name, u.Age)
}`,
			Solution: `package main
import ("fmt"; "encoding/json")
type User struct { Name string; Age int }
func main() {
	data := []byte("{\"Name\":\"Bob\",\"Age\":25}")
	var u User
	json.Unmarshal(data, &u)
	fmt.Println(u.Name, u.Age)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "Bob 25\n", Description: "Should decode JSON", Hidden: false}},
			Order:     45,
		},
	}
}
