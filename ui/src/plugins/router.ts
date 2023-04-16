import { createRouter, createWebHashHistory } from 'vue-router';
import {MailHedgehog} from "@/main";

const Dashboard = () => import(/* webpackChunkName: "page-email" */ '../layouts/Dashboard.vue');
const Inbox = () => import(/* webpackChunkName: "page-email" */ '../pages/Inbox.vue');
const Email = () => import(/* webpackChunkName: "page-email" */ '../pages/Email.vue');
const Login = () => import(/* webpackChunkName: "page-login" */ '../pages/Login.vue');



function makeRoutes(app: MailHedgehog) {
  const routes: any = [
    { path: '/:pathMatch(.*)*', redirect: '/', name: 'notFound' },
  ];

  if(app.configValue('http.auth') === 'internal') {
    routes.push( { path: '/login', component: Login, name: 'login' });
  }

  routes.push({
    path: '/',
    component: Dashboard,
    // notice how only the child route has a name
    children: [
      {
        "path": '', "component": Inbox, "alias": 'emails', "name": 'emails',
      },
      { path: 'emails/:id', component: Email, name: 'email' },
    ],
  },)

  return routes;
}

// eslint-disable-next-line import/prefer-default-export
export function setupRouter(app: MailHedgehog) {
  return createRouter({
    history: createWebHashHistory(),
    routes: makeRoutes(app),
  });
}
