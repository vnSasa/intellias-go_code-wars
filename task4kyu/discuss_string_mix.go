package task4kyu

// https://www.codewars.com/kata/5629db57620258aa9d000014/go

import (
	"fmt"
	"sort"
	"strings"
)

func Mix(s1, s2 string) string {
	countLetters := func(s string) map[rune]int {
		count := make(map[rune]int)
		for _, char := range s {
			if char >= 'a' && char <= 'z' {
				count[char]++
			}
		}
		return count
	}

	/*
		Цей фрагмент коду визначає внутрішню функцію countLetters, яка приймає рядок s як параметр і повертає карту (map) типу map[rune]int.
		Основна мета функції countLetters полягає в підрахунку кількості входжень кожного символу в рядку s. Функція перебирає кожен символ char у рядку s за допомогою циклу for _, char := range s.
		Усередині циклу перевіряється, чи належить символ char до діапазону від 'a' до 'z', включно, за допомогою умови char >= 'a' && char <= 'z'. Це забезпечує обробку лише латинських малих літер.
		Якщо символ char відповідає умові, то збільшується лічильник входжень цього символу в карті count за допомогою оператора count[char]++. Значення лічильника збільшується на 1 для кожного входження символу.
		Після завершення циклу, функція повертає карту count, де ключами є символи (тип rune), а значеннями - кількість входжень кожного символу в рядок s.
		Цей фрагмент коду використовується в основній функції Mix для підрахунку кількості букв у рядках s1 і s2, застосовуючи функцію countLetters(s1) і countLetters(s2) відповідно.
	*/

	count1 := countLetters(s1)
	count2 := countLetters(s2)

	result := ""

	for char := 'a'; char <= 'z'; char++ {
		maxCount := max(count1[char], count2[char])

		if maxCount > 1 {
			if count1[char] > count2[char] {
				result += fmt.Sprintf("1:%s/", strings.Repeat(string(char), maxCount))
			} else if count2[char] > count1[char] {
				result += fmt.Sprintf("2:%s/", strings.Repeat(string(char), maxCount))
			} else {
				result += fmt.Sprintf("=:%s/", strings.Repeat(string(char), maxCount))
			}
		}
	}

	/*
		Цей фрагмент коду виконує змішування символів з рядків s1 і s2 та формує рядок result, використовуючи кількість входжень кожної букви у вхідних рядках.
		1. count1 := countLetters(s1): Функція countLetters викликається для рядка s1, і результат зберігається в змінній count1. Це дає кількість входжень кожної букви у рядку s1.
		2. count2 := countLetters(s2): Функція countLetters викликається для рядка s2, і результат зберігається в змінній count2. Це дає кількість входжень кожної букви у рядку s2.
		3. result := "": Ініціалізується порожній рядок result, який буде використовуватись для збереження результату змішування символів.
		4. Цикл for char := 'a'; char <= 'z'; char++: Цикл проходить по всіх буквах від 'a' до 'z', включно.
		5. maxCount := max(count1[char], count2[char]): Знаходиться максимальна кількість входжень букви char у рядках s1 і s2 за допомогою функції max. Результат зберігається в змінній maxCount.
		6. if maxCount > 1 { ... }: Перевіряється, чи maxCount більше 1. Це означає, що дана буква зустрічається не менше двох разів у одному з рядків s1 або s2.
		7. У внутрішньому блоку if, залежно від кількості входжень букви у s1 і s2, формується частинка результату і додається до рядка result за допомогою result += ....
			- Якщо кількість входжень букви char у s1 більше, ніж у s2, то додається рядок у форматі "1:буква/кількість_разів/". Наприклад, якщо char = 'a' і maxCount = 3, то додається "1:aaa/" до result.
			- Якщо кількість входжень букви char у s2 більше, ніж у s1, то додається рядок у форматі "2:буква/кількість_разів/". Наприклад, якщо char = 'b' і maxCount = 2, то додається "2:bb/" до result.
			- Якщо кількість входжень букви char у s1 і s2 однакова, то додається рядок у форматі "=:буква/кількість_разів/". Наприклад, якщо char = 'c' і maxCount = 4, то додається "=:cccc/" до result.
		8. Завершивши проходження циклу, перевіряється довжина рядка result за допомогою len(result). Якщо довжина більша за 0, виконується наступний блок коду.
	*/

	if len(result) > 0 {
		resultSlice := strings.Split(result[:len(result)-1], "/")
		sort.Slice(resultSlice, func(i, j int) bool {
			if len(resultSlice[i]) == len(resultSlice[j]) {
				return resultSlice[i] < resultSlice[j]
			}
			return len(resultSlice[i]) > len(resultSlice[j])
		})

		return strings.Join(resultSlice, "/")
	}

	/*
		Цей фрагмент коду перевіряє, чи рядок result має довжину більше 0. Якщо це так, то виконується наступний блок коду:
		1. resultSlice := strings.Split(result[:len(result)-1], "/"): Рядок result вирізається, щоб вилучити останній символ "/", потім використовується функція strings.Split для розділення рядка на частинки за роздільником "/". Результат зберігається в змінній resultSlice, яка стає срізом рядків.
		2. sort.Slice(resultSlice, func(i, j int) bool { ... }): Сріз resultSlice сортується за допомогою функції sort.Slice. Використовується анонімна функція, яка порівнює елементи resultSlice[i] та resultSlice[j] згідно з наступними умовами:
			- Якщо довжина resultSlice[i] рівна довжині resultSlice[j], то порівнюються самі рядки за допомогою оператора <. Тобто, рядки сортуються в лексикографічному порядку.
			- Якщо довжина resultSlice[i] більша за довжину resultSlice[j], то resultSlice[i] розташовується перед resultSlice[j].
		3. return strings.Join(resultSlice, "/"): Відсортований сріз resultSlice з'єднується за допомогою роздільника "/" за допомогою функції strings.Join. Результат є змішаним рядком з буквами, кількістю входжень та відповідними індексами, який повертається як результат функції Mix.
		Цей фрагмент коду виконує сортування та форматування результату перед поверненням його як вихідного значення функції Mix.
	*/

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
