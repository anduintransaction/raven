import * as React from 'react';
import { FaChevronLeft, FaChevronRight } from 'react-icons/lib/fa';
import IconButton from './IconButton';

interface PaginationBoxProps {
    page: number;
    itemsPerPage: number;
    count: number;
}

class PaginationBox extends React.Component<PaginationBoxProps> {

    constructor(props: PaginationBoxProps) {
        super(props);
    }

    render() {
        let start = (this.props.page - 1) * this.props.itemsPerPage + 1;
        let end = this.props.page * this.props.itemsPerPage;
        if (end > this.props.count) {
            end = this.props.count;
        }
        return (
            <div>
                <IconButton text="Previous">
                    <FaChevronLeft />
                </IconButton>
                <IconButton text="Previous">
                    <FaChevronRight />
                </IconButton>
                <span className="ml3 f6">
                    <strong>{start}</strong> - <strong>{end}</strong> of <strong>{this.props.count}</strong>
                </span>
            </div>
        );
    }
}

export default PaginationBox;
