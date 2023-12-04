package main

import (
	"testing"
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
