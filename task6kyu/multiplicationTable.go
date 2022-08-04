package task6kyu

// Multiplication table
// https://www.codewars.com/kata/534d2f5b5371ecf8d2000a08

func MultiplicationTable(size int) [][]int {
	multiTable := make([][]int, size)
	for i := range multiTable {
		multiTable[i] = make([]int, size)
	}
	counter := 1

	for r := 1; r <= size; r++ {
		for c := 1; c <= size; c++ {
			multiTable[r-1][c-1] = counter * c * r
		}

	}

	return multiTable
}