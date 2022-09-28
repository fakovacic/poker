package poker

type Deck interface {
	SetShuffler(Shuffler)
	Shuffle()
	Cards() []*Card
	Switch([]*Card)
	Draw() *Card
	Discard()
	RemoveCards([]*Card)
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
		shuffler: &defaultShuffler{},
		cards:    cards,
	}
}

type deck struct {
	shuffler Shuffler
	cards    []*Card
}

func (s *deck) SetShuffler(sh Shuffler) {
	s.shuffler = sh
}

func (s *deck) Shuffle() {
	s.cards = s.shuffler.Shuffle(s.cards)
}

func (s *deck) Switch(cards []*Card) {
	copy(s.cards, cards)
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

func (s *deck) RemoveCards(removeCards []*Card) {
	removeCount := len(removeCards)
	cards := make([]*Card, (maxDeckCards - removeCount))

	removeCardsMap := make(map[string]struct{}, removeCount)
	for x := range removeCards {
		removeCardsMap[removeCards[x].String()] = struct{}{}
	}

	i := 0

	for s := range suites {
		for r := range ranksList {
			c := &Card{
				Suite: suites[s],
				Rank:  ranksList[r],
			}

			_, ok := removeCardsMap[c.String()]
			if ok {
				continue
			}

			cards[i] = c
			i++
		}
	}

	s.cards = removeCards
}
