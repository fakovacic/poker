package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestResult(t *testing.T) {
	cases := []struct {
		it             string
		cards          []*poker.Card
		expectedResult string
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
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.Result(tc.cards)
			if tc.expectedResult != res {
				t.Errorf("expected result: '%s' got: '%s'", tc.expectedResult, res)
			}
		})
	}
}

func benchmarkResult(cards []*poker.Card, b *testing.B) {
	for n := 0; n < b.N; n++ {
		poker.Result(cards)
	}
}

func BenchmarkResultHighCard(b *testing.B) {
	benchmarkResult([]*poker.Card{
		{
			Rank: poker.Ace,
		},
		{
			Rank: poker.King,
		},
		{
			Rank: poker.Ace,
		},
		{
			Rank: poker.King,
		},
		{
			Rank: poker.King,
		},
	}, b)
}
func BenchmarkResultPair(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}
func BenchmarkResultTwoPair(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultThreeOfAKind(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultStraight(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultFlush(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultFullHouse(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultFourOfAKind(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultStraightFlush(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}

func BenchmarkResultRoyalFlush(b *testing.B) {
	benchmarkResult([]*poker.Card{
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
	}, b)
}
