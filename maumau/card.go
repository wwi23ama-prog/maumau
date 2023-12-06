package maumau

// Card repräsentiert eine Spielkarte.
type Card struct {
	suit Suit
	rank Rank
}

// String gibt eine Karte als String zurück.
func (c Card) String() string {
	return string(c.suit) + string(c.rank)
}
