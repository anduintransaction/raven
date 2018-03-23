import * as React from 'react';

interface SearchBoxProps {
    search: string;
    onSubmit?: (search: string) => void;
}

interface SearchBoxState {
    search: string;
}

class SearchBox extends React.Component<SearchBoxProps, SearchBoxState> {

    constructor(props: SearchBoxProps) {
        super(props);
        this.state = { search: this.props.search };
    }

    render() {
        return (
            <form onSubmit={this.onSubmit}>
                <input
                    autoComplete="off"
                    name="search"
                    className="input-reset ba b--black-20 pa2 mb2 db w-100"
                    type="text"
                    placeholder="Search"
                    value={this.state.search}
                    onChange={this.onChange}
                />
            </form>
        );
    }

    componentWillReceiveProps(props: SearchBoxProps) {
        this.setState({ search: props.search });
    }

    onSubmit = (event: React.SyntheticEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (this.props.onSubmit !== undefined) {
            this.props.onSubmit(this.state.search);
        }
    }

    onChange = (event: React.SyntheticEvent<HTMLInputElement>) => {
        this.setState({ search: event.currentTarget.value });
    }
}

export default SearchBox;
