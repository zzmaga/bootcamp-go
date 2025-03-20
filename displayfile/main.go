package main

import (
	"fmt"
	"os"
)

func main() {
	// Проверка количества аргументов
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	// Получаем имя файла из аргументов
	filename := os.Args[1]

	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Чтение файла с использованием небольшого буфера
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Error reading file: %v\n", err)
			}
			break
		}
		// Выводим считанные данные
		fmt.Print(string(buffer[:n]))
	}
}
