import React from 'react';

export default class Player extends React.Component{
    render(){
       
        //boucle for pour les cartes qui creer une liste avec les cartes
        var listeCards = this.props.Cards
        return (
            <div>
                <div className={this.props.winner ? 'background-green' : 'background-basic'}>
                    <p>ID : {this.props.ID}</p>
                    <p>Jetons : {this.props.nbToken}</p>
                    <p>Mise : {this.props.bet}</p>
                    <p>Action : {this.props.action}</p>
                    <p>{listeCards}</p>
                </div>
            </div>
        )
    }
}