package main

import (
	"strings"
	"testing"
)

func TestDecodeCalibrationDocument(t *testing.T) {
	in := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	want := 142

	got := decodeCalibrationDocument(strings.NewReader(in))
	if got != want {
		t.Errorf("decodeCalibrationDocument(%q) = %d; want %d", in, got, want)
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
	}

	for _, tc := range tests {
		got := decodeCalibration(tc.in)
		if got != tc.want {
			t.Errorf("decodeCalibration(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}
