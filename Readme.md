# Poker functions

- golang poker helper functions for use to determine poker cards result, odds and more

## Defined constants

### Ranks

```
const (
	Ace   string = "A"
	King  string = "K"
	Queen string = "Q"
	Jack  string = "J"
	Ten   string = "10"
	Nine  string = "9"
	Eight string = "8"
	Seven string = "7"
	Six   string = "6"
	Five  string = "5"
	Four  string = "4"
	Three string = "3"
	Two   string = "2"
)
```
### Suits

```
const (
	Clubs    string = "C"
	Diamonds string = "D"
	Hearts   string = "H"
	Spades   string = "S"
)
```

### Colors

```
const (
	Red   string = "R"
	Black string = "B"
)
```

### Results

```
const (
	HighCard      string = "high-card"
	Pair          string = "pair"
	TwoPairs      string = "two-pairs"
	ThreeOfAKind  string = "three-of-a-kind"
	Straight      string = "straight"
	Flush         string = "flush"
	FullHouse     string = "full-house"
	FourOfAKind   string = "four-of-a-kind"
	StraightFlush string = "straight-flush"
	RoyalFlush    string = "royal-flush"
)
```

## Deck

- deck is used to generate/shuffle all cards in deck, and also draw or discard cards

```
func main(){
    deck := poker.NewDeck()
    deck.Shuffle()

    card := deck.Draw()
    fmt.Println(card)

    deck.Discard()

    leftCards := deck.Cards()
    fmt.Println(leftCards)
}
```

- deck by default uses default shuffler
```
rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(cards), func(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
})
```

- but you can inject your own implementing Shuffler interface

```
type newShuffler struct{}

func (ns newShuffler) Shuffle(cards []*poker.Card) []*poker.Card{
    // shuffle cards
}

func main(){
    deck := poker.NewDeck().SetShuffler(
        &newShuffler{},
    )
    deck.Shuffle()
}
```

## Hand
- hand mimic player cards hand, you can deal card to player hand or get cards from player

```
func main(){
	d := poker.NewDeck()
	d.Shuffle()

	hand := poker.NewHand()

	for i := 0; i < 5; i++ {
		hand.Deal(d.Draw())
	}

	cards := hand.Cards()

	fmt.Println(cards)
}

```

## Card
- card represent one card from deck

```
type Card struct {
	Suite string `json:"suite"`
	Rank  string `json:"rank"`
}
```

## Check certain result
- there are set of functions to check possible result from list of cards  

```
func main(){
	cards := []*poker.Card{
		{
			Suite: poker.Hearts,
			Rank:  poker.Two,
		},
		{
			Suite: poker.Diamonds,
			Rank:  poker.Two,
		},
	}

	// fmt.Println(poker.IsRoyalFlush(cards))
	// fmt.Println(poker.IsStraightFlush(cards))
	// fmt.Println(poker.IsFourOfAKind(cards))
	// fmt.Println(poker.IsFullHouse(cards))
	// fmt.Println(poker.IsFlush(cards))
	// fmt.Println(poker.IsStraight(cards))
	// fmt.Println(poker.IsThreeOfAKind(cards))
	// fmt.Println(poker.IsTwoPairs(cards))
	fmt.Println(poker.IsPair(cards))
}

```

## Get result
- result will check best possible result from set of cards

```
func main(){
	cards := []*poker.Card{
		{
			Suite: poker.Diamonds,
			Rank:  poker.Queen,
		},
		{
			Suite: poker.Hearts,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Clubs,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Diamonds,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.Four,
		},
	}

	fmt.Println(poker.Result(cards)) // print poker
}

```

## Get best result cards
- check which cards give most stronger result

```
func main(){
	cards := []*poker.Card{
		{
			Suite: poker.Diamonds,
			Rank:  poker.Queen,
		},
		{
			Suite: poker.Hearts,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Clubs,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Diamonds,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.Four,
		},
	}

	result, cards := poker.BestResult(cards)

	fmt.Println(result) // print poker
	fmt.Println(cards) // print [1,2,3,4]
}

```

## Odds
- you can calculate odds of each result from set of cards

```
func main(){
	cards := []*poker.Card{
		{
			Suite: poker.Diamonds,
			Rank:  poker.Queen,
		},
		{
			Suite: poker.Hearts,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Clubs,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Diamonds,
			Rank:  poker.Four,
		},
		{
			Suite: poker.Spades,
			Rank:  poker.Four,
		},
	}

	combinations, odds := poker.OddsDeck(cards)

	fmt.Println(combinations)
	fmt.Println(odds)
}

```

## Rank odds
- compare ranks between 2 cards

```
func main(){
	firstCard := &poker.Card{
		Suite: poker.Diamonds,
		Rank:  poker.Queen,
	}

	secondCard := &poker.Card{
		Suite: poker.Diamonds,
		Rank:  poker.Six,
	}

	fmt.Println(poker.CompareCardsRank(firstCard, secondCard)) // lower
}

```

## TODO

- cards score
- find best match between multiple hands
- winning odds for multiple hands

- types