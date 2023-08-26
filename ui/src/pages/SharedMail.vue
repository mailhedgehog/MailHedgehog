<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import {
  inject, onMounted, Ref, ref,
} from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
import {
  ArrowDownTrayIcon,
  DocumentArrowDownIcon,
  ExclamationTriangleIcon, EyeIcon,
  EyeSlashIcon,
  PaperClipIcon,
} from '@heroicons/vue/24/outline';
import { MailHedgehog } from '@/main';
import { useShowHeader } from '@/composables/email-headers';
import { useEmailTabs } from '@/composables/email-tabs';
import { useEmailIframe } from '@/composables/email-iframe';
import { Email } from '@/interfaces/Email';
import { useEmailAttachments } from '@/composables/email-attachments';
import ColorModeSelect from '@/components/Header/ColorModeSelect.vue';
import LangModeSelect from '@/components/Header/LangModeSelect.vue';
import Copyright from '@/components/Footer/Copyright.vue';
import FooterRightContent from '@/components/Footer/FooterRightContent.vue';
import Logo from '@/components/Header/Logo.vue';

const { t } = useI18n();
const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
const router = useRouter();
const route = useRoute();
const store = useStore();

const isRequesting = ref(false);
const notFound = ref(false);
const email = ref<Email|null>(null);

const { isAllHeaders, headers } = useShowHeader(email);
const { currentTab, tabs } = useEmailTabs();
const { iframeHeight, resizeIframe } = useEmailIframe();
const { downloadEmailAttachment, downloadEmail } = useEmailAttachments(email, {
  urlBuilder: (index: number, email: Ref<Email|null>) => `shared-email/${route.params.id}/attachment/${index}`,
});

const getSharedEmail = (sharedEmailId: string) => {
  isRequesting.value = true;
  mailHedgehog?.request()
    .get(`shared-email/${sharedEmailId}`)
    .then((response) => {
      if (response.data?.data) {
        email.value = response.data?.data;
      } else {
        email.value = null;
      }
    })
    .catch(() => {
      notFound.value = true;
    })
    .finally(() => {
      isRequesting.value = false;
    });
};

onMounted(() => {
  getSharedEmail(route.params.id as string);
});
</script>

