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

// demande si ça joue
type Request struct {
	Info string `json:"Info"`
}

// incrémente un token
type Response struct {
	Token int `json:"Token"`
}

type RequestUpdate struct {
	Update string `json:"Update"`
}

type ResponseUpdate struct {
	NbTables int `json:"NbTables"`
	NbGames  int `json:"NbGames"`
	//Players  []playeragent.PlayerAgent `json:"Players"`
	//Tables   []tableagent.TableAgent   `json:"Tables"`
}
