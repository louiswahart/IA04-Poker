import React from 'react';

export default class State extends React.Component{
    render(){
       
        return (
            <div>
                <p>Etat de la Table :</p>
                <p>Tour actuel : {this.props.Turn}</p>
                <p>Partie actuelle : {this.props.Game}</p>
                <p>Pot : {this.props.Pot}</p>
            </div>
        )
    }
}