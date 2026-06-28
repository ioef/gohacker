package challenges

import "gohacker/models"

// GetBeginnerChallenges returns all beginner-level challenges
func GetBeginnerChallenges() []models.Challenge {
	return []models.Challenge{
		{
			ID:            "001",
			Title:         "First Breach: Variable Declaration",
			Description:   "Declare a variable called 'password' with the value \"admin123\"",
			Story:         "🎯 Welcome, newbie hacker! Your first mission is simple: inject a password variable into the system. This is your entry point!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      10,
			RequiredLevel: 1,
			Hints: []string{
				"Use the 'var' keyword to declare a variable",
				"Syntax: var name type = value",
				"var password string = \"admin123\"",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Declare a variable called 'password' with value "admin123"
	
	fmt.Println(password)
}`,
			Solution: `package main

import "fmt"

func main() {
	var password string = "admin123"
	fmt.Println(password)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "admin123\n",
					Description:    "Should print the password",
					Hidden:         false,
				},
			},
			Order: 1,
		},
		{
			ID:            "002",
			Title:         "Data Types: Integer Operations",
			Description:   "Create two variables: 'x' with value 10 and 'y' with value 20, then print their sum",
			Story:         "💻 The system requires numerical calculations. Show your understanding of integers!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      15,
			RequiredLevel: 1,
			Hints: []string{
				"Declare two int variables",
				"Use the + operator to add them",
				"Print the result with fmt.Println",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create x = 10 and y = 20, then print their sum
	
}`,
			Solution: `package main

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println(x + y)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "30\n",
					Description:    "Should print 30",
					Hidden:         false,
				},
			},
			Order: 2,
		},
		{
			ID:            "003",
			Title:         "Control Flow: If Statement",
			Description:   "Write a program that checks if a number is greater than 10 and prints \"Big\" or \"Small\"",
			Story:         "🔍 Security systems use conditional logic. Master the if statement!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
			Hints: []string{
				"Use an if-else statement",
				"Compare the number with 10 using >",
				"Print different messages based on the condition",
			},
			StarterCode: `package main

import "fmt"

func main() {
	number := 15
	// TODO: Check if number > 10, print "Big" or "Small"
	
}`,
			Solution: `package main

import "fmt"

func main() {
	number := 15
	if number > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Big\n",
					Description:    "Should print Big for 15",
					Hidden:         false,
				},
			},
			Order: 3,
		},
		{
			ID:            "004",
			Title:         "Loops: For Loop Basics",
			Description:   "Use a for loop to print numbers from 1 to 5",
			Story:         "🔄 Automation is key! Learn to repeat actions with loops.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 2,
			Hints: []string{
				"Use a for loop with initialization, condition, and increment",
				"for i := 1; i <= 5; i++ { ... }",
				"Print i in each iteration",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Print numbers 1 to 5 using a for loop
	
}`,
			Solution: `package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\n3\n4\n5\n",
					Description:    "Should print 1 through 5",
					Hidden:         false,
				},
			},
			Order: 4,
		},
		{
			ID:            "005",
			Title:         "Functions: Create Your First Function",
			Description:   "Create a function called 'greet' that takes a name and returns \"Hello, [name]!\"",
			Story:         "🎭 Functions are reusable code blocks. Master them to become efficient!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 2,
			Hints: []string{
				"Define a function with: func greet(name string) string { ... }",
				"Use fmt.Sprintf to format the string",
				"Return the formatted string",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create a greet function here

func main() {
	result := greet("Hacker")
	fmt.Println(result)
}`,
			Solution: `package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	result := greet("Hacker")
	fmt.Println(result)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Hello, Hacker!\n",
					Description:    "Should greet the hacker",
					Hidden:         false,
				},
			},
			Order: 5,
		},
		{
			ID:            "006",
			Title:         "Arrays: Data Storage",
			Description:   "Create an array of 3 strings and print the second element",
			Story:         "📦 Arrays store multiple values. Learn to access them!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 3,
			Hints: []string{
				"Declare an array: var arr [3]string",
				"Assign values: arr[0] = \"first\"",
				"Access with index: arr[1]",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create array with "alpha", "beta", "gamma" and print second element
	
}`,
			Solution: `package main

import "fmt"

func main() {
	arr := [3]string{"alpha", "beta", "gamma"}
	fmt.Println(arr[1])
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "beta\n",
					Description:    "Should print beta",
					Hidden:         false,
				},
			},
			Order: 6,
		},
		{
			ID:            "007",
			Title:         "Slices: Dynamic Arrays",
			Description:   "Create a slice, append values 10, 20, 30 to it, and print the length",
			Story:         "🔧 Slices are more flexible than arrays. Master dynamic data!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 3,
			Hints: []string{
				"Create a slice: numbers := []int{}",
				"Use append: numbers = append(numbers, 10)",
				"Get length with len(numbers)",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create slice, append 10, 20, 30, print length
	
}`,
			Solution: `package main

import "fmt"

func main() {
	numbers := []int{}
	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)
	fmt.Println(len(numbers))
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "3\n",
					Description:    "Should print 3",
					Hidden:         false,
				},
			},
			Order: 7,
		},
		{
			ID:            "008",
			Title:         "Maps: Key-Value Storage",
			Description:   "Create a map with string keys and int values, add \"score\": 100, and print it",
			Story:         "🗺️ Maps are like dictionaries. Perfect for storing related data!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 4,
			Hints: []string{
				"Create a map: data := make(map[string]int)",
				"Add entry: data[\"score\"] = 100",
				"Print the value: fmt.Println(data[\"score\"])",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create map, add score: 100, print the value
	
}`,
			Solution: `package main

import "fmt"

func main() {
	data := make(map[string]int)
	data["score"] = 100
	fmt.Println(data["score"])
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "100\n",
					Description:    "Should print 100",
					Hidden:         false,
				},
			},
			Order: 8,
		},
		{
			ID:            "009",
			Title:         "String Manipulation",
			Description:   "Take a string \"golang\" and print it in uppercase",
			Story:         "✨ Text processing is essential. Learn string operations!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 4,
			Hints: []string{
				"Import the strings package",
				"Use strings.ToUpper()",
				"Print the result",
			},
			StarterCode: `package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "golang"
	// TODO: Convert to uppercase and print
	
}`,
			Solution: `package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "golang"
	fmt.Println(strings.ToUpper(text))
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "GOLANG\n",
					Description:    "Should print GOLANG",
					Hidden:         false,
				},
			},
			Order: 9,
		},
		{
			ID:            "010",
			Title:         "Multiple Return Values",
			Description:   "Create a function that returns both quotient and remainder of division",
			Story:         "🎯 Go functions can return multiple values. Use this power!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 5,
			Hints: []string{
				"Function signature: func divide(a, b int) (int, int)",
				"Return two values: return a/b, a%b",
				"Capture both: q, r := divide(10, 3)",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create divide function that returns quotient and remainder

func main() {
	q, r := divide(10, 3)
	fmt.Println(q, r)
}`,
			Solution: `package main

import "fmt"

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divide(10, 3)
	fmt.Println(q, r)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "3 1\n",
					Description:    "Should print quotient and remainder",
					Hidden:         false,
				},
			},
			Order: 10,
		},
	}
}
