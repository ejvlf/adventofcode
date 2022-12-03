package fileOperations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const puzzleFilesDir string = "puzzleInput"
const dayPrefix string = "day"
const testFilename string = "test.txt"
const mainFilename string = "input.txt"

func getPuzzleData(day int, testRun bool, retrievalType string) interface{} {

	if retrievalType == "number" {
		return getNumbersFromFile(day, testRun)
	}
	if retrievalType == "list" {
		return getRowsFromFile(day, testRun)
	}

	return nil
}

func formFileName(day int, testRun bool) string {

	var fileName string

	if testRun {
		fileName = testFilename
	} else {
		fileName = mainFilename
	}

	fname := "./puzzleInput/day" + strconv.FormatInt(int64(day), 10) + "/" + fileName

	return fname

}

func getRowsFromFile(day int, testRun bool) []string {

	var listOfStrings []string

	fileName := formFileName(day, testRun)

	// Read file in current directory
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	// Read files with \n
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Iterate through the rows
	for scanner.Scan() {
		listOfStrings = append(listOfStrings, scanner.Text())
	}

	return listOfStrings

}

func getNumbersFromFile(day int, testRun bool) []int {

	var listOfInts []int

	fileName := formFileName(day, testRun)

	// Read file in current directory
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	// Read files with \n
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Iterate through the rows
	for scanner.Scan() {
		if s, err := strconv.Atoi(scanner.Text()); err == nil {

			listOfInts = append(listOfInts, s)

		}

	}

	return listOfInts

}
