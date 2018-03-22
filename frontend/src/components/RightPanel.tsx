import * as React from 'react';
import * as moment from 'moment';
import { FaEnvelope, FaUser } from 'react-icons/lib/fa';
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
            content = (
                <div className="pv3 ph4">
                    <h2 className="bb black-70 b--black-10 pb3">{email.Subject}</h2>
                    <div className="cf">
                        <div className="fl w-60">
                            <div className="fl w-10 f2 mid-gray">
                                <FaUser />
                            </div>
                            <div className="fl w-90">
                                <b className="f6">{email.FromName}</b>
                                <br />
                                <span className="f7 mid-gray">To: {email.ToName} &lt;{email.ToEmail}&gt;</span>
                            </div>
                        </div>
                        <div className="fl w-40 tr mid-gray f7">
                            {moment(email.CreatedAt).calendar()} ({moment(email.CreatedAt).fromNow()})
                        </div>
                    </div>
                    <div dangerouslySetInnerHTML={{ __html: email.EmailContent.HTML }} />
                </div>
            );
        }
        return content;
    }
}

export default RightPanel;
