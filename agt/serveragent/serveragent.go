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
	players  []playeragent.PlayerAgent
	tables   []tableagent.TableAgent
}

// ------ CONSTRUCTOR ------
func NewServerAgent(addr string, id int, nbTables int, nbGames int, wg *sync.WaitGroup) *ServerAgent {
	return &ServerAgent{addr: addr, id: id, nbTables: nbTables, nbGames: nbGames, c: make([]chan int, nbTables), wg: wg, players: make([]playeragent.PlayerAgent, nbTables*PlayersPerTable), tables: make([]tableagent.TableAgent, nbTables)}
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
	log.Printf("[Serveur] Création des joueurs")
	for i := 0; i < server.nbTables*PlayersPerTable; i++ {
		server.players[i] = *RandomPlayerAgent(i)
	}

	// Créer les TableAgents et y assigner les PlayerAgents
	log.Printf("[Serveur] Création des tables")
	for i := 0; i < server.nbTables; i++ {
		// Génération des joueurs de chaque table
		players := server.players[i*PlayersPerTable : (i+1)*PlayersPerTable]
		server.c[i] = make(chan int)
		server.tables[i] = *tableagent.NewTableAgent(i, server.c[i], server.wg, players)
	}

	// Démarrer serveur http (pour l'affichage web)
	// Lancer les TableAgents
	log.Printf("[Serveur] Lancement des tables")
	for i := range server.tables {
		server.wg.Add(1)
		go server.tables[i].Start()
	}
	server.wg.Wait()

	log.Printf("[Serveur] Lancement des parties")
	for i := 0; i < server.nbGames; i++ {
		// Synchroniser les TableAgents tour après tour + envoyer requêtes web pour affichage
		for turn := 0; turn < 4; turn++ {
			// On envoie le signal du tour aux tables à travers le channel associé
			for j := 0; j < server.nbTables; j++ {
				server.c[j] <- turn
				server.wg.Add(1)
			}
			server.wg.Wait()
		}
		// On attend que toutes les tables ait fini leur tour
		server.wg.Wait()
	}

	log.Printf("[Serveur] Parties terminées, fermeture des tables")
	// On ferme toutes les tables
	for i := 0; i < server.nbTables; i++ {
		server.c[i] <- -1
	}
	return
}

func RandomPlayerAgent(id int) *playeragent.PlayerAgent {
	return playeragent.NewPlayerAgent(id, nil, rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100))
}

// partie web
func (*ServerAgent) decodeRequest(r *http.Request) (req agt.Request, err error) {
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

// Play  = envoyer un nombre
func (serv *ServerAgent) play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("PLAY")

	// Décode de la requête pour vérifier que correspond à la bonne action
	req, err := serv.decodeRequest(r)
	if err != nil {
		return
	}
	if req.Info != "play" {
		return
	}

	// Envoyer le nombre de table
	send := agt.Response{
		Token: serv.nbTables,
	}
	data, _ := json.Marshal(send)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Fonction de mise à jour des informations des tables et joueurs
func (serv *ServerAgent) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("UPDATE")

	// Décode de la requête pour vérifier que correspond à la bonne action
	req, err := serv.decodeRequestUpdate(r)
	if err != nil {
		return
	}
	if req.Update != "update" {
		return
	}

	fmt.Printf(req.Update)

	// Fournir l'état actuel du serveur casino
	send := agt.ResponseUpdate{
		NbTables: serv.nbTables,
		NbGames:  serv.nbGames,
		//Players:  serv.players,
		//Tables:   serv.tables,
	}
	log.Println(serv.nbTables)
	data, _ := json.Marshal(send)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (serv *ServerAgent) Sleeping() {
	time.Sleep(5 * time.Minute)
}

func (serv *ServerAgent) StartServer() {
	// Création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/play", serv.play)
	mux.HandleFunc("/update", serv.update)

	// Création du serveur
	s := &http.Server{
		Addr:           serv.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// Lancement du serveur
	log.Println("Listening on", serv.addr)
	go log.Fatal(s.ListenAndServe())
}
