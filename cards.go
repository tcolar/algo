// History: Dec 09 13 tcolar Creation

package algo

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck of playing cads
type CardDeck struct {
	Cards []string
}

// Build a 52 card deck (not sorted)
func CardDeck52() CardDeck {
	cards := []string{}
	for kind := 0; kind != 4; kind++ {
		nm := ""
		switch kind {
		case 0:
			nm = "♣ " //club
		case 1:
			nm = "♦ " //heart
		case 2:
			nm = "♥ " //spade
		case 3:
			nm = "♠ " // carreaux
		}
		for i := 1; i != 14; i++ {
			switch i {
			case 1:
				cards = append(cards, "A"+nm)
			case 11:
				cards = append(cards, "J"+nm)
			case 12:
				cards = append(cards, "Q"+nm)
			case 13:
				cards = append(cards, "K"+nm)
			default:
				cards = append(cards, fmt.Sprintf("%d", i)+nm)
			}
		}
	}
	return CardDeck{Cards: cards}
}

func (deck *CardDeck) Shuffle() {
	// If not seeded then it's deterministic (not random)
	rand.Seed(time.Now().UTC().UnixNano())
	sz := len(deck.Cards)
	for i := 0; i != sz; i++ {
		// pick a random card between i and end of deck (start at 0-51)
		rd := i + rand.Int()%(sz-i)
		// and swap it to current pos from front of deck
		deck.Cards[i], deck.Cards[rd] = deck.Cards[rd], deck.Cards[i]
		// continue on with whole deck
	}
	// Note: Maybe we could stop at 51 since last card would already be in place ?
}
