package main

import "os"

const (
	maxInt64 = 9223372036854775807
	minInt64 = -9223372036854775808
)

func stringToInt(str string) (int64, bool) {
	var result int64
	var sign int64 = 1
	i := 0

	if str[0] == '-' {
		sign = -1
		i = 1
	}

	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0, false
		}
		result = result*10 + int64(str[i]-'0')
	}

	return result * sign, true
}

func checkOverflow(val1, val2 int64, operator string) bool {
	switch operator {
	case "+":
		if val1 > 0 && val2 > 0 && val1+val2 < 0 {
			return true
		}
		if val1 < 0 && val2 < 0 && val1+val2 > 0 {
			return true
		}
	case "-":
		if val1 < 0 && val2 > 0 && val1-val2 > 0 {
			return true
		}
		if val1 > 0 && val2 < 0 && val1-val2 < 0 {
			return true // Переполнение
		}
	case "*":
		// Проверка переполнения при умножении
		if val1 > 0 && val2 > 0 && val1 > maxInt64/val2 {
			return true // Переполнение
		}
		if val1 < 0 && val2 < 0 && val1 < minInt64/val2 {
			return true // Переполнение
		}
		if val1 > 0 && val2 < 0 && val1 < minInt64/val2 {
			return true // Переполнение
		}
		if val1 < 0 && val2 > 0 && val1 > maxInt64/val2 {
			return true // Переполнение
		}
	case "/":
		// Проверка деления на ноль
		if val2 == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return true // Деление на ноль
		}
	case "%":
		// Проверка деления на ноль для остатка
		if val2 == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return true // Деление на ноль
		}
	}
	return false
}

// Функция для вывода числа в стандартный вывод
func printInt(n int64) {
	// Для работы с отрицательными числами
	if n < 0 {
		os.Stdout.Write([]byte("-"))
		n = -n
	}

	// Преобразование числа в строку
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte(n%10) + '0'}, digits...)
		n /= 10
	}

	// Если число равно нулю, выводим 0
	if len(digits) == 0 {
		digits = append(digits, '0')
	}

	// Выводим результат
	os.Stdout.Write(digits)
}

func main() {
	// Проверка количества аргументов
	if len(os.Args) != 4 {
		return // Если аргументов не 4, ничего не делаем
	}

	// Чтение первого и третьего аргументов как строки
	val1, valid1 := stringToInt(os.Args[1])
	val2, valid2 := stringToInt(os.Args[3])

	// Если не удалось преобразовать хотя бы одно из чисел в int64, ничего не выводим
	if !valid1 || !valid2 {
		return
	}

	// Оператор из второго аргумента
	operator := os.Args[2]

	// Проверка переполнения
	if checkOverflow(val1, val2, operator) {
		// Если переполнение или деление на ноль, не выводим ничего
		return
	}

	// Выполнение арифметической операции в зависимости от оператора
	var result int64
	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		// Если деление на ноль прошло проверку, делаем деление
		result = val1 / val2
	case "%":
		// Если деление на ноль прошло проверку, делаем остаток от деления
		result = val1 % val2
	default:
		// Если оператор не поддерживается, ничего не выводим
		return
	}

	// Выводим результат с новой строкой, если не произошло переполнения или ошибок
	printInt(result)
	os.Stdout.Write([]byte("\n"))
}
