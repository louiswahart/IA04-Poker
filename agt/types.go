package agt

// Structure de carte
type Card struct {
	Value int
	Color string
}

// Structure des requêtes dans les messages
type RequestMessage struct {
	Instruction string
	Cards       []Card
	CurrentBet  int
	Order       int
	NbTokens    int
}

// Structure des messages entre la table et le joueur
type PlayerMessage struct {
	Request  RequestMessage
	Response int
}
