package main

import (
	"bytes"
	"fmt"
	"hash/maphash"
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
	seen := make(map[uint64]int)
	seen2 := make(map[string]int)
	scores := make(map[int]int)
	var h maphash.Hash
	// for i := 0; i < 1_000_000_000; i++ {
	for i := 0; i < 1000; i++ {
		fmt.Println("round", i, "to go", 1_000_000_000-i)
		cycle(pattern)
		v := hash(h, pattern)
		if j, ok := seen[v]; ok {
			fmt.Println("seen pattern before in", j)
			return 0, nil
		}
		v2 := boardToString(pattern)
		if j, ok := seen2[v2]; ok {
			scoreId := j - 1 + (1_000_000_000-j)%(len(seen)-j)
			res := scores[scoreId]
			fmt.Println("seen pattern before in", j)
			return res, nil
		}
		seen[v] = i
		seen2[v2] = i
		scores[i] = calcScore(pattern)
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

func calcScore(in [][]byte) int {
	var sum int
	rows := len(in)
	for row, line := range in {
		for _, r := range line {
			if r == 'O' {
				sum += rows - row
			}
		}
	}
	return sum
}

func hash(z maphash.Hash, in [][]byte) uint64 {
	var h maphash.Hash
	for _, row := range in {
		h.Write(row)
	}
	res := h.Sum64()
	fmt.Println(res)
	h.Reset()
	return res
}

func boardToString(in [][]byte) string {
	var res strings.Builder
	for i, row := range in {
		res.WriteString(string(row))
		if i < len(in)-1 {
			res.WriteRune('\n')
		}
	}
	return res.String()
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
