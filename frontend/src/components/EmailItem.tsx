import * as React from 'react';
import * as moment from 'moment';
import Email from '../models/Email';

interface EmailItemProps {
    email: Email;
}

class EmailItem extends React.Component<EmailItemProps> {

    constructor(props: EmailItemProps) {
        super(props);
    }

    render() {
        let email = this.props.email;
        return (
            <div className="cl cf f6">
                <div className="pv1">
                    <b className="blue">{this.renderEmailAddress(email.FromEmail, email.FromName)}</b> &raquo; <b className="green f6">{this.renderEmailAddress(email.ToEmail, email.ToName)}</b>
                </div>
                <div className="mid-gray">
                    <div className="fl w-60 pv1 truncate">
                        {email.Subject}
                    </div>
                    <div className="fl w-40 pv1 pl2 tr">
                        {moment(email.CreatedAt).calendar()}
                    </div>
                </div>
            </div>
        );
    }

    renderEmailAddress(email: string, name: string): string {
        return name === '' ? email : name;
    }
}

export default EmailItem;
