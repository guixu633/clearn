package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type Card struct {
	Rank int
	Suit string // s(spades), h(hearts), d(diamonds), c(clubs)
}

type Hand struct {
	Hand []Card
}

type HandRank int8

const (
	HighCard HandRank = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func (r HandRank) String() string {
	switch r {
	case HighCard:
		return "HighCard"
	case Pair:
		return "Pair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case StraightFlush:
		return "StraightFlush"
	case RoyalFlush:
		return "RoyalFlush"
	}
	return "Unknown"
}

func parseCard(card string) (Card, error) {
	if len(card) != 2 {
		return Card{}, errors.New("invalid card")
	}
	rank := card[:len(card)-1]
	suit := card[len(card)-1:]
	var rankInt int
	switch rank {
	case "A":
		rankInt = 14
	case "K":
		rankInt = 13
	case "Q":
		rankInt = 12
	case "J":
		rankInt = 11
	case "T":
		rankInt = 10
	default:
		var err error
		rankInt, err = strconv.Atoi(rank)
		if err != nil {
			return Card{}, errors.New("invalid card")
		}
	}
	if rankInt < 2 || rankInt > 14 {
		return Card{}, errors.New("invalid card")
	}
	return Card{Rank: rankInt, Suit: suit}, nil
}

func parseHands(cards []string) (Hand, error) {
	hand := Hand{}
	for _, card := range cards {
		card, err := parseCard(card)
		if err != nil {
			return hand, err
		}
		hand.Hand = append(hand.Hand, card)
	}
	hand.Sort()
	return hand, nil
}

func (h *Hand) Sort() {
	sort.Slice(h.Hand, func(i, j int) bool {
		return h.Hand[i].Rank > h.Hand[j].Rank
	})
}

func (h *Hand) FindFlush() Hand {
	count := [4]int{}
	list := [4][]Card{}
	for _, card := range h.Hand {
		switch card.Suit {
		case "s":
			count[0]++
			list[0] = append(list[0], card)
		case "h":
			count[1]++
			list[1] = append(list[1], card)
		case "d":
			count[2]++
			list[2] = append(list[2], card)
		case "c":
			count[3]++
			list[3] = append(list[3], card)
		}
	}
	max := 0
	for i := 1; i < 4; i++ {
		if count[i] > count[max] {
			max = i
		}
	}
	return Hand{Hand: list[max]}
}

func (h *Hand) FindPair() []Hand {
	m := make(map[int][]Card)
	for _, card := range h.Hand {
		m[card.Rank] = append(m[card.Rank], card)
	}
	result := []Hand{}
	for _, cards := range m {
		result = append(result, Hand{Hand: cards})
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i].Hand) != len(result[j].Hand) {
			return len(result[i].Hand) > len(result[j].Hand)
		}
		return result[i].Hand[0].Rank > result[j].Hand[0].Rank
	})
	return result
}

func (h *Hand) MaxStraight() (int, int) {
	hand := h.Hand
	if hand[0].Rank == 14 && hand[len(hand)-1].Rank == 2 {
		hand = append(hand, Card{Rank: 1, Suit: hand[0].Suit})
	}

	maxLen := 1
	maxStart := hand[0].Rank
	currentLen := 1
	for i := 1; i < len(hand); i++ {
		if hand[i].Rank == hand[i-1].Rank-1 {
			currentLen++
			if currentLen == 5 {
				return 5, hand[i-4].Rank
			}
		} else {
			if maxLen < currentLen {
				maxLen = currentLen
				maxStart = hand[i-1].Rank + currentLen - 1
			}
			currentLen = 1
		}
	}
	if maxLen < currentLen {
		maxLen = currentLen
		maxStart = hand[len(hand)-1].Rank + currentLen - 1
	}
	return maxLen, maxStart
}

func (h *Hand) CalHighCardScore() int {
	score := h.Hand[0].Rank
	for _, card := range h.Hand[1:6] {
		score <<= 4
		score += card.Rank
	}
	return score
}

func (h *Hand) CalRank() (HandRank, int) {
	flush := h.FindFlush()
	if len(flush.Hand) >= 5 {
		maxLen, maxStart := flush.MaxStraight()
		if maxLen == 5 {
			if maxStart == 14 {
				return RoyalFlush, 0
			}
			return StraightFlush, maxStart
		}
	}
	pairs := h.FindPair()
	if len(pairs[0].Hand) == 4 {
		return FourOfAKind, pairs[0].Hand[0].Rank<<4 + pairs[1].Hand[0].Rank
	}
	if len(pairs[0].Hand) == 3 && len(pairs[1].Hand) == 2 {
		return FullHouse, pairs[0].Hand[0].Rank<<4 + pairs[1].Hand[0].Rank
	}
	if len(flush.Hand) >= 5 {
		return Flush, flush.Hand[0].Rank
	}
	if maxLen, maxStart := h.MaxStraight(); maxLen == 5 {
		return Straight, maxStart
	}
	if len(pairs[0].Hand) == 3 {
		return ThreeOfAKind, pairs[0].Hand[0].Rank<<8 + pairs[1].Hand[0].Rank<<4 + pairs[2].Hand[0].Rank
	}
	if len(pairs[0].Hand) == 2 && len(pairs[1].Hand) == 2 {
		return TwoPair, pairs[0].Hand[0].Rank<<8 + pairs[1].Hand[0].Rank<<4 + pairs[2].Hand[0].Rank
	}
	if len(pairs[0].Hand) == 2 {
		return Pair, pairs[0].Hand[0].Rank<<12 + pairs[1].Hand[0].Rank<<8 + pairs[2].Hand[0].Rank<<4 + pairs[3].Hand[0].Rank
	}
	return HighCard, h.CalHighCardScore()
}

func Calculate(hands []string) (HandRank, int, error) {
	if len(hands) != 7 {
		return 0, 0, errors.New("invalid hands")
	}
	hs, err := parseHands(hands)
	if err != nil {
		return 0, 0, err
	}
	rank, score := hs.CalRank()
	return rank, score, nil
}

func Compare(hands1, hands2 []string) (int, error) {
	rank1, score1, err := Calculate(hands1)
	if err != nil {
		return 0, err
	}
	fmt.Println(rank1, score1)
	rank2, score2, err := Calculate(hands2)
	if err != nil {
		return 0, err
	}
	fmt.Println(rank2, score2)
	if rank1 > rank2 {
		return 1, nil
	}
	if rank1 < rank2 {
		return -1, nil
	}
	if score1 > score2 {
		return 1, nil
	}
	if score1 < score2 {
		return -1, nil
	}
	return 0, nil
}
