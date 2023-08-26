export interface Attachment {
  filename: string;
  contentType: string;
  data?: string;
  index: number;
}
export interface Email {
  id: string;
  html: string;
  plain: string;
  source: string;
  headers: {[key: string]: Array<string>};
  attachments: Array<Attachment>;
}
