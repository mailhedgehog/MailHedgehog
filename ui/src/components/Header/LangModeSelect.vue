<script setup lang="ts">
import {
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
} from '@headlessui/vue';
import { inject } from 'vue';
import { useI18n } from 'vue-i18n';

const ALLOWED_LOCALES = ['en', 'fr', 'uk'];
const SetLocale = inject<(locale: string) => void>('SetLocale');

const { t } = useI18n();
</script>

<template>
  <Menu
    as="div"
    class="icon-select ml-3"
  >
    <div>
      <MenuButton
        class="icon-select__menu-btn"
        :title="t('locales.title')"
      >
        <span class="sr-only">{{ t('locales.title') }}</span>
        <div class="w-6 h-6 flex items-center justify-center font-medium uppercase">
          {{ $i18n.locale }}
        </div>
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
        <MenuItem
          v-for="locale in ALLOWED_LOCALES"
          :key="locale"
          v-slot="{ active }"
        >
          <a
            :class="[active ? 'icon-select__menu-item--active' : '']"
            class="icon-select__menu-item"
            @click="(!!SetLocale)?SetLocale(locale):null"
          >
            <div class="font-medium uppercase  mr-2">
              {{ locale }}
            </div>
            <span class="">{{ t(`locales.${locale}`) }}</span>
          </a>
        </MenuItem>
      </MenuItems>
    </transition>
  </Menu>
</template>
