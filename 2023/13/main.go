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

	var sum int
	for i, pattern := range patterns {
		fmt.Println("pattern", i)
		for _, line := range pattern {
			fmt.Println(string(line))
		}
		sum += horizontalMirrors(pattern)
		sum += verticalMirrors(pattern)
	}

	return sum, nil
}

func horizontalMirrors(pattern [][]byte) int {
	var sum int
	reflectionLine := 1
	for reflectionLine < len(pattern) {
		isMirror := true
		for i, j := reflectionLine-1, reflectionLine; i >= 0 && j < len(pattern); i, j = i-1, j+1 {
			if string(pattern[i]) != string(pattern[j]) {
				isMirror = false
				break
			}
		}
		if isMirror {
			// adding the elements above the reflection line*100
			sum += reflectionLine * 100
		}
		reflectionLine++
	}
	return sum
}

func verticalMirrors(pattern [][]byte) int {
	var sum int
	reflectionLine := 1
	n := len(pattern[0])
	for reflectionLine < n {
		isMirror := true
		for i, j := reflectionLine-1, reflectionLine; i >= 0 && j < n; i, j = i-1, j+1 {
			for _, row := range pattern {
				if row[i] != row[j] {
					isMirror = false
					break
				}
			}
		}
		if isMirror {
			// adding the elements left oft the reflection line
			sum += reflectionLine
		}
		reflectionLine++
	}
	return sum
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
