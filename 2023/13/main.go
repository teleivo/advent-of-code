package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
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
		sum += horizontalMirrorsPartTwo(pattern)
		sum += verticalMirrorsPartTwo(pattern)
	}

	return sum, nil
}

func horizontalMirrorsPartTwo(pattern [][]byte) int {
	locations := horizontalMirrorsLocation(pattern)

	var sum int
	reflectionLine := 1
	for reflectionLine < len(pattern) {
		var smudgeRow, smudgeCol int
		var smudge int
		isMirror := true
		for i, j := reflectionLine-1, reflectionLine; i >= 0 && j < len(pattern); i, j = i-1, j+1 {
			for z := 0; z < len(pattern[i]); z++ {
				if pattern[i][z] != pattern[j][z] {
					smudge++
					smudgeRow = i
					smudgeCol = z
				}
				// only one smudge can be fixed
				if smudge > 1 {
					isMirror = false
					break
				}
			}
		}
		if isMirror {
			if smudge != 0 {
				fmt.Println("reflectionLine", reflectionLine, "via smudge in", smudgeRow, smudgeCol, "diff", smudge)
				old := pattern[smudgeRow][smudgeCol]
				if old == '.' {
					pattern[smudgeRow][smudgeCol] = '#'
				} else {
					pattern[smudgeRow][smudgeCol] = '.'
				}
			} else {
				fmt.Println("reflectionLine without smudge", reflectionLine)
				if slices.Contains(locations, reflectionLine) {
					fmt.Println("skipping old line")
					reflectionLine++
					continue
				}
			}
			// adding the elements above the reflection line*100
			sum += reflectionLine * 100
			fmt.Println("sum", sum)
		}
		reflectionLine++
	}
	return sum
}

// finds old mirror locations so I can exclude them in part 2 smudges
func horizontalMirrorsLocation(pattern [][]byte) []int {
	var locations []int
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
			locations = append(locations, reflectionLine)
		}
		reflectionLine++
	}
	return locations
}

func verticalMirrorsPartTwo(pattern [][]byte) int {
	var sum int
	reflectionLine := 1
	n := len(pattern[0])
	for reflectionLine < n {
		var diff int
		isMirror := true
		for i, j := reflectionLine-1, reflectionLine; i >= 0 && j < n; i, j = i-1, j+1 {
			for _, row := range pattern {
				if row[i] != row[j] {
					diff++
				}
				// only one smudge can be fixed
				if diff > 1 {
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
