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

	cal, err = solvePartTwo(file)
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
func solvePartTwo(file string) (int, error) {
	f1, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer f1.Close()
	s := bufio.NewScanner(f1)

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
	}

	f2, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer f2.Close()
	s = bufio.NewScanner(f2)

	patterns = nil
	pattern.Reset()
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
		sum += verticalMirrorsPartTwo(pattern)
	}

	return sum, nil
}

func horizontalMirrorsPartTwo(pattern [][]byte) int {
	locations := horizontalMirrorLocations(pattern)

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
func horizontalMirrorLocations(pattern [][]byte) []int {
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
	locations := verticalMirrorLocations(pattern)
	fmt.Println("mirror locations old", locations)

	var sum int
	reflectionLine := 1
	n := len(pattern[0])
	for reflectionLine < n {
		var smudgeRow, smudgeCol int
		var smudge int
		isMirror := true
		for i, j := reflectionLine-1, reflectionLine; i >= 0 && j < n; i, j = i-1, j+1 {
			for z, row := range pattern {
				if row[i] != row[j] {
					smudge++
					smudgeRow = z
					smudgeCol = i
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
			// adding the elements to the left of the reflection line
			sum += reflectionLine
			fmt.Println("sum", sum)
		}
		reflectionLine++
	}
	return sum
}

func verticalMirrorLocations(pattern [][]byte) []int {
	var locations []int
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
			locations = append(locations, reflectionLine)
		}
		reflectionLine++
	}
	return locations
}
