import { createRouter, createWebHistory } from 'vue-router';
import Inbox from '../pages/Inbox.vue';
import Email from '../pages/Email.vue';
import Settings from '../pages/Settings.vue';

const routes = [
  { path: '/:pathMatch(.*)*', name: 'notFound', redirect: '/' },
  { path: '/', component: Inbox, alias: '/emails' },
  { path: '/emails/:id', component: Email, name: 'email' },
  { path: '/settings', component: Settings },
];

export function setupRouter() {
  return createRouter({
    history: createWebHistory(),
    routes,
  });
}
