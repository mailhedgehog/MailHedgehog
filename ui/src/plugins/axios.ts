import axios from 'axios';

// eslint-disable-next-line import/prefer-default-export
export function setupAxios() {
  const instance = axios.create();

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
