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
	// fmt.Println("galaxies", galaxies)
	distances := bfs(grid, galaxies)
	sum := sumDistances(galaxies, distances)
	return sum, nil
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

	// fmt.Println("unexpanded rows", rows, "cols", cols)
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

func bfs(grid [][]rune, galaxies map[int][2]int) map[int][][]int {
	rows := len(grid)
	cols := len(grid[0])
	distances := make(map[int][][]int)

	for galaxy, source := range galaxies {
		distance := make([][]int, rows)
		visited := make([][]bool, rows)

		sRow, sCol := source[0], source[1]
		if distance[sRow] == nil {
			distance[sRow] = make([]int, cols)
		}
		distance[sRow][sCol] = 0

		if visited[sRow] == nil {
			visited[sRow] = make([]bool, cols)
		}
		visited[sRow][sCol] = true
		queue := [][2]int{source}

		for len(queue) != 0 {
			// dequeue
			// fmt.Println("queue before dequeu", queue)
			v := queue[0]
			vRow, vCol := v[0], v[1]
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				queue = nil
			}
			// fmt.Println("queue after dequeu", queue)

			for _, w := range neighbors(rows, cols, v) {
				wRow, wCol := w[0], w[1]
				// fmt.Println("rows", rows, "cols", cols, "v", v, "neighbor", w)

				if visited[wRow] == nil {
					visited[wRow] = make([]bool, cols)
				}
				if visited[wRow][wCol] {
					continue
				}
				visited[wRow][wCol] = true

				if distance[wRow] == nil {
					distance[wRow] = make([]int, cols)
				}
				distance[wRow][wCol] = distance[vRow][vCol] + 1
				queue = append(queue, w)
				// fmt.Println("queue", queue, "distanceTo", distanceTo)
			}
		}
		distances[galaxy] = distance
	}

	return distances
}

func neighbors(rows, cols int, cell [2]int) [][2]int {
	moves := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var legalMoves [][2]int
	for _, move := range moves {
		nextRow := cell[0] + move[0]
		nextCol := cell[1] + move[1]

		if nextCol >= 0 && nextCol < cols && nextRow >= 0 && nextRow < rows {
			legalMoves = append(legalMoves, [2]int{nextRow, nextCol})
		}
	}
	return legalMoves
}

func sumDistances(galaxies map[int][2]int, distances map[int][][]int) int {
	var sum int

	for i := range galaxies {
		for j := range galaxies {
			if j < i+1 {
				continue
			}
			// fmt.Printf("galaxy %d to %d\n", i, j)
			distance := distances[i]
			destination := galaxies[j]
			dRow, dCol := destination[0], destination[1]
			sum += distance[dRow][dCol]
		}
	}

	return sum
}

// solvePartOne solves part two of the puzzle.
func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		_ = line
	}
	return 0, nil
}
