package task4kyu

import (
	"sort"
	"strconv"
)

func NextBigger(n int) int {
    numStr := strconv.Itoa(n)

    index := -1
    for i := len(numStr) - 2; i >= 0; i-- {
        if numStr[i] < numStr[i+1] {
            index = i
            break
        }
    }

    if index == -1 {
        return -1
    }

    nextIndex := -1
    for i := len(numStr) - 1; i > index; i-- {
        if numStr[i] > numStr[index] {
            nextIndex = i
            break
        }
    }

    numStrBytes := []byte(numStr)
    numStrBytes[index], numStrBytes[nextIndex] = numStrBytes[nextIndex], numStrBytes[index]

    sort.Slice(numStrBytes[index+1:], func(i, j int) bool {
        return numStrBytes[index+1+i] < numStrBytes[index+1+j]
    })

    result, _ := strconv.Atoi(string(numStrBytes))
    return result
}