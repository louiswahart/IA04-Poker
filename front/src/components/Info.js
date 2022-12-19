import React from 'react';

export default class Info extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
          displayTables: false
      };
    }

    openTables () {
      this.setState((state,props) => ({displayTables: true}));
    }
  
    closeTables () {
      this.setState((state,props) => ({displayTables: false}));
    }
    
    render() {
        let listButton = []

        for (let i=0;i<this.props.nbTable;i++) {
            listButton.push(
              <div>
                <button onClick={() => this.props.onTableChanged(i)}>
                    Table {i}
                </button>
                <br/>
            </div>
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