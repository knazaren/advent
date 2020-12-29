package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}

// ReadInputAsInts returns the contents of filename as a slice of int64s.
func ReadInputAsInts(filename string) []int64 {
	input := strings.Fields(ReadInput(filename))
	inputInt := make([]int64, len(input))
	var err error
	for i, v := range input {
		inputInt[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return inputInt
}

//BoolToInt simple conversion routine
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

//ReadFileLines reads all the lines from the file and returns them as []string
func ReadFileLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
