package maumau

import "fmt"

// ExampleNewDeck32 zeigt die Verwendung von NewDeck32.
func ExampleNewDeck32() {
	// Kartenstapel erzeugen
	deck := NewDeck32()

	// Anzahl der Karten ausgeben
	fmt.Println(len(deck.cards))

	// Alle Karten ausgeben
	for _, card := range deck.cards {
		fmt.Println(card)
	}

	// Output:
	// 32
	// ♠7
	// ♠8
	// ♠9
	// ♠10
	// ♠J
	// ♠Q
	// ♠K
	// ♠A
	// ♥7
	// ♥8
	// ♥9
	// ♥10
	// ♥J
	// ♥Q
	// ♥K
	// ♥A
	// ♦7
	// ♦8
	// ♦9
	// ♦10
	// ♦J
	// ♦Q
	// ♦K
	// ♦A
	// ♣7
	// ♣8
	// ♣9
	// ♣10
	// ♣J
	// ♣Q
	// ♣K
	// ♣A
}

// ExampleDeckShuffle zeigt die Verwendung von Deck.Shuffle().
func ExampleDeck_Shuffle() {
	// Kartenstapel erzeugen
	deck := NewDeck32()

	// Kartenstapel mischen
	deck.Shuffle()

	// Prüfen, ob die Kartenstapel sich an wenigstens 5 Stellen unterscheiden.
	counter := 0
	for i := range deck.cards {
		if deck.cards[i] != NewDeck32().cards[i] {
			counter++
		}
	}
	fmt.Println(counter >= 5)

	// Output:
	// true
}
