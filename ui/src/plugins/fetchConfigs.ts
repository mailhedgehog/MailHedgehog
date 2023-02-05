// eslint-disable-next-line import/prefer-default-export
export function fetchConfigs() {
  let conf = {};
  const xhr = new XMLHttpRequest();
  xhr.open('GET', import.meta.env.VITE_MH_CONFIG_URL, false);
  xhr.onload = () => {
    if (xhr.readyState === 4) {
      if (xhr.status === 200) {
        conf = JSON.parse(xhr.responseText);
        return;
      }
      // eslint-disable-next-line no-console
      console.error(xhr.statusText);
    }
  };
  xhr.onerror = () => {
    // eslint-disable-next-line no-console
    console.error(xhr.statusText);
  };
  xhr.send(null);

  return conf;
}
