package challenges

import "gohacker/models"

// GetBeginnerChallenges returns all beginner-level challenges
func GetBeginnerChallenges() []models.Challenge {
	return []models.Challenge{
		{
			ID:            "001",
			Title:         "First Breach: Variable Declaration",
			Description:   "Hack the entry node by declaring a secret password variable set to \"admin123\".",
			Story:         "🛡️ Your first quest begins at the login gate. Plant the secret password variable to unlock the system and prove you belong in the crew.",
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
			Description:   "Breach the vault by computing the access code from two secret values, x=10 and y=20.",
			Story:         "🔐 The vault's cipher is simple but lethal. Combine the secret integers to reveal the code and advance deeper.",
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
			Description:   "Probe the firewall and print \"Big\" if the signal is above 10, otherwise print \"Small\".",
			Story:         "🕵️‍♂️ Your code is scanning a security trigger. Use an if-statement to decide whether the signal is strong enough to breach.",
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
			Title:         "If Statements",
			Description:   "Inspect a signal, use an if-statement to mark it positive, then output your verdict.",
			Story:         "⚡ The system logs a packet. Use logic to judge whether it is a positive trigger and send back the right response.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
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
			Description:   "Forge a greeting function that injects a name into the welcome message: \"Hello, [name]!\".",
			Story:         "🧩 Build a helper function like a hacker's utility; it should craft the perfect greeting for any alias.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
			Description:   "Cache three codewords in an array and exfiltrate the second one.",
			Story:         "📡 Your loot cache holds three secrets. Pull out the second entry to continue your reconnaissance.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
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
			Description:   "Build a packet queue, append 10, 20, 30, then report how many entries are staged.",
			Story:         "🧠 The network buffer grows as payloads arrive. Append the values and report the queue size.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
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
			Description:   "Craft a key-value exploit map, stash \"score\" at 100, and retrieve it.",
			Story:         "🗝️ In the hacker vault, keys open hidden values. Store the score and fetch it like a pro.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
			Description:   "Elevate the secret codeword \"golang\" to uppercase and unlock the next stage.",
			Story:         "🧪 Decode the signature by transforming the codeword into uppercase — the signal must be loud and clear.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
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
			Description:   "Send a division function that returns quotient and remainder for the extraction.",
			Story:         "🔧 Craft a utility that splits a value cleanly and returns both parts, just like a hacker dividing loot.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 1,
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
			Description:   "Sweep through the numeric slice with range and log every value.",
			Story:         "🧭 Launch a range scan over the data stream. Each number is a signal to capture.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
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
			Description:   "Decrypt day code 3 with a switch statement and print its name.",
			Story:         "🧠 The scheduler uses code numbers for days. Branch your logic and decode day 3.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
			Description:   "Write a sum function that ingests any number of values and returns the total.",
			Story:         "⚡ Build a flexible summoner function that can absorb any number of integers and return the total payload.",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 1,
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
			Description:   "Use fmt.Sprintf to forge a formatted string containing name and age.",
			Story:         "✨ String formatting is essential for output. Master it!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
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
			Description:   "Convert the intercepted float 3.14 into an integer and print the result.",
			Story:         "🔄 Type conversion is crucial. Learn to transform data types!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
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
			Description:   "Declare a constant PI = 3.14159 and print its value to lock in the numeric key.",
			Story:         "🔒 Constants are immutable values. Perfect for fixed data!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      15,
			RequiredLevel: 1,
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
			Description:   "Use iota to define weekday constants and print the Tuesday code.",
			Story:         "🔢 Iota creates sequential constants automatically!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
			Description:   "Check whether a signal is between 10 and 20 inclusive and print true or false.",
			Story:         "✅ Boolean logic is fundamental to programming!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
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
			Description:   "Loop from 1 to 10 and cut the scan short when you hit 5.",
			Story:         "🛑 Break lets you exit loops early!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
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
			Description:   "Print 1 through 5 but bypass signal 3 with continue.",
			Story:         "⏭️ Continue skips to next iteration!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
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
			Description:   "Build a function with named return values that computes rectangle area and perimeter.",
			Story:         "📛 Named returns make code more readable!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 1,
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
			Description:   "Define and immediately invoke an anonymous function to print a mission message.",
			Story:         "🎭 Anonymous functions are functions without names!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
			Description:   "Implement a recursive factorial function and compute 5!.",
			Story:         "🔄 Recursion: a function calling itself!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      40,
			RequiredLevel: 1,
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
			Description:   "Create a slice with length 3 and capacity 5, then print both len and cap.",
			Story:         "📏 Understanding capacity is key to slice performance!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
			Description:   "Copy one slice into another and print the copied slice.",
			Story:         "📋 Copying slices prevents unwanted modifications!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
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
		{
			ID:            "026-b",
			Title:         "Multi-dimensional Arrays",
			Description:   "Create a 2x3 array and print the element at row 1, column 2.",
			Story:         "🎯 Multi-dimensional arrays store matrix data!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	// TODO: Create 2x3 array, set [1][2] = 6, print it
}`,
			Solution: `package main
import "fmt"
func main() {
	arr := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr[1][2])
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "6\n", Description: "Should print 6", Hidden: false}},
			Order:     26,
		},
		{
			ID:            "027-b",
			Title:         "Nested Maps",
			Description:   "Construct a nested map for user scores and print Alice's math score.",
			Story:         "🗺️ Nested maps handle complex data!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	// TODO: Create map[string]map[string]int, add user "alice" with score "math": 90
}`,
			Solution: `package main
