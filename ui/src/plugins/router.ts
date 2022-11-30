import { createRouter, createWebHistory } from 'vue-router';
import Inbox from '../pages/Inbox.vue';
import Message from '../pages/Message.vue';
import Settings from '../pages/Settings.vue';

const routes = [
  { path: '/:pathMatch(.*)*', name: 'NotFound', redirect: '/' },
  { path: '/', component: Inbox, alias: '/emails' },
  { path: '/emails/:id', component: Message },
  { path: '/settings', component: Settings },
];

export function setupRouter() {
  return createRouter({
    history: createWebHistory(),
    routes,
  });
}
