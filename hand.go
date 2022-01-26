package poker

type Hand interface {
	Deal(*Card)
	Cards() []*Card
}

func NewHand() Hand {
	cards := make([]*Card, 0)

	return &hand{
		cards: cards,
	}
}

type hand struct {
	cards []*Card
}

func (s *hand) Deal(card *Card) {
	s.cards = append(s.cards, card)
}

func (s *hand) Cards() []*Card {
	return s.cards
}
