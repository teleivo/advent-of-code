package main

import (
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
	// var sum int
	// sc := bufio.NewScanner(r)
	// for sc.Scan() {
	// 	vals, err := parseLine(sc.Text())
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	for _, v := range vals {
	// 		sum += v
	// 	}
	// }
	// return sum, nil
	return 0, nil
}

type line struct {
	Numbers []number
	Symbols map[int]struct{}
}

type number struct {
	Value int
	Start int
	End   int
}

func parseLine(in string) (line, error) {
	res := line{Symbols: map[int]struct{}{}}
	var num *number

	for i, c := range in {
		if unicode.IsDigit(c) {
			if num == nil {
				num = &number{Start: i, End: i}
			}
		} else {
			if num != nil {
				num.End = i - 1
				v, err := strconv.Atoi(in[num.Start : num.End+1])
				if err != nil {
					return res, fmt.Errorf("failed to parse number: %v", err)
				}
				num.Value = v
				res.Numbers = append(res.Numbers, *num)
				num = nil
			}

			// if c != '.' && unicode.IsSymbol(c) {
			if c == '*' {
				res.Symbols[i] = struct{}{}
			}
		}
	}

	return res, nil
}
