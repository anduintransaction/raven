import * as React from 'react';
import Email from '../models/Email';
import EmailItem from './EmailItem';

interface EmailListProps {
    emails: Array<Email>;
}

class EmailList extends React.Component<EmailListProps> {

    constructor(props: EmailListProps) {
        super(props);
    }

    render() {
        let emails = this.props.emails.map((email) => {
            return (
                <li className="pa3 bb b--black-10 pointer" key={email.ID}>
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
}

export default EmailList;
