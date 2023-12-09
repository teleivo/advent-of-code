package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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
	s := bufio.NewScanner(r)
	var input [][]int
	for s.Scan() {
		line := s.Text()
		numsStr := strings.Split(line, " ")
		nums := make([]int, len(numsStr))
		for i, s := range numsStr {
			num, _ := strconv.Atoi(s)
			nums[i] = num
		}
		input = append(input, nums)
	}

	var sum int
	for _, v := range input {
		sum += predict(v)
	}

	return sum, nil
}

func predict(in []int) int {
	fmt.Println("in", in)
	var lasts []int
	cur := in
	for {
		var last int
		var next []int
		allZeroes := true
		for i := 1; i < len(cur); i++ {
			d := cur[i] - cur[i-1]
			if d != 0 {
				allZeroes = false
			}
			last = cur[i]
			next = append(next, d)
		}
		lasts = append(lasts, last)

		if allZeroes {
			break
		}
		cur = next
	}
	lasts = append(lasts, 0)

	// could range over it in reverse instead. doing this to solve the problem without getting hung
	// up on order issues
	slices.Reverse(lasts)

	var prediction int
	predictions := make([]int, len(lasts))
	for i := 1; i < len(lasts); i++ {
		predictions[i] = lasts[i] + predictions[i-1]
		prediction = predictions[i]
	}
	fmt.Println("predictions", predictions, "prediction", prediction)

	return prediction
}

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	var input [][]int
	for s.Scan() {
		line := s.Text()
		numsStr := strings.Split(line, " ")
		nums := make([]int, len(numsStr))
		for i, s := range numsStr {
			num, _ := strconv.Atoi(s)
			nums[i] = num
		}
		input = append(input, nums)
	}

	var sum int
	for _, v := range input {
		sum += predictPartTwo(v)
	}

	return sum, nil
}

func predictPartTwo(in []int) int {
	fmt.Println("in", in)
	var firsts []int
	cur := in
	for {
		var next []int
		allZeroes := true
		for i := 1; i < len(cur); i++ {
			d := cur[i] - cur[i-1]
			if d != 0 {
				allZeroes = false
			}
			next = append(next, d)
		}
		firsts = append(firsts, cur[0])

		if allZeroes {
			break
		}
		cur = next
	}
	firsts = append(firsts, 0)

	// could range over it in reverse instead. doing this to solve the problem without getting
	// hung
	// up on order issues
	slices.Reverse(firsts)

	fmt.Println("firsts", firsts)

	var prediction int
	predictions := make([]int, len(firsts))
	for i := 1; i < len(firsts); i++ {
		predictions[i] = firsts[i] - predictions[i-1]
		prediction = predictions[i]
	}
	fmt.Println("predictions", predictions, "prediction", prediction)

	return prediction
}
