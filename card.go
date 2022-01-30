package poker

import "fmt"

type Card struct {
	Suite Suite `json:"suite"`
	Rank  Rank  `json:"rank"`
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suite)
}
