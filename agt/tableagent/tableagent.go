package tableagent

import (
	"log"
	"math/rand"
	"sync"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
	"gitlab.utc.fr/nivoixpa/ia04-poker/rules"
)

var baseBlind = 5

// ------ STRUCT ------
type TableAgent struct {
	id               int
	c                <-chan int                // Permet de savoir quel est le prochain tour à jouer
	wg               *sync.WaitGroup           // Permet de dire au serveur quand on a fini chaque tour
	players          []playeragent.PlayerAgent // Peut être besoin d'utiliser des pointeurs ?
	currentTurn      int                       // Le tour actuel
	cp               []chan agt.PlayerMessage  // La liste des channels vers les joueurs
	gameNb           int                       // Le numéro de la partie actuelle
	currentBet       int                       // La mise actuelle de la table
	currentTableBets []int                     // La liste des mises actuelles de chaque joueur
	smallBlindIndex  int                       // L'indice auquel se trouve la small blind (augmente de 1 à chaque nouvelle partie)
	auxPots          []int                     // Pots annexes
	gameInProgress   bool
}

// ------ CONSTRUCTOR ------
func NewTableAgent(id int, c <-chan int, wg *sync.WaitGroup, players []playeragent.PlayerAgent) *TableAgent {
	return &TableAgent{id: id, c: c, wg: wg, players: players, currentTurn: 0, cp: make([]chan agt.PlayerMessage, len(players)), gameNb: 0, currentBet: 0, currentTableBets: make([]int, len(players)), smallBlindIndex: 0, auxPots: make([]int, len(players)), gameInProgress: true}
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
	log.Printf("[Table %v] Lancement de la table %v, channel {%v}", table.id, table.id, table.c)
	for i := range table.players {
		table.cp[i] = make(chan agt.PlayerMessage)
		table.players[i].SetC(table.cp[i])
		go table.players[i].Start()
		table.players[i].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "nouvelle",
			Cards: nil, CurrentBet: 0, Order: 1, NbTokens: 10000}, Response: 0}
	}
	table.wg.Done()

	deck := []agt.Card{}

	for turnNb := range table.c {
		if turnNb < 0 {
			return
		}

		// On met à jour le tour actuel récupéré à travers le channel
		table.currentTurn = turnNb
		winners := make([]int, 0)
		if turnNb == 0 {
			table.gameNb++
			table.gameInProgress = true
			deck = table.startNewPot(table.gameNb)
			log.Printf("\n[Table %v] Preflop", table.id)
			table.doRoundTable(nil)
			winner := table.checkDefaultWinner()
			if winner != -1 {
				winners = append(winners, winner)
				table.gameInProgress = false
				table.distribEarnings(winners)
			}
		} else if table.gameInProgress {
			table.doTurn(deck)
			winner := table.checkDefaultWinner()
			if winner != -1 {
				winners = append(winners, winner)
				table.gameInProgress = false
				table.distribEarnings(winners)
			} else if turnNb == 3 {
				winners = table.checkWinnersByScore(deck)
				table.distribEarnings(winners)
			}
		}
		table.wg.Done()
	}
}

func (table *TableAgent) doTurn(deck []agt.Card) {
	log.Printf("[Table %v] Tour %v", table.id, table.currentTurn)
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

	table.doRoundTable(cards)

	table.currentTurn++
}

func (table *TableAgent) startNewPot(roundNb int) []agt.Card {
	deck := table.newShuffledDeck()

	//table.totalPot = 0
	for i := range table.players {
		table.auxPots[i] = 0
		table.currentTableBets[i] = 0
	}

	// On garde uniquement les cartes des joueurs (par 2), et les cartes qui seront posées sur la table
	selection := deck[:len(table.players)*2+5]

	smallBlind := baseBlind * (roundNb/10 + 1)
	bigBlind := 2 * smallBlind

	log.Printf("[Table %v] Paiement des blindes", table.id)
	// Mises obligatoires des petites et grosses blindes
	table.players[table.smallBlindIndex].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "mise",
		Cards: nil, CurrentBet: smallBlind, Order: 0, NbTokens: 0}, Response: 0}
	table.auxPots[table.smallBlindIndex] += smallBlind
	resp := <-table.players[table.smallBlindIndex].C()
	table.currentTableBets[table.smallBlindIndex] = resp.Response
	table.currentBet = resp.Response

	table.players[(table.smallBlindIndex+1)%len(table.players)].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "mise",
		Cards: nil, CurrentBet: bigBlind, Order: 1, NbTokens: 0}, Response: 0}
	table.auxPots[table.smallBlindIndex+1] += bigBlind
	resp = <-table.players[(table.smallBlindIndex+1)%len(table.players)].C()
	table.currentTableBets[(table.smallBlindIndex+1)%len(table.players)] = resp.Response
	if resp.Response > table.currentBet {
		table.currentBet = resp.Response
	}

	// Distribuer les cartes aux joueurs
	log.Printf("[Table %v] Distribution des cartes", table.id)
	for i := range table.players {
		table.players[i].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "distrib",
			Cards: selection[i*2 : (i+1)*2], CurrentBet: 0, Order: (i - 2) % len(table.players), NbTokens: 0}, Response: 0}
	}
	return selection[len(selection)-5:]
}

