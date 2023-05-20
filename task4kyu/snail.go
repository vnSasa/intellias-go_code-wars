package task4kyu

func Snail(snaipMap [][]int) []int {

	if len(snaipMap[0]) == 0 {
		return []int{}
	}
    
	var result []int

	top := 0
	bottom := len(snaipMap) - 1
	left := 0
	right := len(snaipMap[0]) - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, snaipMap[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, snaipMap[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, snaipMap[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, snaipMap[i][left])
			}
			left++
		}
	}

	return result
}