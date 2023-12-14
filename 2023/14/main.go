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
			// fmt.Println("tilt after", "row", row, "col", col)
			// for _, row := range in {
			// fmt.Println(string(row))
			// }
		}
	}
}

// solvePartTwo solves part one of the puzzle.
func solvePartTwo(in []byte) (int, error) {
	pattern := bytes.Fields(in)
	// fmt.Println("before tilt")
	for _, row := range pattern {
		fmt.Println(string(row))
	}
	tiltNorth(pattern)
	// fmt.Println()
	// fmt.Println("after tilt")
	var sum int
	rows := len(pattern)
	for row, line := range pattern {
		// fmt.Println(string(line))
		for _, r := range line {
			if r == 'O' {
				sum += rows - row
			}
		}
	}

	return sum, nil
}
