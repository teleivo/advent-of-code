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

	m, err := parseMap(br)
	if err != nil {
		return 0, err
	}
	fmt.Println("map", m)

	// m, err = parseMap(r)
	// if err != nil {
	// 	return 0, err
	// }
	// fmt.Println("map", m)

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

func parseMap(br *bufio.Reader) ([][]int, error) {
	var all [][]int
	var nums []int
	var pending bool
	for {
		in, err := br.ReadString('\n')
		fmt.Printf("%q\n", in)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("failed to parse map line %q: %v", in, err)
		}

		if len(in) > 0 {
			if unicode.IsLetter(rune(in[0])) {
				fmt.Println("map description: ", in)
				// skip map description
				pending = true
				continue
			} else if in == "\n" && pending {
				fmt.Println("newline while pending: ", in)
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
				return all, nil
			} else {
				return nil, fmt.Errorf("failed to parse map line %q: %v", in, err)
			}
		}
		nums = append(nums, dest, src, length)
	}
}

func solvePartTwo(r io.Reader) (int, error) {
	return 0, nil
}
