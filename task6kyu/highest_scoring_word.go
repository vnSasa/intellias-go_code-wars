package task6kyu

// Highest Scoring Word
// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272/train/go

import (
	"fmt"
	"strings"
)

func wordScore(s string) (score byte) {
	for i := range s {
		score += s[i] - 96
	}
	return
}

func High(s string) string {
	var score, newScore byte
	var word string

	for _, splitWord := range strings.Split(s, " ") {
		newScore = wordScore(splitWord)
		if newScore > score {
			score = newScore
			word = splitWord
		}
	}
	return word
}

func HigestScoringWordTest() {
	arr := [...][2]string{
		{"man i need a taxi up to ubud", "ubud"},
		{"what time are we climbing up the volcano", "volcano"},
		{"take me to semynak", "semynak"},
		{"aa b", "aa"},
		{"b aa", "b"},
		{"bb d", "bb"},
		{"d bb", "d"},
		{"aaa b", "aaa"},
	}
	for _, v := range arr {
		input := v[0]
		output := v[1]
		fmt.Printf("Input: %v\n", v[0])
		if strings.Compare(High(input), output) == 0 {
			fmt.Printf("Highest scoring word is %v\nExpected %v\nPassed\n\n", High(input), output)
		} else {
			fmt.Printf("Expected: %v --- But highest scoring word is: %v\nNOT PASSED\n\n", output, High(input))
		}
	}
}
