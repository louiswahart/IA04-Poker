package tableagent

import (
	"math/rand"
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
)

var baseBlind = 100

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

func (table *TableAgent) Start(roundNb int) {
	for i, p := range table.players {
		table.cp[i] = make(chan agt.PlayerMessage)
		p.SetC(table.cp[i])
	}

	// Définir montant petite et grosse blindes
	smallBlind := baseBlind * roundNb / 10

	deck := table.startNewPot(roundNb)

	// Faire le tour de table pour les mises / Preflop
	for i, p := range table.players {
		p.C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "joue",
			Cards: nil, CurrentBet: smallBlind * 2, Order: i + 2, NbTokens: p.CurrentTokens()}, Response: 0}
	}

	// Attendre prochains tours
	for <-table.c > -1 {
		table.doNextTurn(deck)
	}
}

func (table *TableAgent) doNextTurn(deck []agt.Card) {
	cards := []agt.Card{}

	switch table.currentTurn {
	case 1: // Premier tour, retourner 3 cartes
		cards = deck[:3]
	case 2: // Deuxième tour, retourner 1 carte de plus
		cards = deck[:4]
	case 3: // Troisième tour, retourner 1 carte de plus
		cards = deck[:5]
	}

	for i, p := range table.players {
		p.C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "joue",
			Cards: cards, CurrentBet: 0, Order: i + 2, NbTokens: p.CurrentTokens()}, Response: 0}
	}

	table.currentTurn++
	table.wg.Done()
}

func (table *TableAgent) startNewPot(roundNb int) []agt.Card {
	deck := table.newShuffledDeck()

	// On garde uniquement les cartes des joueurs (par 2), et les cartes qui seront posées sur la table
	selection := deck[:len(table.players)*2+5]

	smallBlind := baseBlind * roundNb / 10
	bigBlind := 2 * smallBlind

	// Distribuer les cartes aux joueurs
	for i, p := range table.players {
		if i == 0 {
			p.SetCurrentTokens(p.CurrentTokens() - smallBlind)
		}
		if i == 1 {
			p.SetCurrentTokens(p.CurrentTokens() - bigBlind)
		}
		p.C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "distrib",
			Cards: selection[i*2 : (i+1)*2], CurrentBet: 0, Order: i + 2, NbTokens: p.CurrentTokens()}, Response: 0}
	}
	return selection[len(selection)-5:]
}

func (table *TableAgent) newShuffledDeck() (deck []agt.Card) {
	// Instancier le deck
	deck = make([]agt.Card, 52)

	// Initialisation du deck de cartes
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			deck[i+13*j] = agt.Card{Value: i, Color: agt.CardSuit(j)}
		}
	}

	// Bubble swaps, pour mélanger le paquet de cartes
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	return deck
}
