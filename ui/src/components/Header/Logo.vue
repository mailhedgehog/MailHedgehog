<template>
  <div class="flex flex-shrink-0 justify-center items-center">
    <img
      class="h-8 w-auto mr-2"
      :src="logoUrl"
      :style="logoStyle"
      :alt="title || t('app.title')"
    >
    <div class="text-context-100 font-bold text-sm sm:text-base">
      {{ title || t('app.title') }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { computed, inject } from 'vue';
import { MailHedgehog } from '@/main';

const { t } = useI18n();
const mailHedgehog = inject<MailHedgehog>('MailHedgehog');

const title = computed(() => mailHedgehog?.configValue('ui.appName', null));
const logoUrl = computed(() => mailHedgehog?.configValue('ui.appLogoUrl', '/MailHedgehog.svg'));
const logoStyle = computed(() => mailHedgehog?.configValue('ui.appLogoStyle', []));

</script>
