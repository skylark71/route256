package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Структура для представления директории
type Directory struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files"`
	Folders []Directory `json:"folders"`
}

// Функция для проверки зараженных файлов
func countInfectedFiles(dir Directory) int {
	infectedCount := 0
	// Проверяем файлы в текущей директории
	for _, file := range dir.Files {
		if strings.HasSuffix(file, ".hack") {
			infectedCount++
		}
	}
	// Проверяем вложенные директории рекурсивно
	for _, folder := range dir.Folders {
		infectedCount += countInfectedFiles(folder)
	}
	return infectedCount
}

func main() {
	var t int
	// Сканируем количество тестов
	fmt.Scan(&t)

	// Сканер для многострочного ввода
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < t; i++ {
		var n int
		// Сканируем количество строк с описаниями директорий
		fmt.Scan(&n)

		var directories []Directory

		// Читаем все директории и их содержимое
		for j := 0; j < n; j++ {
			var input string
			// Читаем многострочную строку JSON, поддерживаем пробелы и табуляции
			for scanner.Scan() {
				line := scanner.Text()
				input += line + "\n"
				// Если строка завершена и содержит все данные, выходим
				if strings.Count(input, "{") == strings.Count(input, "}") {
					break
				}
			}

			// Убираем лишние пробелы и табуляции
			input = strings.TrimSpace(input)

			// Парсим JSON строку в структуру
			var directory Directory
			if err := json.Unmarshal([]byte(input), &directory); err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}
			directories = append(directories, directory)
		}

		// Подсчитываем количество зараженных файлов для всех директорий
		infectedCount := 0
		for _, dir := range directories {
			infectedCount += countInfectedFiles(dir)
		}

		// Выводим результат
		fmt.Println(infectedCount)
	}
}
