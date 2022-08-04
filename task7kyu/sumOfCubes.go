package task7kyu

// Sum of Cubes
// https://www.codewars.com/kata/59a8570b570190d313000037

import(
	"fmt"
)

func SumOfCube(n int)int {
	var sum int
	for i := 1; i <= n; i++ {
		cube := i*i*i 
		sum += cube
	}

	fmt.Printf("\nSum of Cubes\n")

	return sum
}