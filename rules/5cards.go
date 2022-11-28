package rules

import (
	"sort"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

func Check5value(Player []agt.Card, Table []agt.Card) (sc int) {
	var hand []agt.Card
	hand = append(hand, Player...)
	hand = append(hand, Table...)
	sort.Slice(hand, func(i int, j int) bool {
		return hand[i].Value > hand[j].Value
	})
	if isRoyalFlush5(hand) {
		sc = Check5RoyalFlush(hand)
	} else if isStraightFlush5(hand) {
		sc = Check5StraightFlush(hand)
	} else if isFourOfAKind5(hand) {
		sc = Check5FourofAKind(hand)
	} else if isFullHouse5(hand) {
		sc = Check5FullHouse(hand)
	} else if isFlush5(hand) {
		sc = Check5Flush(hand)
	} else if isStraight5(hand) {
		sc = Check5Straight(hand)
	} else if isThreeOfAKind5(hand) {
		sc = Check5ThreeofAKind(hand)
	} else if isTwoPair5(hand) {
		sc = Check5TwoPair(hand)
	} else if isPair5(hand) {
		sc = Check5Pair(hand)
	} else {
		sc = Check5HighCard(hand)
	}
	return sc
}
