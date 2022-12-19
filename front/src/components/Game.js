import React from 'react';
import Player from './Player'
import Info from './Info'
import Interactions from './Interactions'

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
            playersId : [],
            playersToken: [],
            playersBet: [],
            playersCards: [[]],
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
                        .then(data => this.setState({ playersId: data.PlayersID,  playersToken: data.PlayersToken, playersBet: data.PlayersBet}));
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
            .then(data => this.setState({ playersId: data.PlayersID,  playersToken: data.PlayersToken, playersBet: data.PlayersBet}));
        this.setState({idTable : i})
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
                    <div className="jeu">
                        <div className="plateau">
                            <div className='rectangle'>
                                <Player 
                                    ID={this.state.playersId[0]} 
                                    nbToken={this.state.playersToken[0]} 
                                    bet={this.state.playersBet[0]} 
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[1]} 
                                    nbToken={this.state.playersToken[1]} 
                                    bet={this.state.playersBet[1]} 
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[2]} 
                                    nbToken={this.state.playersToken[2]} 
                                    bet={this.state.playersBet[2]} 
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[3]} 
                                    nbToken={this.state.playersToken[3]} 
                                    bet={this.state.playersBet[3]} 
                                    listeCards={20}
                                />
                                <Player 
                                    ID={this.state.playersId[4]} 
                                    nbToken={this.state.playersToken[4]} 
                                    bet={this.state.playersBet[4]}  
                                    listeCards={20}
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
                            onTableChanged={i => this.changedTable(i)}
                        />
                    </div>
                </div>
                : null
            }
        </div>
        );
    }
}