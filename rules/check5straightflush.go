package rules

import "gitlab.utc.fr/nivoixpa/ia04-poker/agt"

func isStraightFlush5(hand []agt.Card) bool {
	if !isStraight5(hand) {
		return false
	}
	if !isFlush5(hand) {
		return false
	}
	return true
}

func Check5StraightFlush(hand []agt.Card) (sc int) {
	var temp [][]int
	for i := 0; i < 10; i++ {
		temp = append(temp, []int{})
	}
	temp[0] = []int{13, 1, 2, 3, 4}
	for i := 1; i < 10; i++ {
		temp[i] = append(temp[i], i, i+1, i+2, i+3, i+4)
	}
	s := maxRange5FourofAKind + 1
	for i := range temp {
		temp[i] = append(temp[i], s)
		s++
	}
	var tab []int
	for _, v := range hand {
		tab = append(tab, v.Value)
	}
	for _, v := range temp {
		if EqualHand(tab, v[:5]) {
			sc = v[5]
			return
		}
	}
	return
}
