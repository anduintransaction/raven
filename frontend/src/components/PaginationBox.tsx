import * as React from 'react';
import { FaChevronLeft, FaChevronRight } from 'react-icons/lib/fa';
import IconButton from './IconButton';

interface PaginationBoxProps {
    disabled?: boolean;
    page: number;
    itemsPerPage: number;
    count: number;
    onPreviousClick?: () => void;
    onNextClick?: () => void;
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
        let maxPage = Math.ceil(this.props.count / this.props.itemsPerPage);
        let pageString: JSX.Element;
        if (this.props.count > 0) {
            pageString = (
                <span className="ml3 f6">
                    <strong>{start}</strong> - <strong>{end}</strong> of < strong > {this.props.count}</strong >
                </span>
            );
        } else {
            pageString = <span />;
        }
        return (
            <div>
                <IconButton text="Previous" disabled={this.props.disabled || this.props.page <= 1} onClick={this.props.onPreviousClick}>
                    <FaChevronLeft />
                </IconButton>
                <IconButton text="Next" disabled={this.props.disabled || this.props.page >= maxPage} onClick={this.props.onNextClick}>
                    <FaChevronRight />
                </IconButton>
                <span className="ml3 f6">
                    {pageString}
                </span>
            </div>
        );
    }
}

export default PaginationBox;
