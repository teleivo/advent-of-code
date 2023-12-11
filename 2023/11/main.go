package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	grid := parseAndExpandGrid(r)
	galaxies := enumerateGalaxies(grid)
	fmt.Println("galaxies", galaxies)

	return 0, nil
}

func parseAndExpandGrid(r io.Reader) [][]rune {
	s := bufio.NewScanner(r)
	var grid [][]rune
	var galaxyCount int
	var rows, cols int
	rowsWithoutGalaxy := make(map[int]struct{})
	colsWithGalaxy := make(map[int]struct{})
	for s.Scan() {
		line := s.Text()
		cols = len(line)
		var rowHasGalaxy bool
		var row []rune
		for col, r := range line {
			if r == '#' {
				galaxyCount++
				rowHasGalaxy = true
				colsWithGalaxy[col] = struct{}{}
				_ = col
			}
			row = append(row, r)
		}
		if !rowHasGalaxy {
			rowsWithoutGalaxy[rows] = struct{}{}
		}

		grid = append(grid, row)
		rows++
	}

	fmt.Println("unexpanded rows", rows, "cols", cols)
	colsWithoutGalaxy := make(map[int]struct{})
	for i := 0; i < cols; i++ {
		if _, ok := colsWithGalaxy[i]; !ok {
			colsWithoutGalaxy[i] = struct{}{}
		}
	}

	var expandedRows [][]rune
	for i := 0; i < rows; i++ {
		expandedRows = append(expandedRows, grid[i])
		if _, ok := rowsWithoutGalaxy[i]; ok {
			expandedRows = append(expandedRows, grid[i])
		}
	}

	var expandedGrid [][]rune
	for i := 0; i < len(expandedRows); i++ {
		expandedRow := expandedRows[i]
		var row []rune
		for j := 0; j < len(expandedRow); j++ {
			row = append(row, expandedRow[j])
			if _, ok := colsWithGalaxy[j]; !ok {
				row = append(row, expandedRow[j])
			}
		}
		expandedGrid = append(expandedGrid, row)
	}

	return expandedGrid
}

func enumerateGalaxies(grid [][]rune) map[int][2]int {
	// map of x,y coordinates
	var galaxyCount int
	galaxies := make(map[int][2]int)
	for i, row := range grid {
		for j, r := range row {
			if r == '#' {
				galaxies[galaxyCount] = [2]int{i, j}
				galaxyCount++
			}
		}
	}

	return galaxies
}

// func bfs(grid [][]rune, galaxies map[int][2]int) {
// 	marked := make([]bool, len(graph))
// 	distanceTo := make([]int, len(graph))
// 	for i := 0; i < len(distanceTo); i++ {
// 		distanceTo[i] = -1
// 	}
// 	distanceTo[source] = 0
// 	marked[source] = true
// 	queue := []int{source}
//
// 	for len(queue) != 0 {
// 		// dequeue
// 		// fmt.Println("queue before dequeu", queue)
// 		v := queue[0]
// 		if len(queue) > 1 {
// 			queue = queue[1:]
// 		} else {
// 			queue = nil
// 		}
// 		// fmt.Println("queue after dequeu", queue)
//
// 		for w := range graph[v].neighbors {
// 			if marked[w] {
// 				continue
// 			}
//
// 			distanceTo[w] = distanceTo[v] + 1
// 			marked[w] = true
// 			queue = append(queue, w)
// 			// fmt.Println("queue", queue, "distanceTo", distanceTo)
// 		}
// 	}
//
// 	fmt.Println("distanceTo", distanceTo)
// }

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		_ = line
	}
	return 0, nil
}
