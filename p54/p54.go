package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// High Card: Highest value card.
// One Pair: Two cards of the same value.
// Two Pairs: Two different pairs.
// Three of a Kind: Three cards of the same value.
// Straight: All cards are consecutive values.
// Flush: All cards of the same suit.
// Full House: Three of a kind and a pair.
// Four of a Kind: Four cards of the same value.
// Straight Flush: All cards are consecutive values of same suit.
// Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
const (
	HIGH          = 0
	PAIR          = 1
	TWO           = 2
	THREE         = 3
	STRAIGHT      = 4
	FLUSH         = 5
	FULLHOUSE     = 6
	FOUR          = 7
	STRAIGHTFLUSH = 8
	ROYALFLUSH    = 9
)

type combination int

// suits
const (
	DIAMONDS = "D"
	CLUBS    = "C"
	HEARTS   = "H"
	SPADES   = "S"
)

// values: 2-9, T, J, Q, K, A
const (
	ACE   = "A"
	KING  = "K"
	QUEEN = "Q"
	JACK  = "J"
	TEN   = "T"
)

var ranks map[string]int = map[string]int{
	ACE:   14,
	KING:  13,
	QUEEN: 12,
	JACK:  11,
	TEN:   10,
}

// Card structure with suit and value
type Card struct {
	s string
	v string
}

// Card constructor
func NewCard(s, v string) *Card {
	return &Card{s, v}
}

func (c *Card) Rank() int {
	r, ok := ranks[c.v]
	if !ok {
		r, _ = strconv.Atoi(c.v)
	}

	return r
}

// Hand is array of 5 cards
type Hand []Card

func (h Hand) Len() int           { return len(h) }
func (h Hand) Less(i, j int) bool { return h[i].Rank() < h[j].Rank() }
func (h Hand) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func GetCombination(hand Hand) (combination, Hand) {
	// default combination
	var result combination = HIGH
	var output Hand

	type Calc struct {
		amount int
		cards  Hand
	}

	values := map[string]Calc{}
	isStraight := true
	isFlush := true

	var lastCard Card

	for _, card := range hand {
		if lastCard.s != "" {
			if lastCard.Rank()+1 != card.Rank() {
				isStraight = false
			}

			if lastCard.s != card.s {
				isFlush = false
			}
		}

		_, ok := values[card.v]
		if !ok {
			values[card.v] = Calc{0, []Card{}}
		}

		current := values[card.v]
		current.amount++
		current.cards = append(current.cards, card)
		values[card.v] = current

		lastCard = card
	}

	sameValues := make([]Calc, 0)
	for _, value := range values {
		sameValues = append(sameValues, value)
	}

	sort.Slice(sameValues, func(i, j int) bool {
		return sameValues[i].amount < sameValues[j].amount
	})

	fmt.Println(sameValues)

	switch sameValues[len(sameValues)-1].amount {
	case 4:
		result = FOUR
		output = sameValues[len(sameValues)-1].cards
	case 3:
		result = THREE
		output = sameValues[len(sameValues)-1].cards
		if len(sameValues) == 2 {
			result = FULLHOUSE
			output = hand
		}
	case 2:
		result = PAIR
		output = sameValues[len(sameValues)-1].cards
		if len(sameValues) == 3 {
			result = TWO
			output = make(Hand, 0)
			output = append(output, sameValues[len(sameValues)-1].cards...)
			output = append(output, sameValues[len(sameValues)-2].cards...)
		}
	}

	if isStraight && result < STRAIGHT {
		result = STRAIGHT
		output = hand
	}

	if isFlush && result < FLUSH {
		result = FLUSH
		output = hand
	}

	if isStraight && isFlush {
		result = STRAIGHTFLUSH

		if hand[4].v == ACE {
			result = ROYALFLUSH
		}
	}

	return result, output
}

func ParseHand(input string) Hand {
	hand := make(Hand, 5)

	cards := strings.Split(input, " ")
	for i, c := range cards {
		hand[i] = *NewCard(string(c[1]), string(c[0]))
	}

	sort.Sort(hand)

	return hand
}

func CompareHands(h1, h2 Hand) int {
	for i := len(h1) - 1; i >= 0; i-- {
		if h1[i].Rank() != h2[i].Rank() {
			if h1[i].Rank() > h2[i].Rank() {
				return 1
			}

			return -1
		}
	}

	return 0
}

// how many hands does player 1 win?
func main() {
	result := 0

	file, err := os.Open("./p54/p054_poker.txt")
	if err != nil {
		fmt.Println("ERROR OPEN FILE", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands := scanner.Text()
		hand1 := hands[0:14]
		hand2 := hands[15:len(hands)]

		h1 := ParseHand(hand1)
		h2 := ParseHand(hand2)
		c1, ph1 := GetCombination(h1)
		c2, ph2 := GetCombination(h2)

		if c1 > c2 {
			result++
			continue
		}

		if c1 == c2 {
			r := CompareHands(ph1, ph2)
			if r > 0 {
				result++
			} else {
				if r == 0 {
					if CompareHands(h1, h2) > 0 {
						result++
					}
				}
			}
		}
	}

	fmt.Println(result)
}
