package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

	b1, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read file %q: %v", file, err)
	}
	cal, err := solvePartOne(b1)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle one is: %d\n", cal)

	b2, err := os.ReadFile(file)
	cal2, err := solvePartTwo(b2)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle two is: %d\n", cal2)

	return nil
}

// solvePartOne solves part one of the puzzle.
func solvePartOne(in []byte) (int, error) {
	board := bytes.Fields(in)
	_ = board
	return 0, nil
}

// solvePartTwo solves part one of the puzzle.
func solvePartTwo(in []byte) (int, error) {
	board := bytes.Fields(in)
	_ = board
	return 0, nil
}
