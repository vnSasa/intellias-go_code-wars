package task4kyu

import (
	"sort"
	"strconv"
)

func NextSmaller(n int) int {
	numStr := strconv.Itoa(n)
	length := len(numStr)

	// Знайдемо індекс першої цифри, яка менша за наступну праворуч від неї
	index := -1
	for i := length - 2; i >= 0; i-- {
		if numStr[i] > numStr[i+1] {
			index = i
			break
		}
	}

	if index == -1 {
		return -1 // Немає менших чисел
	}

	// Знайдемо індекс найбільшої цифри праворуч від index, яка менша за numStr[index]
	nextIndex := -1
	for i := length - 1; i > index; i-- {
		if numStr[i] < numStr[index] {
			nextIndex = i
			break
		}
	}

	// Обміняємо цифри numStr[index] і numStr[nextIndex]
	numStrBytes := []byte(numStr)
	numStrBytes[index], numStrBytes[nextIndex] = numStrBytes[nextIndex], numStrBytes[index]

	// Сортуємо цифри праворуч від index в порядку спадання
	sort.Slice(numStrBytes[index+1:], func(i, j int) bool {
		return numStrBytes[index+1+i] > numStrBytes[index+1+j]
	})

	// Перевіряємо, щоб перше число було ненульовим
	if numStrBytes[0] == '0' {
		return -1
	}

	result, _ := strconv.Atoi(string(numStrBytes))
	return result
}
