package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestParseColor(t *testing.T) {
	cases := []struct {
		it             string
		color          string
		expectedOK     bool
		expectedResult poker.Color
	}{
		{
			it:             "red",
			color:          "R",
			expectedResult: poker.Red,
			expectedOK:     true,
		},
		{
			it:             "black",
			color:          "B",
			expectedResult: poker.Black,
			expectedOK:     true,
		},
		{
			it:             "invalid color",
			color:          "ok",
			expectedResult: "",
			expectedOK:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res, ok := poker.ParseColor(tc.color)
			if tc.expectedOK != ok {
				t.Errorf("expected ok: '%v' got: '%v'", tc.expectedOK, ok)
			}

			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
