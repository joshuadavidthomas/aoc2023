package main

import (
	"fmt"

	"github.com/joshuadavidthomas/aoc2023-go/utils"
)

func main() {
	lines := utils.ReadInput("day01")
	part1testlines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	fmt.Println("Part 1 Test -- expecting 142:", Part1(part1testlines))
	fmt.Println("Part 1:", Part1(lines))
}
