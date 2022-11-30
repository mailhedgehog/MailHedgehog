import './assets/scss/app.scss';
// @ts-ignore
import enMessages from './assets/locales/en.js'
import uaMessages from './assets/locales/ua.js'
import frMessages from './assets/locales/fr.js'
import {createApp, provide} from 'vue';
import {setupRouter} from './plugins/router'
import {setupI18n, setI18nLanguage} from './plugins/i18n'
import App from './App.vue';


const i18n = setupI18n({
    legacy: false,
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
        en: enMessages,
        ua: uaMessages,
        fr: frMessages,
    }
})

const router = setupRouter(i18n);

const app = createApp(App);

app.use(router);
app.use(i18n)
app.provide('SetLocale', (locale: string) => setI18nLanguage(i18n, locale))

app.mount('#app');
