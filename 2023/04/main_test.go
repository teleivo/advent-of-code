package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestComputePoints(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: 8,
		},
		{
			in:   "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			want: 2,
		},
		{
			in:   "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			want: 2,
		},
		{
			in:   "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			want: 1,
		},
		{
			in:   "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			want: 0,
		},
		{
			in:   "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			want: 0,
		},
	}

	for _, tc := range tests {
		got := computePoints(tc.in)

		if got != tc.want {
			t.Errorf("computePoints(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}

func TestComputeMatches(t *testing.T) {
	tests := []struct {
		line string
		id   int
		want map[int]int
	}{
		{
			line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			id:   1,
			want: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
		},
		{
			line: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			id:   2,
			want: map[int]int{2: 1, 3: 1, 4: 1},
		},
		{
			line: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			id:   3,
			want: map[int]int{3: 1, 4: 1, 5: 1},
		},
		{
			line: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			id:   5,
			want: map[int]int{5: 1},
		},
	}

	for _, tc := range tests {
		got := make(map[int]int)
		computeMatches(tc.line, tc.id, got)

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("computMatches() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestSolvePartTwo(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			want: 30,
		},
	}

	for _, tc := range tests {
		got, err := solvePartTwo(strings.NewReader(tc.in))
		if err != nil {
			t.Fatalf("expected no error instead got %v", err)
		}

		if got != tc.want {
			t.Errorf("solvePartTwo(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}
