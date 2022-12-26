import React from 'react';

export default class State extends React.Component{
    render(){
        return (
            <div>
                <h1>Etat de la Table :</h1>
                <p className='pTable'>Tour actuel : {this.props.Turn}</p>
                <p className='pTable'>Partie actuelle : {this.props.Game}</p>
                <p className='pTable'>Pot : {this.props.Pot}</p>
                <p className='pTable'>Etat : {this.props.Etat}</p>
            </div>
        )
    }
}