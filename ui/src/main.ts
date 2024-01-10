import 'floating-vue/dist/style.css';
import './assets/scss/app.scss';
import Toast, { POSITION, useToast } from 'vue-toastification';
import _, { isObject } from 'lodash';
import FloatingVue from 'floating-vue';
import { createApp, App } from 'vue';
import mitt, { Handler, WildcardHandler } from 'mitt';
import { Axios, AxiosError, AxiosResponse } from 'axios';
import { Router } from 'vue-router';
import copy from 'copy-to-clipboard';
import enMessages from './assets/locales/en';
import ukMessages from './assets/locales/uk';
import frMessages from './assets/locales/fr';
import { setupAxios } from './plugins/axios';
import { setupRouter } from './plugins/router';
import { setupStore } from './plugins/store';
import { fetchConfigs } from './plugins/fetchConfigs';
import { setI18nLanguage, setupI18n } from './plugins/i18n';
import AppView from './App.vue';

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

const store = setupStore();
const emitter = mitt();

// eslint-disable-next-line import/prefer-default-export
export class MailHedgehog {
  mhConf: object;

  $toast: any;

  $t: any;

  $axios: Axios;

  app: App | null = null;

  router: Router | null = null;

  constructor(mhConf = {}) {
    this.mhConf = mhConf;

    this.$toast = useToast();
    this.$axios = setupAxios(this, {
      baseUrl: this.configValue('http.baseUrl', ''),
    });
  }

  configValue(key: string, defaultValue: any = null): any {
    return _.get(this.mhConf, key, defaultValue);
  }

  request(): Axios {
    return this.$axios;
  }

  setAuthToken(token?: string) {
    if (token) {
      sessionStorage.setItem('mailHedgehogToken', token);
      this.$axios.defaults.headers.common.Authorization = `Bearer ${token}`;
    } else {
      sessionStorage.removeItem('mailHedgehogToken');
      delete this.$axios.defaults.headers.common.Authorization;
    }

    return this;
  }

  getAuthToken(token?: string) {
    const header = this.$axios.defaults.headers.common.Authorization as string;

    return header?.split('Bearer ').pop();
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

  // eslint-disable-next-line class-methods-use-this
  $on(type: any, handler: WildcardHandler): void {
    emitter.on(type, handler);
  }

  // eslint-disable-next-line class-methods-use-this
  $off(type: any, handler: WildcardHandler): void {
    emitter.off(type, handler);
  }

  // eslint-disable-next-line class-methods-use-this
  $emit(type: string, data:any = null) {
    emitter.emit(type, data);
  }

  goTo(to: any) {
    this.router?.push(to);
  }

  copyToClipboard(text: string = '', options = {}) {
    if (copy(text, options)) {
      this.success(this.$t('clipboard.success'));
      return;
    }
    this.error(this.$t('clipboard.error'));
  }

  onResponseError(error: AxiosError, defaultError: string|null = null) {
    if (error.response && error.response.data && isObject(error.response.data)) {
      // @ts-ignore
      if (error.response.data?.messages && Array.isArray(error.response.data?.messages) && error.response.data?.messages.length > 0) {
        // @ts-ignore
        const errorMessage = error.response.data?.messages[0];
        this.error(`Error: ${errorMessage.tag}, ${errorMessage.failed_field}`);
        return;
      }

      if (defaultError) {
        this.error(defaultError);
      }
    }

    console.error(error);
    this.error(defaultError || 'Response returned error.');
  }

  init() {
    this.app = createApp(AppView);
    this.router = setupRouter(this);

    this.app.use(this.router);
    this.app.use(store);
    this.app.use(i18n);
    this.app.use(FloatingVue);
    this.app.use(Toast, {
      position: POSITION.BOTTOM_RIGHT,
    });

    this.$t = i18n.global.t;

    this.app.provide('SetLocale', (locale: string) => setI18nLanguage(i18n, locale));
    this.app.provide('emitter', emitter);

    this.router.beforeEach((to) => {
      if (
        to.name === 'login'
        || to.name === 'sharedMail'
      ) {
        return;
      }
      this.request()
        .get('user')
        .then((response: AxiosResponse) => {
          store.dispatch('setUser', response.data.data || null);
        })
        .catch(() => store.dispatch('setUser', null));
    });

    this.app.provide('MailHedgehog', this);
    window.MailHedgehog = this;

    this.app.mount('#app');
  }
}

declare global {
  // eslint-disable-next-line no-unused-vars
  interface Window {
    MailHedgehog: MailHedgehog;
  }
}

const mailHedgehog = new MailHedgehog(fetchConfigs());
mailHedgehog.init();
