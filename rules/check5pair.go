package rules

import (
	"sort"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gonum.org/v1/gonum/stat/combin"
)

func Combi3Cards() [][]int {
	temp := combin.Combinations(14, 3)
	list := temp[78:]
	for _, v := range list {
		sort.Slice(v, func(i int, j int) bool {
			return v[i] > v[j]
		})
	}
	sort.Slice(list, func(i int, j int) bool {
		if list[i][0] != list[j][0] {
			return list[i][0] < list[j][0]
		}
		if list[i][1] != list[j][1] {
			return list[i][1] < list[j][1]
		}
		return list[i][2] < list[j][2]
	})
	return list
}

func isPair5(hand []agt.Card) bool {
	var mapRep = make(map[int]int)
	for i := range hand {
		mapRep[hand[i].Value] = 0
	}
	for _, v := range hand {
		mapRep[v.Value]++
	}
	for i := range mapRep {
		if mapRep[i] == 2 {
			return true
		}
	}
	return false
}

func Check5Pair(hand []agt.Card) (sc int) {
	var list [][]int
	for i := 1; i < 14; i++ {
		list = append(list, []int{i, i})
	}
	var b bool
	var index int
	combin := Combi3Cards()
	var temp [][]int
	for i := 0; i < len(combin)*10; i++ {
		temp = append(temp, []int{})
	}
	for j := range list {
		for _, v := range combin {
			for i := range v {
				if v[i] == list[j][0] {
					b = true
				}
			}
			if !b {
				temp[index] = append(temp[index], list[j][0], list[j][0])
				for i := range v {
					temp[index] = append(temp[index], v[i])
				}
				index++
			}
			b = false
		}
	}
	s := maxRange5HighCard + 1
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
