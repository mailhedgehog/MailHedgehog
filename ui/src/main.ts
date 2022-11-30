import './assets/scss/app.scss';
import Toasted from 'toastedjs';
import VTooltip from 'v-tooltip';
import { createApp } from 'vue';
import { AxiosRequestConfig } from 'axios';
import enMessages from '@/assets/locales/en.js';
import uaMessages from '@/assets/locales/ua.js';
import frMessages from '@/assets/locales/fr.js';
import { setupAxios } from '@/plugins/axios';
import { setupRouter } from '@/plugins/router';
import { setupStore } from '@/plugins/store';
import { setupI18n, setI18nLanguage } from './plugins/i18n';
import App from './App.vue';

const i18n = setupI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  messages: {
    en: enMessages,
    ua: uaMessages,
    fr: frMessages,
  },
});

const router = setupRouter();
const store = setupStore();

const app = createApp(App);

app.use(VTooltip);
app.use(router);
app.use(store);
app.use(i18n);
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
  request(options: AxiosRequestConfig) {
    const axios = setupAxios();

    if (options !== undefined) {
      return axios(options);
    }

    return axios;
  }

  info(message: string) {
    this.$toasted.show(message, { type: 'info' })
  }

  error(message: string) {
    this.$toasted.show(message, { type: 'error' })
  }

  success(message: string) {
    this.$toasted.show(message, { type: 'success' })
  }

  warning(message: string) {
    this.$toasted.show(message, { type: 'warning' })
  }
}

// @ts-ignore
window.MailHedgehog = new MailHedgehog();
