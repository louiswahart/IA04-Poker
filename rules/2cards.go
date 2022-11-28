package rules

import (
	"sort"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gonum.org/v1/gonum/stat/combin"
)

func Combi2Cards() [][]int {
	temp := combin.Combinations(14, 2)
	list := temp[13:]
	for _, v := range list {
		sort.Slice(v, func(i int, j int) bool {
			return v[i] > v[j]
		})
	}
	sort.Slice(list, func(i int, j int) bool {
		if list[i][0] != list[j][0] {
			return list[i][0] < list[j][0]
		}
		return list[i][1] < list[j][1]
	})
	return list
}

func Check2Pair(val int) (score int) {
	var scores = make([][]int, 13)
	for i := range scores {
		scores[i] = []int{0}
	}
	v := 1
	s := maxRange2HighCard + 1
	for i := 0; i < 13; i++ {
		scores[i] = []int{v, v, s}
		v++
		s++
	}
	for _, v := range scores {
		if v[0] == val {
			score = v[2]
			return
		}
	}
	return
}

func Check2HighCard(p []agt.Card, t []agt.Card) (score int) {
	list := Combi2Cards()
	s := 1
	for i := range list {
		list[i] = append(list[i], s)
		s++
	}
	var tab []int
	for _, v := range p {
		tab = append(tab, v.Value)
	}
	sort.Slice(tab, func(i int, j int) bool {
		return tab[i] > tab[j]
	})
	for _, v := range list {
		if EqualHand(tab, v[:2]) {
			score = v[2]
			return
		}
	}
	return
}

func Check2value(Player []agt.Card, Table []agt.Card) (sc int) {
	if Player[0].Value == Player[1].Value {
		sc = Check2Pair(Player[0].Value)
	} else {
		sc = Check2HighCard(Player, Table)
	}
	return
}
