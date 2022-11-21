package playeragent

import (
	"log"
	"math"
	"math/rand"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
)

// ------ STRUCT ------
type PlayerAgent struct {
	id             int
	c              chan (agt.PlayerMessage)
	bluff          int // Caractéristique décrivant la tendance d'un joueur à jouer (continuer de miser) alors qu'il ne devrait peut être pas
	risk           int // Caractéristique décrivant la tendance d'un joueur à jouer (continuer de miser) selon la puissance de sa main (plus risk est elevé, plus il jouera même avec une main faible)
	aggressiveness int // Caractéristique décrivant à quel point le joueur fait monter la mise quand il joue
	timidity       int // Caractéristique décrivant la tendance d'un joueur à juste suivre la mise actuelle ou à augmenter la mise.
	currentTokens  int
	totalTokens    int
	cards          []agt.Card
	currentBet     int
}

// ------ CONSTRUCTOR ------
func NewPlayerAgent(id int, c chan (agt.PlayerMessage), bluff int, risk int, aggressiveness int, timidity int) *PlayerAgent {
	return &PlayerAgent{id: id, c: c, bluff: bluff, risk: risk, aggressiveness: aggressiveness, timidity: timidity}
}

// ------ GETTER ------
func (player *PlayerAgent) Id() int {
	return player.id
}

func (player *PlayerAgent) C() chan (agt.PlayerMessage) {
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

func (player *PlayerAgent) Card() []agt.Card {
	return player.cards
}

// ------ SETTER ------
func (player *PlayerAgent) SetId(id int) {
	player.id = id
}

func (player *PlayerAgent) SetC(c chan (agt.PlayerMessage)) {
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
func (player *PlayerAgent) SetCards(c []agt.Card) {
	player.cards = c
}

// ------ UTILITAIRES ------

// Aléatoire
func getRandom(n int) int {
	return rand.Intn(n)
}

// Choix fait lorsque l'on joue une main
func (player *PlayerAgent) play(currentBet int) int {
	log.Printf("[Joueur %v] Je joue\n", player.id)
	var mise int
	if player.currentTokens < currentBet {
		mise = player.currentTokens
		log.Printf("[Joueur %v] Je fais tapis, pas assez de jeton pour suivre %v\n", player.id, mise)
		return mise
	} else {
		mise = currentBet - player.currentBet
	}

	// Usage de la timidité, si trop timide alors joue la mise minimum
	r := getRandom(100)
	log.Printf("[Joueur %v] Est ce que je vais augmenter la mise, nombre aléatoire de %v\n", player.id, r)
	if player.timidity > r {
		log.Printf("[Joueur %v] J'augmente la mise\n", player.id)

		// Usage de l'agressivité pour savoir de combien on augmente la mise
		// Utilisation de l'aléatoire pour trouver un nombre pour savoir si tapis ou non
		// Utilisation de l'aléatoire pour attribuer une modification aléatoire sur l'agressivité
		r = getRandom(100)
		modif := getRandom(20)
		signe := getRandom(2)
		if signe == 1 {
			modif = -modif
		}

		log.Printf("[Joueur %v] Vais je faire tapis ? nombre aléatoire de %v | modif de %v\n", player.id, r, modif)
		// Verification si le joueur fait un tapis ou pas
		if player.aggressiveness+modif > r {
			log.Printf("[Joueur %v] Je fais tapis, aggresivité de %v\n", player.id, player.aggressiveness+modif)
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
		instruction := m.Request.Instruction
		log.Printf("[Joueur %v] Instruction reçu : %v\n", player.id, instruction)
		switch instruction {

		// Cas de la distribution
		// Récupération des deux cartes
		// Récupération du nombre de jeton pour la nouvelle partie
		case "distrib":
			player.cards = m.Request.Cards
			player.currentTokens = m.Request.NbTokens
			player.currentBet = 0

		// Cas du tour de jeu
		case "joue":
			// De base on se couche
			mise := -1
			isPlayed := false
			// Si la mise en cours est la même que la notre, de base on check
			if player.currentBet == m.Request.CurrentBet {
				mise = 0
			}

			//Si plus de jeton
			if player.currentTokens == 0 {
				m.Response = -1
				player.c <- m
				log.Printf("[Joueur %v] Plus de jeton pour jouer\n", player.id)
				break
			}

			// Si dernier à jouer et que personne à misé (check)
			// Récupération d'un nombre aléatoire, si le bluff est supérieur alors bluff en jouant
			if m.Request.Order == 5 && player.currentBet == m.Request.CurrentBet {
				r := getRandom(100)
				log.Printf("[Joueur %v] Dernier à jouer, tout le monde a check, nombre aléatoire de %v\n", player.id, r)
				if player.bluff > r {
					log.Printf("[Joueur %v] Je bluff\n", player.id)
					mise = player.play(m.Request.CurrentBet)
					isPlayed = true
				}
			}

			// Si on a pas encore trouver la mise
			if !isPlayed {
				// Récupération du score de la main et du score maximal possible (pour le nombre de cartes pour le tour)
				score := CheckCombinations(player.cards, m.Request.Cards)
				max := MaxRange(len(player.cards) + len(m.Request.Cards))
				log.Printf("[Joueur %v] Mon score actuel : %v | Le max que je peux avoir : %v\n", player.id, score, max)

				// Vérification de l'attribut de risk pour savoir si on joue ou pas
				min := (1.0 - (float64(player.risk) / 100.0)) * float64(max)
				log.Printf("[Joueur %v] Mon score doit être d'au moins : %v pour que je joue\n", player.id, min)
				if score >= min {
					mise = player.play(m.Request.CurrentBet)
					// Si je ne joue pas, vérification si je bluff, critère de bluff divisé par 4
				} else {
					r := getRandom(100)
					log.Printf("[Joueur %v] Normalement je ne joue pas, est ce que je bluff ? nombre aléatoire de %v\n", player.id, r)
					if player.bluff/4 > r {
						log.Printf("[Joueur %v] Je bluff\n", player.id)
						mise = player.play(m.Request.CurrentBet)
					}
				}
			}

			// Envoi de la mise à la table (-1 = couche, 0 = check, >0 = mise)
			m.Response = mise
			player.currentTokens -= mise
			player.currentBet += mise
			player.c <- m
			log.Printf("[Joueur %v] Mise ajoutée : %v | Mise totale : %v\n", player.id, mise, player.currentBet)

		// Cas d'une mise obligatoire
		case "mise":
			mise := m.Request.CurrentBet
			if mise > player.currentTokens {
				mise = player.currentTokens
			}
			// Envoi de la mise à la table
			m.Response = mise
			player.currentTokens -= mise
			player.currentBet += mise
			player.c <- m
			log.Printf("[Joueur %v] Mise ajoutée : %v | Mise totale : %v\n", player.id, mise, player.currentBet)

		// Cas d'un gain
		case "gain":
			log.Printf("[Joueur %v] Gain reçu : %v\n", player.id, m.Request.NbTokens)
			player.currentTokens += m.Request.NbTokens
			player.currentBet = 0

		// Fin de la partie, ajout des jetons de la partie précédente au total des jetons
		case "fin":
			player.totalTokens += player.currentTokens
			player.currentTokens = 0
			player.currentBet = 0
		}
	}
	// Arret de l'agent
	log.Printf("[Joueur %v] Arret\n", player.id)
}
