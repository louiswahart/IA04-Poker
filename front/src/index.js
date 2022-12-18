import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

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

class Plateau extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            token: 0,
            play: false
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
    }


    componentDidMount(){
        setInterval(() => {
                const options = {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ Update: 'update' })
                };
                fetch('http://localhost:8080/update', options)
                    .then(resp => resp.json())
                    .then(data => this.setState({ nbTables: data.NbTables,nbGames: data.NbGames }));
            }, 3000)
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

    render() {
        return (
        <div className='all'>
            <div className="plateau">
                <div className='rectangle'>
                    <div className='joueur1'>
                        <div className='Card'></div>
                        <div className='Token'>{this.state.token}</div>
                    </div>
                    <div className='joueur2'>
                        <div className='Card'></div>
                        <div className='Token'>{this.state.token}</div>
                    </div>
                    <div className='joueur3'>
                        <div className='Card'></div>
                        <div className='Token'>{this.state.token}</div>
                    </div>
                    <div className='joueur4'>
                        <div className='Card'></div>
                        <div className='Token'>{this.state.token}</div>
                    </div>
                    <div className='joueur5'>
                        <div className='Card'></div>
                        <div className='Token'>{this.state.token}</div>
                    </div>
                </div>
            </div>
            <div className="interactions">
                <Interactions onPlay={() => this.Play()}
                    onPause={() => this.Pause()}
                    onReset={() => this.Reset()}
                />
            </div>
        </div>
        );
    }
}
  
class Info extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        nbTables: 0,
        nbGames: 0,
        displayTables: false
    };
  }

  componentDidMount(){
    setInterval(() => {
            const options = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ Update: 'update' })
            };
            fetch('http://localhost:8080/update', options)
                .then(resp => resp.json())
                .then(data => this.setState({ nbTables: data.NbTables,nbGames: data.NbGames }));
        }, 3000)
}

  openTables () {
    this.setState((state,props) => ({displayTables: true}));
  }

  closeTables () {
    this.setState((state,props) => ({displayTables: false}));
  }
  
    render() {
        let listButton = []

        for (let i=0;i<=5;i++) {
            listButton.push(
            <button>
                Table {i}
            </button>
            )    
        }

        return (
        <div>
            <button onClick={this.openTables.bind(this)}>Tables</button>

            {
              this.state.displayTables ? (
            <div className='Tables'>
                {listButton}
            </div>
              ) : (null)
            }
            <button>Statistiques</button>
        </div>
        );
    }
}

class App extends React.Component {
    render() { 
        return (
        <div className='container'>
            <div className="jeu">
                <Plateau />
            </div>  
            <div className="infos">
                <Info />
            </div>
        </div>
        );
    }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<App />);
