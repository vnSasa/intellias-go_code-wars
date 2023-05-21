package task4kyu

import (
	"fmt"
	"sort"
	"strings"
)

func Mix(s1, s2 string) string {
	countLetters := func(s string) map[rune]int {
		count := make(map[rune]int)
		for _, char := range s {
			if char >= 'a' && char <= 'z' {
				count[char]++
			}
		}
		return count
	}

	count1 := countLetters(s1)
	count2 := countLetters(s2)

	result := ""

	for char := 'a'; char <= 'z'; char++ {
		maxCount := max(count1[char], count2[char])

		if maxCount > 1 {
			if count1[char] > count2[char] {
				result += fmt.Sprintf("1:%s/", strings.Repeat(string(char), maxCount))
			} else if count2[char] > count1[char] {
				result += fmt.Sprintf("2:%s/", strings.Repeat(string(char), maxCount))
			} else {
				result += fmt.Sprintf("=:%s/", strings.Repeat(string(char), maxCount))
			}
		}
	}

	if len(result) > 0 {
		resultSlice := strings.Split(result[:len(result)-1], "/")
		sort.Slice(resultSlice, func(i, j int) bool {
			if len(resultSlice[i]) == len(resultSlice[j]) {
				return resultSlice[i] < resultSlice[j]
			}
			return len(resultSlice[i]) > len(resultSlice[j])
		})

		return strings.Join(resultSlice, "/")
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
