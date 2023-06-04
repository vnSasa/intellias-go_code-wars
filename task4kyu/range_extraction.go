package task4kyu

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(list []int) string {
	var ranges []string
	var start, end int
	for i := 0; i < len(list); i++ {
		if i == 0 || list[i]-list[i-1] != 1 {
			start = list[i]
			end = list[i]
		} else {
			end = list[i]
		}
		if i == len(list)-1 || list[i+1]-list[i] != 1 {
			if end-start >= 2 {
				ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
			} else {
				for j := start; j <= end; j++ {
					ranges = append(ranges, strconv.Itoa(j))
				}
			}
		}
	}
	return strings.Join(ranges, ",")
}
