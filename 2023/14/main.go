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
	cal, err = solvePartTwo(b2)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle two is: %d\n", cal)

	return nil
}

// solvePartOne solves part one of the puzzle.
func solvePartOne(in []byte) (int, error) {
	pattern := bytes.Fields(in)
	for _, row := range pattern {
		fmt.Println(string(row))
	}
	tiltNorth(pattern)
	var sum int
	rows := len(pattern)
	for row, line := range pattern {
		for _, r := range line {
			if r == 'O' {
				sum += rows - row
			}
		}
	}
	return sum, nil
}

type cell struct {
	row int
	col int
}

func tiltNorth(in [][]byte) {
	freeCells := make(map[int][]int)
	for row, line := range in {
		for col, r := range line {
			// keep track of the free cell that is furthest up north
			if r == '.' {
				if minRows, ok := freeCells[col]; ok {
					minRows = append(minRows, row)
					freeCells[col] = minRows
				} else {
					freeCells[col] = []int{row}
				}
			} else if r == '#' {
				delete(freeCells, col)
			} else if r == 'O' {
				// move rock to the cell furthest up north
				if minRows, ok := freeCells[col]; ok {
					minRow := minRows[0]
					in[minRow][col] = 'O'
					in[row][col] = '.'
					minRows = append(minRows, row)
					// remove top as its taken now by the O
					minRows = minRows[1:]
					freeCells[col] = minRows
				}
			}
		}
	}
}

// solvePartTwo solves part one of the puzzle.
func solvePartTwo(in []byte) (int, error) {
	pattern := bytes.Fields(in)
	for i := 0; i < 1_000_000_000; i++ {
		fmt.Println("round", i, "to go", 1_000_000_000-i)
		cycle(pattern)
	}
	var sum int
	rows := len(pattern)
	for row, line := range pattern {
		for _, r := range line {
			if r == 'O' {
				sum += rows - row
			}
		}
	}
	return sum, nil
}

func cycle(in [][]byte) {
	tiltNorth(in)
	tiltWest(in)
	tiltSouth(in)
	tiltEast(in)
}

func tiltSouth(in [][]byte) {
	freeCells := make(map[int][]int)
	for row := len(in) - 1; row >= 0; row-- {
		line := in[row]
		for col, r := range line {
			// keep track of the free cell that is furthest up north
			if r == '.' {
				if minRows, ok := freeCells[col]; ok {
					minRows = append(minRows, row)
					freeCells[col] = minRows
				} else {
					freeCells[col] = []int{row}
				}
			} else if r == '#' {
				delete(freeCells, col)
			} else if r == 'O' {
				// move rock to the cell furthest up north
				if minRows, ok := freeCells[col]; ok {
					minRow := minRows[0]
					in[minRow][col] = 'O'
					in[row][col] = '.'
					minRows = append(minRows, row)
					// remove top as its taken now by the O
					minRows = minRows[1:]
					freeCells[col] = minRows
				}
			}
		}
	}
}

func tiltWest(in [][]byte) {
	for row, line := range in {
		var minCols []int
		for col := 0; col < len(line); col++ {
			r := line[col]
			// keep track of the free cell that is furthest up north
			if r == '.' {
				minCols = append(minCols, col)
			} else if r == '#' {
				minCols = nil
			} else if r == 'O' {
				// move rock to the cell furthest up north
				if len(minCols) > 0 {
					minCol := minCols[0]
					in[row][minCol] = 'O'
					in[row][col] = '.'
					minCols = append(minCols, col)
					// remove top as its taken now by the O
					minCols = minCols[1:]
				}
			}
		}
	}
}

func tiltEast(in [][]byte) {
	for row, line := range in {
		var minCols []int
		for col := len(line) - 1; col >= 0; col-- {
			r := line[col]
			// keep track of the free cell that is furthest up north
			if r == '.' {
				minCols = append(minCols, col)
			} else if r == '#' {
				minCols = nil
			} else if r == 'O' {
				// move rock to the cell furthest up north
				fmt.Println("row", row, "col", col, minCols)
				if len(minCols) > 0 {
					minCol := minCols[0]
					in[row][minCol] = 'O'
					in[row][col] = '.'
					minCols = append(minCols, col)
					// remove top as its taken now by the O
					minCols = minCols[1:]
				}
			}
		}
	}
}
