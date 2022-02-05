package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestParseSuite(t *testing.T) {
	cases := []struct {
		it             string
		suite          string
		expectedOK     bool
		expectedResult poker.Suite
	}{
		{
			it:             "clubs",
			suite:          "C",
			expectedResult: poker.Clubs,
			expectedOK:     true,
		},
		{
			it:             "diamonds",
			suite:          "D",
			expectedResult: poker.Diamonds,
			expectedOK:     true,
		},
		{
			it:             "hearts",
			suite:          "H",
			expectedResult: poker.Hearts,
			expectedOK:     true,
		},
		{
			it:             "spades",
			suite:          "S",
			expectedResult: poker.Spades,
			expectedOK:     true,
		},
		{
			it:             "invalid suite",
			suite:          "ok",
			expectedResult: "",
			expectedOK:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res, ok := poker.ParseSuite(tc.suite)
			if tc.expectedOK != ok {
				t.Errorf("expected ok: '%v' got: '%v'", tc.expectedOK, ok)
			}

			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}

func TestSuiteColor(t *testing.T) {
	cases := []struct {
		it             string
		suite          poker.Suite
		expectedResult poker.Color
	}{
		{
			it:             "clubs",
			suite:          poker.Clubs,
			expectedResult: poker.Black,
		},
		{
			it:             "diamonds",
			suite:          poker.Diamonds,
			expectedResult: poker.Red,
		},
		{
			it:             "hearts",
			suite:          poker.Hearts,
			expectedResult: poker.Red,
		},
		{
			it:             "spades",
			suite:          poker.Spades,
			expectedResult: poker.Black,
		},
		{
			it:             "ok",
			suite:          "ok",
			expectedResult: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := tc.suite.Color()
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