func (table *TableAgent) newShuffledDeck() (deck []agt.Card) {
	// Instancier le deck
	deck = make([]agt.Card, 52)

	// Initialisation du deck de cartes
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			deck[i+13*j] = agt.Card{Value: i + 1, Color: agt.CardSuit(j)}
		}
	}

	// Bubble swaps, pour mélanger le paquet de cartes
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	return deck
}

func (table *TableAgent) doRoundTable(cards []agt.Card) {
	var cnt int = 0
	var i int = (table.smallBlindIndex + 2) % len(table.players) // On commence le tour juste après la grosse blinde
	for cnt+1 < len(table.players) || table.currentTableBets[i] < table.currentBet {
		if table.currentTableBets[i] != -1 {
			log.Printf("[Table %v] --------- %v ---------", table.id, table.currentTableBets)
			//bet := table.currentBet - table.currentTableBets[i]
			table.players[i].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "joue",
				Cards: cards, CurrentBet: table.currentBet, Order: i, NbTokens: table.players[i].CurrentTokens()}, Response: 0}

			resp := <-table.players[i].C()

			if resp.Response == -1 {
				// table.totalPot += table.currentTableBets[i]
				table.auxPots[i] += table.currentTableBets[i]
				table.currentTableBets[i] = -1
			} else {
				table.currentTableBets[i] += resp.Response
				if table.currentTableBets[i] > table.currentBet {
					table.currentBet = table.currentTableBets[i]
				}
			}
		}
		i = (i + 1) % len(table.players)
		cnt++
	}
	log.Printf("[Table %v] --------- %v ---------", table.id, table.currentTableBets)
	for i := range table.currentTableBets {
		if table.currentTableBets[i] != -1 {
			table.auxPots[i] += table.currentTableBets[i]
			table.currentTableBets[i] = 0
		}
	}
	table.currentBet = 0
}

func (table *TableAgent) checkDefaultWinner() int {
	stillIn := 0
	index := -1
	for i := range table.players {
		if table.currentTableBets[i] != -1 {
			index = i
			stillIn++
		}
	}
	if stillIn > 1 {
		log.Printf("[Table %v] Reste %v joueurs dans la partie", table.id, stillIn)
		return -1
	} else {
		log.Printf("[Table %v] Joueur %v a gagné la partie", table.id, index)
		return index
	}
}

func (table *TableAgent) checkWinnersByScore(deck []agt.Card) (winners []int) {
	var maxScore int = -1
	// Should only be executed at the end of the very last turn
	for i := range table.players {
		if table.currentTableBets[i] > -1 {
			score := rules.CheckCombinations(table.players[i].Card(), deck)
			if score > maxScore {
				maxScore = score
				winners = make([]int, 1, len(table.players))
				winners[0] = i

			} else if score == maxScore {
				winners = append(winners, i)
			}
		}
	}
	return winners
}

func (table *TableAgent) distribEarnings(winners []int) {
	reste := make([]int, len(table.players))
	gain := make([]int, len(table.players))
	for player := range table.players {
		reste[player] = table.auxPots[player]
		for winner := range winners {
			gain[winner] += rules.Min(table.auxPots[winner], table.auxPots[player]/len(winners))
			reste[player] -= gain[winner]
		}
		gain[player] += reste[player]
	}
	for player := range table.players {
		if gain[player] > 0 {
			table.players[player].C() <- agt.PlayerMessage{Request: agt.RequestMessage{Instruction: "gain",
				Cards: nil, CurrentBet: 0, Order: 0, NbTokens: gain[player]}, Response: 0}
		}
	}
	table.wg.Done()
}
