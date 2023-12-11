package main

import (
	"bufio"
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

	grid := parseAndExpandGrid(r)
	_ = grid

	// TODO turn galaxies into numbers and add test
	// row = append(row, rune(galaxyCount))

	return 0, nil
}

func parseAndExpandGrid(r io.Reader) [][]rune {
	s := bufio.NewScanner(r)
	var grid [][]rune
	var galaxyCount int
	var rows, cols int
	rowsWithoutGalaxy := make(map[int]struct{})
	colsWithGalaxy := make(map[int]struct{})
	for s.Scan() {
		line := s.Text()
		cols = len(line)
		var rowHasGalaxy bool
		var row []rune
		for col, r := range line {
			if r == '#' {
				galaxyCount++
				rowHasGalaxy = true
				colsWithGalaxy[col] = struct{}{}
				_ = col
			}
			row = append(row, r)
		}
		if !rowHasGalaxy {
			rowsWithoutGalaxy[rows] = struct{}{}
		}

		grid = append(grid, row)
		rows++
	}

	fmt.Println("rows", rows, "cols", cols)
	for row := range rowsWithoutGalaxy {
		fmt.Println("row without galaxy", row)
	}
	colsWithoutGalaxy := make(map[int]struct{})
	for i := 0; i < cols; i++ {
		if _, ok := colsWithGalaxy[i]; !ok {
			colsWithoutGalaxy[i] = struct{}{}
		}
	}
	for col := range colsWithoutGalaxy {
		fmt.Println("col without galaxy", col)
	}

	var expandedRows [][]rune
	for i := 0; i < rows; i++ {
		expandedRows = append(expandedRows, grid[i])
		if _, ok := rowsWithoutGalaxy[i]; ok {
			expandedRows = append(expandedRows, grid[i])
		}
	}

	var expandedGrid [][]rune
	for i := 0; i < len(expandedRows); i++ {
		expandedRow := expandedRows[i]
		var row []rune
		for j := 0; j < len(expandedRow); j++ {
			row = append(row, expandedRow[j])
			if _, ok := colsWithGalaxy[j]; !ok {
				row = append(row, expandedRow[j])
			}
		}
		expandedGrid = append(expandedGrid, row)
	}

	return expandedGrid
}

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		_ = line
	}
	return 0, nil
}
