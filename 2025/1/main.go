package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DIAL_START = 50
const DIAL_THRESHOLD = 100

func getDozens(value int) (int, int) {
	hundreds := value / 100
	dozens := (value % 100)

	return dozens, hundreds
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

func rotation(instruction string, currentPoint int) (int, int) {
	direction := string(instruction[0])
	rotation, err := strconv.Atoi(instruction[1:])
	var newRotationValue int
	turns := 0

	if err == nil {

		if rotation > DIAL_THRESHOLD {
			rotation, turns = getDozens(rotation)
		}

		switch direction {
		case "L":
			workDial := currentPoint - rotation

			if workDial < 0 {

				newRotationValue = DIAL_THRESHOLD - (workDial * -1)
				turns += 1

			} else {
				newRotationValue = workDial
			}
		case "R":
			workDial := currentPoint + rotation
			if workDial >= DIAL_THRESHOLD {

				newRotationValue = workDial - DIAL_THRESHOLD
				turns += 1

			} else {
				newRotationValue = workDial
			}
		default:
			fmt.Fprintln(os.Stdout, "Wrong direction: ", direction)

		}

	}

	return newRotationValue, turns

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
	extraZeros := 0

	if err == nil {

		for i := 0; i < len(workCases); i++ {
			var extraTurns int

			fmt.Println((workCases[i]))

			currentPoint, extraTurns = rotation(workCases[i], currentPoint)
			allRotationValues = append(allRotationValues, currentPoint)

			extraZeros += extraTurns

			fmt.Println("Dial is on", currentPoint, "over 0", extraZeros)

		}

	}

	fmt.Println("Part 1 result is ", part1(allRotationValues))
	fmt.Println("Part 2 result is ", (part1(allRotationValues)+extraZeros)-part1(allRotationValues))

}
