package main

import (
	"testing"

	questions "github.com/cbailey6079/adventOfCodeGO/internal/questions"
)

func TestDay4_Part1_1(t *testing.T) {
	want := 609043

	if got := questions.Day4(questions.Part1, "abcdef", false); got != want {
		t.Errorf("Day4() = %d, want %d", got, want)
	}
}

func TestDay4_Part1_2(t *testing.T) {
	want := 1048970

	if got := questions.Day4(questions.Part1, "pqrstuv", false); got != want {
		t.Errorf("Day4() = %d, want %d", got, want)
	}
}

func TestDay4_Part1(t *testing.T) {
	want := 282749

	if got := questions.Day4(questions.Part1, "./../inputs/day4.txt", true); got != want {
		t.Errorf("Day4() = %d, want %d", got, want)
	}
}

func TestDay4_Part2(t *testing.T) {
	want := 9962624

	if got := questions.Day4(questions.Part2, "./../inputs/day4.txt", true); got != want {
		t.Errorf("Day4() = %d, want %d", got, want)
	}
}
