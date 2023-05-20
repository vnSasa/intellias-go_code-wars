package task4kyu

import (
	"math"
	"sort"
)

func Decompose(n int64) []int64 {
	return decomposeHelper(n, n*n)
}

func decomposeHelper(n, remainder int64) []int64 {
	if remainder == 0 {
		return []int64{}
	}

	for i := n - 1; i >= 1; i-- {
		newRemainder := remainder - int64(math.Pow(float64(i), 2))

		if newRemainder >= 0 {
			result := decomposeHelper(i, newRemainder)
			if result != nil {
				return append(result, i)
			}
		}
	}

	return nil
}

func IsSumOfSquares(arr []int64) bool {
	sum := int64(0)
	for _, num := range arr {
		sum += int64(math.Pow(float64(num), 2))
	}
	return sum == int64(math.Pow(float64(arr[0]), 2))
}

func SortSequence(arr []int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}