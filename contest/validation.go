package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(line))

	for i := 0; i < t; i++ {
		line, _ = reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(line))

		items := make(map[string]int)
		for j := 0; j < n; j++ {
			line, _ = reader.ReadString('\n')
			parts := strings.Fields(strings.TrimSpace(line))
			name := parts[0]
			price, _ := strconv.Atoi(parts[1])
			items[name] = price
		}

		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if validateOutput(items, line) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func validateOutput(items map[string]int, output string) bool {
	seenPrices := make(map[int]bool)
	parts := strings.Split(output, ",")

	for _, part := range parts {
		// Разделяем name:price
		itemParts := strings.Split(part, ":")
		if len(itemParts) != 2 {
			return false // Неверный формат строки
		}

		name := itemParts[0]
		priceStr := itemParts[1]

		if len(priceStr) > 1 && priceStr[0] == '0' {
			return false
		}

		price, err := strconv.Atoi(priceStr)
		if err != nil {
			return false
		}

		expectedPrice, exists := items[name]
		if !exists || expectedPrice != price {
			return false
		}

		if seenPrices[price] {
			return false
		}
		seenPrices[price] = true
	}

	for _, price := range items {
		if !seenPrices[price] {
			return false
		}
	}

	return true
}
