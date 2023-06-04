package task3kyu

import "math/big"

func Fib(n int64) *big.Int {
	if n < 0 {
		n = -n
		if n%2 == 0 {
			return new(big.Int).Neg(fibHelper(n))
		} else {
			return fibHelper(n)
		}
	}

	return fibHelper(n)

	/*
		Функція fib(n int64) *big.Int:

		- Перевіряє, чи n від'ємне.
		- Якщо n від'ємне, змінює його на його абсолютне значення і визначає знак результату в залежності від парності n.
		- Викликає функцію fibHelper зі змінним значенням n.
		- Повертає результат функції fibHelper.
	*/
}

func fibHelper(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	matrix := [][]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(1), big.NewInt(0)}}
	result := matrixPower(matrix, n-1)

	return result[0][0]

	/*
		Функція fibHelper(n int64) *big.Int:

		- Перевіряє, чи n дорівнює 0. Якщо так, повертає 0 як результат.
		- Створює початкову матрицю 2x2, яка представляє початкове значення чисел Фібоначчі.
		- Викликає функцію matrixPower з матрицею та значенням n-1, щоб отримати піднесення матриці до степеня n-1.
		- Повертає перший елемент матриці результату, який відповідає fib(n).
	*/
}

func matrixPower(matrix [][]*big.Int, n int64) [][]*big.Int {
	result := [][]*big.Int{{big.NewInt(1), big.NewInt(0)}, {big.NewInt(0), big.NewInt(1)}}
	for n > 0 {
		if n%2 == 1 {
			result = multiplyMatrices(result, matrix)
		}
		matrix = multiplyMatrices(matrix, matrix)
		n /= 2
	}
	return result

	/*
		Функція matrixPower(matrix [][]*big.Int, n int64) [][]*big.Int:

		- Ініціалізує початкову матрицю-результат, яка містить одиничну матрицю 2x2.
		- Виконує цикл, поки n не стане рівним 0.
		- Якщо n є непарним числом, множить поточну матрицю-результат на вхідну матрицю matrix за допомогою функції multiplyMatrices.
		- Множить вхідну матрицю matrix на саму себе (матричне піднесення до квадрату) за допомогою функції multiplyMatrices.
		- Ділить n на 2, щоб зменшити його на одну половину.
		- Повертає отриману матрицю-результат.
	*/
}

func multiplyMatrices(a, b [][]*big.Int) [][]*big.Int {
	c := [][]*big.Int{{big.NewInt(0), big.NewInt(0)}, {big.NewInt(0), big.NewInt(0)}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] = new(big.Int).Add(c[i][j], new(big.Int).Mul(a[i][k], b[k][j]))
			}
		}
	}
	return c

	/*
		Функція multiplyMatrices(a, b [][]*big.Int) [][]*big.Int:

		- Створює нову порожню матрицю c розміром 2x2 для зберігання результату.
		- Застосовує потрійний цикл для обчислення кожного елемента матриці c шляхом сумування добутків елементів матриць a і b.
		- Повертає отриману матрицю-результат c.
		- Ці функції разом утворюють ефективний алгоритм для обчислення чисел Фібоначчі з великими значеннями.
	*/
}
