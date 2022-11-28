package rules

import (
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

func isThreeOfAKind5(hand []agt.Card) bool {
	var mapRep = make(map[int]int)
	for i := range hand {
		mapRep[hand[i].Value] = 0
	}
	for _, v := range hand {
		mapRep[v.Value]++
	}
	for i := range mapRep {
		if mapRep[i] == 3 {
			return true
		}
	}
	return false
}

func Check5ThreeofAKind(hand []agt.Card) (sc int) {
	var temp [][]int
	for i := 0; i < 13*66; i++ {
		temp = append(temp, []int{})
	}
	var index int
	for h := 1; h < 14; h++ {
		for i := 1; i < 13; i++ {
			for j := i + 1; j < 14; j++ {
				if i != h && j != h {
					temp[index] = append(temp[index], h, h, h, i, j)
					index++
				}
			}
		}
	}
	s := maxRange5TwoPair + 1
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
