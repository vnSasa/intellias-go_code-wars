package task4kyu

// Number of integer partitions
// https://www.codewars.com/kata/546d5028ddbcbd4b8d001254

import (
	"fmt"
)

func Partitions(n int) int {
	modifications := make([]int, n+1)
	modifications[0] = 1

	for i := 1; i < n+1; i++ {
		for sum := 1; sum < len(modifications); sum++ {
			if sum >= i {
				modifications[sum] += modifications[sum-i]

				// i=1 sum=1 mod[1]=1 sum=2 mod[2]=1 sum=3 mod[3]=1 sum=4 mod[4]=1 sum=5 mod[5]=1 sum=6 mod[6]=1
				// i=2 sum=1 mod[1]=1 sum=2 mod[2]=2 sum=3 mod[3]=2 sum=4 mod[4]=3 sum=5 mod[5]=3 sum=6 mod[6]=4
				// i=3 sum=1 mod[1]=1 sum=2 mod[2]=2 sum=3 mod[3]=3 sum=4 mod[4]=4 sum=5 mod[5]=5 sum=6 mod[6]=7
				// i=4 sum=1 mod[1]=1 sum=2 mod[2]=2 sum=3 mod[3]=3 sum=4 mod[4]=5 sum=5 mod[5]=6 sum=6 mod[6]=9
				// i=5 sum=1 mod[1]=1 sum=2 mod[2]=2 sum=3 mod[3]=3 sum=4 mod[4]=5 sum=5 mod[5]=7 sum=6 mod[6]=10
				// i=6 sum=1 mod[1]=1 sum=2 mod[2]=2 sum=3 mod[3]=3 sum=4 mod[4]=5 sum=5 mod[5]=7 sum=6 mod[6]=11
			}
		}
		fmt.Printf("mod %v = %v\n", n, modifications[n])
	}

	return modifications[n]
}
