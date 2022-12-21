import React from 'react';
import Select from 'react-select';

export default class Info extends React.Component {
    handleChangeTable = (e) => {
      this.props.onTableChanged(e.value);
    };

    handleChangePlayer = (e) => {
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
              name="Choix des tables"
              defaultValue={listTables[0]}
              maxMenuHeight={200}
              options = {listTables}
            />
            <Select onChange={this.handleChangePlayer}
              placeholder="Choix des joueurs"
              options = {listPlayers} 
              maxMenuHeight={200}
            />
            <div className='joueurInfo'>
              <p>Timidité : {this.props.Timidity}</p>
              <p>Agressivité : {this.props.Aggressiveness}</p>
              <p>Risque : {this.props.Risk}</p>
              <p>Bluff : {this.props.Bluff}</p>
            </div>
        </div>
        );
    }
  }