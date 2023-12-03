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

	var vals []int
	var err error
	var prev, cur, next *line
	for sc.Scan() {
		if prev == nil && cur == nil { // initial case
			cur, err = parseLine(sc.Text())
			if err != nil {
				return 0, err
			}
			continue // we need to see if there is a line after this one to make a decision on the numbers
		} else if next == nil { // second case
			next, err = parseLine(sc.Text())
			if err != nil {
				return 0, err
			}
		} else { // once two lines have been parsed
			prev = cur
			cur = next
			next, err = parseLine(sc.Text())
			if err != nil {
				return 0, err
			}
		}
		vals = collectPartNumbers(prev, cur, next, vals)
	}
	vals = collectPartNumbers(cur, next, nil, vals)

	var sum int
	for _, v := range vals {
		sum += v
	}
	return sum, nil
}

func collectPartNumbers(prev, cur, next *line, partNumbers []int) []int {
	fmt.Println("collectPartNumbers", prev, cur, next)
	for _, n := range cur.Numbers {
		if (prev != nil && isSymbolAdjacent(n, prev.Symbols)) || isSymbolAdjacent(n, cur.Symbols) || (next != nil && isSymbolAdjacent(n, next.Symbols)) {
			partNumbers = append(partNumbers, n.Value)
			continue
		}
	}

	return partNumbers
}

func isSymbolAdjacent(n number, symbols map[int]struct{}) bool {
	for pos := range symbols {
		fmt.Println("number", n, "symbol pos", pos)
		if pos >= n.Start-1 && pos <= n.End+1 {
			fmt.Println("yes")
			return true
		}
		// TODO bail out of symbols loop if we are out of range
	}
	return false
}

// TODO Symbols can simply be a slice
type line struct {
	Numbers []number
	Symbols map[int]struct{}
}

type number struct {
	Value int
	Start int
	End   int
}

func parseLine(in string) (*line, error) {
	res := line{Symbols: map[int]struct{}{}}
	var num *number

	var x int
	for i, c := range in {
		x = i
		if unicode.IsDigit(c) {
			if num == nil {
				num = &number{Start: i, End: i}
			}
		} else {
			if num != nil {
				num.End = i - 1
				v, err := strconv.Atoi(in[num.Start : num.End+1])
				if err != nil {
					return nil, fmt.Errorf("failed to parse number: %v", err)
				}
				num.Value = v
				res.Numbers = append(res.Numbers, *num)
				num = nil
			}

			if c != '.' && !unicode.IsSpace(c) && !unicode.IsLetter(c) {
				res.Symbols[i] = struct{}{}
			}
		}
	}
	if num != nil {
		num.End = x
		v, err := strconv.Atoi(in[num.Start : num.End+1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse number: %v", err)
		}
		num.Value = v
		res.Numbers = append(res.Numbers, *num)
		num = nil
	}

	return &res, nil
}

func isSymbol(c rune) bool {
	if c != '.' && !unicode.IsSpace(c) && !unicode.IsLetter(c) && !unicode.IsNumber(c) {
		return true
	}
	return false
}
