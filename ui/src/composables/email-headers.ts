import { computed, Ref, ref } from 'vue';
import { Email } from '@/interfaces/Email';

export function useShowHeader(email: Ref<Email|null>) {
  const isAllHeaders = ref(false);

  const importantHeaders = [
    'Subject',
    'From',
    'To',
    'Cc',
  ];

  const getHeaders = (email: Email) => {
    const headersList: {[key: string]: Array<string>} = {};

    if (email && email.headers) {
      for (let i = 0; i < Object.entries(email.headers).length; i += 1) {
        const [headerKey, headerValues] = Object.entries(email.headers)[i];
        if (Array.isArray(headerValues) && headerValues.length > 0) {
          if (isAllHeaders.value) {
            headersList[headerKey] = headerValues;
          } else if (importantHeaders.includes(headerKey)) {
            headersList[headerKey] = headerValues;
          }
        }
      }
    }

    return headersList;
  };

  const headers = computed(() => {
    if (email.value) {
      return getHeaders(email.value);
    }

    return null;
  });

  return {
    isAllHeaders,
    importantHeaders,
    headers,
  };
}
