package main

import (
	"base"
	"fmt"
)

type Node struct {
	row      int
	startCol int
	endCol   int
	val      int
}

func parse(input []string) ([]Node, [][]string) {
	lineLength := len(input[0])

	nodes := make([]Node, 0)

	// Initialize symbol grid
	symbolGrid := make([][]string, len(input))
	for i := range symbolGrid {
		symbolGrid[i] = make([]string, lineLength)
	}

	// Extract all numbers
	for row, line := range input {
		// For each character in line
		var tmpNode Node
		var tmpVal string
		for column := 0; column < lineLength; column++ {
			char := line[column]

			//  Create a new node or add to in-progress node
			if char >= '0' && char <= '9' {
				if tmpVal != "" {
					// In progress node
					tmpNode.endCol = column
				} else {
					// New Node
					tmpNode = Node{row, column, column, 0}
				}
				tmpVal += string(char)
				continue
			}

			// Commit any in-progress nodes if we make it here
			if tmpVal != "" {
				tmpNode.val = base.AtoiOrPanic(tmpVal)
				nodes = append(nodes, tmpNode)
				tmpVal = ""
			}

			if char == '.' {
				// Not a symbol
				continue
			}

			// Its a symbol
			symbolGrid[row][column] = string(char)
		}

		// Commit any in-progress nodes that weren't committed (ie, end of line)
		if tmpVal != "" {
			tmpNode.val = base.AtoiOrPanic(tmpVal)
			nodes = append(nodes, tmpNode)
			tmpVal = ""
		}

	}
	return nodes, symbolGrid
}

func run1(input []string) int {

	nodes, symbolGrid := parse(input)

	// Print all nodes
	if base.IsDebug() {
		for _, node := range nodes {
			base.Debug("%v", node)
		}
		for _, row := range symbolGrid {
			base.Debug("%v", row)
		}
	}

	sum := 0

	for _, node := range nodes {
		var hasSymbol bool
		for row := node.row - 1; row <= node.row+1; row++ {
			for col := node.startCol - 1; col <= node.endCol+1; col++ {
				if row < 0 || row >= len(symbolGrid) {
					continue
				}
				if col < 0 || col >= len(symbolGrid[row]) {
					continue
				}
				if symbolGrid[row][col] != "" {
					hasSymbol = true
					base.Debug("Node '%v' has adjacent symbol '%v' at (%v,%v)", node, symbolGrid[row][col], row, col)
					sum += node.val
				}
			}
		}
		if !hasSymbol {
			base.Debug("Node '%v' has no symbols adjacent to it!", node)
		}
	}

	fmt.Println(sum)
	return sum
}

func main() {
	run1(base.ReadExample1Lines())
	run1(base.ReadInputLines())
	// run2(base.ReadInputLines())
}
