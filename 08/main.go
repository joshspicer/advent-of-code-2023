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

func run1(input []string) int {
	count := 0

	steps := input[0] // Either L or R
	nodes := parseMap(input[2:])

	currentNode := nodes["AAA"]
	target := "ZZZ"
	for {
		step := steps[count%len(steps)]
		count++
		switch step {
		case 'L':
			currentNode = nodes[currentNode.left]
			base.Debug("Left to %s", currentNode.from)
		case 'R':
			currentNode = nodes[currentNode.right]
			base.Debug("Right to %s", currentNode.from)
		default:
			panic("Unknown step:")
		}

		if currentNode.from == target {
			break
		}
	}

	fmt.Println(count)
	return count

}

func main() {
	// run1(base.ReadExample2Lines())
	run1(base.ReadInputLines())
	// run2(base.ReadInputLines())
}
