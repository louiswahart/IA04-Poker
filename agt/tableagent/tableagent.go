package tableagent

import (
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
)

// ------ STRUCT ------
type TableAgent struct {
	id          int
	c           <-chan int                // Permet de savoir quel est le prochain tour à jouer
	wg          *sync.WaitGroup           // Permet de dire au serveur quand on a fini chaque tour
	players     []playeragent.PlayerAgent // Peut être besoin d'utiliser des pointeurs ?
	currentTurn int
	cp          []chan agt.PlayerMessage
}

// ------ CONSTRUCTOR ------
func NewTableAgent(id int, c <-chan int, wg *sync.WaitGroup, players []playeragent.PlayerAgent) *TableAgent {
	return &TableAgent{id: id, c: c, wg: wg, players: players, currentTurn: 0, cp: make([]chan agt.PlayerMessage, len(players))}
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

func (table *TableAgent) Start() {
	for i, p := range table.players {
		table.cp[i] = make(chan agt.PlayerMessage)
		p.SetC(table.cp[i])
	}
	// Distribuer les cartes
	// Demander argent petite et grosse blinde
	// Faire le tour de table pour les mises

	// Attendre prochains tours
	for <-table.c > -1 {
		table.doNextTurn()
	}

}

func (table *TableAgent) doNextTurn() {
	switch table.currentTurn {
	case 1: // Premier tour, retourner 3 cartes
	case 2: // Deuxième tour, retourner 1 carte
	case 3: // Troisième tour, retourner 1 carte
	}
	table.currentTurn++
	table.wg.Done()
}
