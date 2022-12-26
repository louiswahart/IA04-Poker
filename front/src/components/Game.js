import React from 'react';
import Player from './Player'
import Info from './Info'
import Interactions from './Interactions'
import State from './State'
import { getCard } from './Cards';

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
            idPlayer: -1,
            tableCards: [],
            menu: true,
            play: false,
            pot:0,
            playersId : [],
            playersToken: [],
            playersBet: [],
            playersTotalBet: [],
            playersActions: [],
            playersWinner: [],
            playersCards: [[], [], [], [], []],
            playerTimidity :0,
            playerAggressiveness:0,
            playerRisk:0,
            playerBluff:0
        };

        this.verifyRequestNeeded = this.verifyRequestNeeded.bind(this);
        this.requestUpdate = this.requestUpdate.bind(this);
        this.changedTable = this.changedTable.bind(this);
        this.PlayPause = this.PlayPause.bind(this);
        this.Reset = this.Reset.bind(this);
        this.changeNbGame = this.changeNbGame.bind(this);
        this.changeNbTable = this.changeNbTable.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    componentDidMount(){
        setInterval(() => {
            if(this.state.play){
                this.setState({currentTurn : (this.state.currentTurn + 1)}, function() {
                        if(this.state.currentTurn>=4){
                            this.setState({
                                currentGame : (this.state.currentGame + 1),
                                currentTurn : 0}, function() {
                                    this.verifyRequestNeeded()
                            })
                        } else{
                            this.verifyRequestNeeded()
                        }
                }
                )
             }}, 5000)
    }

    verifyRequestNeeded(){
        if(this.state.currentGame < this.state.nbGame){
            this.requestUpdate(false);
        } else{
            this.setState({play : false})
            this.requestUpdate(true);
        }
    }

    requestUpdate(fin){
        if(fin){
            const options = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ Req: 'update', Table: Number(this.state.idTable)})
            };
            fetch('http://localhost:8080/update', options)
                .then(resp => console.log(resp.statusText));
        } else{
            const options = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ Req: 'update', Table: Number(this.state.idTable)})
            };
            fetch('http://localhost:8080/update', options)
                .then(resp => resp.json())
                .then(data => this.setState({ playersId: data.PlayersID,  playersToken: data.PlayersToken, playersBet: data.PlayersBet, playersTotalBet : data.PlayersTotalBet, playersActions: data.PlayersActions, playersCards: data.PlayersCards, playersWinner: data.PlayersWinner, pot: data.Pot, tableCards: data.TableCards}));
            }
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
            .then(data => this.setState({ playersId: data.PlayersID,  playersToken: data.PlayersToken, playersBet: data.PlayersBet, playersTotalBet : data.PlayersTotalBet, playersActions: data.PlayersActions, playersCards: data.PlayersCards, playersWinner: data.PlayersWinner, pot: data.Pot, tableCards: data.TableCards}));
        this.setState({idTable : i}, this.setState({idPlayer : (Math.trunc(this.state.idPlayer/5) === i ? this.state.idPlayer : -1)}))
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
            .then(data => this.setState({ playerTimidity: data.Timidity,  playerAggressiveness: data.Aggressiveness, playerRisk: data.Risk,playerBluff: data.Bluff}, this.setState({idPlayer : i}, this.changedTable(data.Table))));
    }

    PlayPause() {
        this.setState({play : !this.state.play});
    }

    Reset(){
        this.setState({
            nbTable: 50,
            nbGame: 50,
            currentTurn: -1,
            currentGame: 0,
            token: 0,
            idTable: 0,
            idPlayer: -1,
            tableCards: [],
            menu: true,
            play: false,
            pot:0,
            playersId : [],
            playersToken: [],
            playersBet: [],
            playersTotalBet: [],
            playersActions: [],
            playersWinner: [],
            playersCards: [[], [], [], [], []],
            playerTimidity :0,
            playerAggressiveness:0,
            playerRisk:0,
            playerBluff:0});
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
        let etat = "En cours"
        if(!this.state.play){
            if(this.state.currentGame < this.state.nbGame){
                etat = "En pause"
            } else{
                etat = "Terminé (toutes les parties sont terminées)"
            }
        }

        let Cards = []
        this.state.tableCards.forEach(element => {
            let color = ""
            if (element.Color === 0 || element.Color === 3){
                color = "black"
            } else {
                color = "red"
            }
            Cards.push(
                <div className={'CardTable ' + color} key = {element.Color.toString() + element.Value.toString()}>{getCard(element.Color, element.Value)}</div> 
              )    
        });

        return (
        <div className='all'>
            {/* Affichage du menu */}
            {this.state.menu ?
                <div className='Menu'>
                    <h1>Configuration de la simulation :</h1>
                    <form onSubmit={this.handleSubmit}>        
                        <label>
                            Nombre de tables (5 joueurs par table)
                            <input type="number" min="1" max="1000" value={this.state.nbTable} onChange={this.changeNbTable} required/>        
                        </label>
                        <br></br>
                        <label>
                            Nombre de parties
                            <input type="number" min="1" max="1000" value={this.state.nbGame} onChange={this.changeNbGame} required/>        
                        </label>
                        <br></br>
                        <br></br>
                        <input type="submit" value="Lancer" />
                    </form>
                </div>
                : null
            }


            {/* Affichage du jeu */}
            {!this.state.menu ?
                <>
                    <div className="jeu">
                        <div className="plateau">
                            <div className='zone1'>
                                <div className='playerSideTop'>
                                    <Player
                                        ID={this.state.playersId[1]} 
                                        nbToken={this.state.playersToken[1]} 
                                        bet={this.state.playersBet[1]}
                                        totalBet={this.state.playersTotalBet[1]}
                                        action={this.state.playersActions[1]}
                                        winner={this.state.playersWinner[1]} 
                                        listeCards={this.state.playersCards[1]}
                                    />
                                </div>
                                <div className='playerSideBottom'>
                                    <Player 
                                        ID={this.state.playersId[0]} 
                                        nbToken={this.state.playersToken[0]} 
                                        bet={this.state.playersBet[0]}
                                        totalBet={this.state.playersTotalBet[0]} 
                                        action={this.state.playersActions[0]} 
                                        winner={this.state.playersWinner[0]}
                                        listeCards={this.state.playersCards[0]}
                                    />
                                </div>
                            </div>
                            <div className='zone2'>
                                <div className='playerTop'>
                                    <Player 
                                        ID={this.state.playersId[2]} 
                                        nbToken={this.state.playersToken[2]} 
                                        bet={this.state.playersBet[2]}
                                        totalBet={this.state.playersTotalBet[2]} 
                                        action={this.state.playersActions[2]} 
                                        winner={this.state.playersWinner[2]}
                                        listeCards={this.state.playersCards[2]}
                                    />
                                </div>
                                <div className='rectangle'>
                                    <div className='tableZone'>
                                        <div className='etat'>
                                            <State
                                                Turn={this.state.currentTurn}
                                                Game={this.state.currentGame}
                                                Pot={this.state.pot}
                                                Etat={etat}
                                            />
                                        </div>
                                    </div>
                                    <div className='tableZone'>
                                        {Cards}
                                    </div>
                                </div>
                            </div>
                            <div className='zone3'>
                                <div className='playerSideTop'>
                                    <Player 
                                        ID={this.state.playersId[3]} 
                                        nbToken={this.state.playersToken[3]} 
                                        bet={this.state.playersBet[3]}
                                        totalBet={this.state.playersTotalBet[3]} 
                                        action={this.state.playersActions[3]} 
                                        winner={this.state.playersWinner[3]}
                                        listeCards={this.state.playersCards[3]}
                                    />
                                </div>
                                <div className='playerSideBottom'>
                                    <Player 
                                        ID={this.state.playersId[4]} 
                                        nbToken={this.state.playersToken[4]} 
                                        bet={this.state.playersBet[4]}
                                        totalBet={this.state.playersTotalBet[4]}  
                                        action={this.state.playersActions[4]} 
                                        winner={this.state.playersWinner[4]}
                                        listeCards={this.state.playersCards[4]}
                                    />
                                </div>
                            </div>
                        </div> 
                        <div className="interactions">
                            <Interactions 
                                onPlayPause={() => this.PlayPause()}
                                onReset={() => this.Reset()}
                            />
                        </div>
                    </div>  
                    <div className="infos">
                        <p>Options de sélection :</p>
                        <Info
                            nbTable={this.state.nbTable} 
                            nbPlayers={5*this.state.nbTable}
                            idTable={this.state.idTable}
                            idPlayer={this.state.idPlayer} 
                            onTableChanged={i => this.changedTable(i)}
                            onPlayerChanged={i => this.changedPlayer(i)}
                            Timidity={this.state.playerTimidity} 
                            Aggressiveness={this.state.playerAggressiveness} 
                            Risk={this.state.playerRisk}  
                            Bluff={this.state.playerBluff} 
                        />
                    </div>
                </>
                : null
            }
        </div>
        );
    }
}