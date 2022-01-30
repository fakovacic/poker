package poker_test

// import (
// 	"testing"

// 	"github.com/fakovacic/poker"
// )

// func TestCompareHands(t *testing.T) {
// 	cases := []struct {
// 		it             string
// 		hands          [][]*poker.Card
// 		expectedResult map[int64]float64
// 	}{
// 		{
// 			it: "compare high card and pair",
// 			hands: [][]*poker.Card{
// 				{
// 					{
// 						Suite: poker.Spades,
// 						Rank:  poker.Ten,
// 					},
// 					{
// 						Suite: poker.Diamonds,
// 						Rank:  poker.Six,
// 					},
// 					{
// 						Suite: poker.Hearts,
// 						Rank:  poker.Queen,
// 					},
// 					{
// 						Suite: poker.Diamonds,
// 						Rank:  poker.Two,
// 					},
// 					{
// 						Suite: poker.Spades,
// 						Rank:  poker.Ace,
// 					},
// 				},
// 				{
// 					{
// 						Suite: poker.Spades,
// 						Rank:  poker.Ten,
// 					},
// 					{
// 						Suite: poker.Diamonds,
// 						Rank:  poker.Six,
// 					},
// 					{
// 						Suite: poker.Hearts,
// 						Rank:  poker.Two,
// 					},
// 					{
// 						Suite: poker.Diamonds,
// 						Rank:  poker.Two,
// 					},
// 					{
// 						Suite: poker.Spades,
// 						Rank:  poker.Ace,
// 					},
// 				},
// 			},
// 			expectedResult: map[int64]float64{
// 				1: 1,
// 				2: 1,
// 			},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.it, func(t *testing.T) {
// 			result := poker.CompareHands(tc.hands...)

// 			for i := range tc.expectedResult {
// 				if tc.expectedResult[i] != result[i] {
// 					t.Errorf("expected card %d: '%v' got: '%v'", i, tc.expectedResult[i], result[i])
// 				}
// 			}
// 		})
// 	}
// }
