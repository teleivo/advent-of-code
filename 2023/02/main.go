package main

import (
	"bufio"
	"strings"
	"errors"
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
	cal, err := solveFeasibleGames(f, [3]int{12, 13, 14})
	if err!=nil{
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle one is: %d\n", cal)

	ff, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer ff.Close()
	cal, err = solveMinimumSetOfCubes(ff)
	if err!=nil{
		return err
	}
	fmt.Fprintf(w, "The solution to puzzle two is: %d\n", cal)

	return nil
}

// solveFeasibleGames solves part one of the puzzle. What is the sum of the game IDs that could be
// played with the given cubes.
func solveFeasibleGames(r io.Reader, cubes[3]int) (int, error) {
	var sum int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		g , err := parseLine(sc.Text())
		if err != nil {
			return 0, err
		}

		if g.cubes[0] <= cubes[0] &&
		g.cubes[1] <= cubes[1] &&
		g.cubes[2] <= cubes[2] {
		sum += g.ID
		}
	}
	return sum, nil
}

// solveMinimumSetOfCubes solves part two of the puzzle. What is the sum of the powers of the minimum set of cubes that must have been present
// for a game to be played. Luckily, I already parse the game sets into the max cubes per rgb
// colors. The max cube per color is the minimum cubes you need of that color to play a game.
func solveMinimumSetOfCubes(r io.Reader) (int, error) {
	var sum int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		g , err := parseLine(sc.Text())
		if err != nil {
			return 0, err
		}

		sum += g.cubes[0]*g.cubes[1]*g.cubes[2]
	}
	return sum, nil
}

func parseLine(line string) (*game, error) {
	id, sets, found:=strings.Cut(line, ":")
	if !found{
		return nil, errors.New( fmt.Sprintf("separator ':' not found in line %q", line))
	}

	var ID int
	n, err:=fmt.Sscanf(id, "Game %d",&ID)
	if err!=nil{
		return nil, fmt.Errorf("failed scanning game ID and sets: %v", err)
	}
	if n!=1{
		return nil, fmt.Errorf("failed scanning game ID and sets expected 1 token instead got %d", n)
	}
	
	return &game{
		ID: ID,
		cubes: maxCubes(sets),
	}, nil
}

type game struct{
	ID int
	cubes [3]int
}

func maxCubes(line string) [3]int {
	cubes := map[rune]int{
		'r': 0,
		'g': 0,
		'b': 0,
	}

	var pending bool
	var num []rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			pending=true
			num = append(num, c)
			continue
		}

		if pending && ( c == 'r' || c == 'g' || c == 'b') {
			if num != nil {
				v, _ := strconv.Atoi(string(num))
				cnt := cubes[c]
				if v > cnt {
					cubes[c] = v
				}
				pending = false
				num = nil
			}
		}
	}

	r := cubes['r']
	g := cubes['g']
	b := cubes['b']
	return [3]int{r, g, b}
}
