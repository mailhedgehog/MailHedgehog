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
            <div class="flex flex-shrink-0 items-center px-4">
              <img
                class="h-8 w-auto"
                src="https://tailwindui.com/img/logos/mark.svg?color=cyan&shade=300"
                alt="Easywire logo"
              >
            </div>
            <nav
              class="mt-5 h-full flex-shrink-0 divide-y divide-primary-800 overflow-y-auto"
              aria-label="Sidebar"
            >
              <div class="space-y-1 px-2">
                <a
                  v-for="item in navigation"
                  :key="item.name"
                  :href="item.href"
                  :class="[item.current ? 'bg-primary-800 text-context-50' : 'text-primary-100 hover:text-context-50 hover:bg-primary-600', 'group flex items-center px-2 py-2 text-base font-medium rounded-md']"
                  :aria-current="item.current ? 'page' : undefined"
                >
                  <component
                    :is="item.icon"
                    class="mr-4 h-6 w-6 flex-shrink-0 text-primary-200"
                    aria-hidden="true"
                  />
                  {{ item.name }}
                </a>
              </div>
            </nav>
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
      <div class="flex flex-shrink-0 justify-center items-center px-4">
        <img
          class="h-8 w-auto mr-2"
          src="/MailHedgehog.svg"
          :alt="t('app.title')"
        >
        <div class="text-context-100 font-bold text-sm sm:text-base">
          {{ t('app.title') }}
        </div>
      </div>
      <nav
        class="mt-5 flex flex-1 flex-col divide-y divide-primary-800 overflow-y-auto"
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
              {{ item.name }}
            </a>
          </router-link>
        </div>
      </nav>
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
          <ColorModeSelect />
          <LangModeSelect />
          <ProfileDropdown />
        </div>
      </div>
    </div>
    <main class="flex-grow flex-1 pb-8 overflow-hidden">
      <router-view />
    </main>
    <div
      class="py-4 px-4 sm:px-6 flex justify-between items-center shadow dark:shadow-context-500 text-context-900 dark:text-context-100"
    >
      <div class="text-sm">
        <span class="mr-0">
          ©
        </span>
        <a
          href="https://think.studio/"
          class="underline transition-colors duration-500 hover:text-primary-600"
          target="_blank"
        >
          Think One Communications Ltd
        </a>
      </div>
      <div class="hidden sm:block text-sm text-context-300 dark:text-context-700 select-none">
        <div class="sr-only">
          #StandWithUkraine
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import {
  Dialog,
  DialogPanel,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue';
import {
  Bars3CenterLeftIcon,
  XMarkIcon,
  InboxArrowDownIcon,
} from '@heroicons/vue/24/outline';
import { useI18n } from 'vue-i18n';
import ColorModeSelect from '@/components/Header/ColorModeSelect.vue';
import LangModeSelect from '@/components/Header/LangModeSelect.vue';
import ProfileDropdown from '@/components/Header/ProfileDropdown.vue';

const { t } = useI18n();

const navigation = [
  { name: t('menu.inbox'), href: '/', icon: InboxArrowDownIcon },
];

const sidebarOpen = ref(false);

</script>
