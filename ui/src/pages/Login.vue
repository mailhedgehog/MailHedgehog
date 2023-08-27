<script setup lang="ts">
import { LockClosedIcon } from '@heroicons/vue/20/solid';
import { useI18n } from 'vue-i18n';
import { computed, inject, ref } from 'vue';
import ColorModeSelect from '@/components/Header/ColorModeSelect.vue';
import { MailHedgehog } from '@/main.js';
import LangModeSelect from '@/components/Header/LangModeSelect.vue';
import Copyright from '@/components/Footer/Copyright.vue';

const { t } = useI18n();
const mailHedgehog = inject<MailHedgehog>('MailHedgehog');

const title = computed(() => mailHedgehog?.configValue('ui.appName', null));
const logoUrl = computed(() => mailHedgehog?.configValue('ui.appLogoUrl', '/MailHedgehog.svg'));
const logoStyle = computed(() => mailHedgehog?.configValue('ui.appLogoStyle', []));

const isRequesting = ref(false);
const loginForm = ref({
  username: null,
  password: null,
});
const loginAction = () => {
  isRequesting.value = true;
  mailHedgehog?.request()
    .post('login', loginForm.value)
    .then((response) => {
      if (response.data?.data?.token) {
        mailHedgehog?.setAuthToken(response.data?.data?.token);
        mailHedgehog.goTo({ name: 'emails' });
      }
    })
    .catch((error) => {
      mailHedgehog.onResponseError(error, 'Response Error');
    })
    .finally(() => {
      isRequesting.value = false;
    });
};
</script>

<template>
  <div class="flex min-h-screen items-center justify-center px-4 py-12 sm:px-6 lg:px-8 bg-context-100 dark:bg-context-900">
    <div class="w-full max-w-md space-y-8 py-8">
      <div>
        <div class="flex justify-center items-center">
          <img
            class="h-8 w-auto mr-2"
            :src="logoUrl"
            :style="logoStyle"
            :alt="title || t('app.title')"
          >
        </div>
        <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-context-800 dark:text-context-50">
          {{ t('login.title') }}
        </h2>
      </div>
      <form
        class="mt-8 space-y-6"
        @submit.prevent="loginAction"
      >
        <div class="-space-y-px rounded-md shadow-sm">
          <div>
            <label
              for="email-address"
              class="sr-only"
            >{{ t('login.form.username.label') }}</label>
            <input
              id="username"
              v-model="loginForm.username"
              name="username"
              type="text"
              autocomplete="username"
              required
              class="form-input rounded-b-none focus:z-10 relative"
              :placeholder="t('login.form.username.label')"
            >
          </div>
          <div>
            <label
              for="password"
              class="sr-only"
            >{{ t('login.form.password.label') }}</label>
            <input
              id="password"
              v-model="loginForm.password"
              name="password"
              type="password"
              autocomplete="current-password"
              required
              class="form-input rounded-t-none focus:z-10 relative"
              :placeholder="t('login.form.password.label')"
            >
          </div>
        </div>
        <div>
          <button
            type="submit"
            class="group relative flex w-full justify-center rounded-md
                  bg-primary-600 px-3 py-2 text-sm font-semibold text-white
                  hover:bg-primary-500 focus-visible:outline
                  focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
          >
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <LockClosedIcon
                class="h-5 w-5 text-primary-500 group-hover:text-primary-400"
                aria-hidden="true"
              />
            </span>
            {{ t('login.form.login.label') }}
          </button>
        </div>
      </form>
    </div>
    <div class="fixed bottom-2 w-full flex justify-center">
      <Copyright class="text-context-800 dark:text-context-50" />
    </div>
    <div class="fixed top-2 right-2 flex items-center space-x-3">
      <ColorModeSelect />
      <LangModeSelect />
    </div>
  </div>
</template>
