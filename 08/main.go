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

func testMove(nodes map[string]Node, currentNode Node, step byte) (bool, Node) {
	base.Debug("   Move %c from %s", step, currentNode.from)

	var nextNode Node
	switch step {
	case 'L':
		nextNode = nodes[currentNode.left]
		base.Debug("  to %s", nextNode.from)
	case 'R':
		nextNode = nodes[currentNode.right]
		base.Debug("     to %s", nextNode.from)
	default:
		panic("Unknown step:")
	}

	return nextNode.from[2] == 'Z', nextNode
}

func getStartingLocations2(nodes map[string]Node) []string {
	var startingLocations []string = make([]string, 0)
	for _, node := range nodes {
		if node.from[2] == 'A' {
			startingLocations = append(startingLocations, node.from)
		}
	}
	return startingLocations
}

func run1(input []string) int {
	count := 0
	steps := input[0] // Either L or R
	nodes := parseMap(input[2:])
	currentNode := nodes["AAA"]
	for {
		step := steps[count%len(steps)]
		count++
		if ok, nextNode := testMove(nodes, currentNode, step); ok {
			break
		} else {
			currentNode = nextNode
		}
	}
	fmt.Println(count)
	return count
}

func run2(input []string) int {
	steps := input[0] // Either L or R
	nodes := parseMap(input[2:])
	startingLocations := getStartingLocations2(nodes)
	base.Debug("Starting locations: %v", startingLocations)

	stepsToReachZ := make([]int, len(startingLocations))

	for idx, startingLocation := range startingLocations {
		count := 0
		location := startingLocation
		for {
			step := steps[count%len(steps)]
			count++
			if done, nextLocation := testMove(nodes, nodes[location], step); done {
				stepsToReachZ[idx] = count
				break
			} else {
				location = nextLocation.from
			}
		}
	}

	lcm := base.LCM(stepsToReachZ[0], stepsToReachZ[1], stepsToReachZ[2:]...)
	fmt.Println(lcm)
	return lcm
}

func main() {
	run1(base.ReadInputLines())
	run2(base.ReadInputLines())
}
