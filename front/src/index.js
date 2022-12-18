import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Game from './components/Game';


class App extends React.Component {
    render() { 
        return (
        <div className='container'>
            <Game />
        </div>
        );
    }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<App />);
