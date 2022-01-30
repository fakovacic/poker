package poker_test

import (
	"testing"

	"github.com/fakovacic/poker"
)

func TestCompareResults(t *testing.T) {
	cases := []struct {
		it             string
		results        []poker.Result
		expectedResult poker.Compare
	}{
		{
			it: "high",
			results: []poker.Result{
				poker.HighCard,
				poker.Flush,
			},
			expectedResult: poker.High,
		},
		{
			it: "low",
			results: []poker.Result{
				poker.FullHouse,
				poker.Flush,
			},
			expectedResult: poker.Low,
		},
		{
			it: "equal",
			results: []poker.Result{
				poker.Flush,
				poker.Flush,
			},
			expectedResult: poker.Equal,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			res := poker.CompareResults(tc.results[0], tc.results[1])
			if tc.expectedResult != res {
				t.Errorf("expected: '%v' got: '%v'", tc.expectedResult, res)
			}
		})
	}
}
