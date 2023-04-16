<template>
  <div class="text-sm">
        <span class="mr-0">
          ©
        </span>
    <span
      v-if="copyrightsText"
      v-html="copyrightsText"
    />
    <a
      v-else
      href="https://github.com/mailhedgehog"
      class="underline transition-colors duration-500 hover:text-primary-600"
      target="_blank"
    >
      MailHedgehog
    </a>
  </div>
</template>

<script lang="ts" setup>
import {computed, inject} from "vue";
import {MailHedgehog} from "@/main.js";

const mailHedgehog = inject<MailHedgehog>('MailHedgehog');

const copyrightsText = computed(() => mailHedgehog?.configValue('ui.copyrightsText', null));
</script>
