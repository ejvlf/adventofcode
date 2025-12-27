package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirs = [][2]int{
	{-1, 0}, {-1, 1}, {0, 1}, {1, 1},
	{1, 0}, {1, -1}, {0, -1}, {-1, -1},
}

func countRolls(diagram [][]rune, x int, y int) int {
	count := 0
	for _, direction := range dirs {
		dx := direction[0]
		dy := direction[1]

		newX := x + dx
		newY := y + dy

		if newY >= 0 && newY < len(diagram) && newX >= 0 && newX < len(diagram[0]) && diagram[newY][newX] == '@' {
			count++
		}
	}

	return count
}

func open_file(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func form_coord_map(mapAsText string, currentCoords [][]string) [][]string {

	totalCoords := strings.Split(mapAsText, "")

	currentCoords = append(currentCoords, totalCoords)

	return currentCoords
}

func calculate_part1(fullMap [][]string) int {
	result := 0

	rows := len(fullMap)
	if rows == 0 {
		return 0
	}
	cols := len(fullMap[0])

	for y, row := range fullMap {
		for x, cell := range row {
			if cell == "@" {
				total := 0

				for _, d := range dirs {
					ny, nx := y+d[0], x+d[1]

					if ny >= 0 && ny < rows && nx >= 0 && nx < cols {
						if fullMap[ny][nx] == "@" {
							total += 1
						}
					}

				}

				if total < 4 {
					result += 1
					continue
				}
			}
		}
	}
	return result

}

func calculate_part2(fullMap [][]string) int {
	diagram := [][]rune{}
	for _, row := range fullMap {
		runeRow := make([]rune, len(row))
		for i, cell := range row {
			if len(cell) > 0 {
				runeRow[i] = []rune(cell)[0]
			} else {
				runeRow[i] = ' '
			}
		}
		diagram = append(diagram, runeRow)
	}

	result := 0
	removed := -1

	for removed != 0 {
		removed = 0
		for y := 0; y < len(diagram); y++ {
			for x := 0; x < len(diagram[0]); x++ {
				if diagram[y][x] != '@' {
					continue
				}

				count := countRolls(diagram, x, y)
				if count < 4 {
					diagram[y][x] = '.'
					result++
					removed++
				}
			}
		}
	}

	return result
}
func main() {

	workCases, err := open_file("main.txt")
	var part1 int
	var part2 int

	if err == nil {
		var totalCoords [][]string

		for _, row := range workCases {

			totalCoords = form_coord_map(row, totalCoords)

		}
		part1 = calculate_part1(totalCoords)
		part2 = calculate_part2(totalCoords)

		fmt.Println("Part 1: ", part1)
		fmt.Println("Part 2: ", part2)

	}

}
