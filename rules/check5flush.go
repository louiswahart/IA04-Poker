package rules

import (
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

func isFlush5(hand []agt.Card) bool {
	b := true
	color := hand[0].Color
	for i := range hand {
		if hand[i].Color != color {
			b = false
		}
	}
	return b
}

func Check5Flush(hand []agt.Card) (sc int) {
	list := Combi5Cards()
	for i, v := range list {
		if Straight(v) {
			RemoveIndex(list, i)
		}
	}
	list = list[:len(list)-9] //enlève les 9 suites de nombre parmi les combinaisons
	s := maxRange5Straight + 1
	for i := range list {
		list[i] = append(list[i], s)
		s++
	}
	var tab []int
	for _, v := range hand {
		tab = append(tab, v.Value)
	}
	for _, v := range list {
		if EqualHand(tab, v[:5]) {
			sc = v[5]
			return
		}
	}
	return
}
