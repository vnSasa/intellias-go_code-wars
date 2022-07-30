package task7kyu

import (
	"strings"
	"sort"
)

type Tuple struct {
	R rune
	Diff int
}

func SolveMost(s string) rune {
	arr := []Tuple{}
	for i, r := range s {
		lastIdx := strings.LastIndex(s, string(r))
		arr = append(arr, Tuple{R: r, Diff: lastIdx - i})
	}
  
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Diff > arr[j].Diff {
			return true
		}
		if arr[i].Diff < arr[j].Diff {
			return false
		}
		return arr[i].R < arr[j].R
	})
	return arr[0].R
}