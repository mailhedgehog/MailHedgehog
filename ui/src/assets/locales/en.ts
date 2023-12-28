export default {
  app: {
    title: 'MailHedgehog',
  },
  menu: {
    inbox: 'Inbox',
    logout: 'Logout',
    users: 'Users',
  },
  sidebar: {
    open: 'Open sidebar',
    close: 'Close sidebar',
  },
  response: {
    error: 'Server response error',
  },
  colorMode: {
    title: 'Color mode',
    light: 'Light',
    dark: 'Dark',
    system: 'System',
  },
  locales: {
    title: 'Language',
    en: 'English',
    uk: 'Ukrainian',
    fr: 'French',
  },
  clipboard: {
    success: 'Copied to clipboard',
    error: 'Error to copy to clipboard',
  },
  pagination: {
    requesting: 'Loading...',
    prev: 'Previous',
    next: 'Next',
    text: 'Showing <span class="font-medium">{from}</span> to <span class="font-medium">{to}</span> of <span class="font-medium">{of}</span> results',
    actions: 'Actions',
  },
  email: {
    notAvailable: 'n/a',
    headersTitle: 'Email headers',
    headersSubtitle: 'An Email Header is metadata that accompanies every email and contains detailed information',
    hintHideHeaders: 'Show only important headers',
    hintShowHeaders: 'Show all headers',
    from: 'From',
    to: 'To',
    subject: 'Subject',
    received_at: 'Received at',
    delete: 'Delete',
    deleteHint: 'Delete email',
    deleted: 'Email deleted',
    download: 'Download',
    downloadHint: 'Download email as file',
    share: 'Share',
    shareHint: 'Allow access email by share link',
    release: 'Release',
    releaseHint: 'Release email',
    released: 'Released',
    linkCreated: 'Link created',
    cancel: 'Cancel',
    back: 'Back',
    tab: {
      html: 'HTML',
      plain: 'Plain text',
      source: 'Source',
      attachments: 'Attachments',
    },
    htmlEmpty: 'Email not contains HTML content',
    plainEmpty: 'Email not contains plain text content',
    noAttachments: 'Email has not attachments',
  },
  sharedEmail: {
    notFound: 'Shared email not found or share link expired.',
    modalTitle: 'Share',
    form: {
      expiration: {
        label: 'Expiration in',
        '1h': '1 hour',
        '3h': '3 hours',
        '1d': '1 day',
        '7d': '7 days',
        '30d': '30 days',
      },
      shareLink: {
        label: 'Link',
      },
      submitButton: {
        label: 'Share',
      },
    },
  },
  inbox: {
    hello: 'Hello, "{msg}"',
    pageTitle: 'Inbox',
    clear: 'Clear inbox',
    cleared: 'Inbox cleared',
    search: 'Search emails',
    empty: 'No emails found',
  },
  release: {
    host: 'Host',
    port: 'Port',
    auth: 'Credentials',
    username: 'Username',
    password: 'Password',
    saveToStorage: 'Save credentials to local storage (insecure)',
    saved: 'Saved',
    deleteSavedToStorage: 'Delete saved credentials',
    deleted: 'Deleted',
  },
  confirmDialog: {
    modalTitle: 'Are you sure?',
    modalText: 'Do you want to execute this action?',
    btnCancel: 'No',
    btnAccept: 'Yes',
  },
  login: {
    title: 'Sign in to your account',
    form: {
      username: {
        label: 'Username',
      },
      password: {
        label: 'Password',
      },
      login: {
        label: 'Sign in',
      },
    },
  },
  users: {
    pageTitle: 'Users',
    search: 'Search users',
    empty: 'No users found',
    create: 'Add user',
    deleted: 'User deleted',
    updated: 'User updated',
    created: 'User created',
    username: 'Username',
    hubPassword: 'UI password',
    smtpPassword: 'SMTP password',
    emptySmtpPasswordHint: 'Left blank to set same password as UI',
    emptyPasswordHint: 'Left blank to keep old password',
    smtpNoPassIps: 'SMTP client IPs to skip password auth',
    ipPlaceholder: '127.0.0.1',
    addIPHint: 'Add IP',
    modal: {
      createTitle: 'Create user',
      editTitle: 'Edit user \'{user}\'',
      cancel: 'Cancel',
      submit: 'Submit',
    },
  },
};
