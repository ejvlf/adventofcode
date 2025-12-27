package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const TOTAL_NUMBERS = 2

func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
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

func parse_as_list_of_ints(line string) []int {
	var list_of_ints []int

	for _, number := range strings.Split(line, "") {

		parsed_number, _ := strconv.Atoi(number)
		list_of_ints = append(list_of_ints, parsed_number)

	}

	return list_of_ints
}

func get_max_value(list_of_numbers []int, start_idx int) (int, int) {
	max_number := 0
	max_idx := 0

	for i := start_idx; i < len(list_of_numbers); i++ {
		if list_of_numbers[i] > max_number {
			max_number = list_of_numbers[i]
			max_idx = i
		}

	}

	return max_idx + 1, max_number
}

func form_from_total_of_bank(maxBank []int) int {

	result := 0
	for _, v := range maxBank {
		result = result*10 + v
	}

	return result

}

func getJoltage(line string, totalBanks int) int {
	result := int(0)
	start := 0

	for i := totalBanks; i > 0; i-- {
		max := byte('0')
		end := len(line) - i

		for j := start; j <= end; j++ {
			if line[j] > max {
				max = line[j]
				start = j + 1
			}
		}

		result += int(max-'0') * int(math.Pow(10.0, float64(i-1)))
	}

	return result
}

func main() {

	workCases, err := open_file("main.txt")
	var part1 int
	var part2 int

	if err == nil {

		for _, bank := range workCases {

			part1 += getJoltage(bank, 2)

		}

		fmt.Println("Part 1: ", part1)

		for _, bank := range workCases {

			part2 += getJoltage(bank, 12)

		}

		fmt.Println("Part 2: ", part2)

	}

}
