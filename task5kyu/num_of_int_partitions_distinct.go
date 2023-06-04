package task5kyu

// Number of integer partitions with distinct parts
// https://www.codewars.com/kata/6267a007e67fba0058725ad2/go

import (
	"fmt"
	"math/big"
)

func Partitions(n int) *big.Int {
	modifications := make([]int, n+1)
	modifications[0] = 1

	for i := 1; i < n+1; i += 2 {
		for sum := 1; sum < len(modifications); sum++ {
			if sum >= i {
				modifications[sum] += modifications[sum-i]
			}
		}
		fmt.Printf("mod %v = %v\n", n, modifications[n])
	}
	return big.NewInt(int64(modifications[n]))
}
