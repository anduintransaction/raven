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
                    onRefreshButtonClick={this.fetch}
                    onClearButtonClick={this.onClearButtonClick}
                    onPreviousButtonClick={this.onPreviousButtonClick}
                    onNextButtonClick={this.onNextButtonClick}
                />
                <EmailList emails={this.state.response.Emails} onEmailItemClick={this.props.onEmailItemClick} />
            </div>
        );
    }

    onSearchBoxSubmit = (search: string) => {
        this.setState(
            (prevState: LeftPanelState, prevProp: LeftPanelProps) => {
                let newState = JSON.parse(JSON.stringify(prevState));
                newState.query.Search = search;
                return newState;
            },
            this.fetch
        );
    }

    onSortButtonClick = (direction: string) => {
        this.setState(
            (prevState, prevProp) => {
                let newState = JSON.parse(JSON.stringify(prevState));
                newState.query.Sorts = [{ Field: 'created_at', Direction: direction }];
                return newState;
            },
            this.fetch
        );
    }

    onClearButtonClick = () => {
        this.setState(JSON.parse(JSON.stringify(this.initialState)), this.fetch);
    }

    onPreviousButtonClick = () => {
        this.setState(
            (prevState, prevProps) => {
                let newState = JSON.parse(JSON.stringify(prevState));
                if (newState.query.Page > 1) {
                    newState.query.Page--;
                }
                return newState;
            },
            this.fetch
        );
    }

    onNextButtonClick = () => {
        this.setState(
            (prevState, prevProps) => {
                let newState = JSON.parse(JSON.stringify(prevState));
                let maxPage = Math.ceil(newState.response.Count / newState.query.ItemsPerPage);
                if (newState.query.Page < maxPage) {
                    newState.query.Page++;
                }
                return newState;
            },
            this.fetch
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
