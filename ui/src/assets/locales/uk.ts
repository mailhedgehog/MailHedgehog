export default {
  app: {
    title: 'MailHedgehog',
  },
  menu: {
    inbox: 'Вхідні',
    logout: 'Вийти',
    users: 'Користувачі',
  },
  sidebar: {
    open: 'Відкрити бічну панель',
    close: 'Закрити бічну панель',
  },
  response: {
    error: 'Помилка відповіді сервера',
  },
  colorMode: {
    title: 'Тема',
    light: 'Світлий',
    dark: 'Темний',
    system: 'Системний',
  },
  locales: {
    title: 'Мова',
    en: 'Англійська',
    uk: 'Українська',
    fr: 'Французька',
  },
  pagination: {
    requesting: 'Завантаження...',
    prev: 'Попередня',
    next: 'Наступна',
    text: 'Показано <span class="font-medium">{from}</span> по <span class="font-medium">{to}</span> з <span class="font-medium">{of}</span> результатів',
    actions: 'Дії',
  },
  email: {
    notAvailable: 'н/д',
    headersTitle: 'Заголовки імейлу',
    headersSubtitle: 'Заголовок імейлу – це метадані, які супроводжують кожен електронний лист і містять детальну інформацію',
    hintHideHeaders: 'Показувати лише важливі заголовки',
    hintShowHeaders: 'Показати всі заголовки',
    from: 'Від',
    to: 'Отримувач',
    subject: 'Тема',
    received_at: 'Отримано о',
    delete: 'Видалити',
    deleteHint: 'Видалити імейл',
    deleted: 'імейл видалено',
    download: 'Завантажити',
    downloadHint: 'Завантажити імейлу файл',
    release: 'Переслати',
    releaseHint: 'Переслати імейл',
    released: 'Переслано',
    cancel: 'Скасувати',
    back: 'Назад',
    tab: {
      html: 'HTML',
      plain: 'Простий текст',
      source: 'Джерело',
      attachments: 'Вкладення',
    },
    htmlEmpty: 'Імейл не містить вмісту HTML',
    plainEmpty: 'Імейл не містить звичайний тексту',
    noAttachments: 'Імейл не має вкладень',
  },
  sharedEmail: {
    notFound: 'Імейл не знайдено або термін дії посилання для доступу минув.',
  },
  inbox: {
    hello: 'Привіт, "{msg}"',
    pageTitle: 'Вхідні',
    clear: 'Видалити всі вхідні',
    cleared: 'Вхідні видалені',
    search: 'Пошук імейлів',
    empty: 'Імейлів не знайдено',
  },
  release: {
    host: 'Хост',
    port: 'Порт',
    auth: 'Облікові дані',
    username: 'Ім\'я користувача',
    password: 'Пароль',
    saveToStorage: 'Зберегти облікові дані в локальному сховищі (небезпечно)',
    saved: 'Збережено',
    deleteSavedToStorage: 'Видалити збережені облікові дані',
    deleted: 'Видалено',
  },
  confirmDialog: {
    modalTitle: 'Ви впевнені?',
    modalText: 'Ви впевнені, що бажаєте виконати цю дію?',
    btnCancel: 'Ні',
    btnAccept: 'Так',
  },
  login: {
    title: 'Увійдіть у свій обліковий запис',
    form: {
      username: {
        label: 'Ім\'я користувача',
      },
      password: {
        label: 'Пароль',
      },
      login: {
        label: 'Увійти',
      },
    },
  },
  users: {
    pageTitle: 'Користувачі',
    search: 'Пошук користувачів',
    empty: 'Користувачів не знайдено',
    create: 'Додати користувача',
    deleted: 'Користувача видалено',
    updated: 'Користувач оновлений',
    created: 'Користувач створено',
    username: 'Ім\'я користувача',
    hubPassword: 'UI пароль',
    smtpPassword: 'SMTP пароль',
    emptySmtpPasswordHint: 'Залиште порожнім, щоб встановити той самий пароль, що й UI',
    emptyPasswordHint: 'Залишити порожнім, щоб зберегти старий пароль',
    modal: {
      createTitle: 'Створити користувача',
      editTitle: 'Редагувати користувача \'{user}\'',
      cancel: 'Скасувати',
      submit: 'Надіслати',
    },
  },
};
