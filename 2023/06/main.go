package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	br := bufio.NewReader(r)
	durations, err := parseLine(br)
	if err != nil {
		return 0, err
	}
	fmt.Println("Time:", durations)

	distances, err := parseLine(br)
	if err != nil {
		return 0, err
	}
	fmt.Println("Distance:", distances)

	var result int
	for i, duration := range durations {
		var sum int
		// only loop through half of the possible hold durations to utilize the symmetry in
		// distance traveled = hold duration * (duration - hold duration)
		maxHoldDuration := duration / 2
		if isEven(duration) {
			maxHoldDuration -= 1
		}
		for j := 0; j <= maxHoldDuration; j++ {
			if isBreakingRecord(j, duration, distances[i]) {
				sum++
			}
		}

		sum *= 2
		// if the total possible hold durations are odd check the extra asymmetric case
		if isEven(duration) {
			if isBreakingRecord(maxHoldDuration+1, duration, distances[i]) {
				sum++
			}
		}

		fmt.Println("duration", duration, "maxHoldDuration", maxHoldDuration, "sum", sum)
		if result == 0 {
			result = sum
		} else {
			result *= sum
		}
	}

	return result, nil
}

func isEven(n int) bool {
	return n%2 == 0
}

func isBreakingRecord(timeHeld, raceDuration, recordDistance int) bool {
	if timeHeld*(raceDuration-timeHeld) > recordDistance {
		return true
	}
	return false

}

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	return 0, nil
}

func parseLine(r *bufio.Reader) ([]int, error) {
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
