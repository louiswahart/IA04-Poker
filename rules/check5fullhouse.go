package rules

import (
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

func isFullHouse5(hand []agt.Card) bool {
	var mapRep = make(map[int]int)
	for i := range hand {
		mapRep[hand[i].Value] = 0
	}
	for _, v := range hand {
		mapRep[v.Value]++
	}
	for i := range mapRep {
		if (mapRep[i] == 2 || mapRep[i] == 3) && len(mapRep) == 2 {
			return true
		}
	}
	return false
}

func Check5FullHouse(hand []agt.Card) (sc int) {
	var temp [][]int
	for i := 0; i < 13*12; i++ {
		temp = append(temp, []int{})
	}
	var index int
	for j := 1; j < 14; j++ {
		for h := 1; h < 14; h++ {
			if j != h {
				temp[index] = append(temp[index], j, j, j, h, h)
				index++
			}
		}
	}
	s := maxRange5Flush + 1
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
