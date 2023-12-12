package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

type Hand struct {
	cards string
	combo int
	score int
	bet   int
}

const CARDS = "J23456789TQKA"
const (
	_ int = iota
	HighCard
	Pair
	TwoPair
	Three
	FullHouse
	Four
	Five
)

func Value(r rune) int {
	return strings.Index(CARDS, string(r)) + 1
}

func (h1 Hand) Less(h2 Hand) bool {
	if h1.combo != h2.combo {
		return h1.combo < h2.combo
	} else {
		return h1.score < h2.score
	}
}

func solve(data string) {
	rows := strings.Split(data, "\n")
	var hands []Hand

	for _, row := range rows {
		parts := strings.Split(row, " ")
		rawCards := parts[0]
		var cards string
		if strings.Contains(parts[0], "J") {
			cards = ParseJokers(parts[0])
		} else {
			cards = parts[0]
		}
		bet := Atoi(parts[1])
		hands = append(hands, Hand{cards: cards, bet: bet, combo: Combo(cards), score: Score(rawCards)})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Less(hands[j])
	})

	res := 0
	for i, hand := range hands {
		fmt.Println(hand)
		res += (i + 1) * hand.bet
	}

	fmt.Println(res)
}

func ParseJokers(cards string) string {
	if cards == "JJJJJ" {
		return cards
	}

	counts := make(map[rune]int)
	var maxCount int
	var maxCard rune

	for _, card := range cards {
		if card == 'J' {
			continue
		}

		counts[card] += 1

		if counts[card] > maxCount {
			maxCard = card
			maxCount = counts[card]
		}
		if counts[card] == maxCount && Value(card) > Value(maxCard) {
			maxCard = card
		}
	}

	mod := strings.ReplaceAll(cards, "J", string(maxCard))
	return mod
}

func Combo(cards string) int {
	var scores []int
	for _, r := range cards {
		scores = append(scores, Value(r))
	}

	sort.Ints(scores)

	switch {
	case scores[0] == scores[4]:
		return Five
	case scores[0] == scores[3] || scores[1] == scores[4]:
		return Four
	case (scores[0] == scores[2] && scores[3] == scores[4]) || (scores[0] == scores[1] && scores[2] == scores[4]):
		return FullHouse
	case scores[0] == scores[2] || scores[1] == scores[3] || scores[2] == scores[4]:
		return Three
	case (scores[0] == scores[1] && scores[2] == scores[3]) || (scores[1] == scores[2] && scores[3] == scores[4]) || (scores[0] == scores[1] && scores[3] == scores[4]):
		return TwoPair
	case scores[0] == scores[1] || scores[1] == scores[2] || scores[2] == scores[3] || scores[3] == scores[4]:
		return Pair
	default:
		return HighCard
	}
}

func Score(cards string) int {
	points := 0
	for i, card := range cards {
		points += Value(rune(card)) * int(math.Pow10(10-i*2))
	}

	return points
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
