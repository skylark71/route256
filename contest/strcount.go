package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	results := make([]int, t)

	// Обрабатываем каждый тест
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		stringsArray := make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&stringsArray[j])
		}

		oddMap := make(map[string]int)
		evenMap := make(map[string]int)

		// Подсчёт похожих строк
		count1 := 0
		count2 := 0
		for _, s := range stringsArray {
			odd := ""
			even := ""
			for idx, ch := range s {
				if (idx+1)%2 == 0 { // Чётная позиция
					even += string(ch)
				} else { // Нечётная позиция
					odd += string(ch)
				}
			}

			// Проверяем, если строка не пуста
			if odd != "" {
				count1 += oddMap[odd] // Похожие по нечётным
				oddMap[odd]++
			}
			if even != "" {
				count2 += evenMap[even] // Похожие по чётным
				evenMap[even]++
			}
		}

		if count1 == 0 && count2 == 0 {
			results[i] = 0
		} else {
			if count1 == count2 {
				results[i] = count1
			} else {
				if count1 > count2 {
					results[i] = count1
				} else {
					results[i] = count2
				}
			}
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
