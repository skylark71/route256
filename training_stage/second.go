package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func validateOutput(n int, array []int, output string) string {
	// Сортируем входной массив
	sort.Ints(array)

	// Преобразуем отсортированный массив в строку через Join
	expectedOutput := arrayToString(array)

	// Убираем переносы строк в выходных данных
	output = removeNewlines(output)

	if expectedOutput == output {
		return "yes"
	} else {
		return "no"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Читаем количество тестов
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading number of tests:", err)
		return
	}
	line = strings.TrimSpace(line)
	t, _ := strconv.Atoi(line)

	for i := 0; i < t; i++ {
		// Читаем размер массива
		line, _ := reader.ReadString('\n')
		if err != nil {
			fmt.Println("no")
			return
		}
		line = strings.TrimSpace(line)
		n, err := strconv.Atoi(line)
		if err != nil || n < 1 || n > 100000 {
			fmt.Println("no")
			return
		}

		// Читаем массив
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("no")
			return
		}
		line = strings.TrimSpace(line)
		arrayStr := strings.Fields(line)
		if len(arrayStr) != n {
			fmt.Println("no")
			return
		}

		array := make([]int, n)
		for j, s := range arrayStr {
			array[j], err = strconv.Atoi(s)
			if err != nil {
				fmt.Println("no")
				return
			}
		}

		// Читаем выходные данные
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("no")
			return
		}

		// Проверяем данные
		fmt.Println(validateOutput(n, array, line))
	}
}

func arrayToString(array []int) string {
	var strArray []string
	for _, num := range array {
		strArray = append(strArray, strconv.Itoa(num)) // Преобразуем число в строку
	}
	return strings.Join(strArray, " ") // Объединяем все элементы с пробелом между ними
}

func removeNewlines(input string) string {
	// Удаляем только символы переноса строки \r и \n
	return strings.ReplaceAll(strings.ReplaceAll(input, "\r", ""), "\n", "")
}
