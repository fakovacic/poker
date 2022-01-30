package poker

import "math"

func CalculateCardRank(rank Rank) (int64, int64) {
	if rank == "" {
		return 0, 0
	}

	rankNum := ranksScore[rank]

	var (
		less int64
		more int64
	)

	for k := range ranksScore {
		switch {
		case rankNum < ranksScore[k]:
			more++
		case rankNum > ranksScore[k]:
			less++
		}
	}

	return less, more
}

func CalculateCardRankOdds(rank Rank) (float64, float64) {
	less, more := CalculateCardRank(rank)

	lessOdds := math.Round(((float64(less*4)/float64(51))*100)*oddsRounding) / oddsRounding
	moreOdds := math.Round(((float64(more*4)/float64(51))*100)*oddsRounding) / oddsRounding

	return lessOdds, moreOdds
}
