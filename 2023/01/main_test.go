package main

import (
	"strings"
	"testing"
)

func TestDecodeCalibrationDocument(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
		{
			in: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
	}

	for _, tc := range tests {
		got := decodeCalibrationDocument(strings.NewReader(tc.in))
		if got != tc.want {
			t.Errorf("decodeCalibrationDocument(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}

func TestDecodeCalibration(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   "1abc2",
			want: 12,
		},
		{
			in:   "pqr3stu8vwx",
			want: 38,
		},
		{
			in:   "a1b2c3d4e5f",
			want: 15,
		},
		{
			in:   "treb7uchet",
			want: 77,
		},
		{
			in:   "two1nine",
			want: 29,
		},
		{
			in:   "eighthree",
			want: 83,
		},
		{
			in:   "eightthree",
			want: 83,
		},
		{
			in:   "eightwothree",
			want: 83,
		},
		{
			in:   "abcone2threexyz",
			want: 13,
		},
		{
			in:   "xtwone3four",
			want: 24,
		},
		{
			in:   "4nineeightseven2",
			want: 42,
		},
		{
			in:   "zoneight234",
			want: 14,
		},
		{
			in:   "7pqrstsixteen",
			want: 76,
		},
		{
			in:   "sevenine",
			want: 79,
		},
	}

	for _, tc := range tests {
		got := decodeCalibration(tc.in)
		if got != tc.want {
			t.Errorf("decodeCalibration(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}
