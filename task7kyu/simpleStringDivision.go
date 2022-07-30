package task7kyu

// Simple string division
// https://www.codewars.com/kata/5b83c1c44a6acac33400009a

import(
	"strconv"
)

func Solve(st string, k int) int {
	
	l := len(st)-k
	res := 0 
	for i := 0; i <= len(st)-l; i++ {
		slice := st[i:i+l]
		check, err := strconv.Atoi(slice)
		if err != nil {
			return 0
		}
		if res < check {
			res = check
		}
	}
	return res
}