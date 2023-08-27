<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import {
  inject, nextTick, onMounted, ref,
} from 'vue';
import { useI18n } from 'vue-i18n';
import {
  ArrowSmallLeftIcon,
  TrashIcon,
  DocumentArrowDownIcon,
  PaperAirplaneIcon,
  EyeIcon,
  EyeSlashIcon,
  PaperClipIcon,
  ArrowDownTrayIcon,
  ExclamationTriangleIcon,
  ShareIcon,
} from '@heroicons/vue/24/outline';
import { useStore } from 'vuex';
import { useShowHeader } from '@/composables/email-headers';
import { useEmailTabs } from '@/composables/email-tabs';
import { useEmailIframe } from '@/composables/email-iframe';
import { useEmailAttachments } from '@/composables/email-attachments';
import { Email } from '@/interfaces/Email';
import { MailHedgehog } from '@/main';
import { useEmailRelease } from '@/composables/email-release';
import { useEmailShare } from '@/composables/email-share';
import ModalReleaseEmail from '@/components/Modal/ReleaseEmail.vue';
import ModalShareEmail from '@/components/Modal/ShareEmail.vue';

const { t } = useI18n();
const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
const router = useRouter();
const route = useRoute();
const store = useStore();

const isRequesting = ref(false);
const email = ref<Email|null>(null);

const { isAllHeaders, headers } = useShowHeader(email);
const { currentTab, tabs } = useEmailTabs();
const { iframeHeight, resizeIframe } = useEmailIframe();
const { downloadEmailAttachment, downloadEmail } = useEmailAttachments(email);
const {
  isOpenReleaseModal,
  isRequestingReleaseEmail,
} = useEmailRelease();
const {
  isSharingEnabled,
  isOpenShareModal,
  isRequestingShareEmail,
} = useEmailShare();

const getEmail = (emailId: string) => {
  isRequesting.value = true;
  mailHedgehog?.request()
    .get(`emails/${emailId}`)
    .then((response) => {
      if (response.data?.data) {
        email.value = response.data?.data;
      } else {
        email.value = null;
      }
    })
    .catch(() => {
      router.push({ path: '/404' });
    })
    .finally(() => {
      isRequesting.value = false;
    });
};

onMounted(() => {
  getEmail(route.params.id as string);
});

const goBack = () => {
  router.push({ name: 'emails', params: {} });
};

const deleteEmail = () => {
  if (!email.value) {
    return;
  }
  store.dispatch('confirmDialog/confirm')
    .then(() => {
      isRequesting.value = true;
      mailHedgehog?.request()
        .delete(`emails/${email.value?.id}`)
        .then(() => {
          mailHedgehog?.success(t('email.deleted'));
          nextTick(() => goBack());
        })
        .catch((error) => {
          mailHedgehog.onResponseError(error, 'Response Error');
        })
        .finally(() => {
          isRequesting.value = false;
        });
    });
};

</script>

<template>
  <div
    class="flex flex-col justify-center"
    :class="{
      'pointer-events-none opacity-75': isRequesting || isRequestingReleaseEmail
    }"
  >
    <div
      class="
        bg-context-50 dark:bg-context-900
        px-4 sm:px-6 lg:px-8
        shadow dark:shadow-context-500
        mb-6 md:mb-8
      "
    >
      <div class="flex justify-between py-3 lg:mx-auto lg:max-w-6xl">
        <!-- Left buttons -->
        <div>
          <div class="isolate inline-flex rounded-md shadow-sm sm:space-x-3 sm:shadow-none">
            <button
              type="button"
              class="btn btn--default"
              @click.prevent="goBack"
            >
              <ArrowSmallLeftIcon class="md:mr-2 h-5 w-5" />
              <span class="hidden md:inline">
                {{ t('email.back') }}
              </span>
            </button>
          </div>
        </div>

        <!-- Right buttons -->
        <div>
          <div class="isolate inline-flex rounded-md shadow-sm sm:space-x-3 sm:shadow-none">
            <span class="inline-flex sm:shadow-sm">
              <button
                v-tooltip="t('email.deleteHint')"
                type="button"
                class="btn btn--default rounded-r-none rounded-l-md"
                @click.prevent="deleteEmail"
              >
                <TrashIcon class="md:mr-2 h-5 w-5" />
                <span class="hidden md:inline">
                  {{ t('email.delete') }}
                </span>
              </button>
              <button
                v-tooltip="t('email.downloadHint')"
                type="button"
                class="btn btn--default rounded-none"
                @click.prevent="downloadEmail()"
              >
                <DocumentArrowDownIcon class="md:mr-2 h-5 w-5" />
                <span class="hidden md:inline">
                  {{ t('email.download') }}
                </span>
              </button>
              <button
                v-if="isSharingEnabled"
                v-tooltip="t('email.shareHint')"
                type="button"
                class="btn btn--default rounded-none"
                @click.prevent="isOpenShareModal = true"
              >
                <ShareIcon class="md:mr-2 h-5 w-5" />
                <span class="hidden md:inline">
                  {{ t('email.share') }}
                </span>
              </button>
              <button
                v-tooltip="t('email.releaseHint')"
                type="button"
                class="btn btn--default rounded-r-md rounded-l-none"
                @click.prevent="isOpenReleaseModal = true"
              >
                <PaperAirplaneIcon class="md:mr-2 h-5 w-5" />
                <span class="hidden md:inline">
                  {{ t('email.release') }}
                </span>
              </button>
            </span>
          </div>
        </div>
      </div>
    </div>
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
      <div>
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
        <div class="hidden sm:block lg:mx-auto lg:max-w-6xl">
          <div class="border-b border-context-200 dark:border-context-400">
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
  <ModalReleaseEmail
    :email="email"
    :show="isOpenReleaseModal"
    :requesting="isRequesting || isRequestingReleaseEmail"
    @close="isOpenReleaseModal = false"
  />
  <ModalShareEmail
    v-if="isSharingEnabled"
    :email="email"
    :show="isOpenShareModal"
    :requesting="isRequesting || isRequestingShareEmail"
    @close="isOpenShareModal = false"
  />
</template>
