package task4kyu

func FindAll(sumDig, digs int) []int {
	var result []int
	findNumbers(&result, []int{}, 1, 0, sumDig, digs)

	if len(result) == 0 {
		return []int{}
	}

	return []int{len(result), result[0], result[len(result)-1]}
}

func findNumbers(result *[]int, current []int, start, currentSum, sumDig, digs int) {
	if digs == 0 && currentSum == sumDig {
		num := 0
		for i := 0; i < len(current); i++ {
			num = num*10 + current[i]
		}
		*result = append(*result, num)
		return
	}

	if digs < 0 || currentSum > sumDig {
		return
	}

	for i := start; i <= 9; i++ {
		findNumbers(result, append(current, i), i, currentSum+i, sumDig, digs-1)
	}
}
