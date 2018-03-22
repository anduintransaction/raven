import * as React from 'react';
import SearchBox from './SearchBox';
import ToolBox from './ToolBox';
import FilterPanel from './FilterPanel';
import PaginationBox from './PaginationBox';
import * as classNames from 'classnames';
import { MessageQuery } from '../models/Messages';

interface FilterBoxProps {
    query: MessageQuery;
    count: number;
}

interface FilterBoxState {
    hideFilterPanel: boolean;
}

class FilterBox extends React.Component<FilterBoxProps, FilterBoxState> {

    constructor(props: FilterBoxProps) {
        super(props);
        this.state = { hideFilterPanel: true };
    }

    handleFilterClick = () => {
        this.setState((prevState, props) => ({
            hideFilterPanel: !prevState.hideFilterPanel
        }));
    }

    render() {
        let classes = classNames('cl w-100', {
            'dn': this.state.hideFilterPanel
        });
        return (
            <div className="bb b--black-30 cf">
                <div className="fl w-40 pa3">
                    <SearchBox />
                </div>
                <div className="fl w-20 pa3 pl0">
                    <ToolBox onFilterClick={this.handleFilterClick} />
                </div>
                <div className="fl w-40 pa3 pl0">
                    <PaginationBox page={this.props.query.Page} itemsPerPage={this.props.query.ItemsPerPage} count={this.props.count} />
                </div>
                <div className={classes}>
                    <FilterPanel />
                </div>
            </div>
        );
    }
}

export default FilterBox;
