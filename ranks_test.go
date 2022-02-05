package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestScorePosition(t *testing.T) {
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
			low, high := tc.rank.ScorePosition()
			if tc.expectedResultLow != low {
				t.Errorf("expected low: '%v' got: '%v'", tc.expectedResultLow, low)
			}

			if tc.expectedResultHigh != high {
				t.Errorf("expected high: '%v' got: '%v'", tc.expectedResultHigh, high)
			}
		})
	}
}

func TestParseRank(t *testing.T) {
	cases := []struct {
		it             string
		rank           string
		expectedOK     bool
		expectedResult poker.Rank
	}{
		{
			it:             "two",
			rank:           "2",
			expectedResult: poker.Two,
			expectedOK:     true,
		},
		{
			it:             "three",
			rank:           "3",
			expectedResult: poker.Three,
			expectedOK:     true,
		},
		{
			it:             "four",
			rank:           "4",
			expectedResult: poker.Four,
			expectedOK:     true,
		},
		{
			it:             "five",
			rank:           "5",
			expectedResult: poker.Five,
			expectedOK:     true,
		},
		{
			it:             "six",
			rank:           "6",
			expectedResult: poker.Six,
			expectedOK:     true,
		},
		{
			it:             "seven",
			rank:           "7",
			expectedResult: poker.Seven,
			expectedOK:     true,
		},
		{
			it:             "eight",
			rank:           "8",
			expectedResult: poker.Eight,
			expectedOK:     true,
		},
		{
			it:             "nine",
			rank:           "9",
			expectedResult: poker.Nine,
			expectedOK:     true,
		},
		{
			it:             "ten",
			rank:           "10",
			expectedResult: poker.Ten,
			expectedOK:     true,
		},
		{
			it:             "jack",
			rank:           "J",
			expectedResult: poker.Jack,
			expectedOK:     true,
		},
		{
			it:             "queen",
			rank:           "Q",
			expectedResult: poker.Queen,
			expectedOK:     true,
		},
		{
			it:             "king",
			rank:           "K",
			expectedResult: poker.King,
			expectedOK:     true,
		},
		{
			it:             "ace",
			rank:           "A",
			expectedResult: poker.Ace,
			expectedOK:     true,
		},
		{
			it:             "invalid suite",
			rank:           "ok",
			expectedResult: "",
			expectedOK:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res, ok := poker.ParseRank(tc.rank)
			if tc.expectedOK != ok {
				t.Errorf("expected ok: '%v' got: '%v'", tc.expectedOK, ok)
			}

			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
