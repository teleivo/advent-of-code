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
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		return 0, errors.New("failed to parse seeds, nothing to parse")
	}
	seeds, err := parseSeeds(scanner.Text())
	if err != nil {
		return 0, err
	}
	fmt.Println("seeds", seeds)

	m, err := parseMap(r)
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

func parseSeeds(in string) ([]int, error) {
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

func parseMap(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var nums []int
	for scanner.Scan() {
		fmt.Println("parseMap", scanner.Text())
		if len(scanner.Text()) > 0 && !unicode.IsDigit(rune(scanner.Text()[0])) {
			// skip map description
			continue
		}
		var src, dest, length int
		_, err := fmt.Sscan(scanner.Text(), &dest, &src, &length)
		if err != nil {
			return nil, fmt.Errorf("failed to parse map line %q: %v", scanner.Text(), err)
		}

		nums = append(nums, dest, src, length)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse map: %v", err)
	}
	fmt.Println("parseMap", nums)
	return nums, nil
}

func solvePartTwo(r io.Reader) (int, error) {
	return 0, nil
}
