package task4kyu

// https://www.codewars.com/kata/52742f58faf5485cae000b9a

import (
	"strconv"
)

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	// describe the structure for convenient storage of time components
	type durationComponent struct {
		value int64
		unit  string
	}

	// Set values for time units in seconds
	const (
		second = 1
		minute = 60 * second
		hour   = 60 * minute
		day    = 24 * hour
		year   = 365 * day
	)

	// calculate time components
	var components []durationComponent

	years := seconds / year
	if years > 0 {
		components = append(components, durationComponent{years, "year"})
	}
	seconds -= years * year

	days := seconds / day
	if days > 0 {
		components = append(components, durationComponent{days, "day"})
	}
	seconds -= days * day

	hours := seconds / hour
	if hours > 0 {
		components = append(components, durationComponent{hours, "hour"})
	}
	seconds -= hours * hour

	minutes := seconds / minute
	if minutes > 0 {
		components = append(components, durationComponent{minutes, "minute"})
	}
	seconds -= minutes * minute

	if seconds > 0 {
		components = append(components, durationComponent{seconds, "second"})
	}

	/*
		Ця функція FormatDuration використовується для форматування вхідного часу у секундах у зрозумілу для користувача строку, що відображає тривалість.
		Давайте розглянемо дії, які відбуваються у цій функції:
		- Спочатку перевіряється, чи вхідне значення seconds дорівнює 0. Якщо так, то функція повертає "now", оскільки тривалість є нульовою.
		- Створюється тип durationComponent, який використовується для зручного зберігання компонентів часу (значення та одиницю виміру).
		- Задаються значення для одиниць часу у секундах: second, minute, hour, day та year.
		- Ініціалізується порожній масив components, який буде містити компоненти тривалості часу.
		- Обчислюються роки, шляхом ділення вхідного значення seconds на year. Якщо значення years більше 0, то створюється новий durationComponent з відповідними значеннями та додається до масиву components. Також віднімається від seconds відповідна кількість секунд за роками.
		- Повторюються кроки 5 для днів, годин, хвилин та секунд. Кожен компонент, якщо він більше 0, додається до масиву components, а відповідна кількість секунд віднімається від seconds.
		- Якщо після всіх обчислень залишилися ще секунди (seconds > 0), то створюється новий durationComponent з залишковими секундами та одиницею виміру "second" і додається до масиву components.
	*/

	var result string
	for i, c := range components {
		if c.value == 1 {
			result += strconv.FormatInt(c.value, 10) + " " + c.unit
		} else {
			result += strconv.FormatInt(c.value, 10) + " " + c.unit + "s"
		}

		if i < len(components)-2 {
			result += ", "
		} else if i == len(components)-2 {
			result += " and "
		}
	}

	/*
		- Формується відповідь, об'єднуючи компоненти тривалості в рядок result. Кожен компонент перетворюється в рядок за допомогою функції strconv.FormatInt, яка конвертує ціле число у його рядкове представлення.
		- У циклі for проходиться по всіх елементах масиву components. Для кожного компонента перевіряється, чи його значення (c.value) дорівнює 1. Якщо так, до рядка result додається рядок зі значенням та одиницею виміру без "s" (наприклад, "1 year"). У противному випадку, до рядка result додається рядок зі значенням та одиницею виміру з "s" (наприклад, "2 days").
		- Додається розділювач між компонентами тривалості: якщо індекс поточного компонента i менше довжини масиву components мінус 2, то додається кома та пробіл (", "). Якщо індекс дорівнює довжині масиву мінус 2, додається " and " як розділювач між останнім і передостаннім компонентами.
		- Після завершення циклу, рядок result містить сформовану тривалість у зрозумілому форматі.
		- Рядок result повертається як вихідний результат функції.

		Отже, ця функція обчислює компоненти тривалості часу (роки, дні, години, хвилини, секунди) на основі вхідного значення seconds і форматує їх у зрозумілу строку з правильними одиницями виміру та розділювачами.
	*/

	return result
}
