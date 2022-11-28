package rules

import (
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gonum.org/v1/gonum/stat/combin"
)

func Check7value(Player []agt.Card, Table []agt.Card) int {
	max := Check5value([]agt.Card{}, Table)
	c := combin.Combinations(5, 4)
	for _, v := range c {
		if Check5value([]agt.Card{Player[0]}, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]], Table[v[3]]}) > max {
			max = Check5value([]agt.Card{Player[0]}, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]], Table[v[3]]})
		}
		if Check5value([]agt.Card{Player[1]}, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]], Table[v[3]]}) > max {
			max = Check5value([]agt.Card{Player[1]}, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]], Table[v[3]]})
		}
	}
	c = combin.Combinations(5, 3)
	for _, v := range c {
		if Check5value(Player, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]]}) > max {
			max = Check5value(Player, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]]})
		}
	}
	return max
}
