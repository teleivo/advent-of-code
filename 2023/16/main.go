package main

import (
	"bytes"
	"errors"
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

var move map[direction]func(beam) beam = map[direction]func(beam) beam{
	left:  moveLeft,
	right: moveRight,
	up:    moveUp,
	down:  moveDown,
}

func moveLeft(in beam) beam {
	return beam{dir: left, pos: point{row: in.pos.row, col: in.pos.col - 1}}
}

func moveRight(in beam) beam {
	return beam{dir: right, pos: point{row: in.pos.row, col: in.pos.col + 1}}
}

func moveUp(in beam) beam {
	return beam{dir: up, pos: point{row: in.pos.row - 1, col: in.pos.col}}

}

func moveDown(in beam) beam {
	return beam{dir: down, pos: point{row: in.pos.row + 1, col: in.pos.col}}
}

type beam struct {
	pos point
	dir direction
}

// solvePartOne solves part one of the puzzle.
func solvePartOne(in []byte) (int, error) {
	pattern := bytes.Fields(in)
	seen := make(map[beam]struct{})

	energized := make(map[point]struct{})
	todo := []beam{{dir: right}}

	for len(todo) > 0 {
		currentBeam := todo[0]
		fmt.Printf("currentBeam %#v\n", currentBeam)
		if len(todo) > 1 { // drop beam we already processed
			todo = todo[1:]
		} else {
			todo = nil
		}
		if _, ok := seen[currentBeam]; ok {
			fmt.Printf("beam already seen %v\n", currentBeam)
			continue
		}
		if !isInBounds(currentBeam, pattern) {
			fmt.Printf("beam is out of bounds %v\n", currentBeam)
			continue
		}
		seen[currentBeam] = struct{}{}
		energized[currentBeam.pos] = struct{}{}

		switch pattern[currentBeam.pos.row][currentBeam.pos.col] {
		case '/':
			if currentBeam.dir == right {
				currentBeam = moveUp(currentBeam)
				fmt.Printf("hit / moving up to %v\n", currentBeam)
			} else if currentBeam.dir == left {
				currentBeam = moveDown(currentBeam)
				fmt.Printf("hit / from left moving down to %v\n", currentBeam)
			} else if currentBeam.dir == up {
				currentBeam = moveRight(currentBeam)
				fmt.Printf("hit / from up moving right to %v\n", currentBeam)
			} else if currentBeam.dir == down {
				currentBeam = moveLeft(currentBeam)
				fmt.Printf("hit / from down moving left to %v\n", currentBeam)
			}
			todo = append(todo, currentBeam)
		case '\\':
			if currentBeam.dir == right {
				currentBeam = moveDown(currentBeam)
				fmt.Printf("hit \\ moving down to %v\n", currentBeam)
			} else if currentBeam.dir == left {
				currentBeam = moveUp(currentBeam)
				fmt.Printf("hit \\ from the left moving up to %v\n", currentBeam)
			} else if currentBeam.dir == up {
				currentBeam = moveLeft(currentBeam)
				fmt.Printf("hit \\ from the up moving left to %v\n", currentBeam)
			} else if currentBeam.dir == down {
				currentBeam = moveRight(currentBeam)
				fmt.Printf("hit \\ from the down moving right to %v\n", currentBeam)
			}
			todo = append(todo, currentBeam)
		case '|':
			if currentBeam.dir == up || currentBeam.dir == down {
				// continue in the same direction
				mf, ok := move[currentBeam.dir]
				if !ok {
					return 0, errors.New("move func not found")
				}
				currentBeam = mf(currentBeam)
				fmt.Printf("hit | from up or down so continue in the same direction%v\n", currentBeam)
				todo = append(todo, currentBeam)
			} else { // split the beam
				downBeam := moveDown(currentBeam)
				if isInBounds(downBeam, pattern) {
					todo = append(todo, downBeam)
				}
				currentBeam = moveUp(currentBeam)
				todo = append(todo, currentBeam)
				fmt.Printf("hit | splitting up %v and down %v\n", currentBeam, downBeam)
			}
		case '-':
			if currentBeam.dir == left || currentBeam.dir == right {
				// continue in the same direction
				mf, ok := move[currentBeam.dir]
				if !ok {
					return 0, errors.New("move func not found")
				}
				currentBeam = mf(currentBeam)
				fmt.Printf("hit - from left or right so continue in the same direction%v\n", currentBeam)
				todo = append(todo, currentBeam)
			} else { // split the beam
				rightBeam := moveRight(currentBeam)
				if isInBounds(rightBeam, pattern) {
					todo = append(todo, rightBeam)
				}
				currentBeam = moveLeft(currentBeam)
				todo = append(todo, currentBeam)
				fmt.Printf("hit - splitting left %v and right %v\n", currentBeam, rightBeam)
			}
		case '.':
			mf, ok := move[currentBeam.dir]
			if !ok {
				return 0, errors.New("move func not found")
			}
			currentBeam = mf(currentBeam)
			todo = append(todo, currentBeam)
			fmt.Printf("empty tile continue moving %v\n", currentBeam)
		}

		st := fieldToString(pattern, energized)
		fmt.Println(st)
		fmt.Println("energized tiles", len(energized))
	}

	st := fieldToString(pattern, energized)
	fmt.Println(st)
	return len(energized), nil
}

func isInBounds(in beam, field [][]byte) bool {
	rows := len(field)
	cols := len(field[0])
	return in.pos.row >= 0 && in.pos.row < rows && in.pos.col >= 0 && in.pos.col < cols
}

func fieldToString(pattern [][]byte, energized map[point]struct{}) string {
	var res strings.Builder
	for row := 0; row < len(pattern); row++ {
		for col := 0; col < len(pattern[0]); col++ {
			if _, ok := energized[point{row: row, col: col}]; ok {
				res.WriteRune('#')
			} else {
				res.WriteRune('.')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

// solvePartTwo solves part two of the puzzle.
func solvePartTwo(in []byte) (int, error) {
	pattern := bytes.Fields(in)
	_ = pattern
	return 0, nil
}
