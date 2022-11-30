import Inbox from "../pages/Inbox.vue";
import Message from "../pages/Message.vue";
import Settings from "../pages/Settings.vue";
import {createRouter, createWebHistory} from "vue-router";
import {I18n} from "vue-i18n";

const routes = [
    {path: '/:pathMatch(.*)*', name: 'NotFound', redirect: '/'},
    {path: '/', component: Inbox, alias: '/emails'},
    {path: '/emails/:id', component: Message},
    {path: '/settings', component: Settings},
];

export function setupRouter(i18n: I18n) {
    return createRouter({
        history: createWebHistory(),
        routes,
    })
}

