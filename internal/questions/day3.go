package questions

import (
	"fmt"
	"os"
	"strings"
)

func Day3(part Part, input string, path bool) int {

	newInput := day3_setup(input, path)

	switch part {
	case Part1:
		return day3_part1(newInput)
	case Part2:
		return day3_part2(newInput)
	}

	return -1
}

func day3_part1(input string) int {
	houses := make(map[string]int)
	x := 0
	y := 0
	houses[fmt.Sprintf("%d,%d", x, y)] = 1

	for _, direction := range strings.Split(input, "") {
		switch direction {
		case "^":
			y++
		case ">":
			x++
		case "v":
			y--
		case "<":
			x--
		}

		houses[fmt.Sprintf("%d,%d", x, y)] = houses[fmt.Sprintf("%d,%d", x, y)] + 1
	}

	return len(houses)
}

func day3_part2(input string) int {
	houses := make(map[string]int)
	x, y, x2, y2 := 0, 0, 0, 0
	var tempx *int
	var tempy *int

	houses[fmt.Sprintf("%d,%d", x, y)] = 1

	for index, direction := range strings.Split(input, "") {
		if index%2 == 0 {
			tempx = &x
			tempy = &y
		} else {
			tempx = &x2
			tempy = &y2
		}

		switch direction {
		case "^":
			*tempy++
		case ">":
			*tempx++
		case "v":
			*tempy--
		case "<":
			*tempx--
		}

		houses[fmt.Sprintf("%d,%d", *tempx, *tempy)] = houses[fmt.Sprintf("%d,%d", *tempx, *tempy)] + 1
	}

	return len(houses)
}

func day3_setup(input string, path bool) string {
	if path {
		input = day3_parseinput(input)
	}

	return input
}

func day3_parseinput(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(data)
}
