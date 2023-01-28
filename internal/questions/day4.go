package questions

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func Day4(part Part, input string, path bool) int {

	input = day4_setup(input, path)

	switch part {
	case Part1:
		return day4_part1(input)
	case Part2:
		return day4_part2(input)
	}

	return -1
}

func day4_part1(input string) int {
	i := 1

	for {
		hasher := md5.New()
		hasher.Write([]byte(fmt.Sprintf("%s%d", input, i)))

		if hex.EncodeToString(hasher.Sum(nil))[:5] == "00000" {
			return i
		}

		i++
	}
}

func day4_part2(input string) int {
	i := 1

	for {
		hasher := md5.New()
		hasher.Write([]byte(fmt.Sprintf("%s%d", input, i)))

		if hex.EncodeToString(hasher.Sum(nil))[:6] == "000000" {
			return i
		}

		i++
	}
}

func day4_setup(input string, path bool) string {
	if path {
		input = parseinput(input)
	}

	return input
}

func day4_parseinput(filePath string) string {

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(data)
}
