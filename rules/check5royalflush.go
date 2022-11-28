package rules

import "gitlab.utc.fr/nivoixpa/ia04-poker/agt"

func isRoyalFlush5(hand []agt.Card) bool {
	if !isStraightFlush5(hand) {
		return false
	}
	if hand[0].Value != 13 {
		return false
	}
	return true
}

func Check5RoyalFlush(hand []agt.Card) (sc int) {
	return maxRange5StraightFlush + 1
}
