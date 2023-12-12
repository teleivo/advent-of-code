package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	t.Skip()
	file := "testdata/example"
	want := 21
	f, err := os.ReadFile(file)

	got, err := solvePartOne(f)

	assertNoError(t, err)
	assertEquals(t, "solvePartOne", file, want, got)
}

func TestFindArrangements(t *testing.T) {
	tests := []struct {
		in      []byte
		damaged []int
		want    int
	}{
		// {
		// 	in:      []byte("#.#.###"),
		// 	damaged: []int{1, 1, 3},
		// 	want:    1,
		// },
		{
			in:      []byte(".??..??...?##."),
			damaged: []int{1, 1, 3},
			want:    4,
		},
		// {
		// 	in:      []byte("?#?#?#?#?#?#?#?"),
		// 	damaged: []int{1, 3, 1, 6},
		// 	want:    1,
		// },
		// {
		// 	in:      []byte("????.#...#..."),
		// 	damaged: []int{4, 1, 1},
		// 	want:    1,
		// },
		// {
		// 	in:      []byte("????.######..#####."),
		// 	damaged: []int{1, 6, 5},
		// 	want:    4,
		// },
		// {
		// 	in:      []byte("?###????????"),
		// 	damaged: []int{3, 2, 1},
		// 	want:    10,
		// },
	}

	for _, tc := range tests {
		got := findArrangements(tc.in, tc.damaged)
		assertEquals(t, "findArrangements", tc.in, got, tc.want)
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
