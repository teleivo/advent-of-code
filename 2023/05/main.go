package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
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
	return parseInput(r)
}

func parseInput(r io.Reader) (int, error) {
	br := bufio.NewReader(r)

	seeds, err := parseSeeds(br)
	if err != nil {
		return 0, err
	}

	// TODO parse them into the tree's directly
	maps, err := parseMaps(br)
	if err != nil {
		return 0, err
	}

	nodes := mapsToNodes(maps)

	smallest := math.MaxInt
	for _, seed := range seeds {
		var target = seed
		for _, n := range nodes {
			target = find(n, target)
		}
		smallest = min(smallest, target)
	}

	return smallest, nil
}

func mapsToNodes(maps [][][3]int) []*node {
	var nodes []*node
	for _, m := range maps {
		var root *node
		for i, entry := range m {
			if i == 0 {
				root = &node{Source: entry[1], Dest: entry[0], RangeLen: entry[2]}
			} else {
				insert(root, entry[1], entry[0], entry[2])
			}
		}
		nodes = append(nodes, root)
	}
	return nodes
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
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("failed to parse map line %q: %v", in, err)
		}

		if len(in) > 0 {
			if unicode.IsLetter(rune(in[0])) {
				// skip map description
				pending = true
				continue
			} else if in == "\n" && pending {
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
	Source   int
	Dest     int
	RangeLen int
	Left     *node
	Right    *node
}

func insert(n *node, source, dest, rangeLen int) {
	if n.Source > source {
		if n.Left == nil {
			n.Left = &node{Source: source, Dest: dest, RangeLen: rangeLen}
		} else {
			insert(n.Left, source, dest, rangeLen)
		}
		return
	}

	if n.Source < source {
		if n.Right == nil {
			n.Right = &node{Source: source, Dest: dest, RangeLen: rangeLen}
		} else {
			insert(n.Right, source, dest, rangeLen)
		}
		return
	}
}

func find(n *node, source int) int {
	if n.Source <= source &&
		source < n.Source+n.RangeLen {
		return source - n.Source + n.Dest
	}
	if n.Source > source {
		if n.Left == nil {
			return source
		}
		return find(n.Left, source)
	}
	if n.Source < source {
		if n.Right == nil {
			return source
		}
		return find(n.Right, source)
	}

	return source
}

func solvePartTwo(r io.Reader) (int, error) {
	br := bufio.NewReader(r)

	seeds, err := parseSeeds(br)
	if err != nil {
		return 0, err
	}

	var seedRanges [][2]int
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, [2]int{seeds[i], seeds[i+1]})
	}

	// TODO parse them into the tree's directly
	maps, err := parseMaps(br)
	if err != nil {
		return 0, err
	}

	nodes := mapsToNodes(maps)

	smallest := math.MaxInt
	for _, seedRange := range seedRanges {
		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]; seed++ {
			var target = seed
			for _, n := range nodes {
				target = find(n, target)
			}
			smallest = min(smallest, target)
		}
	}

	return smallest, nil
}
