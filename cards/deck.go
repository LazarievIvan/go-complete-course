package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Clubs", "Hearts", "Spades", "Diamonds"}
	court_cards := []string{"Jack", "Queen", "King"}
	for _, suit := range suits {
		cards = append(cards, "Ace of "+suit)
		for i := 2; i <= 10; i++ {
			cards = append(cards, strconv.Itoa(i)+" of "+suit)
		}
		for _, card := range court_cards {
			cards = append(cards, card+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	dealtCards := d[:1]
	remainingCards := d[1:]
	for i := 1; i < handSize; i++ {
		dealtCards = append(dealtCards, remainingCards[:1]...)
		remainingCards = remainingCards[1:]
	}
	return dealtCards, remainingCards
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	randSource := rand.NewSource(seed)
	randomizer := rand.New(randSource)
	for i := range d {
		j := randomizer.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func takeDeckFromFile(filename string) deck {
	savedDeck, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	deck := deck(strings.Split(string(savedDeck), ","))
	return deck
}