<template>
  <div class="min-h-screen flex flex-col bg-context-50 dark:bg-context-900">
    <div class="flex justify-between items-center py-4 px-4 sm:px-6 lg:px-8 relative shadow dark:shadow-context-500">
      <Logo class="" />

      <div class="inline-flex sm:space-x-3 z-10">
        <ColorModeSelect />
        <LangModeSelect />
      </div>
    </div>
    <div
      v-if="notFound"
      class="flex-grow p-8 flex justify-center items-center"
    >
      <div class="flex flex-col items-center space-y-6">
        <div class="font-semibold text-4xl text-primary-600">
          404
        </div>
        <div class="text-context-900 dark:text-context-100 text-lg">
          {{ t('sharedEmail.notFound') }}
        </div>
      </div>
    </div>
    <div
      v-else
      class="flex-grow flex flex-col py-8"
      :class="{
        'pointer-events-none opacity-75': isRequesting
      }"
    >
      <div
        class="px-4 sm:px-6 lg:px-8"
      >
        <div
          class="
          overflow-hidden bg-context-50 dark:bg-context-900 border border-context-300
          shadow dark:shadow-context-500 sm:rounded-lg mb-6
          lg:mx-auto lg:max-w-6xl
        "
        >
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6 text-context-900 dark:text-context-100 flex justify-between">
              <div>
                {{ t('email.headersTitle') }}
              </div>
              <div>
                <a
                  v-tooltip="isAllHeaders?t('email.hintHideHeaders'):t('email.hintShowHeaders')"
                  class="cursor-pointer"
                  @click.prevent="isAllHeaders = !isAllHeaders"
                >
                  <EyeSlashIcon
                    v-if="isAllHeaders"
                    class="text-primary-400 h-5 w-5"
                  />
                  <EyeIcon
                    v-else
                    class="text-primary-700 h-5 w-5"
                  />
                </a>
              </div>
            </h3>
            <p class="mt-1 max-w-2xl text-sm text-context-500 dark:text-context-400">
              {{ t('email.headersSubtitle') }}
            </p>
          </div>
          <div class="border-t border-context-200 px-4 py-5 sm:p-0">
            <dl class="sm:divide-y sm:divide-context-200 sm:dark:divide-context-300">
              <div
                v-for="(headerValues, headerName) in headers"
                :key="headerName"
                class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:py-5 sm:px-6"
              >
                <dt class="text-sm font-medium text-context-500 dark:text-context-400">
                  {{ headerName }}
                </dt>
                <dd class="mt-1 text-sm text-context-900 dark:text-context-100 sm:col-span-2 sm:mt-0">
                  <div
                    v-if="headerValues.length === 1"
                  >
                    {{ headerValues[0] }}
                  </div>
                  <ul
                    v-else
                    role="list"
                    class="divide-y divide-context-200 dark:divide-context-200 rounded-md border border-context-200 dark:border-context-200"
                  >
                    <li
                      v-for="(headerValue, index) in headerValues"
                      :key="index"
                      class="py-3 pl-3 pr-4 text-sm"
                    >
                      {{ headerValue }}
                    </li>
                  </ul>
                </dd>
              </div>
            </dl>
          </div>
        </div>
        <div class="flex items-center lg:mx-auto lg:max-w-6xl border-b-0 sm:border-b border-context-200 dark:border-context-400">
          <div class="flex-grow">
            <div class="sm:hidden">
              <select
                id="tabs"
                v-model="currentTab"
                name="tabs"
                class="block w-full rounded border-context-300 dark:border-context-600 py-2 pl-3 pr-10 text-base
            focus:border-primary-500 focus:primary-none focus:ring-primary-500 sm:text-sm"
              >
                <option
                  v-for="tab in tabs"
                  :key="tab.id"
                  :value="tab.id"
                  :selected="currentTab == tab.id"
                >
                  {{ t(`email.tab.${tab.id}`) }}
                </option>
              </select>
            </div>
            <div class="hidden sm:block">
              <nav
                class="-mb-px flex space-x-8"
                aria-label="Tabs"
              >
                <a
                  v-for="tab in tabs"
                  :key="tab.id"
                  href="#"
                  class="whitespace-nowrap flex py-4 px-1 border-b-2 font-medium text-sm transition-colors duration-500"
                  :class="[currentTab === tab.id ? 'border-primary-500 text-primary-600' : 'cursor-pointer border-transparent text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300 hover:border-context-200 hover:dark:border-context-400']"
                  :aria-current="currentTab === tab.id ? 'page' : undefined"
                  @click.prevent="currentTab = tab.id"
                >
                  {{ t(`email.tab.${tab.id}`) }}
                  <span
                    v-if="tab.id==='attachments' && email?.attachments?.length !== undefined"
                    class="inline-block ml-3 py-0.5 px-2.5 rounded-full text-xs font-medium transition-colors duration-500"
                    :class="[currentTab === tab.id ? 'bg-primary-100 text-primary-600' : 'bg-context-200 dark:bg-context-700 text-context-900 dark:text-context-100']"
                  >
                    {{ email?.attachments?.length }}
                  </span>
                </a>
              </nav>
            </div>
          </div>

          <div class="pl-4">
            <button
              v-tooltip="t('email.downloadHint')"
              type="button"
              class="btn btn--default rounded"
              @click.prevent="downloadEmail()"
            >
              <DocumentArrowDownIcon class="md:mr-2 h-5 w-5" />
              <span class="hidden md:inline">
                {{ t('email.download') }}
              </span>
            </button>
          </div>
        </div>
        <div
          v-if="email"
          class="lg:mx-auto lg:max-w-6xl"
        >
          <div
            v-if="currentTab === 'html'"
            class="py-6"
          >
            <iframe
              v-if="email.html"
              id="preview-html"
              :height="iframeHeight"
              :srcdoc="email.html"
              class="w-full border border-context-200 dark:border-context-400"
              @load="resizeIframe"
            />
            <div
              v-else
              class="flex justify-center items-center"
            >
              <ExclamationTriangleIcon
                class="text-primary-500 h-8 w-8"
              />
              <div class="ml-4 text-context-900 dark:text-context-100">
                {{ t('email.htmlEmpty') }}
              </div>
            </div>
          </div>
          <div
            v-if="currentTab === 'plain'"
            class="py-6"
          >
            <div
              v-if="email.plain"
              class="text-context-900 dark:text-context-100"
            >
              <pre class="overflow-auto">{{ email.plain }}</pre>
            </div>
            <div
              v-else
              class="flex justify-center items-center"
            >
              <ExclamationTriangleIcon
                class="text-primary-500 h-8 w-8"
              />
              <div class="ml-4 text-context-900 dark:text-context-100">
                {{ t('email.plainEmpty') }}
              </div>
            </div>
          </div>
          <div
            v-if="currentTab === 'source'"
            class="py-6"
          >
            <div
              class="text-context-900 dark:text-context-100"
            >
              <pre class="overflow-auto">{{ email.source }}</pre>
            </div>
          </div>
          <div
            v-if="currentTab === 'attachments'"
            class="py-6"
          >
            <div
              v-if="email?.attachments?.length > 0"
            >
              <ul
                role="list"
                class="divide-y divide-context-200 rounded-md border border-context-200"
              >
                <li
                  v-for="(attachment, index) in email?.attachments"
                  :key="index"
                  class="flex items-center justify-between py-3 pl-3 pr-4 text-sm"
                >
                  <div class="flex w-0 flex-1 items-center">
                    <PaperClipIcon
                      class="text-context-400 dark:text-context-500 flex-shrink-0 h-5 w-5"
                    />
                    <span class="ml-2 w-0 flex-1 text-context-900 dark:text-context-100 truncate">
                      <span>
                        {{ attachment.filename }}
                      </span>
                      <span
                        v-if="attachment.contentType"
                        class="ml-2"
                      >
                        ({{ attachment.contentType }})
                      </span>
                    </span>
                  </div>
                  <div class="ml-4 flex-shrink-0">
                    <a
                      class="cursor-pointer font-medium text-primary-500 hover:text-primary-600 transition-colors duration-500"
                      @click.prevent="downloadEmailAttachment(attachment)"
                    >
                      <ArrowDownTrayIcon
                        class="md:hidden flex-shrink-0 h-5 w-5"
                      />
                      <span class="hidden md:inline">
                        {{ t('email.download') }}
                      </span>
                    </a>
                  </div>
                </li>
              </ul>
            </div>
            <div
              v-else
              class="flex justify-center items-center"
            >
              <ExclamationTriangleIcon
                class="text-primary-500 h-8 w-8"
              />
              <div class="ml-4 text-context-900 dark:text-context-100">
                {{ t('email.noAttachments') }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="py-4 px-4 sm:px-6 flex justify-between items-center shadow dark:shadow-context-500 text-context-900 dark:text-context-100"
    >
      <Copyright />
      <FooterRightContent />
    </div>
  </div>
</template>
