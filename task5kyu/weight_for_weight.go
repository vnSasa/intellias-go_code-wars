package task5kyu

import (
	"sort"
	"strings"
)

func OrderWeight(strng string) string {
	strNums := strings.Fields(strng)
	nums := make([]string, len(strNums))
	for i, strNum := range strNums {
		nums[i] = strNum
	}

	weight := func(num string) int {
		sum := 0
		for _, digit := range num {
			sum += int(digit - '0')
		}
		return sum
	}

	sort.Slice(nums, func(i, j int) bool {
		wi, wj := weight(nums[i]), weight(nums[j])
		if wi < wj {
			return true
		} else if wi > wj {
			return false
		} else {
			return nums[i] < nums[j]
		}
	})

	return strings.Join(nums, " ")
}
