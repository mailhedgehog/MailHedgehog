import axios, { AxiosInstance } from 'axios';

// eslint-disable-next-line import/prefer-default-export
export function setupAxios(): AxiosInstance {
  const instance = axios.create({
    baseURL: 'http://localhost:8025/api/v1',
  });

  instance.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';

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
        // eslint-disable-next-line no-restricted-globals
        location.reload();
      }

      return Promise.reject(error);
    },
  );

  return instance;
}
