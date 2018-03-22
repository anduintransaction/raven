import * as React from 'react';
import { ReactNode } from 'react';

export interface IconButtonProps {
    text?: string;
    children?: ReactNode;
    onClick?: () => void;
}

class IconButton extends React.Component<IconButtonProps> {

    constructor(props: IconButtonProps) {
        super(props);
    }

    render() {
        return (
            <a className="no-underline inline-flex items-center tc f7 br2 pa2 mt1 ml1 dim bg-mid-gray white pointer" onClick={this.props.onClick} title={this.props.text}>
                {this.props.children}
            </a>
        );
    }
}

export default IconButton;
