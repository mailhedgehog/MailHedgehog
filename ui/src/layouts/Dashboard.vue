<script setup lang="ts">
import { computed, inject, ref } from 'vue';
import {
  Dialog,
  DialogPanel,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue';
import {
  Bars3CenterLeftIcon,
  XMarkIcon,
} from '@heroicons/vue/24/outline';
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n';
import AppNavigation from '@/components/Header/AppNavigation.vue';
import WebsocketConnector from '@/components/Header/WebsocketConnector.vue';
import ColorModeSelect from '@/components/Header/ColorModeSelect.vue';
import LangModeSelect from '@/components/Header/LangModeSelect.vue';
import ProfileDropdown from '@/components/Header/ProfileDropdown.vue';
import Copyright from '@/components/Footer/Copyright.vue';
import Logo from '@/components/Header/Logo.vue';
import ConfirmDialog from '@/components/Confirm/Dialog.vue';
import { User } from '@/plugins/store';
import FooterRightContent from '@/components/Footer/FooterRightContent.vue';
import { MailHedgehog } from '@/main';

const { t } = useI18n();
const mailHedgehog = inject<MailHedgehog>('MailHedgehog');
const store = useStore();

const user = computed<User | null>(() => store.getters.getUser);

const sidebarOpen = ref(false);
</script>

<template>
  <TransitionRoot
    as="template"
    :show="sidebarOpen"
  >
    <Dialog
      as="div"
      class="relative z-40 lg:hidden"
      @close="sidebarOpen = false"
    >
      <TransitionChild
        as="template"
        enter="transition-opacity ease-linear duration-300"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="transition-opacity ease-linear duration-300"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-context-600 bg-opacity-75" />
      </TransitionChild>

      <div class="fixed inset-0 z-40 flex">
        <TransitionChild
          as="template"
          enter="transition ease-in-out duration-300 transform"
          enter-from="-translate-x-full"
          enter-to="translate-x-0"
          leave="transition ease-in-out duration-300 transform"
          leave-from="translate-x-0"
          leave-to="-translate-x-full"
        >
          <DialogPanel class="relative flex w-full max-w-xs flex-1 flex-col bg-primary-700 pt-5 pb-4">
            <TransitionChild
              as="template"
              enter="ease-in-out duration-300"
              enter-from="opacity-0"
              enter-to="opacity-100"
              leave="ease-in-out duration-300"
              leave-from="opacity-100"
              leave-to="opacity-0"
            >
              <div class="absolute top-0 right-0 -mr-12 pt-2">
                <button
                  type="button"
                  class="ml-1 flex h-10 w-10 items-center justify-center rounded-full
                  focus:outline-none focus:ring-2 focus:ring-inset
                  focus:ring-context-50 focus:dark:ring-context-900"
                  @click="sidebarOpen = false"
                >
                  <span class="sr-only">
                    {{ t('sidebar.close') }}
                  </span>
                  <XMarkIcon
                    class="h-6 w-6 text-context-50"
                    aria-hidden="true"
                  />
                </button>
              </div>
            </TransitionChild>
            <Logo
              class="px-4"
              :reverse-color="false"
            />
            <AppNavigation />
          </DialogPanel>
        </TransitionChild>
        <div
          class="w-14 flex-shrink-0"
          aria-hidden="true"
        >
          <!-- Dummy element to force sidebar to shrink to fit close icon -->
        </div>
      </div>
    </Dialog>
  </TransitionRoot>

  <!-- Static sidebar for desktop -->
  <div class="hidden lg:fixed lg:inset-y-0 lg:flex lg:w-64 lg:flex-col">
    <!-- Sidebar component, swap this element with another sidebar if you like -->
    <div class="flex flex-grow flex-col overflow-y-auto bg-primary-700 pt-5 pb-4">
      <Logo
        class="px-4"
        :reverse-color="false"
      />
      <AppNavigation />
    </div>
  </div>

  <div class="min-h-screen flex flex-1 flex-col lg:pl-64 bg-context-100 dark:bg-context-900">
    <div
      class="flex h-16 flex-shrink-0 border-b border-context-200 dark:border-context-700
      bg-context-50 dark:bg-context-900 lg:border-none z-10 shadow dark:shadow-context-500"
    >
      <button
        type="button"
        class="border-r border-context-200 dark:border-context-700
        px-4 text-context-400
        focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500 lg:hidden"
        @click="sidebarOpen = true"
      >
        <span class="sr-only">{{ t('sidebar.open') }}</span>
        <Bars3CenterLeftIcon
          class="h-6 w-6"
          aria-hidden="true"
        />
      </button>
      <!-- Search bar -->
      <div class="flex flex-1 justify-between px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8">
        <div
          id="header-search"
          class="flex flex-1"
        />
        <div class="ml-4 flex items-center md:ml-6 space-x-3">
          <WebsocketConnector v-if="mailHedgehog?.configValue('http.websocket', false)" />
          <ColorModeSelect />
          <LangModeSelect />
          <ProfileDropdown v-if="user" />
        </div>
      </div>
    </div>
    <main class="flex-grow flex-1 pb-8 overflow-hidden">
      <router-view />
    </main>
    <div
      class="py-4 px-4 sm:px-6 flex justify-between items-center shadow dark:shadow-context-500 text-context-900 dark:text-context-100"
    >
      <Copyright />
      <FooterRightContent />
    </div>
  </div>
  <ConfirmDialog />
</template>
