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

        var ID = <p className='pJoueur'>ID : {this.props.ID}</p>
        if(this.props.blind > 0){
            if(this.props.blind === 1){
                ID = <p className='pJoueur'>ID : {this.props.ID} | <b>Small Blind</b></p>
            } else{
                ID = <p className='pJoueur'>ID : {this.props.ID} | <b>Big Blind</b></p>
            }
        }

        return (
            <>
                <div className={this.props.winner ? 'background-green' : 
                this.props.action === "Je me couche" ? 'background-grey' :
                this.props.action === "Je n'ai plus de jeton pour jouer" ? 'background-red' : 
                (this.props.action === "Terminé" && this.props.nbToken === 0) ? 'background-red' : 'background-basic'}>
                    {ID}
                    <p className={'pJoueur'+ (this.props.action === "Terminé" ? ' bold' : '')}>Jetons : {this.props.nbToken}</p>
                    <p className='pJoueur'>Mise du tour : {this.props.bet}</p>
                    <p className='pJoueur'>Mise de la partie : {this.props.totalBet}</p>
                    <p className='pJoueur'>Action : {this.props.action}</p>
                    {Cards}
                    {this.props.lastTurn && <p className='pJoueur'>Bénéfice de la partie : {this.props.result}</p>}
                </div>
            </>
        )
    }
}