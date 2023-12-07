package main

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	file := "testdata/example"
	want := 6440
	f1, err := os.Open(file)
	if err != nil {
		t.Fatalf("failed to open file %q: %v", file, err)
	}
	defer f1.Close()

	got, err := solvePartOne(f1)

	assertNoError(t, err)
	assertEquals(t, "solvePartOne", file, want, got)
}

func TestSolvePartTwo(t *testing.T) {
	t.Skip()
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

func TestCompareHands(t *testing.T) {
	tests := []struct {
		a    hand
		b    hand
		want int
	}{
		{
			a:    hand{Hand: `32T3K`},
			b:    hand{Hand: `32T3K`},
			want: 0,
		},
		{
			a:    hand{Hand: `2AAAA`},
			b:    hand{Hand: `33332`},
			want: -1,
		},
	}

	for _, tc := range tests {
		got := compareHands(tc.a, tc.b)

		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("%s(%q, %q) mismatch (-want +got):\n%s", "CompareHands", tc.a, tc.b, diff)
		}
	}
}

func TestCategorizeHand(t *testing.T) {
	tests := []struct {
		in   string
		want handType
	}{
		{
			in:   `32T3K`,
			want: one,
		},
		{
			in:   `KK677`,
			want: two,
		},
		{
			in:   `KTJJT`,
			want: two,
		},
		{
			in:   `T55J5`,
			want: three,
		},
		{
			in:   `QQQJA`,
			want: three,
		},
		{
			in:   `AAAAA`,
			want: five,
		},
		{
			in:   `AA8AA`,
			want: four,
		},
		{
			in:   `23332`,
			want: fullHouse,
		},
		{
			in:   `23456`,
			want: high,
		},
	}

	for _, tc := range tests {
		got := categorizeHand(tc.in)
		assertDeepEquals(t, "categorizeHand", tc.in, tc.want, got)
	}
}
func TestCardFrequencies(t *testing.T) {
	tests := []struct {
		in   string
		want map[rune]int
	}{
		{
			in: `32T3K`,
			want: map[rune]int{
				'3': 2,
				'2': 1,
				'T': 1,
				'K': 1,
			},
		},
	}

	for _, tc := range tests {
		got := cardFrequencies(tc.in)
		assertDeepEquals(t, "cardFrequencies", tc.in, tc.want, got)
	}
}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Fatal("expected error instead got nil instead", err)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error instead got: %q", err)
	}
}

func assertEquals(t *testing.T, method string, in, want, got any) {
	if got != want {
		t.Errorf("%s(%q) = %d; want %d", method, in, got, want)
	}
}

func assertDeepEquals(t *testing.T, method string, in, want, got any) {
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("%s(%q) mismatch (-want +got):\n%s", method, in, diff)
	}
}
