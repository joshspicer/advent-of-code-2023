package main

import (
	"base"
	"fmt"
	"os"
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
	J Card = iota // (Part 2 Only)
	_2
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	T
	// J (Part 1 Only)
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

// func (d Card) String() string {
// 	return []string{
// 		"2", "3", "4", "5",
// 		"6", "7", "8", "9", "T",
// 		"J", "Q", "K", "A"}[d]
// }

func (d Card) String() string {
	return []string{
		"J", "2", "3", "4", "5",
		"6", "7", "8", "9", "T",
		"Q", "K", "A"}[d]
}

var MapStringToCard = func() map[string]Card {
	m := make(map[string]Card)
	// for c := _2; c <= A; c++ {
	for c := J; c <= A; c++ {
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

	// If 5 jokers, exit early
	for idx, c := range hand {
		if c != J {
			break
		}
		if idx == len(hand)-1 {
			return FiveOfAKind
		}
	}

	piles := make([][]Card, 0)
	pileName := []Card{_2, _3, _4, _5, _6, _7, _8, _9, T, Q, K, A}
	wildCards := 0
	for _, t := range pileName {
		// Grab every card of this type
		pile := make([]Card, 0)
		for _, c := range hand {
			if c == t {
				pile = append(pile, c)
			}
		}
		if len(pile) > 0 {
			piles = append(piles, pile)
		}
	}

	for _, c := range hand {
		if c == J {
			wildCards++
		}
	}

	// Sort the piles by size
	sort.SliceStable(piles, func(l, r int) bool {
		return len(piles[l]) > len(piles[r])
	})

	//If there's wildcards, add to the largest group until 5 card, then next group...
	for wildCards > 0 {
		for idx, pile := range piles {
			pileLength := len(pile)
			for pileLength < 5 && wildCards > 0 {
				piles[idx] = append(piles[idx], J)
				wildCards--
				pileLength++
			}
		}
	}

	base.Debug("Piles: %v", piles)

	if len(piles[0]) == 5 {
		return FiveOfAKind
	}

	if len(piles[0]) == 4 {
		return FourOfAKind
	}

	if len(piles[0]) == 3 {
		if len(piles[1]) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}

	if len(piles[0]) == 2 {
		if len(piles[1]) == 2 {
			return TwoPair
		}
		return OnePair
	}

	return HighCard
}

func run(input []string) int {
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

	// If argument present, pass to computeHandStrength
	if len(os.Args) > 1 {
		hand := parseHand(os.Args[1])
		strength := computeHandStrength(hand)
		base.Debug("%v", strength)
		return
	}

	// run(base.ReadExample1Lines()) // Part 1: 6440, Part 2: 5905
	run(base.ReadInputLines()) // Part 1: 248569531, Part 2: 250382098
}
