import 'floating-vue/dist/style.css';
import './assets/scss/app.scss';
import Toast, { POSITION, useToast } from 'vue-toastification';
import _ from 'lodash';
import FloatingVue from 'floating-vue';
import { createApp, App } from 'vue';
import mitt, {Handler, WildcardHandler} from 'mitt';
import { Axios, AxiosResponse } from 'axios';
import enMessages from './assets/locales/en';
import ukMessages from './assets/locales/uk';
import frMessages from './assets/locales/fr';
import { setupAxios } from './plugins/axios';
import { setupRouter } from './plugins/router';
import { setupStore } from './plugins/store';
import { fetchConfigs } from './plugins/fetchConfigs';
import { setI18nLanguage, setupI18n } from './plugins/i18n';
import AppView from './App.vue';
import {Router} from "vue-router";

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
  $emit(type: string) {
    emitter.emit(type);
  }

  goTo(to: any) {
    this.router?.push(to)
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

    this.app.provide('SetLocale', (locale: string) => setI18nLanguage(i18n, locale));
    this.app.provide('emitter', emitter);


    this.router.beforeEach((to) => {
      if(to.name === 'login') {
        return
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

