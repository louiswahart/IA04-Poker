package rules

import "gitlab.utc.fr/nivoixpa/ia04-poker/agt"

const (
	maxRange2HighCard      int = 78
	maxRange2Pair          int = 91
	maxRange2              int = 91
	maxRange5HighCard      int = 1278
	maxRange5Pair          int = 4138
	maxRange5TwoPair       int = 4996
	maxRange5ThreeofAKind  int = 5854
	maxRange5Straight      int = 5864
	maxRange5Flush         int = 7142
	maxRange5FullHouse     int = 7298
	maxRange5FourofAKind   int = 7454
	maxRange5StraightFlush int = 7464
	maxRange5RoyalFlush    int = 7465
	maxRange5              int = 7465
	maxRange6              int = 7465
	maxRange7              int = 7465
)

func CheckCombinations(PlayerCards []agt.Card, TableCards []agt.Card) (score int) {
	var PlayerCopy = make([]agt.Card, len(PlayerCards))
	var TableCopy = make([]agt.Card, len(TableCards))
	copy(PlayerCopy, PlayerCards)
	copy(TableCopy, TableCards)
	for i := range PlayerCopy {
		if PlayerCopy[i].Value == 1 {
			PlayerCopy[i].Value = 13
		} else {
			PlayerCopy[i].Value -= 1
		}
	}
	for i := range TableCopy {
		if TableCopy[i].Value == 1 {
			TableCopy[i].Value = 13
		} else {
			TableCopy[i].Value -= 1
		}
	}
	switch len(TableCopy) {
	case 0:
		score = Check2value(PlayerCopy, TableCopy)
	case 3:
		score = Check5value(PlayerCopy, TableCopy)
	case 4:
		score = Check6value(PlayerCopy, TableCopy)
	case 5:
		score = Check7value(PlayerCopy, TableCopy)
	default:
		return -1
	}
	return
}

func MaxRange(number int) int {
	switch number {
	case 2:
		return maxRange2
	case 5, 6, 7:
		return maxRange5
	default:
		return -1
	}
}
