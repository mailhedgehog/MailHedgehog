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

<script setup>
import {
  SignalIcon,
  SignalSlashIcon,
} from '@heroicons/vue/24/outline';
import { onMounted, ref, inject } from 'vue';

const emitter = inject('emitter');
const disabled = ref(false);
const soc = ref(null);

const closeSocket = () => {
  disabled.value = true;
  if (soc.value) {
    soc.value.close();
  }
  soc.value = null;
  disabled.value = false;
};

const handleSystemMessage = (msg) => {
  switch (msg.type) {
    case 'new_message':
      emitter.emit('new_message', msg.data);
      window.MailHedgehog.info('Message received');
      break;
    default:
      window.MailHedgehog.error(`Online system flow type ${msg.type} not supported`);
  }
};

const initSocket = () => {
  if (soc.value) {
    return;
  }
  disabled.value = true;
  const url = window.MailHedgehog.congValue('http.baseUrl', '')
    .trim()
    .replace(/\/+$/, '')
    .replace(/(http)(s)?\:\/\//, 'ws$2://');
  soc.value = new WebSocket(`${url}/websocket`);

  soc.value.onopen = () => {
    // nothing to do
  };
  soc.value.onclose = () => {
    closeSocket();
  };
  soc.value.onerror = () => {
    window.MailHedgehog.error('Online connection error');
  };

  soc.value.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data);

      switch (msg.flow) {
        case 'system':
          handleSystemMessage(msg);
          break;
        default:
          window.MailHedgehog.error(`Online message flow ${msg.flow} not supported`);
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
