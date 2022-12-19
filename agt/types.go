package agt

// Structure de carte
type CardSuit int

const (
	Club CardSuit = iota
	Diamond
	Heart
	Spades
)

type Card struct {
	Value int
	Color CardSuit
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

type RequestPlay struct {
	Req      string `json:"Req"`
	NbTables int    `json:"NbTables"`
	NbGames  int    `json:"NbGames"`
}

type RequestUpdate struct {
	Req   string `json:"Req"`
	Table int    `json:"Table"`
}

type ResponseUpdate struct {
	PlayersID    []int `json:"PlayersID"`
	PlayersToken []int `json:"PlayersToken"`
	PlayersBet   []int `json:"PlayersBet"`
}
