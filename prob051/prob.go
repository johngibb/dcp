package prob006

import "math/rand"

// This problem was asked by Facebook.
//
// Given a function that generates perfectly random numbers between 1 and k
// (inclusive), where k is an input, write a function that shuffles a deck of
// cards represented as an array using only swaps.
//
// It should run in O(N) time.
//
// Hint: Make sure each one of the 52! permutations of the deck is equally
// likely.

type Card int

var suits []string = []string{
	"Clubs",
	"Diamonds",
	"Hearts",
	"Spades",
}

var ranks []string = []string{
	"Ace",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"Jack",
	"Queen",
	"King",
}

func (c Card) Suit() string {
	return suits[c%4]
}

func (c Card) Rank() string {
	return ranks[c/4]
}

func Shuffle(deck []Card) {
	// Just fisher-yates shuffle.
	for i := len(deck) - 1; i >= 1; i-- {
		j := random(i+1) - 1
		deck[i], deck[j] = deck[j], deck[i]
	}
}

// random generates perfectly random numbers between 1 and k (inclusive).
func random(k int) int {
	return rand.Intn(k) + 1
}
