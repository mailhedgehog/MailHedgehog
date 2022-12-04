<template>
  {{email}}
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';
import {
  Dialog,
  DialogPanel,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue';
import {
  BellIcon,
} from '@heroicons/vue/24/outline';

const router = useRouter();
const route = useRoute();

const isRequesting = ref(false);
const email = ref(null);

const getEmail = (emailId) => {
  isRequesting.value = true;
  window.MailHedgehog.request()
    .get(`emails/${emailId}`)
    .then((response) => {
      if (response.data?.data) {
        email.value = response.data?.data;
      } else {
        emails.value = null;
      }
    })
    .catch(() => {
      router.push({ path: '/404' });
    })
    .finally(() => {
      isRequesting.value = false;
    });
};

onMounted(() => {
  getEmail(route.params.id);
});

</script>
