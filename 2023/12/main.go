package main

import (
	"bufio"
	"bytes"
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

func findArrangements(springs []byte, groups []int) int {
	if len(springs) == 0 {
		return 1
	}
	if len(groups) == 1 {
		group := groups[0]
		// TODO at least in case springs are all ? there are arrangements: unknowns - group + 1
		if !strings.ContainsAny(string(springs), "#.") && strings.ContainsRune(string(springs), '?') {
			// assuming springs is all ?
			return len(springs) - group + 1
		}
		return 1
	}
	input := strings.Trim(string(springs), ".")
	// TODO I am using all groups, I might need to only use the one that are <= min width
	// assuming springs is all ?
	if !strings.ContainsAny(input, "#.") && strings.ContainsRune(input, '?') {
		n := len(input) - groups[0] - separators(groups)
		arrangements := sumOfN(n)
		return arrangements
	}

	fmt.Println(input)
	// find a delimited chunk of springs
	// either all ????
	// or what are the other special cases?
	var end int
	var unknowns int
	for end < len(input) {
		if input[end] == '?' {
			unknowns++
		} else if input[end] == '.' && unknowns != 0 {
			end++ // as this separator needs to be consumed
			break
		}
		end++
	}
	head := input[0:end]
	head = strings.Trim(head, ".")
	tail := input[end:]
	tail = strings.Trim(tail, ".")
	headGroups, tailGroups := splitGroups(head, groups)
	fmt.Println("chunk", string(head), "rest", string(tail))

	// TODO split groups as well. a group can only be consumed once
	// in case head is all ????

	return findArrangements([]byte(head), headGroups) + findArrangements([]byte(tail), tailGroups)
}

func sumOfN(n int) int {
	return (n * (n + 1)) / 2
}

func minWidth(groups []int) int {
	return sum(groups) + separators(groups)
}

func sum(nums []int) int {
	var res int
	for _, num := range nums {
		res += num
	}
	return res
}

// separators calculates how many working springs need to separate broken groups in a contiguous
// stream of unknown '?' spring conditions
func separators(groups []int) int {
	return len(groups) - 1
}

func splitGroups(unknownSprings string, groups []int) ([]int, []int) {
	width := len(unknownSprings)
	for i := 0; i < len(groups); i++ {
		m := minWidth(groups[:i+1])
		if width == m {
			return groups[:i+1], groups[i+1:]
		}
		if width < m {
			return groups[:i], groups[i:]
		}
	}
	return groups, nil
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
