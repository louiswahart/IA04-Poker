package rules

import (
	"sort"

	"gonum.org/v1/gonum/stat/combin"
)

func Combi5Cards() [][]int {
	temp := combin.Combinations(14, 5)
	list := temp[715:]
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
		if list[i][2] != list[j][2] {
			return list[i][2] < list[j][2]
		}
		if list[i][3] != list[j][3] {
			return list[i][3] < list[j][3]
		}
		return list[i][4] < list[j][4]
	})
	return list
}

func EqualHand(tab1 []int, tab2 []int) bool {
	var mapRep = make(map[int]int)
	for i := range tab1 {
		mapRep[tab1[i]] = 0
	}
	for _, v := range tab1 {
		mapRep[v]++
	}
	var mapRep2 = make(map[int]int)
	for i := range tab2 {
		mapRep2[tab2[i]] = 0
	}
	for _, v := range tab2 {
		mapRep2[v]++
	}
	for i := range mapRep {
		if mapRep[i] != mapRep2[i] {
			return false
		}
	}
	return true
}

func RemoveIndex(c [][]int, index int) {
	for i := index; i < len(c)-1; i++ {
		c[i] = c[i+1]
	}
}

func Straight(combi []int) bool {
	for i := 0; i < len(combi)-1; i++ {
		if combi[i] != combi[i+1]+1 {
			return false
		}
	}
	return true
}
