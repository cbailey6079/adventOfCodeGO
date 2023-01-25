package questions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day2(part Part, input string, path bool) int {

	newInput := day2_setup(input, path)

	switch part {
	case Part1:
		return day2_part1(newInput)
	case Part2:
		return day2_part2(newInput)
	}

	return -1
}

func day2_part1(input []string) int {
	wrappingPaper := 0
	sides := make([]int, 3)

	for _, box := range input {
		sidesS := strings.Split(box, "x")

		for i := range sides {
			sides[i], _ = strconv.Atoi(sidesS[i])
		}

		area1 := sides[0] * sides[1]
		area2 := sides[1] * sides[2]
		area3 := sides[0] * sides[2]

		minArea := math.Min(math.Min(float64(area1), float64(area2)), float64(area3))

		wrappingPaper = wrappingPaper + area1*2 + area2*2 + area3*2 + int(minArea)
	}

	return wrappingPaper
}

func day2_part2(input []string) int {
	wrappingPaper := 0

	sides := make([]int, 3)
	for _, box := range input {
		sidesS := strings.Split(box, "x")

		for i := range sides {
			sides[i], _ = strconv.Atoi(sidesS[i])
		}

		sort.Slice(sides, func(i, j int) bool {
			return sides[i] < sides[j]
		})

		ribbon1 := sides[0]*2 + sides[1]*2
		ribbon2 := sides[0] * sides[1] * sides[2]

		wrappingPaper = wrappingPaper + ribbon1 + ribbon2
	}

	return wrappingPaper
}

func day2_setup(input string, path bool) []string {
	ret := make([]string, 0)
	if path {
		ret = day2_parseinput(input)
	} else {
		ret = append(ret, input)
	}

	return ret
}

func day2_parseinput(filePath string) []string {
	inputStrings := make([]string, 0)
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		inputStrings = append(inputStrings, fileScanner.Text())
	}

	return inputStrings
}
