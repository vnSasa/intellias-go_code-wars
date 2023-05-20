package task7kyu

// If you can't sleep, just count sheep!!
// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077

import (
	"fmt"
)

func Sheep(x int)string {
	var text string
	var i uint
	for i = 1; i <= uint(x); i++ {
		text += fmt.Sprint(i) + " sheep..."
	}

	fmt.Printf("\nIf you can't sleep, just count sheep!!\n")

	return text
}