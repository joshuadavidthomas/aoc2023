package main

import (
	"strconv"
	"unicode"
)

type Digit struct {
	digit string
	idx   int
}

func processLine(line string) int {
	var digits []Digit
	var firstDigit, lastDigit Digit

	for i, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, Digit{string(c), i})
		}
	}

	firstDigit = digits[0]
	if len(digits) == 1 {
		lastDigit = digits[0]
	} else {
		lastDigit = digits[len(digits)-1]
	}

	value, err := strconv.Atoi(firstDigit.digit + lastDigit.digit)
	if err != nil {
		panic(err)
	}

	return value
}

func Part1(lines []string) int {
	var sums []int
	var total int

	for _, line := range lines {
		sum := processLine(line)
		sums = append(sums, sum)
	}

	for i := 0; i < len(sums); i++ {
		total = total + sums[i]
	}

	return total
}
