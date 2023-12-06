package maumau

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
