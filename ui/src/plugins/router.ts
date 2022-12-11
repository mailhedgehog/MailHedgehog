import { createRouter, createWebHashHistory } from 'vue-router';

const Inbox = () => import(/* webpackChunkName: "page-email" */ '../pages/Inbox.vue');
const Email = () => import(/* webpackChunkName: "page-email" */ '../pages/Email.vue');

const routes = [
  { path: '/:pathMatch(.*)*', redirect: '/', name: 'notFound' },
  {
    path: '/', component: Inbox, alias: '/emails', name: 'emails',
  },
  { path: '/emails/:id', component: Email, name: 'email' },
];

// eslint-disable-next-line import/prefer-default-export
export function setupRouter() {
  return createRouter({
    history: createWebHashHistory(),
    routes,
  });
}
