package maumau

import "math/rand"

// Deck repräsentiert einen Kartenstapel.
type Deck struct {
	cards []Card
}

// NewDeck32 erzeugt einen neuen Kartenstapel mit 32 Karten (Skatblatt).
func NewDeck32() Deck {
	deck := Deck{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for _, rank := range []Rank{Seven, Eight, Nine, Ten, Jack, Queen, King, Ace} {
			deck.cards = append(deck.cards, Card{suit: suit, rank: rank})
		}
	}
	return deck
}

// Shuffle mischt den Kartenstapel.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	// Falls der Kartenstapel sich an weniger als 5 Positionen unterscheidet, nochmal mischen.
	counter := 0
	unshuffledeck := NewDeck32()
	for i := range d.cards {
		if d.cards[i] != unshuffledeck.cards[i] {
			counter++
		}
	}
	if counter < 5 {
		d.Shuffle()
	}
}
