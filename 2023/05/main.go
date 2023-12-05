package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	f1, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
	defer f1.Close()
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
func solvePartOne(r io.Reader) (int, error) {
	return 0, nil
}

func parseInput(r io.Reader) (int, error) {
	br := bufio.NewReader(r)

	seeds, err := parseSeeds(br)
	if err != nil {
		return 0, err
	}
	fmt.Println("seeds", seeds)

	maps, err := parseMaps(br)
	if err != nil {
		return 0, err
	}
	fmt.Println("map", maps)

	// TODO insert into tree
	var trees []*node
	for j, m := range maps {
		var root *node
		for i, entry := range m {
			if i == 0 {
				root = &node{Val: entry[1], Dest: entry[0], RangeLen: entry[2]}
			} else {
				insert(root, entry[1], entry[0], entry[2])
			}
		}
		fmt.Println("tree", j, root)
		trees = append(trees, root)
	}

	// TODO find soil in tree
	// todo replace with loop over seeds
	targetSeed := 79
	var smallest int
	for _, n := range trees {

	}

	return 0, nil
}

func parseSeeds(r *bufio.Reader) ([]int, error) {
	in, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to parse seeds in: %q", in)
	}

	_, seeds, found := strings.Cut(in, ": ")
	if !found {
		return nil, fmt.Errorf("failed to parse seeds in: %q", in)
	}

	// TODO I wanted to make parseNumbers reusable but sharing the reader is not smooth so far
	return parseNumbers(strings.NewReader(seeds))
}

func parseNumbers(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse number: %v", err)
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse numbers: %v", err)
	}
	return nums, nil
}

func parseMaps(br *bufio.Reader) ([][][3]int, error) {
	var all [][][3]int
	var nums [][3]int
	var pending bool
	for {
		in, err := br.ReadString('\n')
		fmt.Printf("%q\n", in)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("failed to parse map line %q: %v", in, err)
		}

		if len(in) > 0 {
			if unicode.IsLetter(rune(in[0])) {
				// fmt.Println("map description: ", in)
				// skip map description
				pending = true
				continue
			} else if in == "\n" && pending {
				// fmt.Println("newline while pending: ", in)
				// newlines after a map description terminate the map
				pending = false
				all = append(all, nums)
				nums = nil
				continue
			} else if in == "\n" {
				continue
			}
		}

		var src, dest, length int
		_, err = fmt.Sscan(in, &dest, &src, &length)
		if err != nil {
			if errors.Is(err, io.EOF) {
				all = append(all, nums)
				return all, nil
			} else {
				return nil, fmt.Errorf("failed to parse map line %q: %v", in, err)
			}
		}
		nums = append(nums, [3]int{dest, src, length})
	}
}

type node struct {
	Val      int
	Dest     int
	RangeLen int
	Left     *node
	Right    *node
}

func insert(n *node, val, dest, rangeLen int) {
	if n.Val < val {
		if n.Left == nil {
			n.Left = &node{Val: val, Dest: dest, RangeLen: rangeLen}
		} else {
			insert(n.Left, val, dest, rangeLen)
		}
		return
	}

	if n.Val > val {
		if n.Right == nil {
			n.Right = &node{Val: val, Dest: dest, RangeLen: rangeLen}
		} else {
			insert(n.Right, val, dest, rangeLen)
		}
		return
	}
}

func find(n *node, source int) int {
	if n.Val < source {
		if n.Left == nil {
			return source
		} else {
			insert(n.Left, val, dest, rangeLen)
		}
		return
	}

	if n.Val > val {
		if n.Right == nil {
			n.Right = &node{Val: val, Dest: dest, RangeLen: rangeLen}
		} else {
			insert(n.Right, val, dest, rangeLen)
		}
		return
	}
}

func solvePartTwo(r io.Reader) (int, error) {
	return 0, nil
}
