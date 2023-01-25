package main

import (
	"testing"

	questions "github.com/cbailey6079/adventOfCodeGO/internal/questions"
)

func TestDay2_Part1_1(t *testing.T) {
	want := 58

	if got := questions.Day2(questions.Part1, "2x3x4", false); got != want {
		t.Errorf("Day1() = %d, want %d", got, want)
	}
}

func TestDay2_Part1_2(t *testing.T) {
	want := 43

	if got := questions.Day2(questions.Part1, "1x1x10", false); got != want {
		t.Errorf("Day2() = %d, want %d", got, want)
	}
}

func TestDay2_Part1(t *testing.T) {
	want := 1588178

	if got := questions.Day2(questions.Part1, "./../inputs/day2.txt", true); got != want {
		t.Errorf("Day2() = %d, want %d", got, want)
	}
}

func TestDay2_Part2_1(t *testing.T) {
	want := 34

	if got := questions.Day2(questions.Part2, "2x3x4", false); got != want {
		t.Errorf("Day2() = %d, want %d", got, want)
	}
}

func TestDay2_Part2_2(t *testing.T) {
	want := 14

	if got := questions.Day2(questions.Part2, "1x1x10", false); got != want {
		t.Errorf("Day2() = %d, want %d", got, want)
	}
}

func TestDay2_Part2(t *testing.T) {
	want := 3783758

	if got := questions.Day2(questions.Part2, "./../inputs/day2.txt", true); got != want {
		t.Errorf("Day2() = %d, want %d", got, want)
	}
}
