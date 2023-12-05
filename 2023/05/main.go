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
	sc := bufio.NewScanner(r)

	var sum int
	for sc.Scan() {
		sum += computePoints(sc.Text())
	}

	return sum, nil
}

func parseSeeds(r io.Reader) ([]int, error) {

	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		_, seeds, found := strings.Cut(scanner.Text(), ": ")
		if !found {
			return nil, fmt.Errorf("failed to parse seeds in: %q", scanner.Text())
		}

		return parseNumbers(strings.NewReader(seeds))
	}

	return nil, errors.New("failed to parse seeds: did not find ': '")
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

// computePoints computes the points won in one scratch card game
func computePoints(line string) int {
	winning := make(map[int]struct{})

	var points int
	var collect bool
	var inWin bool
	var pending bool
	var num []rune
	for _, c := range line {
		if c == ':' { // winning numbers are in between ':' and '|'
			collect = true
			inWin = true
		}
		if c == '|' { // numbers are after '|'
			inWin = false
		}

		if !collect {
			continue
		}

		if unicode.IsDigit(c) {
			pending = true
			num = append(num, c)
			continue
		}

		if pending {
			if num != nil {
				v, _ := strconv.Atoi(string(num))
				pending = false
				num = nil

				if inWin {
					winning[v] = struct{}{}
				} else {
					if _, ok := winning[v]; ok {
						if points == 0 { // first winning nr counts as one
							points = 1
						} else { // ever winning nr after the first doubles points
							points *= 2
						}
					}
				}
			}
		}
	}
	// todo how to beautify this pattern
	if num != nil {
		v, _ := strconv.Atoi(string(num))
		pending = false
		num = nil

		if _, ok := winning[v]; ok {
			if points == 0 { // first winning nr counts as one
				points = 1
			} else { // ever winning nr after the first doubles points
				points *= 2
			}
		}
	}

	return points
}

func solvePartTwo(r io.Reader) (int, error) {
	// sc := bufio.NewScanner(r)
	//
	// i := 1
	// won := make(map[int]int)
	// for ; sc.Scan(); i++ {
	// 	computeMatches(sc.Text(), i, won)
	// }
	//
	// var sum int
	// for _, v := range won {
	// 	sum += v
	// }
	// return sum, nil
	return 0, nil
}

func computeMatches(line string, id int, won map[int]int) {
	// every card counts irrespective of whether it has any matches
	if _, ok := won[id]; !ok {
		won[id] = 1
	} else {
		won[id] = won[id] + 1
	}

	winning := make(map[int]struct{})

	var matches int
	var collect bool
	var inWin bool
	var pending bool
	var num []rune
	for _, c := range line {
		if c == ':' { // winning numbers are in between ':' and '|'
			collect = true
			inWin = true
		}
		if c == '|' { // numbers are after '|'
			inWin = false
		}

		if !collect {
			continue
		}

		if unicode.IsDigit(c) {
			pending = true
			num = append(num, c)
			continue
		}

		if pending {
			if num != nil {
				v, _ := strconv.Atoi(string(num))
				pending = false
				num = nil

				if inWin {
					winning[v] = struct{}{}
				} else {
					if _, ok := winning[v]; ok {
						matches += 1
					}
				}
			}
		}
	}
	// todo how to beautify this pattern
	if num != nil {
		v, _ := strconv.Atoi(string(num))
		pending = false
		num = nil

		if _, ok := winning[v]; ok {
			matches += 1
		}
	}

	// you win copies of the scratchcards below the winning card equal to the number of matches
	for i := id + 1; i < id+1+matches; i++ {
		if _, ok := won[i]; !ok {
			won[i] = won[id]
		} else {
			won[i] = won[i] + won[id]
		}
	}
	// fmt.Printf("id: %d, matches: %d, won: %v\n", id, matches, won)
}
