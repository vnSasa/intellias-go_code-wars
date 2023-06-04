package task5kyu

import "math"

func ChooseBestSum(t, k int, ls []int) int {
	return chooseBestSumRecursive(t, k, ls, 0, 0)
}

func chooseBestSumRecursive(t, k int, ls []int, start, sum int) int {
	if k == 0 {
		if sum <= t {
			return sum
		}
		return -1
	}
	if start >= len(ls) {
		return -1
	}
	maxSum := -1
	for i := start; i < len(ls); i++ {
		newSum := chooseBestSumRecursive(t, k-1, ls, i+1, sum+ls[i])
		maxSum = int(math.Max(float64(maxSum), float64(newSum)))
	}
	return maxSum
}

func findMaxSum(t, k int, ls []int, start, sum int, maxSum *int) {
	if k == 0 || sum > t {
		return
	}

	for i := start; i < len(ls); i++ {
		newSum := sum + ls[i]

		if newSum <= t && k-1 <= 0 {
			if newSum > *maxSum {
				*maxSum = newSum
			}
			return
		}

		findMaxSum(t, k-1, ls, i+1, newSum, maxSum)
	}
}
