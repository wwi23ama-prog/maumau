package maumau

// CardsMatch vergleicht zwei Karten daraufhin,
// ob sie nach den Regeln des Mau-Mau-Spiels
// zusammenpassen.
func CardsMatch(c1 Card, c2 Card) bool {
	return c1.suit == c2.suit || c1.rank == c2.rank
}