import "fmt"
func main() {
	scores := make(map[string]map[string]int)
	scores["alice"] = map[string]int{"math": 90}
	fmt.Println(scores["alice"]["math"])
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "90\n", Description: "Should print 90", Hidden: false}},
			Order:     27,
		},
		{
			ID:            "028-b",
			Title:         "Struct Embedding",
			Description:   "Embed one struct inside another and print the embedded field.",
			Story:         "🎁 Embedding promotes composition over inheritance!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      40,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
type Person struct { Name string }
type Employee struct { Person; ID int }
func main() {
	e := Employee{Person{"Alice"}, 123}
	fmt.Println(e.Name, e.ID)
}`,
			Solution: `package main
import "fmt"
type Person struct { Name string }
type Employee struct { Person; ID int }
func main() {
	e := Employee{Person{"Alice"}, 123}
	fmt.Println(e.Name, e.ID)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "Alice 123\n", Description: "Should print name and ID", Hidden: false}},
			Order:     28,
		},
		{
			ID:            "029-b",
			Title:         "String Concatenation",
			Description:   "Concatenate two strings and print the joined result.",
			Story:         "➕ String concatenation builds text!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	s1, s2 := "Hello", "World"
	// TODO: Concatenate and print
}`,
			Solution: `package main
import "fmt"
func main() {
	s1, s2 := "Hello", "World"
	fmt.Println(s1 + " " + s2)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "Hello World\n", Description: "Should concatenate", Hidden: false}},
			Order:     29,
		},
		{
			ID:            "030-b",
			Title:         "String Contains",
			Description:   "Check whether a string contains the substring \"World\" and print the boolean.",
			Story:         "🔍 String searching is essential!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
			StarterCode: `package main
import ("fmt"; "strings")
func main() {
	text := "Hello World"
	// TODO: Check if contains "World", print true/false
}`,
			Solution: `package main
import ("fmt"; "strings")
func main() {
	text := "Hello World"
	fmt.Println(strings.Contains(text, "World"))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "true\n", Description: "Should find substring", Hidden: false}},
			Order:     30,
		},
		{
			ID:            "031-b",
			Title:         "String Builder",
			Description:   "Use strings.Builder to efficiently assemble and print a combined string.",
			Story:         "🏗️ Builder is faster for multiple concatenations!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 1,
			StarterCode: `package main
import ("fmt"; "strings")
func main() {
	var b strings.Builder
	// TODO: Write "Go", "Lang", print result
}`,
			Solution: `package main
import ("fmt"; "strings")
func main() {
	var b strings.Builder
	b.WriteString("Go")
	b.WriteString("Lang")
	fmt.Println(b.String())
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "GoLang\n", Description: "Should build string", Hidden: false}},
			Order:     31,
		},
		{
			ID:            "032-b",
			Title:         "Zero Values",
			Description:   "Print the zero values for int, string, and bool.",
			Story:         "0️⃣ Understanding zero values prevents bugs!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      20,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	var i int; var s string; var b bool
	fmt.Println(i, s, b)
}`,
			Solution: `package main
import "fmt"
func main() {
	var i int; var s string; var b bool
	fmt.Println(i, s, b)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "0  false\n", Description: "Should print zero values", Hidden: false}},
			Order:     32,
		},
		{
			ID:            "033-b",
			Title:         "Variable Scope",
			Description:   "Show how block scope decides which variable is printed.",
			Story:         "📦 Scope determines variable visibility!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	x := 1
	{ x := 2; fmt.Println(x) }
	fmt.Println(x)
}`,
			Solution: `package main
