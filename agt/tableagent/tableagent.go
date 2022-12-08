package tableagent

import (
	"fmt"
	"math/rand"
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
)

var baseBlind = 5

// ------ STRUCT ------
type TableAgent struct {
	id          int
	c           <-chan int                // Permet de savoir quel est le prochain tour à jouer
	wg          *sync.WaitGroup           // Permet de dire au serveur quand on a fini chaque tour
	players     []playeragent.PlayerAgent // Peut être besoin d'utiliser des pointeurs ?
	currentTurn int
	cp          []chan agt.PlayerMessage
	gameNb      int
}

// ------ CONSTRUCTOR ------
func NewTableAgent(id int, c <-chan int, wg *sync.WaitGroup, players []playeragent.PlayerAgent) *TableAgent {
	return &TableAgent{id: id, c: c, wg: wg, players: players, currentTurn: 0, cp: make([]chan agt.PlayerMessage, len(players)), gameNb: 0}
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
	fmt.Println("Start de la table", table.id, ", channel:", table.c)
	for i := range table.players {
		table.cp[i] = make(chan agt.PlayerMessage)
		table.players[i].SetC(table.cp[i])
		go table.players[i].Start()
		table.players[i].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "nouvelle",
			Cards: nil, CurrentBet: 0, Order: 1, NbTokens: 10000}, Response: 0}
	}
	table.wg.Done()

	deck := []agt.Card{}

	// Attendre prochains tours
	fmt.Println("On attend")
	for turnNb := range table.c {
		fmt.Println("On a bien lu le tour:", turnNb)
		if turnNb < 0 {
			return
		}
		// On met à jour le tour actuel récupéré à travers le channel
		table.currentTurn = turnNb
		if turnNb == 0 {
			table.gameNb++
			fmt.Println("Avant création du pot")
			deck = table.startNewPot(table.gameNb)
			fmt.Println("avant preflop")
			table.doPreFlop()
		} else {
			table.doTurn(deck)
		}
		fmt.Println("Juste avant le done")
		table.wg.Done()
	}
}

func (table *TableAgent) doTurn(deck []agt.Card) {
	if len(deck) == 0 {
		return
	}
	cards := []agt.Card{}

	switch table.currentTurn {
	case 1: // Premier tour, retourner 3 cartes
		cards = deck[:3]
	case 2: // Deuxième tour, retourner 1 carte de plus
		cards = deck[:4]
	case 3: // Troisième tour, retourner 1 carte de plus
		cards = deck[:5]
	}

	// Code non fini, le code actuel s'arrête après un tour
	// Mais tant que tout le monde n'a pas misé ou s'est couché, le tour ne doit pas se finir
	for i, p := range table.players {
		p.C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "joue",
			Cards: cards, CurrentBet: 0, Order: (i - 2) % len(table.players), NbTokens: p.CurrentTokens()}, Response: 0}
	}

	table.currentTurn++
}

func (table *TableAgent) startNewPot(roundNb int) []agt.Card {
	deck := table.newShuffledDeck()

	// On garde uniquement les cartes des joueurs (par 2), et les cartes qui seront posées sur la table
	selection := deck[:len(table.players)*2+5]

	smallBlind := baseBlind * roundNb / 10
	bigBlind := 2 * smallBlind

	// Mises obligatoires des petites et grosses blindes
	table.players[0].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "mise",
		Cards: nil, CurrentBet: smallBlind, Order: 0, NbTokens: 0}, Response: 0}
	table.players[1].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "mise",
		Cards: nil, CurrentBet: bigBlind, Order: 1, NbTokens: 0}, Response: 0}

	// Distribuer les cartes aux joueurs
	fmt.Println("on distribue les cartes")
	for i := range table.players {
		table.players[i].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "distrib",
			Cards: selection[i*2 : (i+1)*2], CurrentBet: 0, Order: (i - 2) % len(table.players), NbTokens: 0}, Response: 0}
	}
	fmt.Println("on a fini de distribuer")
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

func (table *TableAgent) doPreFlop() {
	// Définir montant petite et grosse blindes
	bigBlind := 2 * baseBlind * table.gameNb / 10

	// Faire le tour de table pour les mises / Preflop
	for i := range table.players {
		fmt.Println("Envoi de message au joueur")
		table.players[i].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "joue",
			Cards: nil, CurrentBet: bigBlind, Order: i + 2, NbTokens: table.players[i].CurrentTokens()}, Response: 0}
	}
}
