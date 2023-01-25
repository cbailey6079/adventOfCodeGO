package main

import (
	"testing"

	questions "github.com/cbailey6079/adventOfCodeGO/internal/questions"
)

func TestDay1_Part1_1(t *testing.T) {
	want := 0

	if got := questions.Day1(questions.Part1, "(())", false); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay1_Part1_2(t *testing.T) {
	want := 3

	if got := questions.Day1(questions.Part1, "(()(()(", false); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay1_Part1_3(t *testing.T) {
	want := -1

	if got := questions.Day1(questions.Part1, "))(", false); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay1_Part1_4(t *testing.T) {
	want := -3

	if got := questions.Day1(questions.Part1, ")())())", false); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay1_Part1(t *testing.T) {
	want := 74

	if got := questions.Day1(questions.Part1, "./../inputs/day1.txt", true); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay1_Part2_1(t *testing.T) {
	want := 5

	if got := questions.Day1(questions.Part2, "()())", false); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay1_Part2(t *testing.T) {
	want := 1795

	if got := questions.Day1(questions.Part2, "./../inputs/day1.txt", true); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}
