package rules

import (
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

func isTwoPair5(hand []agt.Card) bool {
	var mapRep = make(map[int]int)
	for i := range hand {
		mapRep[hand[i].Value] = 0
	}
	for _, v := range hand {
		mapRep[v.Value]++
	}
	var n int
	for i := range mapRep {
		if mapRep[i] == 2 {
			n++
		}
	}
	return n == 2
}

func Check5TwoPair(hand []agt.Card) (sc int) {
	var temp [][]int
	for i := 0; i < 13*66; i++ {
		temp = append(temp, []int{})
	}
	var index int
	for j := 1; j < 13; j++ {
		for h := j + 1; h < 14; h++ {
			for i := 1; i < 14; i++ {
				if i != j && i != h {
					temp[index] = append(temp[index], j, j, h, h, i)
					index++
				}
			}
		}
	}
	s := maxRange5Pair + 1
	for i := range temp {
		temp[i] = append(temp[i], s)
		s++
	}
	var tab []int
	for _, v := range hand {
		tab = append(tab, v.Value)
	}
	for i := range temp {
		if EqualHand(tab, temp[i][:5]) {
			sc = temp[i][5]
			return
		}
	}
	return
}
