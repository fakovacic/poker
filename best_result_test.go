package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
	"github.com/matryer/is"
)

func TestBestResult(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult string
		expectedCards  []int64
	}{
		{
			it: "high-card",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.HighCard,
			expectedCards: []int64{
				4,
			},
		},
		{
			it: "pair",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.Pair,
			expectedCards: []int64{
				2, 3,
			},
		},
		{
			it: "two-pairs",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.TwoPairs,
			expectedCards: []int64{
				0, 1, 2, 3,
			},
		},
		{
			it: "three-of-a-kind",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.ThreeOfAKind,
			expectedCards: []int64{
				0, 1, 2,
			},
		},
		{
			it: "straight",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Three,
				},
			},
			expectedResult: poker.Straight,
			expectedCards: []int64{
				0, 1, 2, 3, 4,
			},
		},
		{
			it: "flush",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Three,
				},
			},
			expectedResult: poker.Flush,
			expectedCards: []int64{
				0, 1, 2, 3, 4,
			},
		},
		{
			it: "full-house",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Four,
				},
			},
			expectedResult: poker.FullHouse,
			expectedCards: []int64{
				0, 1, 2, 3, 4,
			},
		},
		{
			it: "four-of-a-kind",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Four,
				},
			},
			expectedResult: poker.FourOfAKind,
			expectedCards: []int64{
				1, 2, 3, 4,
			},
		},
		{
			it: "straight-flush",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Three,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
			},
			expectedResult: poker.StraightFlush,
			expectedCards: []int64{
				0, 1, 2, 3, 4,
			},
		},
		{
			it: "royal-flush",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.King,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.RoyalFlush,
			expectedCards: []int64{
				0, 1, 2, 3, 4,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			result, cards := poker.BestResult(tc.cards)
			is.New(t).Equal(tc.expectedResult, result)
			is.New(t).Equal(tc.expectedCards, cards)
		})
	}
}

func TestBestResult7Cards(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult string
		expectedCards  []int64
	}{
		{
			it: "high-card",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.King,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Seven,
				},
			},
			expectedResult: poker.HighCard,
			expectedCards: []int64{
				5,
			},
		},
		{
			it: "pair",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Eight,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.King,
				},
			},
			expectedResult: poker.Pair,
			expectedCards: []int64{
				3, 4,
			},
		},
		{
			it: "two-pairs",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.King,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Queen,
				},
			},
			expectedResult: poker.TwoPairs,
			expectedCards: []int64{
				0, 2, 3, 4,
			},
		},
		{
			it: "three-of-a-kind",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
			},
			expectedResult: poker.ThreeOfAKind,
			expectedCards: []int64{
				0, 2, 6,
			},
		},
		{
			it: "straight",
			cards: []*poker.Card{
				{
					Suite: poker.Spades,
					Rank:  poker.Six,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Three,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Seven,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Queen,
				},
			},
			expectedResult: poker.Straight,
			expectedCards: []int64{
				0, 1, 2, 4, 5,
			},
		},
		{
			it: "flush",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Three,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.Flush,
			expectedCards: []int64{
				0, 1, 2, 4, 6,
			},
		},
		{
			it: "full-house",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.FullHouse,
			expectedCards: []int64{
				0, 1, 2, 3, 5,
			},
		},
		{
			it: "four-of-a-kind",
			cards: []*poker.Card{
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Clubs,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Spades,
					Rank:  poker.Four,
				},
			},
			expectedResult: poker.FourOfAKind,
			expectedCards: []int64{
				1, 2, 3, 6,
			},
		},
		{
			it: "straight-flush",
			cards: []*poker.Card{
				{
					Suite: poker.Hearts,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Five,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Two,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Three,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Four,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Six,
				},
			},
			expectedResult: poker.StraightFlush,
			expectedCards: []int64{
				1, 3, 4, 5, 6,
			},
		},
		{
			it: "royal-flush",
			cards: []*poker.Card{
				{
					Suite: poker.Hearts,
					Rank:  poker.Ace,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ten,
				},
				{
					Suite: poker.Hearts,
					Rank:  poker.King,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Jack,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Queen,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.King,
				},
				{
					Suite: poker.Diamonds,
					Rank:  poker.Ace,
				},
			},
			expectedResult: poker.RoyalFlush,
			expectedCards: []int64{
				1, 3, 4, 5, 6,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			result, cards := poker.BestResult(tc.cards)
			is.New(t).Equal(tc.expectedResult, result)
			is.New(t).Equal(tc.expectedCards, cards)
		})
	}
}
