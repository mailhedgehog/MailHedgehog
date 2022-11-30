import { createI18n, I18n, I18nOptions } from 'vue-i18n';

export function setI18nLanguage(i18n: I18n, locale: string) {
  // @ts-ignore
  i18n.global.locale.value = locale;

  localStorage.setItem('lang', locale);
  document.querySelector('html')?.setAttribute('lang', locale);
}

export function setupI18n(options: I18nOptions = {}) {
  const i18n = createI18n(options);
  setI18nLanguage(i18n, localStorage.getItem('lang') || 'en');
  return i18n;
}
