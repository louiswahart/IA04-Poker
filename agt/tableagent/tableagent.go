package tableagent

import (
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
)

// ------ STRUCT ------
type TableAgent struct {
	id          int
	c           <-chan int                // Permet de savoir quel est le prochain tour à jouer
	wg          *sync.WaitGroup           // Permet de dire au serveur quand on a fini chaque tour
	players     []playeragent.PlayerAgent // Peut être besoin d'utiliser des pointeurs ?
	currentTurn int
}

// ------ CONSTRUCTOR ------
func NewTableAgent(id int, c <-chan int, wg *sync.WaitGroup, players []playeragent.PlayerAgent) *TableAgent {
	return &TableAgent{id: id, c: c, wg: wg, players: players, currentTurn: 0}
}

// ------ GETTER ------
func (table *TableAgent) Id() int {
	return table.id
}

func (table *TableAgent) Players() []playeragent.PlayerAgent {
	return table.players
}

func (table *TableAgent) CurrentTurn() int {
	return table.currentTurn
}
