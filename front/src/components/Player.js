import React from 'react';

export default class Player extends React.Component{
    render(){
        //boucle for pour les cartes qui creer une liste avec les cartes
        var listeCards = this.props.Cards
        return (
            <div>
                <p>ID : {this.props.ID}</p>
                <p>Jetons : {this.props.nbToken}</p>
                <p>Mise : {this.props.bet}</p>
                <p>{listeCards}</p>
            </div>
        )
    }
}