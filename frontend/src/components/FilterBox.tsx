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
    onSearchBoxSubmit?: (search: string) => void;
    onSortButtonClick?: (direction: string) => void;
    onClearButtonClick?: () => void;
}

interface FilterBoxState {
    hideFilterPanel: boolean;
}

class FilterBox extends React.Component<FilterBoxProps, FilterBoxState> {

    constructor(props: FilterBoxProps) {
        super(props);
        this.state = { hideFilterPanel: false };
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
        let sortDirection = 'DESC';
        this.props.query.Sorts.map((sorter) => {
            sortDirection = sorter.Direction;
        });
        return (
            <div className="bb b--black-30 cf">
                <div className="fl w-40 pa3">
                    <SearchBox onSubmit={this.props.onSearchBoxSubmit} search={this.props.query.Search} />
                </div>
                <div className="fl w-20 pa3 pl0">
                    <ToolBox
                        sortDirection={sortDirection}
                        onSortClick={this.props.onSortButtonClick}
                        onFilterClick={this.handleFilterClick}
                        onClearClick={this.props.onClearButtonClick}
                    />
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
