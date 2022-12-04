import 'floating-vue/dist/style.css';
import './assets/scss/app.scss';
// @ts-ignore
import Toasted from 'toastedjs';
import FloatingVue from 'floating-vue';
import { createApp } from 'vue';
import { AxiosInstance, AxiosResponse, AxiosRequestConfig } from 'axios';
import enMessages from './assets/locales/en';
import uaMessages from './assets/locales/ua';
import frMessages from './assets/locales/fr';
import { setupAxios } from './plugins/axios';
import { setupRouter } from './plugins/router';
import { setupStore } from './plugins/store';
import { setupI18n, setI18nLanguage } from './plugins/i18n';
import App from './App.vue';

const i18n = setupI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  warnHtmlMessage: false,
  messages: {
    en: enMessages,
    ua: uaMessages,
    fr: frMessages,
  },
});

const router = setupRouter();
const store = setupStore();

const app = createApp(App);

app.use(router);
app.use(store);
app.use(i18n);
app.use(FloatingVue);

app.provide('SetLocale', (locale: string) => setI18nLanguage(i18n, locale));

app.mount('#app');

class MailHedgehog {
  $toasted: any;

  constructor() {
    this.$toasted = new Toasted({
      theme: 'app',
      position: 'bottom-right',
      duration: 6000,
    });
  }

  // eslint-disable-next-line class-methods-use-this
  request(options?: AxiosRequestConfig): AxiosInstance | Promise<AxiosResponse<any, any>> {
    const axios = setupAxios();

    if (options !== undefined) {
      return axios(options);
    }

    return axios;
  }

  info(message: string) {
    this.$toasted.show(message, { type: 'info' });
  }

  error(message: string) {
    this.$toasted.show(message, { type: 'error' });
  }

  success(message: string) {
    this.$toasted.show(message, { type: 'success' });
  }

  warning(message: string) {
    this.$toasted.show(message, { type: 'warning' });
  }
}

declare global {
  interface Window {
    MailHedgehog: MailHedgehog;
  }
}

window.MailHedgehog = new MailHedgehog();
