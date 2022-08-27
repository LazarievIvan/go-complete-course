package main

import (
	"os"
	"reflect"
	"testing"
)

var deckSize int

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	deckSize = 52

	if len(deck) < deckSize {
		t.Errorf("Expected size of the deck %d, got %d", 52, len(deck))
	}
}

func TestDeal(t *testing.T) {

	deck := newDeck()
	handSize := 6
	hand, _ := deal(deck, handSize)
	if len(hand) != handSize {
		t.Errorf("Expected amount of cards in hand: %d, got %d", handSize, len(hand))
	}
}

// Check uniqueness of shuffled deck elements.
// @todo implement this later
func TestShuffle(t *testing.T) {
	deck := newDeck()
	deck.shuffle()
}

func TestSaveToFileAndTakeDeckFromFile(t *testing.T) {
	os.Remove(".testdeck")
	deck := newDeck()
	err := deck.saveToFile(".testdeck")
	if err != nil {
		t.Errorf("Something went wrong saving deck to file: %s", err)
	}
	deckFromFile := takeDeckFromFile(".testdeck")
	if !reflect.DeepEqual(deck, deckFromFile) {
		t.Errorf("Deck from file is not equal to the initial deck")
	}
}
