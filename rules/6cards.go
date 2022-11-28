package rules

import (
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gonum.org/v1/gonum/stat/combin"
)

func Check6value(Player []agt.Card, Table []agt.Card) int {
	max := Check5value([]agt.Card{Player[0]}, Table)
	if Check5value([]agt.Card{Player[1]}, Table) > max {
		max = Check5value([]agt.Card{Player[1]}, Table)
	}
	c := combin.Combinations(4, 3)
	for _, v := range c {
		if Check5value(Player, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]]}) > max {
			max = Check5value(Player, []agt.Card{Table[v[0]], Table[v[1]], Table[v[2]]})
		}
	}
	return max
}
