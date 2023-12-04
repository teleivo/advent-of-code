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
func solvePartOne(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)

	var vals []int
	var prev, cur, next *line
	for sc.Scan() {
		l, err := parseLine(isAnySymbol, sc.Text())
		if err != nil {
			return 0, err
		}
		if prev == nil && cur == nil { // initial case
			cur = l
			continue // we need to see if there is a line after this one to make a decision on the numbers
		} else if next == nil { // second case
			next = l
		} else { // once two lines have been parsed
			prev = cur
			cur = next
			next = l
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
	for _, n := range cur.Numbers {
		if (prev != nil && isSymbolAdjacent(n, prev.Symbols)) || isSymbolAdjacent(n, cur.Symbols) || (next != nil && isSymbolAdjacent(n, next.Symbols)) {
			partNumbers = append(partNumbers, n.Value)
			continue
		}
	}

	return partNumbers
}

func isSymbolAdjacent(n number, symbols []int) bool {
	for _, pos := range symbols {
		if pos >= n.Start-1 && pos <= n.End+1 {
			return true
		}
		// TODO bail out of symbols loop if we are out of range
	}
	return false
}

// TODO Symbols can simply be a slice
type line struct {
	Numbers []number
	Symbols []int
}

type number struct {
	Value int
	Start int
	End   int
}

func parseLine(isSymbol func(rune) bool, in string) (*line, error) {
	res := line{}
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

			if isSymbol(c) {
				res.Symbols = append(res.Symbols, i)
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

func isAnySymbol(c rune) bool {
	if c != '.' && !unicode.IsSpace(c) && !unicode.IsLetter(c) && !unicode.IsNumber(c) {
		return true
	}
	return false
}

func isGearSymbol(c rune) bool {
	if c == '*' {
		return true
	}
	return false
}

func solvePartTwo(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)

	var vals []int
	var prev, cur, next *line
	for sc.Scan() {
		l, err := parseLine(isGearSymbol, sc.Text())
		if err != nil {
			return 0, err
		}
		if prev == nil && cur == nil { // initial case
			cur = l
			continue // we need to see if there is a line after this one to make a decision on the numbers
		} else if next == nil { // second case
			next = l
		} else { // once two lines have been parsed
			prev = cur
			cur = next
			next = l
		}
		vals = collectGears(prev, cur, next, vals)
	}
	vals = collectGears(cur, next, nil, vals)

	var sum int
	for _, v := range vals {
		sum += v
	}
	return sum, nil
}

// TODO 4082059 is too low
func collectGears(prev, cur, next *line, gears []int) []int {
	var adj []int
	for _, pos := range cur.Symbols {
		for _, num := range prev.Numbers {
			if isNumberAdjacent(pos, num) {
				if len(adj) == 2 {
					return gears
				}
				adj = append(adj, num.Value)
			}
		}
		for _, num := range cur.Numbers {
			if isNumberAdjacent(pos, num) {
				if len(adj) == 2 {
					return gears
				}
				adj = append(adj, num.Value)
			}
		}
		for _, num := range next.Numbers {
			if isNumberAdjacent(pos, num) {
				if len(adj) == 2 {
					return gears
				}
				adj = append(adj, num.Value)
			}
		}
	}

	if len(adj) == 2 {
		// fmt.Printf("gear: %d*%d=%d\n", adj[0], adj[1], adj[0]*adj[1])
		gears = append(gears, adj[0]*adj[1])
	}
	return gears
}

func isNumberAdjacent(symbolPos int, num number) bool {
	if symbolPos >= num.Start-1 && symbolPos <= num.End+1 {
		// fmt.Printf("adjacent: symbol at %d to number %v\n", symbolPos, num)
		return true
	}
	return false
}
