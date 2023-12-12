package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Printf("failed to solve puzzle due to: %v\n", err)
		os.Exit(1)
	}
}

func run(w io.Writer, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected one argument pointing to puzzle input, instead got %d args", len(args)-1)
	}

	file := args[1]

	f1, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
	cal, err := solvePartOne(f1)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle one is: %d\n", cal)

	f2, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer f2.Close()
	cal, err = solvePartTwo(f2)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle two is: %d\n", cal)

	return nil
}

// solvePartOne solves part one of the puzzle.
func solvePartOne(b []byte) (int, error) {
	rows := bytes.Fields(b)
	for _, row := range rows {
		for _, v := range row {
			fmt.Println(string(v))
		}
	}
	return 0, nil
}

func findArrangements(springs []byte, groups []int) int {
	if len(springs) == 0 {
		return 1
	}
	if len(groups) == 1 {
		group := groups[0]
		// TODO at least in case springs are all ? there are arrangements: unknowns - group + 1
		if !strings.ContainsAny(string(springs), "#.") && strings.ContainsRune(string(springs), '?') {
			// assuming springs is all ?
			return len(springs) - group + 1
		}
		return 1
	}

	// var start, end int
	// find a block

	return 1
}

// solvePartTwo solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		_ = line
	}
	return 0, nil
}
