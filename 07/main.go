package main

import (
	"base"
	"fmt"
	"sort"
	"strings"
)

type Hand struct {
	cards    []Card
	Strength Strength
	bid      int
}

type Card int

const (
	_2 Card = iota
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	T
	J
	Q
	K
	A // Highest
)

type Strength int

const (
	HighCard Strength = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind // Highest
)

func (d Strength) String() string {
	return []string{
		"HighCard", "OnePair", "TwoPair", "ThreeOfAKind",
		"FullHouse", "FourOfAKind", "FiveOfAKind"}[d]
}

func (d Card) String() string {
	return []string{
		"2", "3", "4", "5",
		"6", "7", "8", "9", "T",
		"J", "Q", "K", "A"}[d]
}

var MapStringToCard = func() map[string]Card {
	m := make(map[string]Card)
	for c := _2; c <= A; c++ {
		m[c.String()] = c
	}
	return m
}()

func (d Card) Value() int {
	return int(d)
}

func parseHand(hand string) []Card {
	cards := make([]Card, 0)
	for i := 0; i < len(hand); i++ {
		cards = append(cards, MapStringToCard[hand[i:i+1]])
	}
	return cards
}

func parseLine(line string) ([]Card, int) {
	split := strings.Split(line, " ")
	return parseHand(split[0]), base.AtoiOrPanic(split[1])
}

func computeHandStrength(hand []Card) Strength {
	if len(hand) != 5 {
		panic("Invalid hand")
	}
	counts := make(map[Card]int)
	uniqueCards := 0
	for _, card := range hand {

		if _, ok := counts[card]; !ok {
			uniqueCards++
		}

		counts[card]++

		if counts[card] == 5 {
			return FiveOfAKind
		}
	}

	if uniqueCards == 2 {
		for _, count := range counts {
			if count == 4 {
				return FourOfAKind
			}
			if count == 3 {
				return FullHouse
			}
		}
	}

	if uniqueCards == 3 {
		for _, count := range counts {
			if count == 3 {
				return ThreeOfAKind
			}
			if count == 2 {
				return TwoPair
			}
		}
	}

	if uniqueCards == 4 {
		return OnePair
	}

	return HighCard
}

func run1(input []string) int {
	result := 0

	hands := make([]Hand, 0)
	for _, line := range input {
		hand, bid := parseLine(line)
		strength := computeHandStrength(hand)
		// base.Debug("%v (%v) -> %v", hand, bid, strength)
		hands = append(hands, Hand{hand, strength, bid})
	}

	// Map of strengths to hand in stregenth order
	strengths := []Strength{
		FiveOfAKind,
		FourOfAKind,
		FullHouse,
		ThreeOfAKind,
		TwoPair,
		OnePair,
		HighCard,
	}

	// Sort by strength into buckets
	buckets := make(map[Strength][]Hand)
	for _, hand := range hands {
		buckets[hand.Strength] = append(buckets[hand.Strength], hand)
	}

	for _, s := range strengths {
		bucket := buckets[s]
		sort.SliceStable(bucket, func(l, r int) bool {
			left := bucket[l]
			right := bucket[r]
			for k := 0; k < 5; k++ {
				if left.cards[k] != right.cards[k] {
					return left.cards[k] > right.cards[k]
				}
			}
			return false
		})
	}

	// Splice the buckets together
	sorted := make([]Hand, 0)
	for _, s := range strengths {
		sorted = append(sorted, buckets[s]...)
	}

	base.Debug("Sorted:")
	for idx, hand := range sorted {
		rank := len(sorted) - idx
		base.Debug("[%d] %v (%v) -> %s", rank, hand.cards, hand.bid, hand.Strength)
		result += hand.bid * rank
	}

	fmt.Println(result)
	return result
}

func main() {
	// run1(base.ReadExample1Lines())
	run1(base.ReadInputLines()) // 248569531

	// run2(base.ReadInputLines())
}
