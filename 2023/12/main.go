package main

import (
	"bufio"
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

	f1, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
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
func solvePartOne(b []byte) (int, error) {
	rows := bytes.Fields(b)
	for _, row := range rows {
		for _, v := range row {
			fmt.Println(string(v))
		}
	}
	return 0, nil
}

func findArrangements(springs []byte, damaged []int) int {
	arrangements := 1
	var idx int
	fmt.Println(string(springs))
	for _, damage := range damaged {
		needed := damage
		fmt.Println("looking for ", needed, "springs from idx", idx)
		var pending bool
		var unknown int // represents '?' springs
		arrangement := 1
		for idx < len(springs) {
			fmt.Println("idx", idx, "char", string(springs[idx]))
			if springs[idx] == '?' {
				pending = true
				unknown++
			} else if springs[idx] == '#' {
				needed--
			}

			fmt.Println("unkown", unknown, "needed", needed)
			if springs[idx] == '.' && pending {
				// case without any unknown springs or any permutations of arrangement
				if needed == 0 || needed == unknown {
					arrangement = 1
				} else if unknown != 0 {
					fmt.Println("damage", damage, "unknown", unknown, "needed", needed, "arrangement", arrangement)
					arrangement = unknown - needed + 1
				}
				idx++ // since we are on the '.' we only skip that one
				break
			}
			idx++
		}
		arrangements *= arrangement
	}
	return arrangements
}

// solvePartTwo solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		_ = line
	}
	return 0, nil
}
