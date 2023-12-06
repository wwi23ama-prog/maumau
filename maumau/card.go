package maumau

import "fmt"

// Card repräsentiert eine Spielkarte.
type Card struct {
	suit Suit
	rank Rank
}

// String gibt eine Karte als String zurück.
func (c Card) String() string {
	return string(c.suit) + string(c.rank)
}

// Picture gibt einen String zurück, der die Karte als
// Unicode-Grafik darstellt.
//
// Beispiel:
// ┌───────┐
// │♠7     │
// │       │
// │       │
// │     ♠7│
// └───────┘
func (c Card) Picture() string {
	topline := "┌───────┐"
	botline := "└───────┘"
	midline := "│       │"

	symbolleft := fmt.Sprintf("│%s%s     │", c.suit, c.rank)
	symbolright := fmt.Sprintf("│     %s%s│", c.suit, c.rank)

	// Sonderfall: Bei 10 sind es zwei Zeichen, daher muss die Zeile angepasst werden.
	if c.rank == Ten {
		symbolleft = fmt.Sprintf("│%s%s    │", c.suit, c.rank)
		symbolright = fmt.Sprintf("│    %s%s│", c.suit, c.rank)
	}

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", topline, symbolleft, midline, midline, symbolright, botline)
}
