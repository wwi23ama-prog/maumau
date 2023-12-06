package maumau

// Detail-Datentypen für die Karten
// - Farbe
// - Wert

// Farbe
type Suit string

// Wert
type Rank string

// Konkrete Farben
const (
	Spades   Suit = "♠"
	Hearts   Suit = "♥"
	Diamonds Suit = "♦"
	Clubs    Suit = "♣"
)

// Konkrete Werte
const (
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
	Ace   Rank = "A"
)
