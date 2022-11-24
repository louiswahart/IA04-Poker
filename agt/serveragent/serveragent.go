package serveragent

import (
	"math/rand"
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/tableagent"
)

const PlayersPerTable = 5

// ------ STRUCT ------
type ServerAgent struct {
	id       int
	nbTables int
	nbGames  int
	c        chan int
	wg       *sync.WaitGroup
	players  []playeragent.PlayerAgent
	tables   []tableagent.TableAgent
}

// ------ CONSTRUCTOR ------
func NewServerAgent(id int, nbTables int, nbGames int, wg *sync.WaitGroup) *ServerAgent {
	return &ServerAgent{id: id, nbTables: nbTables, nbGames: nbGames, c: make(chan int), wg: wg, players: make([]playeragent.PlayerAgent, nbTables*PlayersPerTable), tables: make([]tableagent.TableAgent, nbTables)}
}

// ------ GETTER ------
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

func (server *ServerAgent) Start() {
	// Créer les PlayerAgents
	for i := 0; i < server.nbTables*PlayersPerTable; i++ {
		server.players[i] = *RandomPlayerAgent(i)
	}

	// Créer les TableAgents et y assigner les PlayerAgents
	for i := 0; i < server.nbTables; i++ {
		// Génération des joueurs de chaque table
		players := server.players[i*PlayersPerTable : (i+1)*PlayersPerTable]
		server.tables[i] = *tableagent.NewTableAgent(i, server.c, server.wg, players)
	}

	// Démarrer serveur http (pour l'affichage web)
	// Lancer les TableAgents
	for _, table := range server.tables {
		table.Start()
		server.wg.Add(1)
	}
	server.wg.Wait()

	// Synchroniser les TableAgents tour après tour + envoyer requêtes web pour affichage
	for turn := 0; turn < 3; turn++ {
		server.c <- turn
		server.wg.Wait()
	}
	server.wg.Wait()
	// Récupérer la liste des TableAgents et recréer les TableAgents
}

func RandomPlayerAgent(id int) *playeragent.PlayerAgent {
	return playeragent.NewPlayerAgent(id, nil, rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100))
}
