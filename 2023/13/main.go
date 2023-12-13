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

	f1, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer f1.Close()
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
func solvePartOne(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	var patterns [][][]byte
	var pattern strings.Builder
	for s.Scan() {
		line := s.Text()
		if line == "" {
			patterns = append(patterns, bytes.Fields([]byte(pattern.String())))
			pattern.Reset()
			continue
		}
		pattern.WriteString(line)
		pattern.WriteRune('\n')
	}

	if pattern.String() != "" {
		patterns = append(patterns, bytes.Fields([]byte(pattern.String())))
	}

	for i, pattern := range patterns {
		fmt.Println("pattern", i)
		for _, line := range pattern {
			fmt.Println(string(line))
		}
	}

	return 0, nil
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
