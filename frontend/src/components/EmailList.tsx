import * as React from 'react';
import * as classnames from 'classnames';
import Email from '../models/Email';
import EmailItem from './EmailItem';

interface EmailListProps {
    emails: Array<Email>;
    onEmailItemClick?: (emailID: number) => void;
}

interface EmailListState {
    activeEmailID?: number;
}

class EmailList extends React.Component<EmailListProps, EmailListState> {

    constructor(props: EmailListProps) {
        super(props);
        this.state = {};
    }

    render() {
        let emails = this.props.emails.map((email) => {
            let classes = classnames('pa3 bb b--black-10 pointer', {
                'bg-light-gray': this.state.activeEmailID === email.ID
            });
            return (
                <li className={classes} key={email.ID} onClick={() => this.onEmailItemClick(email.ID)}>
                    <EmailItem email={email} />
                </li>
            );
        });
        return (
            <div className="cl">
                <ul className="list pa0 ma0">
                    {emails}
                </ul>
            </div>
        );
    }

    onEmailItemClick = (emailID: number) => {
        this.setState({ activeEmailID: emailID });
        if (this.props.onEmailItemClick !== undefined) {
            this.props.onEmailItemClick(emailID);
        }
    }
}

export default EmailList;
