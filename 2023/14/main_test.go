package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	file := "testdata/example"
	want := 136
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("failed to read file %q: %v", file, err)
	}

	got, err := solvePartOne(b)

	assertNoError(t, err)
	assertEquals(t, "solvePartOne", file, got, want)
}

func TestTilt(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in: `O....#....
O.OO#....#`,
			want: `O.OO.#....
O...#....#
`,
		},
		{
			in: `O....#....
O.O.#....#
....O.....
..O.......`,
			want: `O.O..#....
O.O.#....#
....O.....
..........
`,
		},
		{
			in: `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`,
			want: `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
`,
		},
	}

	for _, tc := range tests {
		in := make([]byte, len(tc.in))
		copy(in, tc.in)
		inB := bytes.Fields([]byte(in))
		tiltNorth(inB)
		var got strings.Builder
		for _, row := range inB {
			got.WriteString(string(row))
			got.WriteRune('\n')
		}

		assertDeepEquals(t, "tiltNorth", tc.in, got.String(), tc.want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	file := "testdata/example"
	want := 64
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("failed to read file %q: %v", file, err)
	}

	got, err := solvePartTwo(b)

	assertNoError(t, err)
	assertEquals(t, "solvePartTwo", file, got, want)
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

func assertEquals(t *testing.T, method string, in, got, want any) {
	if got != want {
		t.Errorf("%s(%q) = %d; want %d", method, in, got, want)
	}
}

func assertDeepEquals(t *testing.T, method string, in, got, want any) {
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("%s(%q) mismatch (-want +got):\n%s", method, in, diff)
	}
}
