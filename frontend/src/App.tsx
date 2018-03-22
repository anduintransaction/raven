import * as React from 'react';
import LeftPanel from './components/LeftPanel';
import RightPanel from './components/RightPanel';

interface AppState {
    displayingEmailID?: number;
}

class App extends React.Component<{}, AppState> {

    constructor(props: {}) {
        super(props);
        this.state = {};
    }

    render() {
        return (
            <div className="sans-serif">
                <div className="fl w-40">
                    <LeftPanel onEmailItemClick={this.onEmailItemClick} />
                </div>
                <div className="fl w-60">
                    <RightPanel emailID={this.state.displayingEmailID} />
                </div>
            </div>
        );
    }

    onEmailItemClick = (emailID: number) => {
        this.setState({
            displayingEmailID: emailID
        });
    }
}

export default App;
