package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DIAL_START = 50
const DIAL_THRESHOLD = 100

func getDozens(value int) int {
	dozens := (value % 100)

	return dozens
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

func rotation(instruction string, currentPoint int) int {
	direction := string(instruction[0])
	rotation, err := strconv.Atoi(instruction[1:])
	var newRotationValue int

	if err == nil {

		if rotation > DIAL_THRESHOLD {
			rotation = getDozens(rotation)
		}

		switch direction {
		case "L":
			workDial := currentPoint - rotation

			if workDial < 0 {

				newRotationValue = DIAL_THRESHOLD - (workDial * -1)

			} else {
				newRotationValue = workDial
			}
		case "R":
			workDial := currentPoint + rotation
			if workDial >= DIAL_THRESHOLD {

				newRotationValue = workDial - DIAL_THRESHOLD

			} else {
				newRotationValue = workDial
			}
		default:
			fmt.Fprintln(os.Stdout, "Wrong direction %v", direction)

		}

	}

	return newRotationValue

}

func part1(listOfValues []int) int {
	counter := 0

	for _, value := range listOfValues {

		if value == 0 {
			counter += 1
		}

	}

	return counter
}

func main() {
	workCases, err := open_file("part1.txt")
	currentPoint := DIAL_START
	var allRotationValues []int

	if err == nil {

		for i := 0; i < len(workCases); i++ {

			fmt.Println((workCases[i]))

			currentPoint = rotation(workCases[i], currentPoint)
			allRotationValues = append(allRotationValues, currentPoint)

			fmt.Println("Dial is on ", currentPoint)

		}

	}

	fmt.Println("Part 1 result is ", part1(allRotationValues))

}
