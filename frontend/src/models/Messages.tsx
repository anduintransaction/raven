import Email from './Email';

export interface MessageQuery {
    Filter: MessageFilter;
    Search: string;
    Sorts: Array<MessageSorter>;
    Page: number;
    ItemsPerPage: number;
}

export interface MessageFilter {
    From: string;
    To: string;
}

export interface MessageSorter {
    Field: string;
    Direction: string;
}

export interface MessagesResponse {
    Count: number;
    Emails: Array<Email>;
}
