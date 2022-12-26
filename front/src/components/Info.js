import React from 'react';
import Select from 'react-select';

export default class Info extends React.Component {
      constructor(props) {
        super(props);

        this.state = {
            joueurInfo : false
        };
    }

    handleChangeTable = (e) => {
      if(Math.trunc(this.props.idPlayer/5) !== e.value) this.setState({joueurInfo : false})
      this.props.onTableChanged(e.value);
    };

    handleChangePlayer = (e) => {
      this.setState({joueurInfo : true})
      this.props.onPlayerChanged(e.value);
    };
    
    render() {
        let listTables = []

        for (let i=0;i<this.props.nbTable;i++) {
            listTables.push(
              {value : i,label:'Table ' + i}
            )    
        }

        let listPlayers = []

        for (let i=0;i<this.props.nbPlayers;i++) {
            listPlayers.push(
              {value : i,label:'Joueur ' + i}
            )    
        }

        return (
        <div>
            <Select onChange={this.handleChangeTable}
              placeholder="Choisir une table"
              value={listTables[this.props.idTable]}
              maxMenuHeight={1000}
              options = {listTables}
            />
            <Select onChange={this.handleChangePlayer}
              placeholder="Choisir un joueur"
              value={this.props.idPlayer === -1 ? null : listPlayers[this.props.idPlayer]}
              options = {listPlayers} 
              maxMenuHeight={1000}
            />
            {this.state.joueurInfo ? (<div className='joueurInfo'>
              <b><p>Statistiques du joueur sélectionné :</p></b>
              <p className='pStats'>Timidité : {this.props.Timidity}</p>
              <p className='pStats'>Agressivité : {this.props.Aggressiveness}</p>
              <p className='pStats'>Risque : {this.props.Risk}</p>
              <p className='pStats'>Bluff : {this.props.Bluff}</p>
            </div>) : null}
        </div>
        );
    }
  }