package main

import (
	"testing"

	questions "github.com/cbailey6079/adventOfCodeGO/internal/questions"
)

func TestDay3_Part1_1(t *testing.T) {
	want := 2

	if got := questions.Day3(questions.Part1, ">", false); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part1_2(t *testing.T) {
	want := 4

	if got := questions.Day3(questions.Part1, "^>v<", false); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part1_3(t *testing.T) {
	want := 2

	if got := questions.Day3(questions.Part1, "^v^v^v^v^v", false); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part1(t *testing.T) {
	want := 2572

	if got := questions.Day3(questions.Part1, "./../inputs/day3.txt", true); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part2_1(t *testing.T) {
	want := 3

	if got := questions.Day3(questions.Part2, "^v", false); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part2_2(t *testing.T) {
	want := 3

	if got := questions.Day3(questions.Part2, "^>v<", false); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part2_3(t *testing.T) {
	want := 11

	if got := questions.Day3(questions.Part2, "^v^v^v^v^v", false); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}

func TestDay3_Part2(t *testing.T) {
	want := 2631

	if got := questions.Day3(questions.Part2, "./../inputs/day3.txt", true); got != want {
		t.Errorf("Day3() = %d, want %d", got, want)
	}
}
