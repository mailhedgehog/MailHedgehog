import axios, { AxiosInstance } from 'axios';
import { MailHedgehog } from '@/main';

// eslint-disable-next-line import/prefer-default-export
export function setupAxios(app: MailHedgehog, conf: { baseUrl: string }): AxiosInstance {
  const url = conf.baseUrl.trim().replace(/\/+$/, '');
  const instance = axios.create({
    baseURL: `${url}/api/v1`,
  });

  instance.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';

  const mailHedgehogToken = sessionStorage.getItem('mailHedgehogToken');
  if (mailHedgehogToken) {
    instance.defaults.headers.common.Authorization = `Bearer ${mailHedgehogToken}`;
  }

  instance.interceptors.response.use(
    (response) => response,
    (error) => {
      if (axios.isCancel(error)) {
        return Promise.reject(error);
      }

      const { response } = error;
      const {
        status,
      } = response;

      if (status === 401) {
        switch (app.configValue('http.auth')) {
          case 'basic':
            app.error('Unauthorized, reloading page');
            location.reload();
            break;
          case 'internal':
            app.error('Unauthorized');
            app.goTo({ name: 'login' });
            break;
        }
      }

      return Promise.reject(error);
    },
  );

  return instance;
}
