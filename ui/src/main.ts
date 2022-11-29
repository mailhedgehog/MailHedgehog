import {createApp} from 'vue'
import './assets/scss/app.scss'
import {createRouter, createWebHistory} from 'vue-router'
import App from './App.vue'
import Inbox from './pages/Inbox.vue'
import Message from './pages/Message.vue'
import Settings from './pages/Settings.vue'

const routes = [
    { path: '/:pathMatch(.*)*', name: 'NotFound', redirect: '/' },
    {path: '/', component: Inbox, alias: '/emails'},
    {path: '/emails/:id', component: Message},
    {path: '/settings', component: Settings},
]
const router = createRouter({
    history: createWebHistory(),
    routes,
})

const app = createApp(App);

app.use(router)

app.mount('#app');