import "fmt"
func main() {
	x := 1
	{ x := 2; fmt.Println(x) }
	fmt.Println(x)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "2\n1\n", Description: "Should show scope", Hidden: false}},
			Order:     33,
		},
		{
			ID:            "034-b",
			Title:         "Nested Loops",
			Description:   "Print a 2x2 multiplication table using nested loops.",
			Story:         "🔄 Nested loops handle 2D iteration!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	// TODO: Print 2x2 multiplication table
}`,
			Solution: `package main
import "fmt"
func main() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println()
	}
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "1 2 \n2 4 \n", Description: "Should print table", Hidden: false}},
			Order:     34,
		},
		{
			ID:            "035-b",
			Title:         "Loop Patterns",
			Description:   "Demonstrate a while-style for loop by printing 0, 1, 2.",
			Story:         "🔁 Go has flexible loop syntax!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	i := 0
	for i < 3 { fmt.Println(i); i++ }
}`,
			Solution: `package main
import "fmt"
func main() {
	i := 0
	for i < 3 { fmt.Println(i); i++ }
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "0\n1\n2\n", Description: "Should loop", Hidden: false}},
			Order:     35,
		},
		{
			ID:            "036-b",
			Title:         "Early Returns",
			Description:   "Use an early return in a function so it yields \"negative\" or \"positive\" cleanly.",
			Story:         "↩️ Early returns improve readability!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func check(n int) string {
	if n < 0 { return "negative" }
	return "positive"
}
func main() { fmt.Println(check(5)) }`,
			Solution: `package main
import "fmt"
func check(n int) string {
	if n < 0 { return "negative" }
	return "positive"
}
func main() { fmt.Println(check(5)) }`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "positive\n", Description: "Should return early", Hidden: false}},
			Order:     36,
		},
		{
			ID:            "037-b",
			Title:         "Function as Value",
			Description:   "Assign a function to a variable and invoke it to print the result.",
			Story:         "🎯 Functions are first-class values!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      35,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	add := func(a, b int) int { return a + b }
	fmt.Println(add(2, 3))
}`,
			Solution: `package main
import "fmt"
func main() {
	add := func(a, b int) int { return a + b }
	fmt.Println(add(2, 3))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "5\n", Description: "Should call function", Hidden: false}},
			Order:     37,
		},
		{
			ID:            "038-b",
			Title:         "Higher-Order Functions",
			Description:   "Pass a function as a parameter and use it to transform a value.",
			Story:         "🎓 Higher-order functions enable powerful patterns!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      40,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func apply(f func(int) int, v int) int { return f(v) }
func main() {
	double := func(n int) int { return n * 2 }
	fmt.Println(apply(double, 5))
}`,
			Solution: `package main
import "fmt"
func apply(f func(int) int, v int) int { return f(v) }
func main() {
	double := func(n int) int { return n * 2 }
	fmt.Println(apply(double, 5))
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "10\n", Description: "Should apply function", Hidden: false}},
			Order:     38,
		},
		{
			ID:            "039-b",
			Title:         "Bitwise Operators",
			Description:   "Compute the bitwise AND of 12 and 10 and print the result.",
			Story:         "⚡ Bitwise ops are fast and powerful!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      30,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
func main() {
	a, b := 12, 10
	fmt.Println(a & b)
}`,
			Solution: `package main
import "fmt"
func main() {
	a, b := 12, 10
	fmt.Println(a & b)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "8\n", Description: "Should AND bits", Hidden: false}},
			Order:     39,
		},
		{
			ID:            "040-b",
			Title:         "Type Aliases",
			Description:   "Create a type alias and use it to declare a UserID value.",
			Story:         "📛 Type aliases improve code readability!",
			Difficulty:    models.DifficultyNewbie,
			Category:      models.CategoryBasics,
			XPReward:      25,
			RequiredLevel: 1,
			StarterCode: `package main
import "fmt"
type UserID int
func main() {
	var id UserID = 123
	fmt.Println(id)
}`,
			Solution: `package main
import "fmt"
type UserID int
func main() {
	var id UserID = 123
	fmt.Println(id)
}`,
			TestCases: []models.TestCase{{Input: "", ExpectedOutput: "123\n", Description: "Should use alias", Hidden: false}},
			Order:     40,
		},
	}
}
