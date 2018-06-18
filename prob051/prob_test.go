package prob006

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var deck []Card
	for i := 0; i < 52; i++ {
		deck = append(deck, Card(i))
	}
	Shuffle(deck)
	for _, c := range deck {
		fmt.Printf("%s of %s\n", c.Rank(), c.Suit())
	}
}

func TestRandom(t *testing.T) {
	const k = 1000000
	seen := map[int]bool{}
	for i := 0; i < k; i++ {
		seen[random(52)] = true
	}

	// Make sure every number is present.
	for i := 1; i <= 52; i++ {
		if !seen[i] {
			t.Errorf("missing %d", i)
		}
	}

	// Make sure every number is within [1, k].
	for k := range seen {
		if k < 1 || k > 52 {
			t.Errorf("got %d, want within [1, 52]", k)
		}
	}
}
