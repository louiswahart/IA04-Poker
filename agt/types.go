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

type RequestReset struct {
	Req string `json:"Req"`
}

type RequestgetTable struct {
	Req   string `json:"Req"`
	Table int    `json:"Table"`
}

type ResponsegetTable struct {
	PlayersID       []int    `json:"PlayersID"`
	PlayersBlind    []int    `json:"PlayersBlind"`
	PlayersToken    []int    `json:"PlayersToken"`
	PlayersBet      []int    `json:"PlayersBet"`
	PlayersTotalBet []int    `json:"PlayersTotalBet"`
	PlayersActions  []string `json:"PlayersActions"`
	PlayersCards    [][]Card `json:"PlayersCards"`
	PlayersGain     []int    `json:"PlayersGain"`
	PlayersWinner   []bool   `json:"PlayersWinner"`
	Pot             int      `json:"Pot"`
	TableCards      []Card   `json:"TableCards"`
}

type RequestgetPlayer struct {
	Req    string `json:"Req"`
	Player int    `json:"Player"`
}

type ResponsegetPlayer struct {
	Timidity       int `json:"Timidity"`
	Aggressiveness int `json:"Aggressiveness"`
	Risk           int `json:"Risk"`
	Bluff          int `json:"Bluff"`
	Table          int `json:"Table"`
}

type RequestchangeStats struct {
	Req            string `json:"Req"`
	Player         int    `json:"Player"`
	Timidity       int    `json:"Timidity"`
	Aggressiveness int    `json:"Aggressiveness"`
	Risk           int    `json:"Risk"`
	Bluff          int    `json:"Bluff"`
}

type ResponsechangeStats struct {
	Timidity       int `json:"Timidity"`
	Aggressiveness int `json:"Aggressiveness"`
	Risk           int `json:"Risk"`
	Bluff          int `json:"Bluff"`
}

type RequestUpdate struct {
	Req   string `json:"Req"`
	Table int    `json:"Table"`
}

type ResponseUpdate struct {
	PlayersID       []int    `json:"PlayersID"`
	PlayersBlind    []int    `json:"PlayersBlind"`
	PlayersToken    []int    `json:"PlayersToken"`
	PlayersBet      []int    `json:"PlayersBet"`
	PlayersTotalBet []int    `json:"PlayersTotalBet"`
	PlayersActions  []string `json:"PlayersActions"`
	PlayersGain     []int    `json:"PlayersGain"`
	PlayersCards    [][]Card `json:"PlayersCards"`
	PlayersWinner   []bool   `json:"PlayersWinner"`
	Pot             int      `json:"Pot"`
	TableCards      []Card   `json:"TableCards"`
}
