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
	s := bufio.NewScanner(r)
	var rows, cols int
	var grid [][]rune
	for s.Scan() {
		line := s.Text()
		cols = len(line)
		row := make([]rune, cols)
		grid = append(grid, row)
		for i, r := range line {
			row[i] = r
		}
		rows++
	}

	var start int
	graph := make([]*node, rows*cols)
	for rowID, row := range grid {
		for colID, col := range row {
			nodeID := toNodeID(cols, rowID, colID)
			currentNode := graph[nodeID]
			if currentNode == nil {
				currentNode = &node{neighbors: make(map[int]struct{}), symbol: col}
				graph[nodeID] = currentNode
			} else {
				currentNode.symbol = col
			}

			fmt.Println("rowID", rowID, "colID", colID, string(col), "nodeID", nodeID)
			switch col {
			case 'S':
				start = nodeID
				continue
			case '-': // node connects left and right nodes
				if colID-1 >= 0 { // left
					neighborNodeID := toNodeID(cols, rowID, colID-1)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
				if colID+1 < cols { // right
					neighborNodeID := toNodeID(cols, rowID, colID+1)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
			case '|': // node connects up and down nodes
				if rowID-1 >= 0 { // up
					neighborNodeID := toNodeID(cols, rowID-1, colID)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
				if rowID+1 < rows { // down
					neighborNodeID := toNodeID(cols, rowID+1, colID)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
			case '7': // node connects south and west
				if colID-1 >= 0 { // left
					fmt.Println("7 going left")
					neighborNodeID := toNodeID(cols, rowID, colID-1)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					// neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
				if rowID+1 < rows { // down
					neighborNodeID := toNodeID(cols, rowID+1, colID)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
			case 'F': // node connects south and east
				if colID+1 < cols { // right
					neighborNodeID := toNodeID(cols, rowID, colID+1)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
				if rowID+1 < rows { // down
					neighborNodeID := toNodeID(cols, rowID+1, colID)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
			case 'J': // node connects up and left nodes
				if colID-1 >= 0 { // left
					neighborNodeID := toNodeID(cols, rowID, colID-1)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
				if rowID-1 >= 0 { // up
					neighborNodeID := toNodeID(cols, rowID-1, colID)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
			case 'L': // node connects north and east
				if colID+1 < cols { // right
					neighborNodeID := toNodeID(cols, rowID, colID+1)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
				if rowID-1 >= 0 { // up
					neighborNodeID := toNodeID(cols, rowID-1, colID)
					neighbor := graph[neighborNodeID]
					if neighbor == nil {
						neighbor = &node{neighbors: make(map[int]struct{})}
						graph[neighborNodeID] = neighbor
					}
					neighbor.neighbors[nodeID] = struct{}{}
					currentNode.neighbors[neighborNodeID] = struct{}{}
				}
			}
		}
	}

	for i, n := range graph {
		if n.symbol == '.' {
			n.neighbors = nil
		}
		if n != nil {
			fmt.Println("i", i, "symbol", string(n.symbol), "neighbors", n.neighbors)
		} else {
			fmt.Println("i", i, "symbol", string(n.symbol), "no neighbors")
		}
	}

	c := cycle{}
	result := c.bfs(graph, start)
	fmt.Println("start", start)
	fmt.Println(graph[start])

	return result, nil
}

type cycle struct {
	path []int
}

// 330 is too low
func (c *cycle) bfs(graph []*node, source int) int {
	var maxDistance int
	marked := make([]bool, len(graph))
	distanceTo := make([]int, len(graph))
	for i := 0; i < len(distanceTo); i++ {
		distanceTo[i] = -1
	}
	distanceTo[source] = 0
	marked[source] = true
	queue := []int{source}

	for len(queue) != 0 {
		// dequeue
		// fmt.Println("queue before dequeu", queue)
		v := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = nil
		}
		// fmt.Println("queue after dequeu", queue)

		for w := range graph[v].neighbors {
			if marked[w] {
				continue
			}

			distanceTo[w] = distanceTo[v] + 1
			maxDistance = max(maxDistance, distanceTo[w])
			marked[w] = true
			queue = append(queue, w)
			// fmt.Println("queue", queue, "distanceTo", distanceTo)
		}
	}

	fmt.Println("distanceTo", distanceTo)
	fmt.Println("maxDistance", maxDistance)
	return maxDistance
}

func (c *cycle) dfs(graph []*node, marked []bool, edgeTo []int, u, v int) {
	marked[v] = true
	for w := range graph[v].neighbors {
		if c.path != nil {
			// cycle already found
			return
		}

		if !marked[w] {
			edgeTo[w] = v
			c.dfs(graph, marked, edgeTo, v, w)
		} else if w != u {
			// check for path (but disregard reverse of edge leading to v)
			var path []int
			for x := v; x != w; x = edgeTo[x] {
				path = append(path, x)
			}
			path = append(path, w)
			path = append(path, v)
			c.path = path
		}
	}
}

type node struct {
	neighbors map[int]struct{}
	symbol    rune
}

func toNodeID(cols, rowID, colID int) int {
	return rowID*cols + colID
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
