import { createRouter, createWebHistory } from 'vue-router';
import Inbox from '../pages/Inbox.vue';
import Email from '../pages/Email.vue';
import Settings from '../pages/Settings.vue';

const routes = [
  { path: '/:pathMatch(.*)*', redirect: '/', name: 'notFound' },
  {
    path: '/', component: Inbox, alias: '/emails', name: 'emails',
  },
  { path: '/emails/:id', component: Email, name: 'email' },
  { path: '/settings', component: Settings },
];

// eslint-disable-next-line import/prefer-default-export
export function setupRouter() {
  return createRouter({
    history: createWebHistory(),
    routes,
  });
}
