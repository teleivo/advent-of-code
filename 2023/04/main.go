package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
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

	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer f.Close()
	cal, err := solvePartOne(f)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle one is: %d\n", cal)

	return nil
}

// solvePartOne solves part one of the puzzle.
func solvePartOne(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)

	var sum int
	for sc.Scan() {
		sum += computePoints(sc.Text())
	}

	return sum, nil
}

// computePoints computes the points won in one scratch card game
func computePoints(line string) int {
	winning := make(map[int]struct{})

	var points int
	var collect bool
	var inWin bool
	var pending bool
	var num []rune
	for _, c := range line {
		if c == ':' { // winning numbers are in between ':' and '|'
			collect = true
			inWin = true
		}
		if c == '|' { // numbers are after '|'
			inWin = false
		}

		if !collect {
			continue
		}

		if unicode.IsDigit(c) {
			pending = true
			num = append(num, c)
			continue
		}

		if pending {
			if num != nil {
				v, _ := strconv.Atoi(string(num))
				pending = false
				num = nil

				if inWin {
					winning[v] = struct{}{}
				} else {
					if _, ok := winning[v]; ok {
						if points == 0 { // first winning nr counts as one
							points = 1
						} else { // ever winning nr after the first doubles points
							points *= 2
						}
					}
				}
			}
		}
	}
	// todo how to beautify this pattern
	if num != nil {
		v, _ := strconv.Atoi(string(num))
		pending = false
		num = nil

		if _, ok := winning[v]; ok {
			if points == 0 { // first winning nr counts as one
				points = 1
			} else { // ever winning nr after the first doubles points
				points *= 2
			}
		}
	}

	return points
}
