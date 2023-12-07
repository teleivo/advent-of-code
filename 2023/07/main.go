package main

import (
	"bufio"
	"errors"
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
	hands, err := parseHands(r)
	if err != nil {
		return 0, err
	}
	fmt.Println(hands)

	types := make(map[handType][]hand)
	for _, hand := range hands {
		t := categorizeHand(hand.Hand)
		fmt.Println("hand", hand.Hand, "type", t)
		// if _, ok := types[t]; !ok {
		types[t] = append(types[t], hand)
		// } else {
		// 	types
		// }
	}
	fmt.Println(types)

	return 0, nil
}

type card int

const (
	A card = iota + 10
	J
	K
	Q
	T
)

type handType int

const (
	high handType = iota
	one
	two
	three
	fullHouse
	four
	five
)

type hand struct {
	Hand string
	Bid  int
}

func parseHands(r io.Reader) ([]hand, error) {
	s := bufio.NewScanner(r)
	var hands []hand
	for s.Scan() {
		hd, bid, found := strings.Cut(s.Text(), " ")
		if !found {
			return nil, errors.New(fmt.Sprintf("failed to split %q", s.Text()))
		}
		v, err := strconv.Atoi(bid)
		if !found {
			return nil, fmt.Errorf("failed to parse bid: %v", err)
		}
		hands = append(hands, hand{Hand: hd, Bid: v})
	}
	return hands, nil
}

func categorizeHand(hand string) handType {
	frequencies := cardFrequencies(hand)
	var triple bool
	var pairs int
	for _, freq := range frequencies {
		if freq == 5 {
			return five
		}
		if freq == 4 {
			return four
		}
		if freq == 3 {
			triple = true
			continue
		}
		if freq == 2 {
			pairs++
		}
	}
	if triple && pairs == 1 {
		return fullHouse
	}
	if triple {
		return three
	}
	if pairs == 2 {
		return two
	}
	if pairs == 1 {
		return one
	}
	return high
}

func cardFrequencies(hand string) map[rune]int {
	result := make(map[rune]int)
	for _, card := range hand {
		if _, ok := result[card]; !ok {
			result[card] = 1
			continue
		}
		result[card]++
	}
	return result
}

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	return 0, nil
}
