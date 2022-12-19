import React from 'react';

export default class Interactions extends React.Component {
    render() {
        return (
        <div>
            <div className='boutons'>
                <button onClick={this.props.onPlay}>Jouer</button>
                <button onClick={this.props.onPause}>Pause</button>
                <button onClick={this.props.onReset}>Reset</button>
            </div>
        </div>
        );
    }
}