<script setup lang="ts">
import {
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
} from '@headlessui/vue';
import {
  UserCircleIcon,
  ChevronDownIcon,
} from '@heroicons/vue/24/outline';
import { inject } from 'vue';
import { useI18n } from 'vue-i18n';
import { MailHedgehog } from '@/main';

const mailHedgehog = inject<MailHedgehog>('MailHedgehog');
const { t } = useI18n();

const logout = () => {
  mailHedgehog?.request()
    .post('logout')
    .finally(() => {
      mailHedgehog?.setAuthToken();
      const reloadNow = () => {
        let url = mailHedgehog?.configValue('http.baseUrl', '')
          .trim()
          .replace(/\/+$/, '');
        url = `${url}/`;
        window.open(url, '_self');
      };

      const xhr = new XMLHttpRequest();
      xhr.open('GET', import.meta.env.VITE_MH_CONFIG_URL, false, 'logout', 'logout');
      xhr.onload = () => {
        if (xhr.readyState === 4) {
          reloadNow();
        }
      };
      xhr.onerror = () => reloadNow();
      xhr.send(null);
    });
};
</script>

<template>
  <Menu
    as="div"
    class="icon-select"
  >
    <div>
      <MenuButton
        class="icon-select__menu-btn flex max-w-xs items-center rounded-full text-sm focus:outline-none lg:rounded-md lg:p-2"
      >
        <UserCircleIcon
          class="w-8 h-8"
        />
        <ChevronDownIcon
          class="ml-1 h-5 w-5 flex-shrink-0 text-context-400 dark:text-context-600"
          aria-hidden="true"
        />
      </MenuButton>
    </div>
    <transition
      enter-active-class="transition ease-out duration-100"
      enter-from-class="transform opacity-0 scale-95"
      enter-to-class="transform opacity-100 scale-100"
      leave-active-class="transition ease-in duration-75"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <MenuItems
        class="icon-select__menu w-48"
      >
        <MenuItem v-slot="{ active }">
          <a
            :class="[active ? 'icon-select__menu-item--active' : '']"
            class="icon-select__menu-item"
            @click="logout"
          >
            {{ t('menu.logout') }}
          </a>
        </MenuItem>
      </MenuItems>
    </transition>
  </Menu>
</template>
