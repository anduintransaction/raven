import * as React from 'react';
import { FaSortAmountAsc, FaSortAmountDesc, FaRefresh } from 'react-icons/lib/fa';
import { TiTimes } from 'react-icons/lib/ti';
import IconButton from './IconButton';

export interface ToolBoxProps {
    disabled?: boolean;
    sortDirection: string;
    onSortClick?: (direction: string) => void;
    onRefreshClick?: () => void;
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
                <IconButton text="Sort" onClick={this.onSortButtonClick} disabled={this.props.disabled}>
                    {sortButton}
                </IconButton>
                <IconButton text="Refresh" onClick={this.props.onRefreshClick} disabled={this.props.disabled}>
                    <FaRefresh />
                </IconButton>
                <IconButton text="Clear" onClick={this.props.onClearClick} disabled={this.props.disabled}>
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
