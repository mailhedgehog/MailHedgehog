import { inject, Ref } from 'vue';
import { Attachment, Email } from '@/interfaces/Email';
import { MailHedgehog } from '@/main';

export function useEmailAttachments(email: Ref<Email|null>, options: {
  urlBuilder?: (index: number, email: Ref<Email|null>)=>string
} = {}) {
  const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
  const downloadBlobFile = (data: string, contentType: string, filename: string) => {
    const blob = new Blob([data], { type: contentType });
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = filename;
    link.click();
    URL.revokeObjectURL(link.href);
    link.remove();
  };

  const buildDownloadAttachmentUrl = (index: number) => {
    if (options.urlBuilder) {
      return options.urlBuilder(index, email);
    }

    return `emails/${email.value?.id}/attachment/${index}`;
  };

  const downloadAttachment = (index: number, filename: string) => {
    if (!email.value) {
      return;
    }
    const url = buildDownloadAttachmentUrl(index);
    mailHedgehog?.request()
      .get(url)
      .then(() => {
        const link = document.createElement('a');
        link.download = filename;
        let urlString = `${mailHedgehog?.request().defaults.baseURL}/${url}`;
        if(urlString.startsWith('//')) {
          urlString = `${window.location.protocol}${urlString}`;
        }
        const newUrl = new URL(urlString);
        newUrl.searchParams.set('token', mailHedgehog?.getAuthToken() as string);
        link.href = newUrl.toString();
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      })
      .catch((error) => {
        mailHedgehog.onResponseError(error, 'Response Error');
      })
      .finally(() => {
      });
  };

  const downloadEmailAttachment = (attachment: Attachment) => {
    if (attachment.data && attachment.data.length > 0) {
      downloadBlobFile(attachment.data, attachment.contentType, attachment.filename);
    } else {
      downloadAttachment(attachment.index, attachment.filename);
    }
  };

  const downloadEmail = () => {
    if (email.value) {
      downloadBlobFile(email.value.source, 'text/plain', `${email.value.id}.eml`);
    }
  };

  return {
    downloadEmailAttachment,
    downloadEmail,
  };
}
