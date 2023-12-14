package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePartOne(t *testing.T) {
	file := "testdata/example"
	want := 136
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("failed to open file %q: %v", file, err)
	}
	defer f.Close()

	got, err := solvePartOne(f)

	assertNoError(t, err)
	assertEquals(t, "solvePartOne", file, got, want)
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
