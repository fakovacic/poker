package poker

type Card struct {
	Suite string `json:"suite"`
	Rank  string `json:"rank"`
}

func (c *Card) RankScore() int64 {
	return ranksScore[c.Rank]
}
