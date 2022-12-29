package playeragent

import (
	"log"
	"math"
	"math/rand"

	"gitlab.utc.fr/nivoixpa/ia04-poker/agt"
	"gitlab.utc.fr/nivoixpa/ia04-poker/rules"
)

// ------ STRUCT ------
type PlayerAgent struct {
	id             int                      // Id du joueur
	c              chan (agt.PlayerMessage) // Chanel de discussion avec la table de jeu
	bluff          int                      // Caractéristique décrivant la tendance d'un joueur à jouer (continuer de miser) alors qu'il ne devrait peut être pas
	risk           int                      // Caractéristique décrivant la tendance d'un joueur à jouer (continuer de miser) selon la puissance de sa main (plus risk est elevé, plus il jouera même avec une main faible)
	aggressiveness int                      // Caractéristique décrivant à quel point le joueur fait monter la mise quand il joue
	timidity       int                      // Caractéristique décrivant la tendance d'un joueur à juste suivre la mise actuelle ou à augmenter la mise.
	currentTokens  int                      // Nombre de jetons actuel
	totalTokens    int                      // Nombre de jetons gagnés à la fin des tables
	cards          []agt.Card               // Cartes du joueur
	previousNbCard int                      // Nombre de carte sur la table la dernière fois qu'on lui a demandé de jouer
	currentBet     int                      // Mise actuel du joueur
	nbPlay         int                      // Nombre de fois que le joueur à jouer sur le même tour
	previousBet    int                      // Précédent Bet
	isBlind        bool                     // Indication si le joueur joue une blind
	isAllIn        bool                     // Indication de si le joueur a fait tapis dans la partie en cours
	action         string                   //action du joueur
	lastEarning    int                      // Dernier gain du joueur
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

func (player *PlayerAgent) PreviousNbCard() int {
	return player.previousNbCard
}

func (player *PlayerAgent) CurrentBet() int {
	return player.currentBet
}

func (player *PlayerAgent) NbPlay() int {
	return player.nbPlay
}

func (player *PlayerAgent) PreviousBet() int {
	return player.previousBet
}

func (player *PlayerAgent) IsBlind() bool {
	return player.isBlind
}

func (player *PlayerAgent) IsAllIn() bool {
	return player.isAllIn
}

func (player *PlayerAgent) Action() string {
	return player.action
}

func (player *PlayerAgent) LastEarning() int {
	return player.lastEarning
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

func (player *PlayerAgent) SetCards(c []agt.Card) {
	player.cards = c
}

// ------ UTILITAIRES ------

// Aléatoire
func getRandom(n int) int {
	return rand.Intn(n)
}

// Permet de choisir l'augmentation de mise
func (player *PlayerAgent) augmentationMise(currentBet int) (mise int) {
	log.Printf("[Joueur %v] J'augmente la mise\n", player.id)
	player.action = "Je joue"
	// Usage de l'agressivité pour savoir de combien on augmente la mise
	// Utilisation de l'aléatoire pour trouver un nombre pour savoir si tapis ou non
	// Utilisation de l'aléatoire pour attribuer une modification aléatoire sur l'agressivité
	r := getRandom(100)
	modif := getRandom(20)
	signe := getRandom(2)
	if signe == 1 {
		modif = -modif
	}

	log.Printf("[Joueur %v] Vais je faire tapis ? nombre aléatoire de %v | mon aggressivité avec modif : %v\n", player.id, r, float64(player.aggressiveness+modif)/10.0)
	// Verification si le joueur fait un tapis ou pas
	if float64(player.aggressiveness+modif)/10.0 > float64(r) {
		log.Printf("[Joueur %v] Je fais tapis\n", player.id)
		player.action = "Je fais tapis"
		mise = player.currentTokens
		player.isAllIn = true
		// S'il ne fait pas tapis alors ajout de jetons en fonction de son agresivité (+ modif aléatoire)
	} else {
		log.Printf("[Joueur %v] Je ne fais pas tapis\n", player.id)

		// Choix du bet de référence
		// Current bet si différent de 0
		// Previous bet si current bet = 0
		indiceBet := 0
		isPreviousBet := false
		if currentBet == 0 {
			indiceBet = player.previousBet
			isPreviousBet = true
		} else {
			indiceBet = currentBet
		}

		// Création du coefficient à rajouter par rapport à la mise et au nombre de tokens restants
		// Vérification si le coefficient n'est pas nul ou négatif
		var coeff float64
		if (float64(player.aggressiveness+modif) / 100.0) <= 0 {
			coeff = 0.1
		} else {
			coeff = (float64(player.aggressiveness+modif) / 100.0)
		}
		if (float64(player.currentTokens) / float64(indiceBet)) >= 5 {
			switch {
			case coeff <= 0.33:
				coeff += 1.0
			case coeff <= 0.66 && coeff > 0.33:
				coeff += 2.0
			default:
				coeff += 3.0
			}
		} else {
			coeff += 1.0
		}
		ajout := int(math.Ceil(float64(indiceBet)*(coeff))) - player.currentBet
		if isPreviousBet {
			ajout -= player.previousBet
		}
		// Si le joueur a assez de token alors fait cette mise
		// Sinon tapis
		if ajout+player.currentBet >= currentBet && ajout < player.currentTokens {
			mise = ajout
		} else {
			mise = player.currentTokens
			player.isAllIn = true
		}
	}
	return
}

// Choix fait lorsque l'on joue une main
func (player *PlayerAgent) play(currentBet int) int {
	log.Printf("[Joueur %v] Je joue\n", player.id)
	var mise int
	if player.currentTokens <= currentBet {
		mise = player.currentTokens
		player.isAllIn = true
		log.Printf("[Joueur %v] Je fais tapis, pas assez de jeton pour suivre (ou pile le nombre)\n", player.id)
		player.action = "Je fais tapis"
		return mise
	} else {
		mise = currentBet - player.currentBet
	}

	// Usage de la timidité, si trop timide alors joue la mise minimum
	r := getRandom(100)
	log.Printf("[Joueur %v] Est ce que je vais augmenter la mise, nombre aléatoire de %v | ma timidité modifiée : %v\n", player.id, r, float64(player.timidity)*math.Pow(1.5, float64(player.nbPlay)))
	if float64(player.timidity)*math.Pow(1.5, float64(player.nbPlay)) < float64(r) {
		mise = player.augmentationMise(currentBet)
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

		// Cas d'une nouvelle partie
		// Récupération du nombre de jeton pour la nouvelle partie
		case "nouvelle":
			player.currentTokens = m.Request.NbTokens
			player.currentBet = 0
			player.previousNbCard = 0
			player.previousBet = 0
			player.nbPlay = 0
			player.isBlind = false
			player.isAllIn = false
			player.action = ""
			player.lastEarning = 0
			m.Response = m.Request.NbTokens
			player.c <- m

		// Cas de la distribution
		// Récupération des deux cartes
		case "distrib":
			player.cards = m.Request.Cards
			player.lastEarning = 0
			if player.currentTokens == 0 {
				player.action = "Je n'ai plus de jeton pour jouer"
			}
			if !player.isBlind {
				player.currentBet = 0
				player.previousNbCard = 0
				player.previousBet = 0
				player.nbPlay = 0
				player.isAllIn = false
			}
			player.isBlind = false
			m.Response = m.Request.NbTokens
			player.c <- m

		// Cas du tour de jeu
		case "joue":
			// On verifie si on a une carte de plus sur la table
			// Si le cas on remet notre bet à 0 et notre nombre de fois jouer dans le round à 0
			if len(m.Request.Cards) != player.previousNbCard {
				player.currentBet = 0
				player.nbPlay = 0
			}
			player.previousNbCard = len(m.Request.Cards)

			// De base on se couche
			mise := -1
			player.action = "Je me couche"
			isPlayed := false
			// Si la mise en cours est la même que la notre, de base on check
			if player.currentBet == m.Request.CurrentBet {
				mise = 0
				if player.nbPlay == 0 || player.currentBet == 0 {
					player.action = "Je check"
				}
			}

			//Si plus de jeton
			if player.currentTokens == 0 {
				// Si tapis en cours
				if player.isAllIn {
					m.Response = 0
					player.c <- m
					log.Printf("[Joueur %v] Tapis en cours\n", player.id)
					player.action = "Tapis en cours"
				} else {
					m.Response = -1
					player.c <- m
					log.Printf("[Joueur %v] Plus de jeton pour jouer\n", player.id)
					player.action = "Je n'ai plus de jeton pour jouer"
				}
				continue
			}

			// Si dernier à jouer et que personne à misé (check)
			// Récupération d'un nombre aléatoire, si le bluff est supérieur alors bluff en jouant
			if m.Request.Order == 5 && player.currentBet == m.Request.CurrentBet {
				r := getRandom(100)
				log.Printf("[Joueur %v] Dernier à jouer, tout le monde a check, nombre aléatoire de %v | mon bluff : %v\n", player.id, r, player.bluff)
				if player.bluff > r {
					log.Printf("[Joueur %v] Je bluff\n", player.id)
					player.action = "Je joue"
					mise = player.play(m.Request.CurrentBet)
					isPlayed = true
				} else {
					log.Printf("[Joueur %v] Je ne bluff pas\n", player.id)
				}
			}

			// Si pas jouer et que je peux check
			// Selon ma timidité, est ce que j'augmente la mise
			if player.currentBet == m.Request.CurrentBet && !isPlayed {
				r := getRandom(100)
				log.Printf("[Joueur %v] Est ce que je vais augmenter la mise, nombre aléatoire de %v | ma timidité modifiée : %v\n", player.id, r, float64(player.timidity)*math.Pow(1.5, float64(player.nbPlay)))
				if float64(player.timidity)*math.Pow(1.5, float64(player.nbPlay)) < float64(r) {
					mise = player.augmentationMise(m.Request.CurrentBet)
					player.action = "Je joue"
				} else {
					log.Printf("[Joueur %v] Je check\n", player.id)
					if player.nbPlay == 0 || player.currentBet == 0 {
						player.action = "Je check"
					}
				}
				isPlayed = true
			}

			// Si on a pas encore trouver la mise
			if !isPlayed {
				// Récupération du score de la main et du score maximal possible (pour le nombre de cartes pour le tour)
				score := rules.CheckCombinations(player.cards, m.Request.Cards)
				max := rules.MaxRange(len(player.cards) + len(m.Request.Cards))
				log.Printf("[Joueur %v] Mon score actuel : %v | Le max que je peux avoir : %v\n", player.id, score, max)

				// Vérification de l'attribut de risk pour savoir si on joue ou pas
				min := (1.0 - ((float64(player.risk) / math.Pow(1.125, float64(player.nbPlay))) / 100.0)) * float64(max)
				log.Printf("[Joueur %v] Mon score doit être d'au moins : %v pour que je joue\n", player.id, min)
				if float64(score) >= min {
					mise = player.play(m.Request.CurrentBet)
					player.action = "Je joue"
					// Si je ne joue pas, vérification si je bluff, critère de bluff divisé par 4
				} else {
					r := getRandom(100)
					log.Printf("[Joueur %v] Normalement je ne joue pas, est ce que je bluff ? nombre aléatoire de %v | mon bluff ajusté : %v\n", player.id, r, player.bluff/4)
					if player.bluff/4 > r {
						log.Printf("[Joueur %v] Je bluff\n", player.id)
						player.action = "Je joue"
						mise = player.play(m.Request.CurrentBet)
					} else {
						log.Printf("[Joueur %v] Je ne bluff pas\n", player.id)
					}
				}
			}

			// Envoi de la mise à la table (-1 = couche, 0 = check, >0 = mise)
			m.Response = mise
			if mise == -1 {
				log.Printf("[Joueur %v] Je me couche\n", player.id)
				player.action = "Je me couche"
			} else if mise > 0 {
				player.currentTokens -= mise
				player.currentBet += mise
				player.previousBet = player.currentBet
				log.Printf("[Joueur %v] Mise ajoutée : %v | Mise totale : %v\n", player.id, mise, player.currentBet)
			}
			player.c <- m
			player.nbPlay += 1

		// Cas d'une mise obligatoire = small ou big blind
		case "mise":
			player.isBlind = true
			player.isAllIn = false
			player.currentBet = 0
			player.previousNbCard = 0
			player.previousBet = 0
			player.nbPlay = 0
			player.action = "Je joue"
			mise := m.Request.CurrentBet
			if mise > player.currentTokens {
				log.Printf("[Joueur %v] Je dois faire tapis\n", player.id)
				player.action = "Je fais tapis"
				mise = player.currentTokens
				player.isAllIn = true
			}
			// Envoi de la mise à la table
			m.Response = mise
			player.currentTokens -= mise
			player.currentBet += mise
			player.previousBet = player.currentBet
			player.c <- m
			log.Printf("[Joueur %v] Mise ajoutée : %v | Mise totale : %v\n", player.id, mise, player.currentBet)

		// Cas d'un gain
		case "gain":
			log.Printf("[Joueur %v] Gain reçu : %v\n", player.id, m.Request.NbTokens)
			player.currentTokens += m.Request.NbTokens
			player.lastEarning = m.Request.NbTokens
			m.Response = m.Request.NbTokens
			player.c <- m

		// Fin de la partie, ajout des jetons de la partie précédente au total des jetons
		case "fin":
			player.totalTokens += player.currentTokens
			player.currentTokens = 0
			player.currentBet = 0
			player.previousNbCard = 0
			player.previousBet = 0
			player.nbPlay = 0
			player.isBlind = false
			player.isAllIn = false
			player.action = ""
			player.lastEarning = 0
		}
	}
	// Arret de l'agent
	player.action = "Terminé"
	player.currentBet = 0
	player.cards = nil
	log.Printf("[Joueur %v] Arret\n", player.id)

}
