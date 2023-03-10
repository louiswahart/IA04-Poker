package serveragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/playeragent"
	"gitlab.utc.fr/nivoixpa/ia04-poker/agt/tableagent"
)

const PlayersPerTable = 5

// ------ STRUCT ------
type ServerAgent struct {
	addr     string
	id       int
	nbTables int
	nbGames  int
	c        []chan int
	wg       *sync.WaitGroup
	sync.Mutex
	players []playeragent.PlayerAgent
	tables  []tableagent.TableAgent
	turn    int
	games   int
}

// ------ CONSTRUCTOR ------
func NewServerAgent(addr string, id int) *ServerAgent {
	return &ServerAgent{addr: addr, id: id, nbTables: 50, nbGames: 50, turn: -1, games: 0}
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

func (*ServerAgent) checkMethod(method string, r *http.Request) bool {
	return r.Method == method
}

func (*ServerAgent) decodeRequestPlay(r *http.Request) (req agt.RequestPlay, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*ServerAgent) decodeRequestUpdate(r *http.Request) (req agt.RequestUpdate, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*ServerAgent) decodeRequestgetTable(r *http.Request) (req agt.RequestgetTable, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*ServerAgent) decodeRequestgetPlayer(r *http.Request) (req agt.RequestgetPlayer, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*ServerAgent) decodeRequestchangeStats(r *http.Request) (req agt.RequestchangeStats, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (server *ServerAgent) play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	server.Lock()
	defer server.Unlock()
	// D??tecter le type de requ??te
	// Si OPTION
	if server.checkMethod("OPTIONS", r) {
		log.Println("[Serveur] Requ??te de connexion")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		return
	} else if server.checkMethod("POST", r) {
		// Si POST
		log.Println("[Serveur] Requ??te en play")

		if server.tables != nil {
			server.resetParameters()
		}

		// Mise ?? jour du seed
		rand.Seed(time.Now().UnixNano())

		// D??code de la requ??te pour v??rifier qu'elle correspond ?? la bonne action
		req, err := server.decodeRequestPlay(r)
		if err != nil {
			log.Printf("[Serveur] Err %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "erreur %v", err)
			return
		}
		if req.Req != "play" {
			return
		}

		// R??cup??ration des informations : nbTable et nbGame
		if req.NbGames <= 0 {
			log.Println("[Serveur] Test")
			return
		}
		if req.NbTables <= 0 {
			log.Println("[Serveur] Test")
			return
		}

		server.nbGames = req.NbGames
		server.nbTables = req.NbTables

		log.Printf("[Serveur] Nombre de tables re??ues : %v | Nombre de games re??ues : %v\n", server.nbTables, server.nbGames)

		// Cr??ation des channels
		server.c = make([]chan int, server.nbTables)

		log.Printf("[Serveur] Channels cr????s\n")

		// Cr??ation du waitgroup
		var wg sync.WaitGroup
		server.wg = &wg

		log.Printf("[Serveur] WaitGroup cr????\n")

		// Cr??ation du tableau qui contiendra les joueurs
		server.players = make([]playeragent.PlayerAgent, server.nbTables*PlayersPerTable)

		// Cr??ation du tableau qui contiendra les tables
		server.tables = make([]tableagent.TableAgent, server.nbTables)

		// Cr??er les PlayerAgents
		log.Printf("[Serveur] Cr??ation des joueurs")
		for i := 0; i < server.nbTables*PlayersPerTable; i++ {
			server.players[i] = *RandomPlayerAgent(i)
		}

		// Cr??er les TableAgents et y assigner les PlayerAgents
		log.Printf("[Serveur] Cr??ation des tables")
		for i := 0; i < server.nbTables; i++ {
			// G??n??ration des joueurs de chaque table
			players := server.players[i*PlayersPerTable : (i+1)*PlayersPerTable]
			server.c[i] = make(chan int)
			server.tables[i] = *tableagent.NewTableAgent(i, server.c[i], server.wg, players)
		}

		// Lancer les TableAgents
		log.Printf("[Serveur] Lancement des tables")
		for i := range server.tables {
			server.wg.Add(1)
			go server.tables[i].Start()
		}
		server.wg.Wait()

		// Envoyer que tout est bon
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Cr??ation des joueurs et des tables effectu??e !")
	}
}

func RandomPlayerAgent(id int) *playeragent.PlayerAgent {
	return playeragent.NewPlayerAgent(id, nil, rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100))
}

// Fonction de mise ?? jour lors d'un changement de table
func (server *ServerAgent) getTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	server.Lock()
	defer server.Unlock()
	// D??tecter le type de requ??te
	// Si OPTION
	if server.checkMethod("OPTIONS", r) {
		log.Println("[Serveur] Requ??te de connexion")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		return
	} else if server.checkMethod("POST", r) {
		// D??code de la requ??te pour v??rifier que correspond ?? la bonne action
		req, err := server.decodeRequestgetTable(r)
		if err != nil {
			log.Printf("[Serveur] Err %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "erreur %v", err)
			return
		}
		if req.Req != "getTable" {
			return
		}
		if req.Table < 0 && req.Table >= server.nbTables {
			return
		}

		// R??cup??ration des informations a envoyer
		ids := make([]int, 5)
		blinds := server.tables[req.Table].Blinds()
		tokens := make([]int, 5)
		bets := make([]int, 5)
		totalBets := make([]int, 0)
		actions := make([]string, 5)
		winners := make([]bool, 5)
		pot := 0
		playerCards := make([][]agt.Card, 5)
		tableCards := make([]agt.Card, 0)
		var win bool
		tableCards = append(tableCards, server.tables[req.Table].Cards()...)
		result := make([]int, 0)
		for _, a := range server.tables[req.Table].AuxPots() {
			pot += a
			totalBets = append(totalBets, a)
			result = append(result, -a)
		}
		for i, p := range server.tables[req.Table].Players() {
			ids[i] = p.Id()
			tokens[i] = p.CurrentTokens()
			bets[i] = p.CurrentBet()
			actions[i] = p.Action()
			playerCards[i] = make([]agt.Card, 2)
			playerCards[i] = p.Card()
			result[i] += p.LastEarning()
			for _, l := range server.tables[req.Table].Winners() {
				if l == i {
					win = true
				}
			}
			if win {
				winners[i] = true
			} else {
				winners[i] = false
			}
			win = false
		}

		log.Printf("[Serveur] Envoie des informations demand??es\nIds : %v\nBlinds : %v\nTokens : %v\nBets : %v\nTotal Bets : %v\nActions : %v\nCartes Joueur : %v\nGains : %v\nWinners : %v\nPot : %v\nCartes Table : %v\n", ids, blinds, tokens, bets, totalBets, actions, playerCards, result, winners, pot, tableCards)
		// Fournir l'??tat du tour actuel
		send := agt.ResponsegetTable{
			PlayersID:       ids,
			PlayersBlind:    blinds,
			PlayersToken:    tokens,
			PlayersBet:      bets,
			PlayersTotalBet: totalBets,
			PlayersActions:  actions,
			PlayersCards:    playerCards,
			PlayersGain:     result,
			PlayersWinner:   winners,
			Pot:             pot,
			TableCards:      tableCards,
		}
		data, _ := json.Marshal(send)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		log.Printf("[Serveur] Informations de changement de table bien envoy??es\n")
	}
}

// Fonction de mise ?? jour lors d'un changement de table
func (server *ServerAgent) getPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	server.Lock()
	defer server.Unlock()
	// D??tecter le type de requ??te
	// Si OPTION
	if server.checkMethod("OPTIONS", r) {
		log.Println("[Serveur] Requ??te de connexion")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		return
	} else if server.checkMethod("POST", r) {
		// D??code de la requ??te pour v??rifier que correspond ?? la bonne action
		req, err := server.decodeRequestgetPlayer(r)
		if err != nil {
			log.Printf("[Serveur] Err %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "erreur %v", err)
			return
		}
		if req.Req != "getPlayer" {
			return
		}
		if req.Player < 0 && req.Player >= len(server.players) {
			return
		}

		// R??cup??ration des informations a envoyer
		timidity := server.players[req.Player].Timidity()
		aggressiveness := server.players[req.Player].Aggressiveness()
		risk := server.players[req.Player].Risk()
		bluff := server.players[req.Player].Bluff()
		var table int
		for _, t := range server.tables {
			for _, p := range t.Players() {
				if p.Id() == req.Player {
					table = t.Id()
				}
			}
		}

		log.Printf("[Serveur] Envoie des informations demand??es\nTimidity : %v\nAggressiveness : %v\nRisk : %v\nBluff : %v\nTable : %v\n", timidity, aggressiveness, risk, bluff, table)
		// Fournir l'??tat du tour actuel
		send := agt.ResponsegetPlayer{
			Timidity:       timidity,
			Aggressiveness: aggressiveness,
			Risk:           risk,
			Bluff:          bluff,
			Table:          table,
		}
		data, _ := json.Marshal(send)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		log.Printf("[Serveur] Informations de changement de table bien envoy??es\n")
	}
}

// Fonction de changement de stats d'un joueur
func (server *ServerAgent) changeStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	server.Lock()
	defer server.Unlock()
	// D??tecter le type de requ??te
	// Si OPTION
	if server.checkMethod("OPTIONS", r) {
		log.Println("[Serveur] Requ??te de connexion")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		return
	} else if server.checkMethod("POST", r) {
		// D??code de la requ??te pour v??rifier que correspond ?? la bonne action
		req, err := server.decodeRequestchangeStats(r)
		if err != nil {
			log.Printf("[Serveur] Err %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "erreur %v", err)
			return
		}
		if req.Req != "changeStats" {
			return
		}
		if req.Player < 0 && req.Player >= len(server.players) {
			return
		}

		// Mise ?? jour des stats du joueurs
		server.players[req.Player].SetTimidity(req.Timidity)
		server.players[req.Player].SetAggressiveness(req.Aggressiveness)
		server.players[req.Player].SetRisk(req.Risk)
		server.players[req.Player].SetBluff(req.Bluff)
		log.Printf("[Serveur] Stats mises ?? jour !\n")

		// R??cup??ration des informations a envoyer
		timidity := server.players[req.Player].Timidity()
		aggressiveness := server.players[req.Player].Aggressiveness()
		risk := server.players[req.Player].Risk()
		bluff := server.players[req.Player].Bluff()

		log.Printf("[Serveur] Envoie des informations demand??es\nTimidity : %v\nAggressiveness : %v\nRisk : %v\nBluff : %v\n", timidity, aggressiveness, risk, bluff)
		// Fournir l'??tat du tour actuel
		send := agt.ResponsegetPlayer{
			Timidity:       timidity,
			Aggressiveness: aggressiveness,
			Risk:           risk,
			Bluff:          bluff,
		}
		data, _ := json.Marshal(send)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		log.Printf("[Serveur] Informations de changement de stats bien envoy??es\n")
	}
}

// Fonction de mise ?? jour des informations des tables et joueurs
func (server *ServerAgent) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	server.Lock()
	defer server.Unlock()
	// D??tecter le type de requ??te
	// Si OPTION
	if server.checkMethod("OPTIONS", r) {
		log.Println("[Serveur] Requ??te de connexion")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		return
	} else if server.checkMethod("POST", r) {

		// D??code de la requ??te pour v??rifier que correspond ?? la bonne action
		req, err := server.decodeRequestUpdate(r)
		if err != nil {
			log.Printf("[Serveur] Err %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "erreur %v", err)
			return
		}
		if req.Req != "update" {
			return
		}
		if req.Table < 0 && req.Table >= server.nbTables {
			return
		}

		log.Println("[Serveur] Demande d'update")
		server.turn++

		if server.turn >= 4 {
			server.games++
			if server.games >= server.nbGames {
				log.Printf("[Serveur] Parties termin??es, fermeture des tables\n")
				// On ferme toutes les tables
				for i := 0; i < len(server.c); i++ {
					close(server.c[i])
				}
				server.c = nil
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Parties termin??es, tables ferm??es !")
				return
			} else {
				server.turn = 0
			}
		}

		log.Printf("[Serveur] Lancement du tour : %v\n", server.turn)
		// On envoie le signal du tour aux tables ?? travers le channel associ??
		for j := 0; j < server.nbTables; j++ {
			server.wg.Add(1)
			server.c[j] <- server.turn
		}
		server.wg.Wait()

		// R??cup??ration des informations ?? envoyer
		ids := make([]int, 5)
		blinds := server.tables[req.Table].Blinds()
		tokens := make([]int, 5)
		bets := make([]int, 5)
		totalBets := make([]int, 0)
		actions := make([]string, 5)
		winners := make([]bool, 5)
		pot := 0
		playerCards := make([][]agt.Card, 5)
		tableCards := make([]agt.Card, 0)
		var win bool
		tableCards = append(tableCards, server.tables[req.Table].Cards()...)
		result := make([]int, 0)
		for _, a := range server.tables[req.Table].AuxPots() {
			pot += a
			totalBets = append(totalBets, a)
			result = append(result, -a)
		}
		for i, p := range server.tables[req.Table].Players() {
			ids[i] = p.Id()
			tokens[i] = p.CurrentTokens()
			bets[i] = p.CurrentBet()
			actions[i] = p.Action()
			playerCards[i] = make([]agt.Card, 2)
			playerCards[i] = p.Card()
			result[i] += p.LastEarning()
			for _, l := range server.tables[req.Table].Winners() {
				if l == i {
					win = true
				}
			}
			if win {
				winners[i] = true
			} else {
				winners[i] = false
			}
			win = false
		}

		log.Printf("[Serveur] Envoie des informations demand??es\nIds : %v\nBlinds : %v\nTokens : %v\nBets : %v\nTotal Bets : %v\nActions : %v\nCartes Joueur : %v\nGains : %v\nWinners : %v\nPot : %v\nCartes Table : %v\n", ids, blinds, tokens, bets, totalBets, actions, playerCards, result, winners, pot, tableCards)
		// Fournir l'??tat du tour actuel
		send := agt.ResponseUpdate{
			PlayersID:       ids,
			PlayersBlind:    blinds,
			PlayersToken:    tokens,
			PlayersBet:      bets,
			PlayersTotalBet: totalBets,
			PlayersActions:  actions,
			PlayersCards:    playerCards,
			PlayersGain:     result,
			PlayersWinner:   winners,
			Pot:             pot,
			TableCards:      tableCards,
		}

		data, _ := json.Marshal(send)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		log.Printf("[Serveur] Informations bien envoy??es\n")
	}
}

func (server *ServerAgent) resetParameters() {
	server.nbTables = 50
	server.nbGames = 50
	server.turn = -1
	server.games = 0
	log.Printf("[Serveur] Param??tres reset")

	//Fermeture des es TableAgents
	log.Printf("[Serveur] Fermeture des tables")
	for i := 0; i < len(server.c); i++ {
		close(server.c[i])
	}

	// Reset channel
	server.c = nil
	log.Printf("[Serveur] Channels reset\n")

	// Reset waitgroup
	server.wg = nil
	log.Printf("[Serveur] WaitGroup reset\n")

	// Reset tableau de joueurs
	server.players = nil
	log.Printf("[Serveur] Joueurs reset\n")

	// Reset tableau de tables
	server.tables = nil
	log.Printf("[Serveur] Tables reset\n")
}

func (serv *ServerAgent) Start() {
	// Cr??ation du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/play", serv.play)
	mux.HandleFunc("/update", serv.update)
	mux.HandleFunc("/getTable", serv.getTable)
	mux.HandleFunc("/getPlayer", serv.getPlayer)
	mux.HandleFunc("/changeStats", serv.changeStats)
	// Cr??ation du serveur
	s := &http.Server{
		Addr:           serv.addr,
		Handler:        mux,
		ReadTimeout:    -1,
		WriteTimeout:   -1,
		MaxHeaderBytes: 1 << 20}

	// Lancement du serveur
	log.Println("[Serveur] Listening on", serv.addr)
	go log.Fatal(s.ListenAndServe())
}
