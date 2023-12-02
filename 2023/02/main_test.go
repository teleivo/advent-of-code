package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			want: 8,
		},
	}

	for _, tc := range tests {
		got, err := solveFeasibleGames(strings.NewReader(tc.in), [3]int{12, 13, 14})
		if err!=nil {
			t.Fatalf("expected no error instead got %v", err)
		}

		if got != tc.want {
			t.Errorf("solve(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}

func TestParseLine(t *testing.T) {
	t.Skip()
	tests := []struct {
		in   string
		want game
	}{
		{
			in:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: game{
				ID: 1,
				cubes: [3]int{ 4, 3, 6 },
			},
		},
	}

	for _, tc := range tests {
		got, err := parseLine(tc.in)
		if err!=nil {
			t.Fatalf("expected no error instead got %v", err)
		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("parseLine() mismatch (-want +got):\n%s", diff)
		}
	}

}

func TestMaxCubes(t *testing.T) {
	tests := []struct {
		in   string
		want [3]int
	}{
		{
			in:   "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: [3]int{4, 2, 6},
		},
		{
			in:   " 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: [3]int{1, 3, 4},
		},
		{
			in:   "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: [3]int{20, 13, 6},
		},
		{
			in:   "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: [3]int{14, 3, 15},
		},
		{
			in:   "6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: [3]int{6, 3, 2},
		},
	}

	for _, tc := range tests {
		got := maxCubes(tc.in)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("maxCubes() mismatch (-want +got):\n%s", diff)
		}
	}
}
