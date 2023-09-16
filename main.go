package main

import (
	"fmt"
	"strings"
	"unicode"
)

func checkIfStringsFormSameSet(array []string) (bool, error) {
	n := len(array)
	// check if length if array is valid
	if n < 1 || n > 500 {
		return false, fmt.Errorf("length of array is not valid")
	}

	// for simplicity we will store columns in map, while validating rows
	var colums = make(map[int]*strings.Builder, n)

	for j := 0; j < n; j++ {
		colums[j] = &strings.Builder{}
	}

	// check if length of each string is valid
	for _, row := range array {
		// We don't need to check if length of string is less than 1 or more than 500
		// because we already checked length of array
		// so we can check only if length of each string is not equal to the length of an array
		if len(row) != n {
			return false, fmt.Errorf("array is not square")
		}
		// check if string contains only lowercase letters
		for j, char := range row {
			if !unicode.IsLower(char) {
				// ch is not a lowercase letter
				return false, fmt.Errorf("string contains non-lowercase letter")
			}
			colums[j].WriteRune(char)
		}
	}

	// iterate and check
	for i, row := range array {
		if row != colums[i].String() {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	strs := []string{"abcd", "bnrt", "crmy", "dtye"}
	ok, err := checkIfStringsFormSameSet(strs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok)
}
