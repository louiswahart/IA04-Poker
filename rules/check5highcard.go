package rules

import "gitlab.utc.fr/nivoixpa/ia04-poker/agt"

func Check5HighCard(hand []agt.Card) (sc int) {
	list := Combi5Cards()
	for i, v := range list {
		if Straight(v) {
			RemoveIndex(list, i)
		}
	}
	list = list[:len(list)-9] //enlève les 9 suites de nombre parmi les combinaisons
	s := 1
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
