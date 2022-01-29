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


## Card


## Check certain result


## Get result


## Get best result


## Odds


## Rank odds



## TODO

- find best match between multiple hands
- winning odds for multiple hands