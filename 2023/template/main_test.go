package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   ``,
			want: 30,
		},
	}

	for _, tc := range tests {
		got, err := solvePartTwo(strings.NewReader(tc.in))
		assertNoError(t, err)
		assertEquals(t, "solvePartOne", tc.in, tc.want, got)
	}
}

func TestSolvePartTwo(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   ``,
			want: 30,
		},
	}

	for _, tc := range tests {
		got, err := solvePartTwo(strings.NewReader(tc.in))
		assertNoError(t, err)
		assertEquals(t, "solvePartTwo", tc.in, tc.want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error instead got %v", err)
	}
}

func assertEquals(t *testing.T, method string, in, want, got any) {
	if got != want {
		t.Errorf("%s(%q) = %d; want %d", method, in, got, want)
	}
}

func assertDeepEquals(t *testing.T, method string, in, want, got any) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%s(%q) mismatch (-want +got):\n%s", in, method, diff)
	}
}
