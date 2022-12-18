import React from 'react';
import Player from './Player'
import Info from './Info'

class Interactions extends React.Component {
    render() {
        return (
        <div>
            <div className='boutons'>
                <button onClick={this.props.onPlay}>Jouer</button>
                <button onClick={this.props.onPause}>Pause</button>
                <button onClick={this.props.onReset}>Reset</button>
            </div>
        </div>
        );
    }
}

export default class Game extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            nbTable: 50,
            nbGame: 50,
            token: 0,
            idTable: 0,
            play: false,
            playersId : [1, 2, 3, 8, 5],
            playersToken: [],
            playersBet: [],
            playersCards: [[]],
        };
    }

    componentDidMount(){
        setInterval(() => {
            if(this.state.play){
                const options = {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ Info: 'play' })
                };
                fetch('http://localhost:8080/play', options)
                    .then(resp => resp.json())
                    .then(data => this.setState({ token: data.Token }));
            }
          }, 3000)

        /*setInterval(() => {
            const options = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ Update: 'update' })
            };
            fetch('http://localhost:8080/update', options)
                .then(resp => resp.json())
                .then(data => this.setState({ nbTables: data.NbTables,nbGames: data.NbGames }));
        }, 3000)*/
    }

    changedTable(i){
        console.log("Table : " + i);
        /*const options = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ Update: 'update', nbTable : i })
        };
        fetch('http://localhost:8080/getTable', options)
            .then(resp => resp.json())
            .then(data => this.setState({ nbTables: data.NbTables,nbGames: data.NbGames }));*/
    }

    Play() {
        this.setState({playersId : [4,8,7,0,2]})
        //this.setState({play : true});
    }

    Pause(){
        this.setState({play : false});
    }

    Reset(){
        this.setState({token : 0});
    }

    render() {
        return (
        <div className='all'>
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
                    <Interactions onPlay={() => this.Play()}
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
        );
    }
}