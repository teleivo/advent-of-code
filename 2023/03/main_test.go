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
			in: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
		},
	}

	for _, tc := range tests {
		got, err := solvePartOne(strings.NewReader(tc.in))
		if err != nil {
			t.Fatalf("expected no error instead got %v", err)
		}

		if got != tc.want {
			t.Errorf("solvePartOne(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		in   string
		want *line
	}{
		{
			in: "467..114..",
			want: &line{
				Numbers: []number{
					{Value: 467, Start: 0, End: 2},
					{Value: 114, Start: 5, End: 7},
				},
				Symbols: map[int]struct{}{},
			},
		},
		{
			in: "617*......",
			want: &line{
				Numbers: []number{
					{Value: 617, Start: 0, End: 2},
				},
				Symbols: map[int]struct{}{
					3: {},
				},
			},
		},
		{
			in: ".....+.58.",
			want: &line{
				Numbers: []number{
					{Value: 58, Start: 7, End: 8},
				},
				Symbols: map[int]struct{}{
					5: {},
				},
			},
		},
		{
			in: "...$.*....",
			want: &line{
				Numbers: nil,
				Symbols: map[int]struct{}{
					3: {},
					5: {},
				},
			},
		},
	}

	for _, tc := range tests {
		got, err := parseLine(tc.in)
		if err != nil {
			t.Fatalf("expected no error instead got %v", err)
		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("parseLine() mismatch (-want +got):\n%s", diff)
		}
	}

}
