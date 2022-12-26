import React from 'react';
import { getCard } from './Cards';

export default class Player extends React.Component{
    render(){
       
        let Cards = []

        if(this.props.listeCards != null){
            this.props.listeCards.forEach(element => {
                let color = ""
                if (element.Color === 0 || element.Color === 3){
                    color = "black"
                } else {
                    color = "red"
                }
                Cards.push(
                    <div className={'Card ' + color} key = {element.Color.toString() + element.Value.toString()}>{getCard(element.Color, element.Value)}</div> 
                  )    
            });
        }

        return (
            <>
                <div className={this.props.winner ? 'background-green' : 
                this.props.action === "Je me couche" ? 'background-grey' :
                this.props.action === "Je n'ai plus de jeton pour jouer" ? 'background-red' : 'background-basic'}>
                    <p className='pJoueur'>ID : {this.props.ID}</p>
                    <p className='pJoueur'>Jetons : {this.props.nbToken}</p>
                    <p className='pJoueur'>Mise du tour : {this.props.bet}</p>
                    <p className='pJoueur'>Mise de la partie : {this.props.totalBet}</p>
                    <p className='pJoueur'>Action : {this.props.action}</p>
                    {Cards}
                </div>
            </>
        )
    }
}