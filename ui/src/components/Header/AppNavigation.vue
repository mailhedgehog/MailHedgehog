<template>
  <nav
    class="mt-5 flex-shrink-0 divide-y divide-primary-800 overflow-y-auto"
    aria-label="Sidebar"
  >
    <div class="space-y-1 px-2">
      <router-link
        v-for="item in navigation"
        v-slot="{ href, navigate, isActive }"
        :key="item.name"
        custom
        :to="item.href"
      >
        <a
          :href="href"
          :class="[isActive ? 'bg-primary-800 text-context-50' : 'text-primary-100 hover:text-context-50 hover:bg-primary-600', 'group flex items-center px-2 py-2 text-sm leading-6 font-medium rounded-md']"
          :aria-current="isActive ? 'page' : undefined"
          @click="navigate"
        >
          <component
            :is="item.icon"
            class="mr-4 h-6 w-6 flex-shrink-0 text-primary-200"
            aria-hidden="true"
          />
          {{ t(`menu.${item.name}`) }}
        </a>
      </router-link>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { InboxArrowDownIcon } from '@heroicons/vue/24/outline';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const navigation = [
  { name: 'inbox', href: '/', icon: InboxArrowDownIcon },
];

</script>
