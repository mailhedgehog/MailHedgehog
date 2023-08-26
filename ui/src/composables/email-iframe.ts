import { ref } from 'vue';

export function useEmailIframe() {
  const iframeHeight = ref('0rem');
  const resizeIframe = (event: Event) => {
    const obj = event.currentTarget as HTMLIFrameElement;
    const newHeight = (obj.contentWindow?.document.documentElement.scrollHeight as number) + 2;

    iframeHeight.value = `${newHeight}px`;
  };

  return {
    iframeHeight,
    resizeIframe,
  };
}
