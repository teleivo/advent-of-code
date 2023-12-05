package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
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

		if got != tc.want {
			t.Errorf("solvePartOne(%q) = %d; want %d", tc.in, got, tc.want)
		}
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		in   string
		want []int
	}{
		{
			in: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			want: []int{79, 14, 55, 13},
		},
	}

	for _, tc := range tests {
		_, err := parseInput(strings.NewReader(tc.in))
		assertNoError(t, err)

		// assertDeepEquals(t, "parseInput", tc.in, tc.want, got)
	}
}

func TestParseSeeds(t *testing.T) {
	tests := []struct {
		in   string
		want []int
	}{
		// TODO without the newline I get the io.EOF back in this test. how can I still use a
		// scanner while sharing the same (buffered) reader
		{
			in: `seeds: 79 14 55 13
			`,
			want: []int{79, 14, 55, 13},
		},
	}

	for _, tc := range tests {
		got, err := parseSeeds(bufio.NewReader(strings.NewReader(tc.in)))
		assertNoError(t, err)

		assertDeepEquals(t, "parseSeeds", tc.in, tc.want, got)
	}
}

func TestParseNumbers(t *testing.T) {
	tests := []struct {
		in   string
		want []int
	}{
		{
			in:   `79 14 55 13`,
			want: []int{79, 14, 55, 13},
		},
	}

	for _, tc := range tests {
		got, err := parseNumbers(strings.NewReader(tc.in))
		assertNoError(t, err)

		assertDeepEquals(t, "parseNumbers", tc.in, tc.want, got)
	}

	errTests := []struct {
		in   string
		want []int
	}{
		{
			in: `79 a14 55 13`,
		},
	}

	for _, tc := range errTests {
		_, err := parseNumbers(strings.NewReader(tc.in))
		assertError(t, err)
	}
}

func TestParseMap(t *testing.T) {
	tests := []struct {
		in   string
		want [][][3]int
	}{
		{
			in: `
humidity-to-location map:
1136439539 28187015 34421000
4130684560 3591141854 62928737
`,
			want: [][][3]int{
				{
					{1136439539,
						28187015,
						34421000,
					},
					{
						4130684560,
						3591141854,
						62928737,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		got, err := parseMaps(bufio.NewReader(strings.NewReader(tc.in)))
		assertNoError(t, err)

		assertDeepEquals(t, "parseMap", tc.in, tc.want, got)
	}
}

func TestInsert(t *testing.T) {

	root := &node{Source: 98, Dest: 50, RangeLen: 2}
	want := &node{Source: 98, Dest: 50, RangeLen: 2, Left: &node{Source: 50, Dest: 52, RangeLen: 48}}

	insert(root, 50, 52, 48)

	if diff := cmp.Diff(want, root); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		mapNode *node
		source  int
		want    int
	}{
		{
			mapNode: &node{
				Source:   98,
				Dest:     50,
				RangeLen: 2,
				Left: &node{
					Source:   50,
					Dest:     52,
					RangeLen: 48,
				},
			},
			source: 79,
			want:   81,
		},
		{
			mapNode: &node{
				Source:   98,
				Dest:     50,
				RangeLen: 2,
				Left: &node{
					Source:   50,
					Dest:     52,
					RangeLen: 48,
				},
			},
			source: 14,
			want:   14,
		},
		{
			mapNode: &node{
				Source:   98,
				Dest:     50,
				RangeLen: 2,
				Left: &node{
					Source:   50,
					Dest:     52,
					RangeLen: 48,
				},
			},
			source: 55,
			want:   57,
		},
	}

	for _, tc := range tests {
		got := find(tc.mapNode, tc.source)
		assertEquals(t, "find", tc.mapNode, tc.want, got)
	}
}

//	func TestSolvePartTwo(t *testing.T) {
//		tests := []struct {
//			in   string
//			want int
//		}{
//			{
//				in:   ``,
//				want: 30,
//			},
//		}
//
//		for _, tc := range tests {
//			got, err := solvePartTwo(strings.NewReader(tc.in))
//			assertNoError(t, err)
//
//			if got != tc.want {
//				t.Errorf("solvePartTwo(%q) = %d; want %d", tc.in, got, tc.want)
//			}
//		}
//	}

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
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%s(%q) mismatch (-want +got):\n%s", in, method, diff)
	}
}
