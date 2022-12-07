<template>
  <div
    class="flex flex-col justify-center"
    :class="{
      'pointer-events-none opacity-75': isRequesting
    }"
  >
    <div
      class="
        bg-white dark:bg-gray-900
        px-4 sm:px-6 lg:px-8
        shadow dark:shadow-gray-500
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
                class="btn btn--default rounded-r-none rounded-l-none"
                @click.prevent="downloadEmail"
              >
                <DocumentArrowDownIcon class="md:mr-2 h-5 w-5" />
                <span class="hidden md:inline">
                  {{ t('email.download') }}
                </span>
              </button>
              <button
                v-tooltip="t('email.releaseHint')"
                type="button"
                class="btn btn--default rounded-r-md rounded-l-none"
                @click.prevent="releaseEmail"
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
      <div class="overflow-hidden bg-white dark:bg-gray-900 border border-gray-300 shadow dark:shadow-gray-500 sm:rounded-lg mb-6">
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900 dark:text-gray-100 flex justify-between">
            <div>
              Email headers
            </div>
            <div>
              <a
                v-tooltip="isAllHeaders?'Show only important headers':'Show all headers'"
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
          <p class="mt-1 max-w-2xl text-sm text-gray-500 dark:text-gray-400">
            An Email Header is metadata that accompanies every email and contains detailed information
          </p>
        </div>
        <div class="border-t border-gray-200 px-4 py-5 sm:p-0">
          <dl class="sm:divide-y sm:divide-gray-200 sm:dark:divide-gray-300">
            <div
              v-for="(headerValues, headerName) in headers"
              :key="headerName"
              class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:py-5 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500 dark:text-gray-400">
                {{ headerName }}
              </dt>
              <dd class="mt-1 text-sm text-gray-900 dark:text-gray-100 sm:col-span-2 sm:mt-0">
                <div
                  v-if="headerValues.length === 1"
                >
                  {{ headerValues[0] }}
                </div>
                <ul
                  v-else
                  role="list"
                  class="divide-y divide-gray-200 dark:divide-gray-200 rounded-md border border-gray-200 dark:border-gray-200"
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
    </div>
  </div>
  <div
    class="bg-white px-4 sm:px-6 lg:px-8"
  >
    <pre>{{ JSON.stringify(email, null, 2) }}</pre>
  </div>
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router';
import { computed, onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import {
  ArrowSmallLeftIcon,
  TrashIcon,
  DocumentArrowDownIcon,
  PaperAirplaneIcon,
  EyeIcon,
  EyeSlashIcon,
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

const deleteEmail = () => {
  alert('TODO: delete');
};

const downloadEmail = () => {
  alert('TODO: download');
};

const releaseEmail = () => {
  alert('TODO: release');
};

const goBack = () => {
  router.push({ name: 'emails', params: {} });
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

</script>
