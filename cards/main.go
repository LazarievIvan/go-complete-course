package main

import "fmt"

func main() {

	cards := newDeck()
	cards.saveToFile("unshuffled_cards")
	cards.shuffle()
	// Dealing cards to the players.
	correctAmount := false
	var playersCount int
	var handSize int
	for !correctAmount {
		fmt.Print("\nNumber of players:")
		fmt.Scanln(&playersCount)
		fmt.Print("Hand size:")
		fmt.Scanln(&handSize)
		if len(cards) >= handSize*playersCount {
			correctAmount = true
		} else {
			fmt.Printf("Cards in deck: %d\nCards needed for %d players: %d\nPlease hange number of players or handsize", len(cards), playersCount, handSize*playersCount)
		}
	}
	playersHands := []deck{}
	for i := 0; i < playersCount; i++ {
		hand, remainingCards := deal(cards, handSize)
		playersHands = append(playersHands, hand)
		cards = remainingCards
	}
	for i, hand := range playersHands {
		fmt.Println()
		fmt.Printf("Player %d hand:\n", i+1)
		hand.print()
	}
	// After playing
	discardPile := deck{}
	for _, hand := range playersHands {
		discardPile = append(discardPile, hand...)
	}
	cards = append(cards, discardPile...)
	cards.saveToFile("used_deck")
	usedDeck := takeDeckFromFile("used_deck")
	fmt.Println("\nUsed deck:")
	usedDeck.print()
}
