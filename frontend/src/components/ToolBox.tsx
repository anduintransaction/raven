import * as React from 'react';
import { FaSortAmountAsc, FaSortAmountDesc, FaFilter, FaRefresh } from 'react-icons/lib/fa';
import { TiTimes } from 'react-icons/lib/ti';
import IconButton from './IconButton';

export interface ToolBoxProps {
    sortDirection: string;
    onSortClick?: (direction: string) => void;
    onFilterClick?: () => void;
    onClearClick?: () => void;
}

export interface ToolBoxState {
    sortDirection: string;
}

class ToolBox extends React.Component<ToolBoxProps, ToolBoxState> {

    constructor(props: ToolBoxProps) {
        super(props);
        this.state = {
            sortDirection: this.props.sortDirection
        };
    }

    render() {
        let sortButton = this.state.sortDirection === 'DESC' ? <FaSortAmountDesc /> : <FaSortAmountAsc />;
        return (
            <div>
                <IconButton text="Sort" onClick={this.onSortButtonClick}>
                    {sortButton}
                </IconButton>
                <IconButton text="Filter" onClick={this.props.onFilterClick}>
                    <FaFilter />
                </IconButton>
                <IconButton text="Refresh">
                    <FaRefresh />
                </IconButton>
                <IconButton text="Clear" onClick={this.props.onClearClick}>
                    <TiTimes />
                </IconButton>
            </div>
        );
    }

    onSortButtonClick = () => {
        this.setState(
            (prevState, prevProps) => {
                return {
                    sortDirection: prevState.sortDirection === 'DESC' ? 'ASC' : 'DESC'
                };
            },
            () => {
                if (this.props.onSortClick !== undefined) {
                    this.props.onSortClick(this.state.sortDirection);
                }
            }
        );
    }
}

export default ToolBox;
