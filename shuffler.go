package poker

import (
	"math/rand"
	"time"
)

type Shuffler interface {
	Shuffle([]*Card) []*Card
}

type defaultShuffler struct{}

func (s defaultShuffler) Shuffle(cards []*Card) []*Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}
