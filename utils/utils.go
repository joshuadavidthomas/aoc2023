package utils

import (
	"bufio"
	"os"
)

func ReadInput(folder string) []string {
	var lines []string

	file, err := os.Open("./" + folder + "/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
