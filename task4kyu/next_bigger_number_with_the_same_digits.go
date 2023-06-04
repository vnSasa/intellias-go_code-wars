package task4kyu

// https://www.codewars.com/kata/55983863da40caa2c900004e

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

	/*
		Функція NextBigger приймає ціле число n і повертає наступне більше число, яке можна утворити шляхом переставлення цифр у числі n. Якщо таке число не існує, функція повертає -1.
		Опишемо дії, які відбуваються в цій функції:
		- Число n перетворюється на рядок numStr за допомогою функції strconv.Itoa, яка конвертує ціле число у його рядкове представлення.
		- Змінна index ініціалізується значенням -1. Ця змінна буде використовуватись для збереження індексу, де знаходиться перший непарний символ з кінця рядка numStr, який задовольняє умову numStr[i] < numStr[i+1].
		- Запускається цикл for, який проходиться по символам рядка numStr зправа наліво, починаючи з передостаннього символу. Для кожного символу numStr[i] перевіряється, чи виконується умова numStr[i] < numStr[i+1]. Якщо так, зберігається індекс i в змінну index, і цикл переривається.
		- Якщо значення змінної index дорівнює -1, це означає, що немає непарних символів у рядку numStr, для яких numStr[i] < numStr[i+1]. У такому випадку повертається -1, означаючи, що наступного більшого числа не існує.
		- Змінна nextIndex ініціалізується значенням -1. Ця змінна буде використовуватись для збереження індексу, де знаходиться найменше число з рядку numStr (починаючи зправа від index), яке є більшим за numStr[index].
		- Запускається цикл for, який проходиться по символам рядка numStr зправа наліво, починаючи з останнього символу. Для кожного символу numStr[i] перевіряється, чи виконується умова numStr[i] > numStr[index]. Якщо так, зберігається індекс i в змінну nextIndex, і цикл переривається.
		- Рядок numStr перетворюється на масив байтів numStrBytes за допомогою конвертації типу.
		- За допомогою обміну елементів масиву numStrBytes між індексами index і nextIndex, значеннями цих елементів здійснюється перестановка. Це виконується за допомогою наступного рядка коду: numStrBytes[index], numStrBytes[nextIndex] = numStrBytes[nextIndex], numStrBytes[index].
		- Застосовується сортування зростанням до підмасиву numStrBytes[index+1:], використовуючи функцію sort.Slice та передаючи у неї функцію-компаратор, яка порівнює символи за їх значенням.
		- Рядок numStrBytes конвертується назад в ціле число result за допомогою функції strconv.Atoi.
		- Число result повертається як вихідний результат функції NextBigger. Це число представляє наступне більше число за порядком, яке можна утворити шляхом перестановки цифр у вхідному числі n.
	*/
}
