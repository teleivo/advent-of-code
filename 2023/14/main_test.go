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

func TestTiltNorth(t *testing.T) {
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

func TestTiltSouth(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in: `O.OO#O...#
O...O#....`,
			want: `O...#O...#
O.OOO#....
`,
		},
		{
			in: `O...O#....
O.O.#....#
....O.....
..O.......`,
			want: `....O#....
....#....#
O.O.......
O.O.O.....
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
			want: `.....#....
....#....#
...O.##...
...#......
O.O....O#O
O.#..O.#.#
O....#....
OO....OO..
#OO..###..
#OO.O#...O
`,
		},
	}

	for _, tc := range tests {
		in := make([]byte, len(tc.in))
		copy(in, tc.in)
		inB := bytes.Fields([]byte(in))
		tiltSouth(inB)
		var got strings.Builder
		for _, row := range inB {
			got.WriteString(string(row))
			got.WriteRune('\n')
		}

		assertDeepEquals(t, "tiltSouth", tc.in, got.String(), tc.want)
	}
}

func TestTiltWest(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   `O....#....`,
			want: `O....#....`,
		},
		{
			in:   `O.OO#....#`,
			want: `OOO.#....#`,
		},
		{
			in:   `.....##...`,
			want: `.....##...`,
		},
		{
			in:   `OO.#O....O`,
			want: `OO.#OO....`,
		},
		{
			in:   `.O.....O#.`,
			want: `OO......#.`,
		},
		{
			in:   `O.#..O.#.#`,
			want: `O.#O...#.#`,
		},
	}

	for _, tc := range tests {
		in := make([]byte, len(tc.in))
		copy(in, tc.in)
		inB := bytes.Fields([]byte(in))
		tiltWest(inB)
		var got strings.Builder
		for i, row := range inB {
			got.WriteString(string(row))
			if i < len(inB)-1 {
				got.WriteRune('\n')
			}
		}

		assertDeepEquals(t, "tiltWest", tc.in, got.String(), tc.want)
	}
}

func TestTiltEast(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   `O....#....`,
			want: `....O#....`,
		},
		{
			in:   `O.OO#....#`,
			want: `.OOO#....#`,
		},
		{
			in:   `.....##...`,
			want: `.....##...`,
		},
		{
			in:   `OO.#O....O`,
			want: `.OO#....OO`,
		},
		{
			in:   `.O.....O#.`,
			want: `......OO#.`,
		},
		{
			in:   `O.#..O.#.#`,
			want: `.O#...O#.#`,
		},
	}

	for _, tc := range tests {
		in := make([]byte, len(tc.in))
		copy(in, tc.in)
		inB := bytes.Fields([]byte(in))
		tiltEast(inB)
		var got strings.Builder
		for i, row := range inB {
			got.WriteString(string(row))
			if i < len(inB)-1 {
				got.WriteRune('\n')
			}
		}

		assertDeepEquals(t, "tiltEast", tc.in, got.String(), tc.want)
	}
}

func TestSolvePartTwo(t *testing.T) {
	t.Skip()
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

func TestCycle(t *testing.T) {
	tests := []struct {
		in    string
		times int
		want  string
	}{
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
			times: 1,
			want: `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`,
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
			times: 2,
			want: `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`,
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
			times: 3,
			want: `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`,
		},
	}

	for _, tc := range tests {
		in := make([]byte, len(tc.in))
		copy(in, tc.in)
		inB := bytes.Fields([]byte(in))
		for i := 0; i < tc.times; i++ {
			cycle(inB)
		}
		var got strings.Builder
		for i, row := range inB {
			got.WriteString(string(row))
			if i < len(inB)-1 {
				got.WriteRune('\n')
			}
		}

		assertDeepEquals(t, "cycle", tc.in, got.String(), tc.want)
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
