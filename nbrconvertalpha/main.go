package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Считываем все аргументы, кроме имени программы

	// Если аргументов нет, просто напечатаем пустую строку и завершим программу
	if len(args) == 0 {
		return
	}

	// Проверяем, включен ли флаг --upper
	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:] // Убираем флаг из списка аргументов
	}

	// Проходим по каждому аргументу
	for _, arg := range args {
		n := 0 // Для хранения числа
		valid := true

		// Преобразуем строку в число вручную
		for _, char := range arg {
			if char < '0' || char > '9' { // Если символ не является цифрой
				valid = false
				break
			}
			n = n*10 + int(char-'0') // Конвертируем символ в число
		}

		// Проверяем, допустимое ли число
		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ') // Выводим пробел для недопустимых значений
			continue
		}

		// Определяем букву
		var letter rune
		if upper {
			letter = 'A' + rune(n-1) // Верхний регистр
		} else {
			letter = 'a' + rune(n-1) // Нижний регистр
		}

		// Печатаем букву
		z01.PrintRune(letter)
	}

	// Переход на новую строку в любом случае
	z01.PrintRune('\n')
}
