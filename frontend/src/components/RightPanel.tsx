import * as React from 'react';
import * as moment from 'moment';
import { FaEnvelope, FaUser } from 'react-icons/lib/fa';
import * as classnames from 'classnames';
import Email from '../models/Email';

interface RightPanelProps {
    emailID?: number;
}

interface RightPanelState {
    email?: Email;
}

class RightPanel extends React.Component<RightPanelProps, RightPanelState> {

    constructor(props: RightPanelProps) {
        super(props);
        this.state = {};
    }

    componentWillReceiveProps(nextProps: RightPanelProps) {
        if (nextProps.emailID !== this.props.emailID) {
            fetch(`/api/message/${nextProps.emailID}`)
                .then((response: Response) => {
                    return response.json();
                }).then((email: Email) => {
                    this.setState({ email: email });
                });
        }
    }

    render() {
        let email = this.state.email;
        let content: JSX.Element;
        if (email === undefined) {
            content = (
                <div className="vh-100 dt w-100">
                    <div className="dtc v-mid tc ph3 ph4-l">
                        <div className="tc f1 f-subheadline-l mid-gray"><FaEnvelope /></div>
                        <div className="tc pt3 mid-gray">Click item on the left to display email content</div>
                    </div>
                </div>
            );
        } else {
            let attachmentContent: JSX.Element;
            if (email.Attachments == null || email.Attachments.length === 0) {
                attachmentContent = (
                    <div />
                );
            } else {
                let attachments = email.Attachments;
                let attachmentList = attachments.map((attachment, index) => {
                    let classes = classnames('pa2', {
                        'bb b--black-10': (index < attachments.length - 1)
                    });
                    return (
                        <li key={attachment.ID} className={classes}>
                            <a className="link blue dim" href={`/api/attachment/${attachment.ID}/download`}>
                                {attachment.Filename}
                            </a>
                        </li>
                    );
                });
                attachmentContent = (
                    <div className="bg-light-gray">
                        <ul className="list pa0 ma0 f7">
                            {attachmentList}
                        </ul>
                    </div>
                );
            }
            content = (
                <div className="pv3 ph4">
                    <h2 className="bb black-70 b--black-10 pb3">{email.Subject}</h2>
                    <div className="cf">
                        <div className="fl w-60">
                            <div className="fl w-10 f2 mid-gray">
                                <FaUser />
                            </div>
                            <div className="fl w-90">
                                <b className="f6">{this.renderEmailAddress(email.FromEmail, email.FromName)}</b>
                                <br />
                                <span className="f7 mid-gray">To: {this.renderEmailAddress(email.ToEmail, email.ToName)}</span>
                            </div>
                        </div>
                        <div className="fl w-40 tr mid-gray f7">
                            {moment(email.CreatedAt).calendar()} ({moment(email.CreatedAt).fromNow()})
                        </div>
                    </div>
                    <div className="cf pv4" dangerouslySetInnerHTML={{ __html: email.EmailContent.HTML }} />
                    {attachmentContent}
                </div>
            );
        }
        return content;
    }

    renderEmailAddress(email: string, name: string): string {
        return name !== '' ? `${name} <${email}>` : `${email}`;
    }
}

export default RightPanel;
