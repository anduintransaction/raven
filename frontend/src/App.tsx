import * as React from 'react';
import LeftPanel from './components/LeftPanel';
import RightPanel from './components/RightPanel';

class App extends React.Component {
    render() {
        return (
            <div className="sans-serif">
                <div className="fl w-40">
                    <LeftPanel />
                </div>
                <div className="fl w-60">
                    <RightPanel />
                </div>
            </div>
        );
    }
}

export default App;
