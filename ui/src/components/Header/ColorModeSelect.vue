<script setup lang="ts">
import {
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
} from '@headlessui/vue';
import {
  SunIcon,
  MoonIcon,
  ComputerDesktopIcon,
} from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const mode = ref('');
const setMode = (newMode: string) => {
  mode.value = newMode;
  localStorage.setItem('ColorMode', mode.value);

  switch (mode.value) {
    case 'light':
      document.querySelector('html')?.classList.remove('dark');
      break;
    case 'dark':
      document.querySelector('html')?.classList.add('dark');
      break;
    default:
      if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        document.querySelector('html')?.classList.add('dark');
      } else {
        document.querySelector('html')?.classList.remove('dark');
      }
      break;
  }
};

onMounted(() => {
  const ALLOWED_COLOR_MODES = ['system', 'light', 'dark'];
  let storedMode = localStorage.getItem('ColorMode') || '';
  if (!ALLOWED_COLOR_MODES.includes(storedMode)) {
    // eslint-disable-next-line prefer-destructuring
    storedMode = ALLOWED_COLOR_MODES[0];
  }

  setMode(storedMode);
});

</script>

<template>
  <Menu
    as="div"
    class="icon-select ml-3"
  >
    <div>
      <MenuButton
        class="icon-select__menu-btn"
        :title="$t('colorMode.title')"
      >
        <span class="sr-only">{{ t('colorMode.title') }}</span>
        <SunIcon
          v-if="mode === 'light'"
          class="h-6 w-6"
          aria-hidden="true"
        />
        <MoonIcon
          v-else-if="mode === 'dark'"
          class="h-6 w-6"
          aria-hidden="true"
        />
        <ComputerDesktopIcon
          v-else
          class="h-6 w-6"
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
        class="icon-select__menu"
      >
        <MenuItem v-slot="{ active }">
          <a
            :class="[active ? 'icon-select__menu-item--active' : '']"
            class="icon-select__menu-item"
            @click="setMode('light')"
          >
            <SunIcon
              class="h-6 w-6 mr-2"
              aria-hidden="true"
            />
            <span class="">{{ t('colorMode.light') }}</span>
          </a>
        </MenuItem>
        <MenuItem v-slot="{ active }">
          <a
            :class="[active ? 'icon-select__menu-item--active' : '']"
            class="icon-select__menu-item"
            @click="setMode('dark')"
          >
            <MoonIcon
              class="h-6 w-6 mr-2"
              aria-hidden="true"
            />
            <span class="">{{ t('colorMode.dark') }}</span>
          </a>
        </MenuItem>
        <MenuItem v-slot="{ active }">
          <a
            :class="[active ? 'icon-select__menu-item--active' : '']"
            class="icon-select__menu-item"
            @click="setMode('system')"
          >
            <ComputerDesktopIcon
              class="h-6 w-6 mr-2"
              aria-hidden="true"
            />
            <span class="">{{ t('colorMode.system') }}</span>
          </a>
        </MenuItem>
      </MenuItems>
    </transition>
  </Menu>
</template>
