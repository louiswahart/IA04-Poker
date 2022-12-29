import React from 'react';
import Select from 'react-select';

export default class Info extends React.Component {
      constructor(props) {
        super(props);

        this.state = {
            joueurInfo : false,
            modifStats : false,
            idPlayer : -1,
            timidity : 0,
            aggressiveness : 0,
            risk : 0,
            bluff : 0
        };
    }

    handleChangeTable = (e) => {
      if(Math.trunc(this.props.idPlayer/5) !== e.value){
        this.setState({joueurInfo : false});
        this.setState({modifStats: false}); 
      } 
      this.props.onTableChanged(e.value);
    };

    handleChangePlayer = (e) => {
      this.setState({idPlayer : e.value});
      this.setState({joueurInfo : true});
      this.setState({modifStats : false});  
      this.props.onPlayerChanged(e.value);
    };

    timidityModified = (e) => {    
      if(!this.state.modifStats){
        this.setState({aggressiveness : this.props.Aggressiveness}); 
        this.setState({risk : this.props.Risk}); 
        this.setState({bluff : this.props.Bluff});
        this.setState({modifStats : true});
      }
      this.setState({timidity : e.target.value}); 
    }

    aggressivenessModified = (e) => {    
      if(!this.state.modifStats){
        this.setState({timidity : this.props.Timidity});
        this.setState({risk : this.props.Risk}); 
        this.setState({bluff : this.props.Bluff}); 
        this.setState({modifStats : true});
      }
      this.setState({aggressiveness : e.target.value}); 
    }

    riskModified = (e) => {   
      if(!this.state.modifStats){
        this.setState({timidity : this.props.Timidity}); 
        this.setState({aggressiveness : this.props.Aggressiveness});
        this.setState({bluff : this.props.Bluff});
        this.setState({modifStats : true}); 
      }
      this.setState({risk : e.target.value}); 
    }

    bluffModified = (e) => { 
      if(!this.state.modifStats){
        this.setState({timidity : this.props.Timidity}); 
        this.setState({aggressiveness : this.props.Aggressiveness}); 
        this.setState({risk : this.props.Risk}); 
        this.setState({modifStats : true}); 
      }
      this.setState({bluff : e.target.value}); 
    }

    handleChangeStat = (e) => {
      e.preventDefault();
      this.setState({modifStats : false}, this.props.onStatsChanged(this.state.idPlayer, this.state.timidity, this.state.aggressiveness, this.state.risk, this.state.bluff));
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
            {this.props.loading ? <b><p>Chargement en cours</p></b> :
              (this.state.joueurInfo ? (<div className='joueurInfo'>
              <b><p>Statistiques du joueur sélectionné :</p></b>
              <form onSubmit={this.handleChangeStat}>        
                <label>
                  Timidité 
                  <input type="number" min="0" max="99" value={!this.state.modifStats ? this.props.Timidity : this.state.timidity} onChange={this.timidityModified} required/>        
                </label>
                <br></br>
                <label>
                  Agressivité
                  <input type="number" min="0" max="99" value={!this.state.modifStats ? this.props.Aggressiveness : this.state.aggressiveness} onChange={this.aggressivenessModified} required/>        
                </label>
                <br></br>
                <label>
                  Risque
                  <input type="number" min="0" max="99" value={!this.state.modifStats ? this.props.Risk : this.state.risk} onChange={this.riskModified} required/>        
                </label>
                <br></br>
                <label>
                  Bluff
                  <input type="number" min="0" max="99" value={!this.state.modifStats ? this.props.Bluff : this.state.bluff} onChange={this.bluffModified} required/>        
                </label>
                <br></br>
                <br></br>
                {this.state.modifStats && <input type="submit" value="Sauvegarder" />}
              </form>
              </div>) : null)}
        </div>
        );
    }
  }