package poker

import "math"

func OddsRank(rank Rank) (float64, float64) {
	less, more := rank.ScorePosition()

	lessOdds := math.Round(((float64(less*4)/float64(51))*100)*oddsRounding) / oddsRounding
	moreOdds := math.Round(((float64(more*4)/float64(51))*100)*oddsRounding) / oddsRounding

	return lessOdds, moreOdds
}
