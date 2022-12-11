<template>
  <div class="flex flex-shrink-0 justify-center items-center">
    <img
      class="h-8 w-auto mr-2"
      src="/MailHedgehog.svg"
      :alt="t('app.title')"
    >
    <div class="text-context-100 font-bold text-sm sm:text-base">
      {{ t('app.title') }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
</script>
