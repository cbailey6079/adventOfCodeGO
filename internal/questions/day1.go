package questions

import (
	"os"
	"strings"
)

type Part int64

const (
	Part1 Part = 0
	Part2      = 1
)

func Day1(part Part, input string, path bool) int {

	input = setup(input, path)

	switch part {
	case Part1:
		return part1(input)
	case Part2:
		return part2(input)
	}

	return -1
}

func part1(input string) int {
	endingMap := make(map[string]int)

	for _, character := range strings.Split(input, "") {
		endingMap[character]++
	}

	return endingMap["("] - endingMap[")"]
}

func part2(input string) int {
	endingMap := make(map[string]int)

	for index, character := range strings.Split(input, "") {
		endingMap[character]++

		if endingMap["("]-endingMap[")"] < 0 {
			return index + 1
		}
	}

	return -1
}

func setup(input string, path bool) string {
	if path {
		input = parseinput(input)
	}

	return input
}

func parseinput(filePath string) string {

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(data)
}
