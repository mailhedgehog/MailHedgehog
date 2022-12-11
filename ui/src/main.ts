import 'floating-vue/dist/style.css';
import './assets/scss/app.scss';
import Toast, { PluginOptions, POSITION, useToast } from 'vue-toastification';
import FloatingVue from 'floating-vue';
import { createApp } from 'vue';
import mitt from 'mitt';
import { Axios, AxiosResponse } from 'axios';
import enMessages from './assets/locales/en';
import ukMessages from './assets/locales/uk';
import frMessages from './assets/locales/fr';
import { setupAxios } from './plugins/axios';
import { setupRouter } from './plugins/router';
import { setupStore } from './plugins/store';
import { setI18nLanguage, setupI18n } from './plugins/i18n';
import App from './App.vue';

const i18n = setupI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  warnHtmlMessage: false,
  messages: {
    en: enMessages,
    uk: ukMessages,
    fr: frMessages,
  },
});

const router = setupRouter();
const store = setupStore();
const emitter = mitt();

const app = createApp(App);

app.use(router);
app.use(store);
app.use(i18n);
app.use(FloatingVue);

const options: PluginOptions = {
  position: POSITION.BOTTOM_RIGHT,
};

app.use(Toast, options);

app.provide('SetLocale', (locale: string) => setI18nLanguage(i18n, locale));
app.provide('emitter', emitter);

class MailHedgehog {
  $toast: any;

  $axios: Axios;

  constructor() {
    this.$toast = useToast();
    this.$axios = setupAxios();
  }

  request(): Axios {
    return this.$axios;
  }

  info(message: string) {
    this.$toast.info(message);
  }

  error(message: string) {
    this.$toast.error(message);
  }

  success(message: string) {
    this.$toast.success(message);
  }

  warning(message: string) {
    this.$toast.warning(message);
  }
}

declare global {
  // eslint-disable-next-line no-unused-vars
  interface Window {
    MailHedgehog: MailHedgehog;
  }
}

window.MailHedgehog = new MailHedgehog();

router.beforeEach(() => {
  window.MailHedgehog.request()
    .get('user')
    .then((response: AxiosResponse) => {
      store.dispatch('setUser', response.data.data || null);
    })
    .catch(() => store.dispatch('setUser', null));
});

app.mount('#app');
