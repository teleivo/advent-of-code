package main

import (
	"bufio"
	"bytes"
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

	f1, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}
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
func solvePartOne(b []byte) (int, error) {
	rows := bytes.Fields(b)
	for _, row := range rows {
		for _, v := range row {
			fmt.Println(string(v))
		}
	}
	return 0, nil
}

func findArrangements(springs []byte, damagedGroups []int) int {
	// if len(damagedGroups) == 0 {
	// 	return 0
	// }
	// if len(damagedGroups) == 0 {
	// 	return 0
	// }

	arrangements := 1
	groupIdx := 0
	group := damagedGroups[groupIdx]
	startIdx := 0
	endIdx := 0
	var unknowns int
	var damaged int
	fmt.Println()
	fmt.Println("springs", string(springs), "groups", damagedGroups)
	for endIdx < len(springs) {
		if springs[endIdx] == '?' {
			unknowns++
		} else if springs[endIdx] == '#' {
			damaged++
		}

		fmt.Println("unknowns", unknowns, "damaged", damaged, "group", group)
		if (springs[endIdx] == '.' || springs[endIdx] == '?') && group == damaged {
			fmt.Println("easy group consumed")
			// ? could not be another damaged spring as this would mean that the group count is
			// wrong
			groupIdx++
			group = damagedGroups[groupIdx]
			unknowns = 0
			damaged = 0
			startIdx = endIdx + 1 // consume this current spring
		} else if springs[endIdx] == '.' && unknowns != 0 {
			fmt.Println("we have", unknowns, "unknowns in substring", startIdx, ":", endIdx+1, "=", string(springs[startIdx:endIdx+1]))
			// TODO compute permutations on substring using all possible remaining groups; greedy

			// TODO there are 2 cases right? what if there are not all consecutive ?
			in := strings.Trim(string(springs[startIdx:endIdx]), ".")
			fmt.Println("working on in", in, "and group", groupIdx)
			if !strings.ContainsRune(in, '#') {
				// assuming all ? in the in string
				width := len(in)
				var sum int
				var containedGroups []int
				// assuming at least one group should still fit
				for i := groupIdx; i < len(damagedGroups); i++ {
					sum += damagedGroups[i] + 1
					if sum <= width {
						containedGroups = append(containedGroups, i)
					} else {
						break
					}
				}
				fmt.Println("groups", containedGroups)
			}

			unknowns = 0
			damaged = 0
			startIdx = endIdx + 1 // consume this current spring
		}

		endIdx++
	}

	// TODO handle the ???? case here as well

	return arrangements
}

func consumeGroups(width int, group []int) {

}

// solvePartTwo solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		_ = line
	}
	return 0, nil
}
