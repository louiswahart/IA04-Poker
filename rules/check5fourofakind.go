package rules

import (
	"fmt"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

func isFourOfAKind5(hand []agt.Card) bool {
	var mapRep = make(map[int]int)
	for i := range hand {
		mapRep[hand[i].Value] = 0
	}
	for _, v := range hand {
		mapRep[v.Value]++
	}
	for i := range mapRep {
		if mapRep[i] == 4 {
			return true
		}
	}
	return false
}

func Check5FourofAKind(hand []agt.Card) (sc int) {
	var temp [][]int
	for i := 0; i < 13*12; i++ {
		temp = append(temp, []int{})
	}
	var index int
	for j := 1; j < 14; j++ {
		for i := 1; i < 14; i++ {
			if i != j {
				temp[index] = append(temp[index], j, j, j, j, i)
				index++
			}
		}
	}
	s := maxRange5FullHouse + 1
	for i := range temp {
		temp[i] = append(temp[i], s)
		s++
	}
	var tab []int
	for _, v := range hand {
		tab = append(tab, v.Value)
	}
	fmt.Println(temp)
	for _, v := range temp {
		if EqualHand(tab, v[:5]) {
			sc = v[5]
			return
		}
	}
	return
}
