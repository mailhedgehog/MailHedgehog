export interface Attachment {
  filename: string;
  contentType: string;
  data?: string;
  index: number;
}
export interface EmailAddress {
  name: string;
  email: string;
}

export interface Email {
  id: string;
  html: string;
  plain: string;
  source: string;
  headers: {[key: string]: Array<string>};
  attachments: Array<Attachment>;
}
export interface EmailItem {
  id: string;
  from: EmailAddress;
  to: Array<EmailAddress>;
  subject: string;
  received_at: string;
  size: number;
}
