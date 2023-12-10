package main

import (
	"base"
	"fmt"
	"regexp"
)

type Node struct {
	from  string
	left  string
	right string
}

// AAA = (BBB, CCC)

var nodeRegex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)`)

func parseMap(input []string) map[string]Node {
	var m map[string]Node = make(map[string]Node, 0)

	for _, line := range input {
		matches := nodeRegex.FindStringSubmatch(line)
		if len(matches) == 0 {
			panic("The regex is broken! Input:   " + line)
		}

		from := matches[1]
		left := matches[2]
		right := matches[3]

		m[from] = Node{
			from:  from,
			left:  left,
			right: right,
		}
	}

	return m

}

func run(steps string, startingNode string, nodes map[string]Node) int {
	count := 0
	currentNode := nodes[startingNode]
	target := "ZZZ" // Part 1
	for {
		step := steps[count%len(steps)]
		count++
		if ok, nextNode := testMove(nodes, currentNode, step, target); ok {
			break
		} else {
			currentNode = nextNode
		}
	}
	return count
}

func testMove(nodes map[string]Node, currentNode Node, step byte, target string) (bool, Node) {
	base.Debug("Move %c from %s", step, currentNode.from)

	var nextNode Node
	switch step {
	case 'L':
		nextNode = nodes[currentNode.left]
		base.Debug("  to %s", nextNode.from)
	case 'R':
		nextNode = nodes[currentNode.right]
		base.Debug("  to %s", nextNode.from)
	default:
		panic("Unknown step:")
	}

	return nextNode.from == target, nextNode
}

func run1(input []string) int {
	count := 0
	steps := input[0] // Either L or R
	nodes := parseMap(input[2:])
	count = run(steps, "AAA", nodes)

	fmt.Println(count)
	return count
}

// func run2(input []string) int {
// }

func main() {
	// run1(base.ReadExample2Lines())
	run1(base.ReadInputLines()) // 17287

	// run2(base.ReadExample3Lines())
	// run2(base.ReadInputLines())
}
