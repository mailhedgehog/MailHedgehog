<template>
  <div
    class="flex flex-col justify-center"
    :class="{
      'pointer-events-none opacity-75': isRequesting
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
      <div class="flex justify-between py-3">
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
                class="btn btn--default rounded-r-md rounded-l-none"
                @click.prevent="downloadEmail"
              >
                <DocumentArrowDownIcon class="md:mr-2 h-5 w-5" />
                <span class="hidden md:inline">
                  {{ t('email.download') }}
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
        class="overflow-hidden bg-context-50 dark:bg-context-900 border border-context-300 shadow dark:shadow-context-500 sm:rounded-lg mb-6"
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
              :selected="currentTab"
            >
              {{ tab.name }}
            </option>
          </select>
        </div>
        <div class="hidden sm:block">
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
                {{ tab.name }}
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
      <div v-if="email">
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
</template>
<script setup>
import { useRouter, useRoute } from 'vue-router';
import {
  computed, nextTick, onMounted, ref,
} from 'vue';
import { useI18n } from 'vue-i18n';
import {
  ArrowSmallLeftIcon,
  TrashIcon,
  DocumentArrowDownIcon,
  EyeIcon,
  EyeSlashIcon,
  PaperClipIcon,
  ArrowDownTrayIcon,
  ExclamationTriangleIcon,
} from '@heroicons/vue/24/outline';

const { t } = useI18n();
const router = useRouter();
const route = useRoute();

const isRequesting = ref(false);
const email = ref(null);

const getEmail = (emailId) => {
  isRequesting.value = true;
  window.MailHedgehog.request()
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
  getEmail(route.params.id);
});

const goBack = () => {
  router.push({ name: 'emails', params: {} });
};

const deleteEmail = () => {
  isRequesting.value = true;
  window.MailHedgehog.request()
    .delete(`emails/${email.value.id}`)
    .then(() => {
      window.MailHedgehog.success(t('email.deleted'));
      nextTick(() => goBack());
    })
    .catch(() => {
      window.MailHedgehog.error(t('response.error'));
    })
    .finally(() => {
      isRequesting.value = false;
    });
};

const downloadBlobFile = (data, contentType, filename) => {
  const blob = new Blob([data], { type: contentType });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.download = filename;
  link.click();
  URL.revokeObjectURL(link.href);
  link.remove();
};

const downloadEmail = () => {
  downloadBlobFile(email.value.source, 'text/plain', `${email.value.id}.eml`);
};

const downloadEmailAttachment = (attachment) => {
  downloadBlobFile(attachment.data, attachment.contentType, attachment.filename);
};

const isAllHeaders = ref(false);
const importantHeaders = [
  'Subject',
  'From',
  'To',
  'Cc',
];
const headers = computed(() => {
  const headersList = {};

  if (email.value && email.value.headers) {
    for (let i = 0; i < Object.entries(email.value.headers).length; i += 1) {
      const [headerKey, headerValues] = Object.entries(email.value.headers)[i];
      if (Array.isArray(headerValues) && headerValues.length > 0) {
        if (isAllHeaders.value) {
          headersList[headerKey] = headerValues;
        } else if (importantHeaders.includes(headerKey)) {
          headersList[headerKey] = headerValues;
        }
      }
    }
  }

  return headersList;
});

const currentTab = ref('html');
const tabs = [
  {
    id: 'html',
    name: t('email.tab.html'),
  },
  {
    id: 'plain',
    name: t('email.tab.plain'),
  },
  {
    id: 'source',
    name: t('email.tab.source'),
  },
  {
    id: 'attachments',
    name: t('email.tab.attachments'),
  },
];

const iframeHeight = ref('0rem');
const resizeIframe = (event) => {
  const obj = event.currentTarget;
  const newHeight = obj.contentWindow.document.documentElement.scrollHeight + 2;
  iframeHeight.value = `${newHeight}px`;
};

</script>
