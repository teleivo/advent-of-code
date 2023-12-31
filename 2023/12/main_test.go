package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	file := "testdata/example"
	want := 21
	f, err := os.Open(file)
	assertNoError(t, err)
	defer f.Close()

	got, err := solvePartOne(f)

	assertNoError(t, err)
	assertEquals(t, "solvePartOne", file, want, got)
}

func TestFindArrangements(t *testing.T) {
	tests := []struct {
		in     []byte
		groups []int
		want   int
	}{
		// {
		// 	in:     nil,
		// 	groups: []int{1, 1, 3},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("#"),
		// 	groups: []int{1},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("#."),
		// 	groups: []int{1},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("#.#."),
		// 	groups: []int{1, 1},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("???.###"),
		// 	groups: []int{1, 1, 3},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("???"),
		// 	groups: []int{1},
		// 	want:   3,
		// },
		// {
		// 	in:     []byte("???"),
		// 	groups: []int{1, 1},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("????"),
		// 	groups: []int{2, 1},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("?????"),
		// 	groups: []int{2, 1},
		// 	want:   3,
		// },
		// {
		// 	in:     []byte("??????"),
		// 	groups: []int{2, 1},
		// 	want:   6,
		// },
		// {
		// 	in:     []byte("???????"),
		// 	groups: []int{2, 1},
		// 	want:   10,
		// },
		// {
		// 	in:     []byte(".??..??."),
		// 	groups: []int{1, 1},
		// 	want:   4,
		// },
		// {
		// 	in:     []byte(".??..??...?##."),
		// 	groups: []int{1, 1, 3},
		// 	want:   4,
		// },
		// {
		// 	in:     []byte("#.#.###"),
		// 	groups: []int{1, 1, 3},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte(".??..??...?##."),
		// 	groups: []int{1, 1, 3},
		// 	want:   4,
		// },
		// {
		// 	in:     []byte("?#?#?#?#?#?#?#?"),
		// 	groups: []int{1, 3, 1, 6},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("????.#...#..."),
		// 	groups: []int{4, 1, 1},
		// 	want:   1,
		// },
		// {
		// 	in:     []byte("????.######..#####."),
		// 	groups: []int{1, 6, 5},
		// 	want:   4,
		// },
		// {
		// 	in:     []byte("?###????????"),
		// 	groups: []int{3, 2, 1},
		// 	want:   10,
		// },
		{
			in:     []byte("????#??##.?????"),
			groups: []int{1, 5, 2},
			want:   12,
		},
	}

	for _, tc := range tests {
		got := findArrangements(tc.in, tc.groups)
		assertEquals(t, "findArrangements", tc.in, got, tc.want)
	}
}

func TestSplitGroups(t *testing.T) {
	tests := []struct {
		in       []byte
		groups   []int
		wantHead []int
		wantTail []int
	}{
		{
			in:       []byte("??"),
			groups:   []int{1, 1},
			wantHead: []int{1},
			wantTail: []int{1},
		},
		{
			in:       []byte("????"),
			groups:   []int{2, 1},
			wantHead: []int{2, 1},
			wantTail: []int{},
		},
		{
			in:       []byte("??????"),
			groups:   []int{2, 1},
			wantHead: []int{2, 1},
			wantTail: nil,
		},
		{
			in:       []byte("????"),
			groups:   []int{4, 1, 1},
			wantHead: []int{4},
			wantTail: []int{1, 1},
		},
	}

	for _, tc := range tests {
		gotHead, gotTail := splitGroups(string(tc.in), tc.groups)
		assertDeepEquals(t, "splitGroups Head", tc.in, gotHead, tc.wantHead)
		assertDeepEquals(t, "splitGroups Tail", tc.in, gotTail, tc.wantTail)
	}
}

func TestMinWidth(t *testing.T) {
	tests := []struct {
		groups []int
		want   int
	}{
		{
			groups: []int{2},
			want:   2,
		},
		{
			groups: []int{2, 1},
			want:   4,
		},
	}

	for _, tc := range tests {
		got := minWidth(tc.groups)
		assertEquals(t, "minWidth", tc.groups, got, tc.want)
	}
}
func TestSolvePartTwo(t *testing.T) {
	t.Skip()
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
