import React from 'react';
import Player from './Player'
import Info from './Info'
import Interactions from './Interactions'
import State from './State'

export default class Game extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            nbTable: 50,
            nbGame: 50,
            currentTurn: -1,
            currentGame: 0,
            token: 0,
            idTable: 0,
            menu: true,
            play: false,
            pot:0,
            playersId : [],
            playersToken: [],
            playersBet: [],
            playersActions: [],
            playersWinner: [],
            playersCards: [[]],
            playerTimidity :0,
            playerAggressiveness:0,
            playerRisk:0,
            playerBluff:0
        };

        this.changedTable = this.changedTable.bind(this);
        this.Play = this.Play.bind(this);
        this.Pause = this.Pause.bind(this);
        this.Reset = this.Reset.bind(this);
        this.changeNbGame = this.changeNbGame.bind(this);
        this.changeNbTable = this.changeNbTable.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    componentDidMount(){
        setInterval(() => {
            if(this.state.play){

                this.setState({currentTurn : this.state.currentTurn + 1})
                if(this.state.currentTurn>=4){
                    this.setState({currentGame : this.state.currentGame + 1})
                    this.setState({currentTurn : 0})
                }

                if(this.state.currentGame < this.state.nbGame){
                    const options = {
                        method: 'POST',
                        headers: { 'Content-Type': 'application/json' },
                        body: JSON.stringify({ Req: 'update', Table: Number(this.state.idTable)})
                    };
                    fetch('http://localhost:8080/update', options)
                        .then(resp => resp.json())
                        .then(data => this.setState({ playersId: data.PlayersID,  playersToken: data.PlayersToken, playersBet: data.PlayersBet,playersActions: data.PlayersActions,playersWinner: data.PlayersWinner,pot: data.Pot}));
                } else{
                    this.setState({play : false})
                }
             }}, 5000)
    }

    changedTable(i){
        console.log("Table : " + i);
        const options = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ Req: 'getTable', Table: Number(i)})
        };
        fetch('http://localhost:8080/getTable', options)
            .then(resp => resp.json())
            .then(data => this.setState({ playersId: data.PlayersID,  playersToken: data.PlayersToken, playersBet: data.PlayersBet,playersActions: data.PlayersActions,pot: data.Pot}));
        this.setState({idTable : i})
    }

    changedPlayer(i){
        console.log("Joueur : " + i);
        const options = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ Req: 'getPlayer', Player: Number(i)})
        };
        console.log("avant " + this.state.idTable)
        fetch('http://localhost:8080/getPlayer', options)
            .then(resp => resp.json())
            .then(data => this.setState({ playerTimidity: data.Timidity,  playerAggressiveness: data.Aggressiveness, playerRisk: data.Risk,playerBluff: data.Bluff, idTable: data.Table},this.changedTable(data.Table)));
    }

    Play() {
        this.setState({play : true});
    }

    Pause(){
        this.setState({play : false});
    }

    Reset(){
        this.setState({token : 0});
    }

    changeNbGame(event) {    
        this.setState({nbGame: event.target.value});  
    }

    changeNbTable(event) {
        this.setState({nbTable: event.target.value});  
    }

    handleSubmit(event) {
        console.log(this.state.nbGame)
        console.log(this.state.nbTable)

        const options = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ Req: 'play', NbTables : Number(this.state.nbTable),  NbGames : Number(this.state.nbGame)})
        };
        fetch('http://localhost:8080/play', options)
            .then(resp => console.log(resp.statusText));

        this.setState({menu : false})
        this.setState({play : true})
    }

    render() {
        return (
        <div className='all'>
            {/* Affichage du menu */}
            {this.state.menu ?
                <div>
                    <form onSubmit={this.handleSubmit}>        
                        <label>
                            Nombre de tables (5 joueurs par table)
                            <input type="number" min="1" max="1000" value={this.state.nbTable} onChange={this.changeNbTable} required/>        
                        </label>
                        <label>
                            Nombre de parties
                            <input type="number" min="1" max="1000" value={this.state.nbGame} onChange={this.changeNbGame} required/>        
                        </label>
                        <input type="submit" value="Lancer" />
                    </form>
                </div>
                : null
            }


            {/* Affichage du jeu */}
            {!this.state.menu ?
                <div>
                    <div className='all'>
                        <div className="jeu">
                            <div className="plateau">
                                <div className='player1'>
                                    <Player
                                        ID={this.state.playersId[0]} 
                                        nbToken={this.state.playersToken[0]} 
                                        bet={this.state.playersBet[0]}
                                        action={this.state.playersActions[0]}
                                        winner={this.state.playersWinner[0]} 
                                        listeCards={20}
                                    />
                                </div>
                                <Player 
                                    ID={this.state.playersId[1]} 
                                    nbToken={this.state.playersToken[1]} 
                                    bet={this.state.playersBet[1]} 
                                    action={this.state.playersActions[1]} 
                                    winner={this.state.playersWinner[1]}
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[2]} 
                                    nbToken={this.state.playersToken[2]} 
                                    bet={this.state.playersBet[2]} 
                                    action={this.state.playersActions[2]} 
                                    winner={this.state.playersWinner[2]}
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[3]} 
                                    nbToken={this.state.playersToken[3]} 
                                    bet={this.state.playersBet[3]} 
                                    action={this.state.playersActions[3]} 
                                    winner={this.state.playersWinner[3]}
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[4]} 
                                    nbToken={this.state.playersToken[4]} 
                                    bet={this.state.playersBet[4]}  
                                    action={this.state.playersActions[4]} 
                                    winner={this.state.playersWinner[4]}
                                    listeCards={20}
                                />
                            </div> 
                                <div className='rectangle'>
                                    <div className='etat'>
                                        <State
                                            Turn={this.state.currentTurn}
                                            Game={this.state.currentGame}
                                            Pot={this.state.pot}
                                        />
                                    </div>
                                </div>
                            <div className="interactions">
                                <Interactions 
                                    onPlay={() => this.Play()}
                                    onPause={() => this.Pause()}
                                    onReset={() => this.Reset()}
                                />
                            </div>
                        </div>  
                        <div className="infos">
                            <Info
                                nbTable= {this.state.nbTable} 
                                nbPlayers= {5*this.state.nbTable}
                                onTableChanged={i => this.changedTable(i)}
                                onPlayerChanged={i => this.changedPlayer(i)}
                                Timidity={this.state.playerTimidity} 
                                Aggressiveness={this.state.playerAggressiveness} 
                                Risk={this.state.playerRisk}  
                                Bluff={this.state.playerBluff} 
                            />
                        </div>
                    </div>
                </div>
                : null
            }
        </div>
        );
    }
}