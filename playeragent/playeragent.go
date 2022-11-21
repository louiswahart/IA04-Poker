package playeragent

import (
	"log"
	"math"
	"math/rand"
	"time"
)

type Cards struct {
	value int
	color string
}

type RequestMessage struct {
	instruction string
	cards       []Cards
	currentBet  int
	order       int
	nbTokens    int
}

type PlayerMessage struct {
	request  RequestMessage
	response int
}

// ------ STRUCT ------
type PlayerAgent struct {
	ID             int
	c              chan (PlayerMessage)
	bluff          int
	risk           int
	aggressiveness int
	timidity       int
	currentTokens  int
	totalTokens    int
	cards          []Cards
	currentBet     int
}

// ------ CONSTRUCTOR ------
func NewPlayerAgent(id int, c chan (PlayerMessage), bluff int, risk int, aggressiveness int, timidity int) *PlayerAgent {
	return &PlayerAgent{ID: id, c: c, bluff: bluff, risk: risk, aggressiveness: aggressiveness, timidity: timidity}
}

// ------ GETTER ------
func (player *PlayerAgent) Id() int {
	return player.ID
}

func (player *PlayerAgent) C() chan (PlayerMessage) {
	return player.c
}

func (player *PlayerAgent) Bluff() int {
	return player.bluff
}

func (player *PlayerAgent) Risk() int {
	return player.risk
}

func (player *PlayerAgent) Aggressiveness() int {
	return player.aggressiveness
}

func (player *PlayerAgent) Timidity() int {
	return player.timidity
}

func (player *PlayerAgent) CurrentTokens() int {
	return player.currentTokens
}

func (player *PlayerAgent) TotalTokens() int {
	return player.totalTokens
}

func (player *PlayerAgent) Cards() []Cards {
	return player.cards
}

// ------ SETTER ------
func (player *PlayerAgent) SetId(id int) {
	player.ID = id
}

func (player *PlayerAgent) SetC(c chan (PlayerMessage)) {
	player.c = c
}

func (player *PlayerAgent) SetBluff(b int) {
	player.bluff = b
}

func (player *PlayerAgent) SetRisk(r int) {
	player.risk = r
}

func (player *PlayerAgent) SetAggressiveness(a int) {
	player.aggressiveness = a
}

func (player *PlayerAgent) SetTimidity(t int) {
	player.timidity = t
}

func (player *PlayerAgent) SetCurrentTokens(c int) {
	player.currentTokens = c
}

func (player *PlayerAgent) SetTotalTokens(t int) {
	player.totalTokens = t
}

// COPIE OU OK
func (player *PlayerAgent) SetCards(c []Cards) {
	player.cards = c
}

// ------ UTILITAIRES ------

// Aléatoire avec seed qui change constamment avec le temps
func getRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Choix fait lorsque l'on joue une main
func (player *PlayerAgent) play(currentBet int) int {
	log.Printf("[Joueur %v] Je joue\n", player.ID)
	var mise int
	if player.currentTokens < currentBet {
		mise = player.currentTokens
		log.Printf("[Joueur %v] Je fais tapis, pas assez de jeton pour suivre %v\n", player.ID, mise)
		return mise
	} else {
		mise = currentBet - player.currentBet
	}

	// Usage de la timidité, si trop timide alors joue la mise minimum
	r := getRandom(100)
	log.Printf("[Joueur %v] Est ce que je vais augmenter la mise, nombre aléatoire de %v\n", player.ID, r)
	if player.timidity > r {
		log.Printf("[Joueur %v] J'augmente la mise\n", player.ID)

		// Usage de l'agressivité pour savoir de combien on augmente la mise
		// Utilisation de l'aléatoire pour trouver un nombre pour savoir si tapis ou non
		// Utilisation de l'aléatoire pour attribuer une modification aléatoire sur l'agressivité
		r = getRandom(100)
		modif := getRandom(20)
		signe := getRandom(2)
		if signe == 1 {
			modif = -modif
		}

		log.Printf("[Joueur %v] Vais je faire tapis ? nombre aléatoire de %v | modif de %v\n", player.ID, r, modif)
		// Verification si le joueur fait un tapis ou pas
		if player.aggressiveness+modif > r {
			log.Printf("[Joueur %v] Je fais tapis, aggresivité de %v\n", player.ID, player.aggressiveness+modif)
			mise = player.currentTokens
			// S'il ne fait pas tapis alors ajout de jetons en fonction de son agresivité (+ modif aléatoire)
		} else {
			// Création du coefficient de jeton restant ajouté à la mise
			// Vérification si le coefficient n'est pas nul ou négatif
			var coeff float64
			if (float64(player.aggressiveness+modif) / 100.0) <= 0 {
				coeff = 0.1
			} else {
				coeff = (float64(player.aggressiveness+modif) / 100.0)
			}
			ajout := int(math.Ceil(float64(player.currentTokens) * coeff))
			// Si l'ajout est bien supérieur ou égale à la mise actuelle alors mise = ajout
			if ajout+player.currentBet >= currentBet {
				mise = ajout
			}
		}
	}

	return mise
}

