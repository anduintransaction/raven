import Attachment from './Attachment';
import EmailContent from './EmailContent';

interface Email {
    ID: number;
    MessageID: number;
    FromEmail: string;
    FromName: string;
    ToEmail: string;
    ToName: string;
    RCPT: string;
    ReplyTo: string;
    Subject: string;
    EmailContent: EmailContent;
    Attachments: Array<Attachment>;
    CreatedAt: string;
}

export default Email;
