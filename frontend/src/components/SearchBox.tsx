import * as React from 'react';

class SearchBox extends React.Component {
    render() {
        return (
            <input
                name="search"
                className="input-reset ba b--black-20 pa2 mb2 db w-100"
                type="text"
                placeholder="Search"
            />
        );
    }
}

export default SearchBox;