// ------ START ------

func (player *PlayerAgent) Start() {

	// Attente de la reception d'un message de la table
	for m := range player.c {

		// Réalisation du traitement selon l'instruction
		instruction := m.request.instruction
		log.Printf("[Joueur %v] Instruction reçu : %v\n", player.ID, instruction)
		switch instruction {

		// Cas de la distribution
		// Récupération des deux cartes
		// Ajout des jetons de la partie précédente au total des jetons
		// Récupération du nombre de jeton pour la nouvelle partie
		case "distrib":
			player.cards = m.request.cards
			player.totalTokens += player.currentTokens
			player.currentTokens = m.request.nbTokens
			player.currentBet = 0

		// Cas du tour de jeu
		case "joue":
			// De base on se couche
			mise := -1
			isPlayed := false
			// Si la mise en cours est la même que la notre, de base on check
			if player.currentBet == m.request.currentBet {
				mise = 0
			}

			//Si plus de jeton
			if player.currentTokens == 0 {
				m.response = -1
				player.c <- m
				log.Printf("[Joueur %v] Plus de jeton pour jouer\n", player.ID)
				break
			}

			// Si dernier à jouer et que personne à misé (check)
			// Récupération d'un nombre aléatoire, si le bluff est supérieur alors bluff en jouant
			if m.request.order == 5 && player.currentBet == m.request.currentBet {
				r := getRandom(100)
				log.Printf("[Joueur %v] Dernier à jouer, tout le monde a check, nombre aléatoire de %v\n", player.ID, r)
				if player.bluff > r {
					log.Printf("[Joueur %v] Je bluff\n", player.ID)
					mise = player.play(m.request.currentBet)
					isPlayed = true
				}
			}

			// Si on a pas encore trouver la mise
			if !isPlayed {
				// Récupération du score de la main et du score maximal possible (pour le nombre de cartes pour le tour)
				score := CheckCombinations(player.cards, m.request.cards)
				max := MaxRange(len(player.cards) + len(m.request.cards))
				log.Printf("[Joueur %v] Mon score actuel : %v | Le max que je peux avoir : %v\n", player.ID, score, max)

				// Vérification de l'attribut de risk pour savoir si on joue ou pas
				min := (1.0 - (float64(player.risk) / 100.0)) * float64(max)
				log.Printf("[Joueur %v] Mon score doit être d'au moins : %v pour que je joue\n", player.ID, min)
				if score >= min {
					mise = player.play(m.request.currentBet)
					// Si je ne joue pas, vérification si je bluff, critère de bluff divisé par 4
				} else {
					r := getRandom(100)
					log.Printf("[Joueur %v] Normalement je ne joue pas, est ce que je bluff ? nombre aléatoire de %v\n", player.ID, r)
					if player.bluff/4 > r {
						log.Printf("[Joueur %v] Je bluff\n", player.ID)
						mise = player.play(m.request.currentBet)
					}
				}
			}

			// Envoi de la mise à la table (-1 = couche, 0 = check, >0 = mise)
			m.response = mise
			player.currentTokens -= mise
			player.currentBet += mise
			player.c <- m
			log.Printf("[Joueur %v] Mise ajoutée : %v | Mise totale : %v\n", player.ID, mise, player.currentBet)

		// Cas d'une mise obligatoire
		case "mise":
			mise := m.request.currentBet
			if mise > player.currentTokens {
				mise = player.currentTokens
			}
			// Envoi de la mise à la table
			m.response = mise
			player.currentTokens -= mise
			player.currentBet += mise
			player.c <- m
			log.Printf("[Joueur %v] Mise ajoutée : %v | Mise totale : %v\n", player.ID, mise, player.currentBet)

		// Cas d'un gain
		case "gain":
			log.Printf("[Joueur %v] Gain reçu : %v\n", player.ID, m.request.nbTokens)
			player.currentTokens += m.request.nbTokens
			player.currentBet = 0
		}
	}
	// Arret de l'agent
	log.Printf("[Joueur %v] Arret\n", player.ID)
}
