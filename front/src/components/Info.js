import React from 'react';
import Select from 'react-select';

export default class Info extends React.Component {
    constructor(props) {
      super(props);
    }

    handleChange = (e) => {
      this.props.onTableChanged(e.value);
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
            <Select onChange={this.handleChange}
              name="Choix des tables"
              defaultValue={listTables[0]}
              maxMenuHeight={200}
              options = {listTables}
            />
            <Select 
              placeholder="Choix des joueurs"
              options = {listPlayers} 
              maxMenuHeight={200}
            />
        </div>
        );
    }
  }