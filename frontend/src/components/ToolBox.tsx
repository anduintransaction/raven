import * as React from 'react';
import { FaSortAmountAsc, FaFilter, FaRefresh } from 'react-icons/lib/fa';
import IconButton from './IconButton';

export interface ToolBoxProps {
    onFilterClick?: () => void;
}

class ToolBox extends React.Component<ToolBoxProps> {

    constructor(props: ToolBoxProps) {
        super(props);
    }

    render() {
        return (
            <div>
                <IconButton text="Sort">
                    <FaSortAmountAsc />
                </IconButton>
                <IconButton text="Filter" onClick={this.props.onFilterClick}>
                    <FaFilter />
                </IconButton>
                <IconButton text="Refresh">
                    <FaRefresh />
                </IconButton>
            </div>
        );
    }
}

export default ToolBox;
