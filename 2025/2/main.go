package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func open_file(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = strings.Split(scanner.Text(), ",")
	}
	return lines, scanner.Err()
}

func parse_sequence(sequence string) (int, int) {

	startAndEnd := strings.Split(sequence, "-")

	start, _ := strconv.Atoi(startAndEnd[0])
	end, _ := strconv.Atoi(startAndEnd[1])

	return start, end

}

func main() {

	workCases, err := open_file("main.txt")
	part1Sum := 0
	part2Sum := 0

	if err == nil {
		for _, sequence := range workCases {

			start, end := parse_sequence(sequence)

			for i := start; i <= end; i++ {
				s := strconv.Itoa(i)
				if len(s) == 1 {
					continue
				}

				if len(s)%2 == 0 {
					right := s[len(s)/2:]
					left := s[:len(s)/2]
					if right == left {
						part1Sum += i
					}
				}
			}

		}
	}
	fmt.Println("Part 1", part1Sum)

	if err == nil {
		for _, sequence := range workCases {
			start, end := parse_sequence(sequence)

			for i := start; i <= end; i++ {
				s := strconv.Itoa(i)

				for size := 1; size < len(s); size++ {
					var relSize int
					elements := []string{}

					if len(s)%size != 0 {
						continue
					}
					for pos := 0; pos < len(s); pos += size {
						relSize += size
						if relSize > len(s) {
							relSize = len(s)
						}
						elements = append(elements, s[pos:relSize])
					}
					invalid := true

					for e := range elements {
						if elements[e] != elements[0] {
							invalid = false
							break
						}
					}
					if invalid {
						part2Sum += i
						break
					}
				}
			}

		}
	}
	fmt.Println("Part 2", part2Sum)
}
