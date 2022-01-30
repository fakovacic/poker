package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestValidateColor(t *testing.T) {
	cases := []struct {
		it             string
		color          poker.Color
		expectedResult bool
	}{
		{
			it:             "red",
			color:          poker.Red,
			expectedResult: true,
		},
		{
			it:             "black",
			color:          poker.Black,
			expectedResult: true,
		},
		{
			it:             "ok",
			color:          "ok",
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.ValidateColor(tc.color)
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
