package main

import (
	"encoding/csv"
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
	cr := csv.NewReader(r)
	seqs, err := cr.Read()
	if err != nil {
		return 0, err
	}
	var sum int
	for _, seq := range seqs {
		sum += hash(seq)
	}
	return sum, nil
}

func hash(in string) int {
	var current int
	for _, r := range in {
		current += int(r)
		current *= 17
		current %= 256
	}
	return current
}

// solvePartTwo solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	cr := csv.NewReader(r)
	seqs, err := cr.Read()
	if err != nil {
		return 0, err
	}

	boxes := make(map[int][]*lens)
	for _, seq := range seqs {
		label, focalLength, add := strings.Cut(seq, "=")
		if add {
			boxId := hash(label)
			val, err := strconv.Atoi(focalLength)
			if err != nil {
				return 0, err
			}

			lenses, ok := boxes[boxId]
			l := &lens{
				label:       label,
				focalLength: val,
			}
			if !ok {
				boxes[boxId] = []*lens{l}
				continue
			}
			id := slices.IndexFunc(lenses, func(l *lens) bool {
				if l.label == label {
					return true
				}
				return false
			})

			if id != -1 { // replace existing lens
				lenses[id] = l
			} else { // add lens
				boxes[boxId] = append(boxes[boxId], l)
			}
		} else {
			label, _, _ := strings.Cut(seq, "-")
			boxId := hash(label)
			lenses, ok := boxes[boxId]
			if !ok {
				continue
			}
			id := slices.IndexFunc(lenses, func(l *lens) bool {
				if l.label == label {
					return true
				}
				return false
			})

			if id != -1 { // remove existing lens and shift other lenses left
				boxes[boxId] = slices.Delete(lenses, id, id+1)
			}
		}
	}

	var sum int
	for i, lenses := range boxes {
		for j, lens := range lenses {
			val := (i + 1) * (j + 1) * lens.focalLength
			sum += val
		}
	}
	return sum, nil
}

type lens struct {
	label       string
	focalLength int
}
