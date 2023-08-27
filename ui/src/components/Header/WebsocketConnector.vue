<script setup lang="ts">
import {
  SignalIcon,
  SignalSlashIcon,
} from '@heroicons/vue/24/outline';
import { onMounted, ref, inject } from 'vue';
import { MailHedgehog } from '@/main';

const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
const disabled = ref(false);
const soc = ref<WebSocket|null>(null);

const closeSocket = () => {
  disabled.value = true;
  if (soc.value) {
    soc.value.close();
  }
  soc.value = null;
  disabled.value = false;
};

const handleSystemMessage = (msg: {
  flow? : string,
  type? : string,
  data? : any,
}) => {
  switch (msg.type) {
    case 'new_message':
      mailHedgehog?.$emit('new_message', msg.data);
      mailHedgehog?.info('Message received');
      break;
    default:
      mailHedgehog?.error(`Online system flow type ${msg.type} not supported`);
  }
};

const initSocket = () => {
  if (soc.value) {
    return;
  }
  disabled.value = true;
  let url = mailHedgehog?.configValue('http.baseUrl', '')
    .trim()
    .replace(/\/+$/, '')
    .replace(/(http)(s)?\:\/\//, 'ws$2://');
  if (url.startsWith('//')) {
    url = `ws:${url}`;
  }
  soc.value = new WebSocket(`${url}/websocket`);

  soc.value.onopen = () => {
    // nothing to do
  };
  soc.value.onclose = () => {
    closeSocket();
  };
  soc.value.onerror = () => {
    mailHedgehog?.error('Online connection error');
  };

  soc.value.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data);

      switch (msg.flow) {
        case 'system':
          handleSystemMessage(msg);
          break;
        default:
          mailHedgehog?.error(`Online message flow ${msg.flow} not supported`);
      }
    } catch (e) {
      console.error(e);
      console.error(event.data);
    }
  };

  disabled.value = false;
};

const toggleSocket = () => {
  if (soc.value) {
    closeSocket();
  } else {
    initSocket();
  }
};

onMounted(() => {
  initSocket();
});
</script>

<template>
  <button
    v-tooltip="soc?'Online':'Offline'"
    class="cursor-pointer text-context-400 dark:text-context-600 flex items-center justify-center font-medium uppercase"
    :disabled="disabled"
    @click.prevent="toggleSocket"
  >
    <SignalIcon
      v-if="soc"
      class="h-6 w-6"
      aria-hidden="true"
    />
    <SignalSlashIcon
      v-else
      class="h-6 w-6"
      aria-hidden="true"
    />
  </button>
</template>
