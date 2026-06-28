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
		{
			ID:            "011-b",
			Title:         "Range: Iterate Over Slices",
			Description:   "Use range to iterate over a slice of numbers and print each one",
			Story:         "🔁 Range makes iteration easy. Master this Go idiom!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 5,
			Hints: []string{
				"Use for i, value := range slice",
				"Print each value",
				"Range returns index and value",
			},
			StarterCode: `package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	// TODO: Use range to iterate and print each number
	
}`,
			Solution: `package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	for _, num := range numbers {
		fmt.Println(num)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "10\n20\n30\n40\n50\n",
					Description:    "Should print all numbers",
					Hidden:         false,
				},
			},
			Order: 11,
		},
		{
			ID:            "012-b",
			Title:         "Switch Statement",
			Description:   "Use a switch statement to print the day name for day number 3",
			Story:         "🎯 Switch statements make multiple conditions cleaner!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 5,
			Hints: []string{
				"Use switch day { case 1: ... case 2: ... }",
				"Day 3 should print \"Wednesday\"",
				"Don't forget the default case",
			},
			StarterCode: `package main

import "fmt"

func main() {
	day := 3
	// TODO: Use switch to print day name
	
}`,
			Solution: `package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Wednesday\n",
					Description:    "Should print Wednesday",
					Hidden:         false,
				},
			},
			Order: 12,
		},
		{
			ID:            "013-b",
			Title:         "Variadic Functions",
			Description:   "Create a sum function that accepts any number of integers",
			Story:         "📊 Variadic functions accept variable arguments. Very powerful!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 6,
			Hints: []string{
				"Use ...int for variadic parameter",
				"func sum(numbers ...int) int",
				"Loop through numbers to calculate sum",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create sum function that accepts variadic int parameters

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}`,
			Solution: `package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "15\n",
					Description:    "Should print sum of 1+2+3+4+5",
					Hidden:         false,
				},
			},
			Order: 13,
		},
		{
			ID:            "014-b",
			Title:         "String Formatting",
			Description:   "Use fmt.Sprintf to create a formatted string with name and age",
			Story:         "✨ String formatting is essential for output. Master it!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 6,
			Hints: []string{
				"Use fmt.Sprintf with %s for strings and %d for integers",
				"Format: \"Name: %s, Age: %d\"",
				"Print the formatted string",
			},
			StarterCode: `package main

import "fmt"

func main() {
	name := "Bob"
	age := 30
	// TODO: Create formatted string "Name: Bob, Age: 30" and print it
	
}`,
			Solution: `package main

import "fmt"

func main() {
	name := "Bob"
	age := 30
	message := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(message)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Name: Bob, Age: 30\n",
					Description:    "Should print formatted string",
					Hidden:         false,
				},
			},
			Order: 14,
		},
		{
			ID:            "015-b",
			Title:         "Type Conversion",
			Description:   "Convert a float to an int and print the result",
			Story:         "🔄 Type conversion is crucial. Learn to transform data types!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 6,
			Hints: []string{
				"Use int() to convert float to int",
				"value := 3.14",
				"converted := int(value)",
			},
			StarterCode: `package main

import "fmt"

func main() {
	value := 3.14
	// TODO: Convert to int and print
	
}`,
			Solution: `package main

import "fmt"

func main() {
	value := 3.14
	converted := int(value)
	fmt.Println(converted)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "3\n",
					Description:    "Should print 3",
					Hidden:         false,
				},
			},
			Order: 15,
		},
		{
			ID:            "016-b",
			Title:         "Constants Declaration",
			Description:   "Declare a constant PI with value 3.14159 and print it",
			Story:         "🔒 Constants are immutable values. Perfect for fixed data!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      15,
			RequiredLevel: 7,
			Hints: []string{
				"Use const keyword",
				"const PI = 3.14159",
				"Constants cannot be changed",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Declare constant PI = 3.14159 and print it
	
}`,
			Solution: `package main

import "fmt"

func main() {
	const PI = 3.14159
	fmt.Println(PI)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "3.14159\n",
					Description:    "Should print PI",
					Hidden:         false,
				},
			},
			Order: 16,
		},
		{
			ID:            "017-b",
			Title:         "Iota Enumerator",
			Description:   "Use iota to create constants for days of week (Sunday=0, Monday=1, Tuesday=2)",
			Story:         "🔢 Iota creates sequential constants automatically!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 7,
			Hints: []string{
				"const ( Sunday = iota; Monday; Tuesday )",
				"iota starts at 0 and increments",
				"Print Tuesday to verify",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create constants using iota and print Tuesday
	
}`,
			Solution: `package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
)

