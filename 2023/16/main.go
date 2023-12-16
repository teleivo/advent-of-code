package main

import (
	"bytes"
	"errors"
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

type direction int

const (
	left direction = iota
	right
	up
	down
)

type point struct {
	row int
	col int
}

var move map[direction]func(*beam) = map[direction]func(*beam){
	left:  moveLeft,
	right: moveRight,
	up:    moveUp,
	down:  moveDown,
}

func moveLeft(in *beam) {
	in.dir = left
	in.pos = point{row: in.pos.row, col: in.pos.col - 1}
}
func moveRight(in *beam) {
	in.dir = right
	in.pos = point{row: in.pos.row, col: in.pos.col + 1}
}
func moveUp(in *beam) {
	in.dir = up
	in.pos = point{row: in.pos.row - 1, col: in.pos.col}
}
func moveDown(in *beam) {
	in.dir = down
	in.pos = point{row: in.pos.row + 1, col: in.pos.col}
}

type beam struct {
	pos point
	dir direction
}

// solvePartOne solves part one of the puzzle.
func solvePartOne(in []byte) (int, error) {
	pattern := bytes.Fields(in)

	energized := make(map[point]struct{})
	beams := []*beam{{dir: right}}

	for len(beams) > 0 {
		currentBeam := beams[0]
		fmt.Printf("currentBeam %v\n", currentBeam)

		switch pattern[currentBeam.pos.row][currentBeam.pos.col] {
		case '/':
			energized[currentBeam.pos] = struct{}{}
			if currentBeam.dir == right {
				moveUp(currentBeam)
				fmt.Printf("hit / moving up to %v\n", currentBeam)
			} else if currentBeam.dir == left {
				moveDown(currentBeam)
				fmt.Printf("hit / from left moving down to %v\n", currentBeam)
			} else { // currentBeam.dir == up || currentBeam.dir == down
				// continue in the same beam.direction
				mf, ok := move[currentBeam.dir]
				if !ok {
					return 0, errors.New("move func not found")
				}
				mf(currentBeam)
				fmt.Printf("hit / from up or down so continue in the same direction%v\n", currentBeam)
			}
		case '\\':
			energized[currentBeam.pos] = struct{}{}
			if currentBeam.dir == right {
				moveDown(currentBeam)
				fmt.Printf("hit \\ moving down to %v\n", currentBeam)
			} else if currentBeam.dir == left {
				moveUp(currentBeam)
				fmt.Printf("hit \\ from the left moving up to %v\n", currentBeam)
			} else { // if currentBeam.dir == up || currentBeam.dir == down
				// continue in the same direction
				mf, ok := move[currentBeam.dir]
				if !ok {
					return 0, errors.New("move func not found")
				}
				mf(currentBeam)
				fmt.Printf("hit \\ from left or right so continue in the same direction%v\n", currentBeam)
			}
		case '|':
			energized[currentBeam.pos] = struct{}{}
			if currentBeam.dir == up || currentBeam.dir == down {
				// continue in the same direction
				mf, ok := move[currentBeam.dir]
				if !ok {
					return 0, errors.New("move func not found")
				}
				mf(currentBeam)
				fmt.Printf("hit | from up or down so continue in the same direction%v\n", currentBeam)
			} else { // split the beam
				downBeam := &beam{pos: currentBeam.pos}
				moveDown(downBeam)
				if isInBounds(downBeam, pattern) {
					beams = append(beams, downBeam)
				}
				moveUp(currentBeam)
				fmt.Printf("hit | splitting up %v and down %v\n", currentBeam, downBeam)
			}
		case '-':
			energized[currentBeam.pos] = struct{}{}
			if currentBeam.dir == left || currentBeam.dir == right {
				// continue in the same direction
				mf, ok := move[currentBeam.dir]
				if !ok {
					return 0, errors.New("move func not found")
				}
				mf(currentBeam)
				fmt.Printf("hit - from left or right so continue in the same direction%v\n", currentBeam)
			} else { // split the beam
				rightBeam := &beam{pos: currentBeam.pos}
				moveRight(rightBeam)
				if isInBounds(rightBeam, pattern) {
					beams = append(beams, rightBeam)
				}
				moveLeft(currentBeam)
				fmt.Printf("hit - splitting left %v and right %v\n", currentBeam, rightBeam)
			}
		case '.':
			energized[currentBeam.pos] = struct{}{}
			mf, ok := move[currentBeam.dir]
			if !ok {
				return 0, errors.New("move func not found")
			}
			mf(currentBeam)
			fmt.Printf("empty tile continue moving %v\n", currentBeam)
		}

		if !isInBounds(currentBeam, pattern) {
			fmt.Printf("current beam out of bounds %v\n", currentBeam)
			if len(beams) > 1 {
				beams = beams[1:]
			} else {
				beams = nil
			}
		}
	}

	for row := 0; row < len(pattern); row++ {
		for col := 0; col < len(pattern[0]); col++ {
			if _, ok := energized[point{row: row, col: col}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return len(energized), nil
}

func isInBounds(in *beam, field [][]byte) bool {
	rows := len(field)
	cols := len(field[0])
	return in.pos.row >= 0 && in.pos.row < rows && in.pos.col >= 0 && in.pos.col < cols
}

// solvePartTwo solves part two of the puzzle.
func solvePartTwo(in []byte) (int, error) {
	pattern := bytes.Fields(in)
	_ = pattern
	return 0, nil
}
