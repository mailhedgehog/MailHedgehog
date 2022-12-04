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
      class="
        px-4 sm:px-6 lg:px-8
      "
    >
      TODO:headers
    </div>
  </div>
  <div
    class="bg-white px-4 sm:px-6 lg:px-8"
  >
    {{ email }}
  </div>
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import {
  Dialog,
  DialogPanel,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue';
import {
  ArrowSmallLeftIcon,
  TrashIcon,
  DocumentArrowDownIcon,
  PaperAirplaneIcon,
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
        emails.value = null;
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

</script>
