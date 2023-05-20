package task7kyu

// Make a function that does arithmetic!
// https://www.codewars.com/kata/583f158ea20cfcbeb400000a

import (
	"fmt"
)

func Calculator(a, b int, operation string)string {
	
	var out int
	switch {
	case operation == "add":
		out = a + b
	case operation == "subtract":
		out = a - b 
	case operation == "multiply":
		out = a * b 
	case operation == "divide":
		out = a / b
	}
	
	fmt.Printf("\nMake a function that does arithmetic!\n")

	return fmt.Sprint(a) + ", " + fmt.Sprint(b) + ", " + operation + " --> " + fmt.Sprint(out)
}