package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestColor(t *testing.T) {
	cases := []struct {
		it             string
		suite          string
		expectedResult string
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
			res := poker.Color(tc.suite)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}

func TestValidateColor(t *testing.T) {
	cases := []struct {
		it             string
		suite          string
		expectedResult bool
	}{
		{
			it:             "red",
			suite:          poker.Red,
			expectedResult: true,
		},
		{
			it:             "black",
			suite:          poker.Black,
			expectedResult: true,
		},
		{
			it:             "ok",
			suite:          "ok",
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.ValidateColor(tc.suite)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
