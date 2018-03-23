import * as React from 'react';
import FilterBox from './FilterBox';
import EmailList from './EmailList';
import { MessageQuery, MessagesResponse } from '../models/Messages';

interface LeftPanelProps {
    onEmailItemClick?: (emailID: number) => void;
}

interface LeftPanelState {
    query: MessageQuery;
    response: MessagesResponse;
}

class LeftPanel extends React.Component<LeftPanelProps, LeftPanelState> {

    initialState: LeftPanelState = {
        query: {
            Filter: { From: '', To: '' },
            Search: '',
            Sorts: [{ Field: 'created_at', Direction: 'DESC' }],
            Page: 1,
            ItemsPerPage: 10
        },
        response: { Count: 0, Emails: [] }
    };

    constructor(props: LeftPanelProps) {
        super(props);
        this.state = JSON.parse(JSON.stringify(this.initialState));
    }

    componentDidMount() {
        this.fetch();
    }

    render() {
        return (
            <div className="br b--black-30">
                <FilterBox
                    query={this.state.query}
                    count={this.state.response.Count}
                    onSearchBoxSubmit={this.onSearchBoxSubmit}
                    onSortButtonClick={this.onSortButtonClick}
                    onClearButtonClick={this.onClearButtonClick}
                />
                <EmailList emails={this.state.response.Emails} onEmailItemClick={this.props.onEmailItemClick} />
            </div>
        );
    }

    onSearchBoxSubmit = (search: string) => {
        this.setState(
            (prevState: LeftPanelState, prevProp: LeftPanelProps) => {
                prevState.query.Search = search;
                return { query: prevState.query };
            },
            this.fetch
        );
    }

    onSortButtonClick = (direction: string) => {
        this.setState(
            (prevState, prevProp) => {
                prevState.query.Sorts = [{ Field: 'created_at', Direction: direction }];
                return prevState;
            },
            this.fetch
        );
    }

    onClearButtonClick = () => {
        this.setState(this.initialState, this.fetch);
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
