package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	file := "testdata/example"
	want := 6440
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("failed to open file %q: %v", file, err)
	}
	defer f.Close()

	got, err := solvePartOne(f)

	assertNoError(t, err)
	assertEquals(t, "solvePartOne", file, want, got)
}

func TestSolvePartTwo(t *testing.T) {
	file := "testdata/example"
	want := 5905
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("failed to open file %q: %v", file, err)
	}
	defer f.Close()

	got, err := solvePartTwo(f)

	assertNoError(t, err)
	assertEquals(t, "solvePartTwo", file, want, got)
}

func TestCompareHandsPartOne(t *testing.T) {
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
		got := compareHandsPartOne(tc.a, tc.b)

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
		got := categorizeHandPartOne(tc.in)
		assertDeepEquals(t, "categorizeHand", tc.in, tc.want, got)
	}
}
func TestCardFrequenciesPartOne(t *testing.T) {
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
		got := cardFrequenciesPartOne(tc.in)
		assertDeepEquals(t, "cardFrequenciesPartTwo", tc.in, tc.want, got)
	}
}

func TestCardFrequenciesPartTwo(t *testing.T) {
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
		{
			in: `T55J5`,
			want: map[rune]int{
				'5': 4,
				'T': 1,
			},
		},
		{
			in: `KTJJT`,
			want: map[rune]int{
				'K': 1,
				'T': 4,
			},
		},
		{
			in: `KKKJK`,
			want: map[rune]int{
				'K': 5,
			},
		},
		{
			in: `T45J3`,
			want: map[rune]int{
				'3': 1,
				'4': 1,
				'5': 1,
				'T': 2,
			},
		},
	}

	for _, tc := range tests {
		got := cardFrequenciesPartTwo(tc.in)
		assertDeepEquals(t, "cardFrequenciesPartTwo", tc.in, tc.want, got)
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
