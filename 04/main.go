package main

import (
	"base"
	"fmt"
	"regexp"
)

type Ticket struct {
	winningNums []int
	myNums      []int
	cardNum     int
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
var ticketRegex = regexp.MustCompile(`Card( )*(?P<cardNum>(\d+)):( )*(?P<winningNums>(\d+( )*)+)\|( )*(?P<myNums>(\d+( )*)+)`)

func parse(input []string) map[int]Ticket {
	results := make(map[int]Ticket, len(input))
	for _, line := range input {
		matches := ticketRegex.FindStringSubmatch(line)
		if len(matches) == 0 {
			panic("The regex is broken! Input:   " + line)
		}

		cardNum := 0
		winningNums := []int{}
		myNums := []int{}
		for i, name := range ticketRegex.SubexpNames() {
			if i != 0 && name != "" {

				if name == "winningNums" {
					// Convert to array of ints
					tmp := base.FilterWhitespace(matches[i])
					winningNums = base.Map(tmp, func(s string) int {
						return base.AtoiOrPanic(s)
					})
					continue
				}

				if name == "myNums" {
					// Convert to array of ints
					tmp := base.FilterWhitespace(matches[i])
					myNums = base.Map(tmp, func(s string) int {
						return base.AtoiOrPanic(s)
					})
					continue
				}

				if name == "cardNum" {
					cardNum = base.AtoiOrPanic(matches[i])
					continue
				}

				panic("Unknown name: " + name)
			}
		}

		results[cardNum] = Ticket{
			winningNums: winningNums,
			myNums:      myNums,
			cardNum:     cardNum,
		}
	}
	base.Debug("%v", results)
	return results
}

func countIntersection(a []int, b []int) int {
	count := 0
	for _, aVal := range a {
		for _, bVal := range b {
			if aVal == bVal {
				count++
			}
		}
	}
	return count
}

func run1(input []string) int {
	tickets := parse(input)

	total := 0
	for key, ticket := range tickets {
		numsMatching := countIntersection(ticket.winningNums, ticket.myNums)
		base.Debug("Ticket '%d' has %d matching numbers!", key, numsMatching)
		if numsMatching == 0 {
			continue
		}
		// The first match makes the card worth one point and each match after the first doubles the point value of that card.
		total += 1 << (numsMatching - 1)
	}

	fmt.Println(total)
	return total
}

func run2(input []string) int {
	originalTicketMap := parse(input)
	memo := make(map[int]int, len(originalTicketMap))

	numTicketsProcessed := 0
	ticketStack := []Ticket{}

	// Add all original tickets to the stack
	for _, ticket := range originalTicketMap {
		ticketStack = append(ticketStack, ticket)
	}

	for len(ticketStack) > 0 {
		numTicketsProcessed++

		// Pop the top ticket off the stack
		ticket := ticketStack[0]
		ticketStack = ticketStack[1:]

		// Count how many numbers match
		var numsMatching int
		if _, ok := memo[ticket.cardNum]; ok {
			// Reuse our past work
			numsMatching = memo[ticket.cardNum]
		} else {
			numsMatching = countIntersection(ticket.winningNums, ticket.myNums)
			memo[ticket.cardNum] = numsMatching
		}

		if numsMatching == 0 {
			continue
		}

		ticketNumber := ticket.cardNum
		for i := 1; i <= numsMatching; i++ {
			// Add a new ticket to the stack for each matching number
			ticketStack = append(ticketStack, originalTicketMap[ticketNumber+i])
		}
	}

	fmt.Println(numTicketsProcessed)
	return numTicketsProcessed
}

func main() {
	run1(base.ReadInputLines())
	run2(base.ReadInputLines())
}
