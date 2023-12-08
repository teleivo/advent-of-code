package main

import (
	"bufio"
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
	s := bufio.NewScanner(r)

	if !s.Scan() {
		return 0, errors.New("no instructions to read")
	}
	instructions := s.Text()

	network := make(map[string]*node)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			// skip empty line
			continue
		}
		label, directions, found := strings.Cut(line, " = ")
		if !found {
			return 0, fmt.Errorf("failed to split node on '=' in line %q", line)
		}
		// fmt.Println("label", label, "directions", directions)
		leftStr, rightStr, found := strings.Cut(directions, ", ")
		if !found {
			return 0, fmt.Errorf("failed to split directions on ', ' in line %q", line)
		}
		left := leftStr[1:]
		right := rightStr[:len(rightStr)-1]
		// fmt.Println("left", left, "right", right)

		var n *node
		if _, ok := network[label]; !ok {
			n = &node{Label: label}
		} else {
			n = network[label]
		}

		var leftNode *node
		if _, ok := network[left]; !ok {
			leftNode = &node{Label: left}
		} else {
			leftNode = network[left]
		}
		n.Left = leftNode

		var rightNode *node
		if _, ok := network[right]; !ok {
			rightNode = &node{Label: right}
		} else {
			rightNode = network[right]
		}
		n.Right = rightNode

		network[label] = n
	}

	for k, v := range network {
		fmt.Printf("%q: label %q, left %q, right %q\n", k, v.Label, v.Left.Label, v.Right.Label)
	}

	start := "AAA"
	steps := findGoal(network, start, instructions)

	return steps, nil
}

func findGoal(network map[string]*node, start, instructions string) int {
	fmt.Println("start at", start)
	var steps int
	cur := start
	for _, instruction := range instructions {
		fmt.Println("at node", cur, "now going", string(instruction))
		n := network[cur]
		if instruction == 'L' {
			cur = n.Left.Label
		} else {
			cur = n.Right.Label
		}
		steps++

		if cur == "ZZZ" {
			fmt.Println("arrived at end in", steps)
			return steps
		}
	}

	if cur != "ZZZ" {
		fmt.Println("TRYING Again", steps)
		return findGoal(network, cur, instructions) + steps
	}

	// Todo?
	return steps
}

type node struct {
	Label string
	Left  *node
	Right *node
}

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	return 0, nil
}
