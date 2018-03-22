import * as React from 'react';
import FilterBox from './FilterBox';
import EmailList from './EmailList';
import { MessageQuery, MessagesResponse } from '../models/Messages';

interface LeftPanelState {
    query: MessageQuery;
    response: MessagesResponse;
}

class LeftPanel extends React.Component<{}, LeftPanelState> {

    constructor(props: {}) {
        super(props);
        this.state = {
            query: {
                Filter: { From: '', To: '' },
                Search: '',
                Sorts: [{ Field: 'created_at', Direction: 'DESC' }],
                Page: 1,
                ItemsPerPage: 10
            },
            response: { Count: 0, Emails: [] }
        };
    }

    componentDidMount() {
        this.fetch();
    }

    render() {
        return (
            <div className="br b--black-30">
                <FilterBox query={this.state.query} count={this.state.response.Count} />
                <EmailList emails={this.state.response.Emails} />
            </div>
        );
    }

    fetch = () => {
        fetch('/api/message', {
            body: JSON.stringify(this.state.query),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        }).then(function(response: Response) {
            return response.json();
        }).then((data: MessagesResponse) => {
            this.setState({ response: data });
        });
    }
}

export default LeftPanel;
