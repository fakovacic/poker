package poker

import (
	"math/rand"
	"time"
)

type Deck interface {
	Shuffle()
	Cards() []*Card
	Switch([]*Card)
	Draw() *Card
	Discard()
}

func NewDeck() Deck {
	cards := make([]*Card, maxDeckCards)

	i := 0

	for s := range suites {
		for r := range ranksList {
			cards[i] = &Card{
				Suite: suites[s],
				Rank:  ranksList[r],
			}
			i++
		}
	}

	return &deck{
		cards: cards,
	}
}

type deck struct {
	cards []*Card
}

func (s *deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.cards), func(i, j int) {
		s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
	})
}

func (s *deck) Switch(cards []*Card) {
	for i := range cards {
		s.cards[i] = cards[i]
	}

	s.cards = s.cards[:len(cards)]
}

func (s *deck) Cards() []*Card {
	return s.cards
}

func (s *deck) Draw() *Card {
	if len(s.cards) == 0 {
		return nil
	}

	card := s.cards[0]
	s.cards = s.cards[1:]

	return card
}

func (s *deck) Discard() {
	s.cards = s.cards[1:]
}
