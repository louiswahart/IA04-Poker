package serveragent

import "sync"

// ------ STRUCT ------
// type PlayerAgent struct {
// 	id             int
// 	c              chan (agt.PlayerMessage)
// 	bluff          int // Caractéristique décrivant la tendance d'un joueur à jouer (continuer de miser) alors qu'il ne devrait peut être pas
// 	risk           int // Caractéristique décrivant la tendance d'un joueur à jouer (continuer de miser) selon la puissance de sa main (plus risk est elevé, plus il jouera même avec une main faible)
// 	aggressiveness int // Caractéristique décrivant à quel point le joueur fait monter la mise quand il joue
// 	timidity       int // Caractéristique décrivant la tendance d'un joueur à juste suivre la mise actuelle ou à augmenter la mise.
// 	currentTokens  int
// 	totalTokens    int
// 	cards          []agt.Card
// 	currentBet     int
// }

type ServerAgent struct {
	id       int
	nbTables int
	nbGames  int
	wg       *sync.WaitGroup
}

// ------ CONSTRUCTOR ------
// func NewPlayerAgent(id int, c chan (agt.PlayerMessage), bluff int, risk int, aggressiveness int, timidity int) *PlayerAgent {
// 	return &PlayerAgent{id: id, c: c, bluff: bluff, risk: risk, aggressiveness: aggressiveness, timidity: timidity}
// }

func NewServerAgent(id int, nbTables int, nbGames int, wg *sync.WaitGroup) *ServerAgent {
	return &ServerAgent{id: id, nbTables: nbTables, nbGames: nbGames, wg: wg}
}

// func (player *PlayerAgent) Id() int {
// 	return player.id
// }

func (server *ServerAgent) Id() int {
	return server.id
}

func (server *ServerAgent) NbTables() int {
	return server.nbTables
}

func (server *ServerAgent) NbGames() int {
	return server.nbGames
}

func (server *ServerAgent) Wg() *sync.WaitGroup {
	return server.wg
}
