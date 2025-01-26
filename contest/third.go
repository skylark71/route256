package main

import (
	"fmt"
)

const (
	EMPTY  = '.'
	WALL   = '#'
	LIGHT  = 'L'
)

var directions = map[rune][3][2]int{
	'D': {{1, -1}, {1, 0}, {1, 1}}, // вниз
	'U': {{-1, -1}, {-1, 0}, {-1, 1}}, // вверх
	'L': {{-1, -1}, {0, -1}, {1, -1}}, // влево
	'R': {{-1, 1}, {0, 1}, {1, 1}}, // вправо
}

func illuminate(grid [][]rune, i, j int, dir rune) {
	// Освещаем начальную клетку
	if grid[i][j] == EMPTY {
		grid[i][j] = LIGHT
	}

	// Рекурсивно освещаем клетки по направлению
	for _, offset := range directions[dir] {
		ni, nj := i+offset[0], j+offset[1]
		// Проверяем границы и стены
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] != WALL {
			if grid[ni][nj] == EMPTY {
				grid[ni][nj] = LIGHT
				illuminate(grid, ni, nj, dir) // Рекурсивный вызов
			}
		}
	}
}

func placeLanterns(grid [][]rune) []string {
	var lanterns []string
	n, m := len(grid), len(grid[0])

	// Сначала устанавливаем фонарик в оптимальное направление
	// Если горизонтальный размер больше, ставим фонарик с направлением вправо (R)
	if m > n {
		illuminate(grid, 0, 0, 'R')
		lanterns = append(lanterns, "1 1 R")
	} else {
		illuminate(grid, 0, 0, 'D')
		lanterns = append(lanterns, "1 1 D")
	}

	// Проверяем правый нижний угол или другие важные клетки
	if grid[n-1][m-1] == EMPTY {
		illuminate(grid, n-1, m-1, 'U')
		lanterns = append(lanterns, fmt.Sprintf("%d %d U", n, m))
	}

	return lanterns
}

func main() {
	var t int
	fmt.Scan(&t)

	// Обрабатываем каждый набор данных
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		// Создаем сетку
		grid := make([][]rune, n)
		for i := 0; i < n; i++ {
			grid[i] = make([]rune, m)
			for j := 0; j < m; j++ {
				grid[i][j] = EMPTY // Изначально все клетки пустые
			}
		}

		// Размещаем фонарики и получаем их координаты и направления
		lanterns := placeLanterns(grid)

		// Выводим количество фонариков и их расположение
		fmt.Println(len(lanterns))
		for _, lantern := range lanterns {
			fmt.Println(lantern)
		}
	}
}
