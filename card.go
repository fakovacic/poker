package poker

type Card struct {
	Suite Suite `json:"suite"`
	Rank  Rank  `json:"rank"`
}
