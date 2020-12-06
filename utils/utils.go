package utils

import (
	"bytes"
	"io/ioutil"
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
