import * as React from 'react';
import * as classnames from 'classnames';
import { ReactNode } from 'react';

export interface IconButtonProps {
    text?: string;
    children?: ReactNode;
    onClick?: () => void;
    disabled?: boolean;
}

class IconButton extends React.Component<IconButtonProps> {

    constructor(props: IconButtonProps) {
        super(props);
    }

    render() {
        let classes = classnames('no-underline inline-flex items-center tc f7 br2 pa2 mt1 ml1 bg-mid-gray white', { 'o-30': this.props.disabled, 'dim': !this.props.disabled, 'pointer': !this.props.disabled });
        return (
            <a className={classes} onClick={this.onClick} title={this.props.text}>
                {this.props.children}
            </a>
        );
    }

    onClick = () => {
        if (this.props.onClick !== undefined && this.isEnabled()) {
            this.props.onClick();
        }
    }

    isEnabled() {
        return this.props.disabled === undefined || this.props.disabled === false;
    }
}

export default IconButton;
