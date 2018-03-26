import * as React from 'react';
import SearchBox from './SearchBox';
import ToolBox from './ToolBox';
import PaginationBox from './PaginationBox';
import { MessageQuery } from '../models/Messages';

interface FilterBoxProps {
    disabled?: boolean;
    query: MessageQuery;
    count: number;
    onSearchBoxSubmit?: (search: string) => void;
    onSortButtonClick?: (direction: string) => void;
    onRefreshButtonClick?: () => void;
    onClearButtonClick?: () => void;
    onPreviousButtonClick?: () => void;
    onNextButtonClick?: () => void;
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
                        disabled={this.props.disabled}
                        sortDirection={sortDirection}
                        onSortClick={this.props.onSortButtonClick}
                        onRefreshClick={this.props.onRefreshButtonClick}
                        onClearClick={this.props.onClearButtonClick}
                    />
                </div>
                <div className="fl w-40 pa3 pl0">
                    <PaginationBox
                        disabled={this.props.disabled}
                        page={this.props.query.Page}
                        itemsPerPage={this.props.query.ItemsPerPage}
                        count={this.props.count}
                        onPreviousClick={this.props.onPreviousButtonClick}
                        onNextClick={this.props.onNextButtonClick}
                    />
                </div>
            </div>
        );
    }
}

export default FilterBox;
