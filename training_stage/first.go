package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)

		// Найти индекс цифры для удаления
		removeIndex := len(s) - 1
		for j := 0; j < len(s)-1; j++ {
			if s[j] < s[j+1] {
				removeIndex = j
				break
			}
		}

		// Удалить выбранную цифру
		newSalary := s[:removeIndex] + s[removeIndex+1:]

		// Удалить ведущие нули
		newSalary = strings.TrimLeft(newSalary, "0")
		// Если строка пуста, вернуть "0"
		if newSalary == "" {
			newSalary = "0"
		}

		fmt.Fprintln(out, newSalary)
	}
}
