package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestCalculateCardRank(t *testing.T) {
	cases := []struct {
		it                 string
		rank               poker.Rank
		expectedResultLow  int64
		expectedResultHigh int64
	}{
		{
			it:                 "two",
			rank:               poker.Two,
			expectedResultLow:  0,
			expectedResultHigh: 12,
		},
		{
			it:                 "three",
			rank:               poker.Three,
			expectedResultLow:  1,
			expectedResultHigh: 11,
		},
		{
			it:                 "four",
			rank:               poker.Four,
			expectedResultLow:  2,
			expectedResultHigh: 10,
		},
		{
			it:                 "five",
			rank:               poker.Five,
			expectedResultLow:  3,
			expectedResultHigh: 9,
		},
		{
			it:                 "six",
			rank:               poker.Six,
			expectedResultLow:  4,
			expectedResultHigh: 8,
		},
		{
			it:                 "seven",
			rank:               poker.Seven,
			expectedResultLow:  5,
			expectedResultHigh: 7,
		},
		{
			it:                 "eight",
			rank:               poker.Eight,
			expectedResultLow:  6,
			expectedResultHigh: 6,
		},
		{
			it:                 "nine",
			rank:               poker.Nine,
			expectedResultLow:  7,
			expectedResultHigh: 5,
		},
		{
			it:                 "ten",
			rank:               poker.Ten,
			expectedResultLow:  8,
			expectedResultHigh: 4,
		},
		{
			it:                 "jack",
			rank:               poker.Jack,
			expectedResultLow:  9,
			expectedResultHigh: 3,
		},
		{
			it:                 "queen",
			rank:               poker.Queen,
			expectedResultLow:  10,
			expectedResultHigh: 2,
		},
		{
			it:                 "king",
			rank:               poker.King,
			expectedResultLow:  11,
			expectedResultHigh: 1,
		},
		{
			it:                 "ace",
			rank:               poker.Ace,
			expectedResultLow:  12,
			expectedResultHigh: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			low, high := poker.CalculateCardRank(tc.rank)
			if tc.expectedResultLow != low {
				t.Errorf("expected low: '%v' got: '%v'", tc.expectedResultLow, low)
			}

			if tc.expectedResultHigh != high {
				t.Errorf("expected high: '%v' got: '%v'", tc.expectedResultHigh, high)
			}
		})
	}
}

func TestCalculateCardRankOdds(t *testing.T) {
	cases := []struct {
		it                 string
		rank               poker.Rank
		expectedResultLow  float64
		expectedResultHigh float64
	}{
		{
			it:                 "two",
			rank:               poker.Two,
			expectedResultLow:  0,
			expectedResultHigh: 94.11765,
		},
		{
			it:                 "three",
			rank:               poker.Three,
			expectedResultLow:  7.84314,
			expectedResultHigh: 86.27451,
		},
		{
			it:                 "four",
			rank:               poker.Four,
			expectedResultLow:  15.68627,
			expectedResultHigh: 78.43137,
		},
		{
			it:                 "five",
			rank:               poker.Five,
			expectedResultLow:  23.52941,
			expectedResultHigh: 70.58824,
		},
		{
			it:                 "six",
			rank:               poker.Six,
			expectedResultLow:  31.37255,
			expectedResultHigh: 62.7451,
		},
		{
			it:                 "seven",
			rank:               poker.Seven,
			expectedResultLow:  39.21569,
			expectedResultHigh: 54.90196,
		},
		{
			it:                 "eight",
			rank:               poker.Eight,
			expectedResultLow:  47.05882,
			expectedResultHigh: 47.05882,
		},
		{
			it:                 "nine",
			rank:               poker.Nine,
			expectedResultLow:  54.90196,
			expectedResultHigh: 39.21569,
		},
		{
			it:                 "ten",
			rank:               poker.Ten,
			expectedResultLow:  62.7451,
			expectedResultHigh: 31.37255,
		},
		{
			it:                 "jack",
			rank:               poker.Jack,
			expectedResultLow:  70.58824,
			expectedResultHigh: 23.52941,
		},
		{
			it:                 "queen",
			rank:               poker.Queen,
			expectedResultLow:  78.43137,
			expectedResultHigh: 15.68627,
		},
		{
			it:                 "king",
			rank:               poker.King,
			expectedResultLow:  86.27451,
			expectedResultHigh: 7.84314,
		},
		{
			it:                 "ace",
			rank:               poker.Ace,
			expectedResultLow:  94.11765,
			expectedResultHigh: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			low, high := poker.CalculateCardRankOdds(tc.rank)
			if tc.expectedResultLow != low {
				t.Errorf("expected low: '%v' got: '%v'", tc.expectedResultLow, low)
			}

			if tc.expectedResultHigh != high {
				t.Errorf("expected high: '%v' got: '%v'", tc.expectedResultHigh, high)
			}
		})
	}
}
