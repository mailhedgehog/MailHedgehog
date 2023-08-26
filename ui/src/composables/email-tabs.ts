import { computed, Ref, ref } from 'vue';

export function useEmailTabs() {
  const currentTab = ref('html');
  const tabs = [
    {
      id: 'html',
    },
    {
      id: 'plain',
    },
    {
      id: 'source',
    },
    {
      id: 'attachments',
    },
  ];

  return {
    currentTab,
    tabs,
  };
}