func main() {
	fmt.Println(Tuesday)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "2\n",
					Description:    "Tuesday should be 2",
					Hidden:         false,
				},
			},
			Order: 17,
		},
		{
			ID:            "018-b",
			Title:         "Boolean Logic",
			Description:   "Check if a number is between 10 and 20 (inclusive) using boolean operators",
			Story:         "✅ Boolean logic is fundamental to programming!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 7,
			Hints: []string{
				"Use && for AND operation",
				"Check num >= 10 && num <= 20",
				"Print true or false",
			},
			StarterCode: `package main

import "fmt"

func main() {
	num := 15
	// TODO: Check if num is between 10 and 20, print result
	
}`,
			Solution: `package main

import "fmt"

func main() {
	num := 15
	result := num >= 10 && num <= 20
	fmt.Println(result)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "true\n",
					Description:    "15 is between 10 and 20",
					Hidden:         false,
				},
			},
			Order: 18,
		},
		{
			ID:            "019-b",
			Title:         "Break Statement",
			Description:   "Loop from 1 to 10 but break when you reach 5",
			Story:         "🛑 Break lets you exit loops early!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 8,
			Hints: []string{
				"Use for loop from 1 to 10",
				"if i == 5 { break }",
				"Print numbers before breaking",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Loop 1-10, print each, break at 5
	
}`,
			Solution: `package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\n3\n4\n",
					Description:    "Should print 1-4 then break",
					Hidden:         false,
				},
			},
			Order: 19,
		},
		{
			ID:            "020-b",
			Title:         "Continue Statement",
			Description:   "Print numbers 1-5 but skip 3 using continue",
			Story:         "⏭️ Continue skips to next iteration!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 8,
			Hints: []string{
				"Use for loop 1 to 5",
				"if i == 3 { continue }",
				"Print after the continue check",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Loop 1-5, skip 3, print others
	
}`,
			Solution: `package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "1\n2\n4\n5\n",
					Description:    "Should skip 3",
					Hidden:         false,
				},
			},
			Order: 20,
		},
		{
			ID:            "021-b",
			Title:         "Named Return Values",
			Description:   "Create a function with named return values that calculates area and perimeter of rectangle",
			Story:         "📛 Named returns make code more readable!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 8,
			Hints: []string{
				"func calc(w, h int) (area, perimeter int)",
				"Assign to area and perimeter directly",
				"Use naked return",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create function with named returns

func main() {
	a, p := calculate(5, 10)
	fmt.Println(a, p)
}`,
			Solution: `package main

import "fmt"

func calculate(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func main() {
	a, p := calculate(5, 10)
	fmt.Println(a, p)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "50 30\n",
					Description:    "Area=50, Perimeter=30",
					Hidden:         false,
				},
			},
			Order: 21,
		},
		{
			ID:            "022-b",
			Title:         "Anonymous Function",
			Description:   "Create and immediately call an anonymous function that prints a message",
			Story:         "🎭 Anonymous functions are functions without names!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 9,
			Hints: []string{
				"func() { ... }()",
				"Define and call in one expression",
				"Use () at the end to call it",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create and call anonymous function that prints "Hello Anonymous"
	
}`,
			Solution: `package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello Anonymous")
	}()
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "Hello Anonymous\n",
					Description:    "Should print message",
					Hidden:         false,
				},
			},
			Order: 22,
		},
		{
			ID:            "023-b",
			Title:         "Recursive Factorial",
			Description:   "Write a recursive function to calculate factorial of 5",
			Story:         "🔄 Recursion: a function calling itself!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      40,
			RequiredLevel: 9,
			Hints: []string{
				"Base case: if n == 0 return 1",
				"Recursive case: return n * factorial(n-1)",
				"factorial(5) = 120",
			},
			StarterCode: `package main

import "fmt"

// TODO: Create recursive factorial function

func main() {
	result := factorial(5)
	fmt.Println(result)
}`,
			Solution: `package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	result := factorial(5)
	fmt.Println(result)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "120\n",
					Description:    "5! = 120",
					Hidden:         false,
				},
			},
			Order: 23,
		},
		{
			ID:            "024-b",
			Title:         "Slice Capacity",
			Description:   "Create a slice with length 3 and capacity 5, print both",
			Story:         "📏 Understanding capacity is key to slice performance!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 10,
			Hints: []string{
				"Use make([]int, length, capacity)",
				"len() returns length",
				"cap() returns capacity",
			},
			StarterCode: `package main

import "fmt"

func main() {
	// TODO: Create slice with len=3, cap=5, print both
	
}`,
			Solution: `package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Println(len(s), cap(s))
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "3 5\n",
					Description:    "Length 3, capacity 5",
					Hidden:         false,
				},
			},
			Order: 24,
		},
		{
			ID:            "025-b",
			Title:         "Copy Slices",
			Description:   "Create a slice, copy it to another slice, and print the copy",
			Story:         "📋 Copying slices prevents unwanted modifications!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 10,
			Hints: []string{
				"Use copy(dst, src)",
				"Create destination slice with make",
				"copy returns number of elements copied",
			},
			StarterCode: `package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	// TODO: Copy to new slice and print it
	
}`,
			Solution: `package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Println(copied)
}`,
			TestCases: []models.TestCase{
				{
					Input:          "",
					ExpectedOutput: "[1 2 3]\n",
					Description:    "Should print copied slice",
					Hidden:         false,
				},
			},
			Order: 25,
		},
	}
}
