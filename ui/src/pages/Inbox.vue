<template>
  <!-- Page header -->
  <div class="lg:-mt-px bg-gray-50 dark:bg-gray-900 shadow dark:shadow-gray-500">
    <div class="px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8">
      <div class="py-6 flex items-center justify-between">
        <h1
          class="ml-3 text-xl md:text-2xl font-bold leading-7 text-gray-900 dark:text-gray-100 truncate sm:leading-9"
        >
          {{ t('inbox.hello', {msg: 'vtp'}) }}
          {{ t('inbox.pageTitle') }}
        </h1>
        <div
          class="flex justify-end ml-4 transition-all duration-200"
          :class="{
            'pointer-events-none opacity-75': isRequesting
          }"
        >
          <button
            v-if="pagination && !pagination.isEmpty()"
            v-tooltip="t('inbox.clear')"
            type="button"
            class="btn btn--default whitespace-nowrap"
            @click.prevent="clearInbox()"
          >
            <TrashIcon class="w-4 h-4 md:mr-2" />
            <span
              class="hidden md:inline"
            >
              {{ t('inbox.clear') }}
            </span>
          </button>
        </div>
      </div>
    </div>
  </div>

  <div class="mt-8">
    <div
      v-if="pagination.isEmpty()"
      class=""
    >
      <h3 class="ml-3 text-lg md:text-xl text-center text-gray-900 dark:text-gray-100 select-none">
        <template v-if="isRequesting">
          {{ t('pagination.requesting') }}
        </template>
        <template v-else>
          {{ t('inbox.empty') }}
        </template>
      </h3>
    </div>
    <template v-else>
      <!-- List (smallest breakpoint only) -->
      <div
        v-if="pagination"
        class="shadow dark:shadow-gray-500 md:hidden"
        :class="{
          'pointer-events-none opacity-75': isRequesting
        }"
      >
        <ul
          role="list"
          class="mt-2 divide-y divide-gray-200 overflow-hidden shadow dark:shadow-gray-500"
        >
          <li
            v-for="email in emails"
            :key="email.id"
          >
            <div
              class="block bg-gray-50 dark:bg-gray-800 px-4 py-4 hover:bg-gray-50"
            >
              <span class="flex items-center space-x-4">
                <span class="flex flex-1 space-x-2 truncate">
                  <span class="flex flex-col truncate text-sm text-gray-500 dark:text-gray-400">
                    <span
                      v-if="email.from"
                      class="truncate"
                    >
                      <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">
                        {{ t('email.from') }}
                      </span>
                      {{ email.from.name }}({{ email.from.email }})
                    </span>
                    <span
                      v-if="email.to[0]"
                      class="truncate"
                    >
                      <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">
                        {{ t('email.to') }}
                      </span>
                      {{ email.to[0].name }}({{ email.to[0].email }})
                    </span>
                    <span>
                      <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">
                        {{ t('email.subject') }}
                      </span>
                      {{ email.subject }}
                    </span>
                    <span>
                      <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">
                        {{ t('email.received_at') }}
                      </span>
                      <time
                        v-if="moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').isValid()"
                        :datetime="email.received_at"
                      >{{ moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').locale(locale).fromNow() }}
                      </time>
                    </span>
                  </span>
                </span>
                <div class="space-y-2">
                  <a
                    class="
                        cursor-pointer block
                        transition-all duration-500
                        text-gray-500 dark:text-gray-400
                        hover:text-gray-700 hover:dark:text-gray-300
                      "
                    @click.prevent="showEmail(email.id)"
                  >
                    <EyeIcon
                      class="h-5 w-5 flex-shrink-0"
                      aria-hidden="true"
                    />
                  </a>
                  <a
                    class="
                        cursor-pointer block
                        transition-all duration-500
                        text-gray-500 dark:text-gray-400
                        hover:text-gray-700 hover:dark:text-gray-300
                      "
                    @click.prevent="deleteEmail(email.id)"
                  >
                    <TrashIcon
                      class="h-5 w-5 flex-shrink-0"
                      aria-hidden="true"
                    />
                  </a>
                </div>
              </span>
            </div>
          </li>
        </ul>

        <nav
          class="flex items-center justify-between border-t border-gray-200 bg-gray-50 dark:bg-gray-800 px-4 py-3"
          aria-label="Pagination"
        >
          <div class="flex flex-1 justify-between">
            <button
              :disabled="pagination.isOnFirst()"
              class="btn btn--default"
              @click.prevent="goToDirection('prev')"
            >
              {{ t('pagination.prev') }}
            </button>
            <button
              :disabled="pagination.isOnLast()"
              class="ml-3 btn btn--default"
              @click.prevent="goToDirection('next')"
            >
              {{ t('pagination.next') }}
            </button>
          </div>
        </nav>
      </div>
      <!-- List (small breakpoint and up) -->
      <div
        v-if="pagination"
        class="hidden md:block"
      >
        <div
          class="mx-auto max-w-6xl px-4 md:px-6 lg:px-8 transition-all duration-200"
          :class="{
            'pointer-events-none opacity-75': isRequesting
          }"
        >
          <div class="mt-2 flex flex-col">
            <div class="min-w-full overflow-hidden overflow-x-auto align-middle shadow md:rounded-lg">
              <table class="min-w-full divide-y divide-gray-200">
                <thead>
                  <tr>
                    <th
                      class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                      scope="col"
                    >
                      {{ t('email.from') }}
                    </th>
                    <th
                      class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                      scope="col"
                    >
                      {{ t('email.to') }}
                    </th>
                    <th
                      class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                      scope="col"
                    >
                      {{ t('email.subject') }}
                    </th>
                    <th
                      class="whitespace-nowrap bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                      scope="col"
                    >
                      {{ t('email.received_at') }}
                    </th>
                    <th
                      class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-sm font-semibold text-gray-900 dark:text-gray-100 text-right"
                      scope="col"
                    >
                      {{ t('pagination.actions') }}
                    </th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-200 bg-gray-50 dark:bg-gray-800">
                  <tr
                    v-for="email in emails"
                    :key="email.id"
                    class="bg-gray-50 dark:bg-gray-800"
                  >
                    <td class="max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-gray-900 dark:text-gray-100">
                      <template v-if="email.from">
                        <div class="truncate">
                          {{ email.from.name }}
                        </div>
                        <div class="truncate">
                          <a
                            :href="`mailto:${email.from.email}`"
                            class="text-gray-500 dark:text-gray-400 hover:text-gray-700 hover:dark:text-gray-200 transition-all duration-500"
                          >
                            {{ email.from.email }}
                          </a>
                        </div>
                      </template>
                      <div v-else>
                        {{ t('email.notAvailable') }}
                      </div>
                    </td>
                    <td class="max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-gray-900 dark:text-gray-100">
                      <template v-if="email.to[0]">
                        <div class="truncate">
                          {{ email.to[0].name }}
                        </div>
                        <div class="truncate">
                          <a
                            :href="`mailto:${email.to[0].email}`"
                            class="text-gray-500 dark:text-gray-400 hover:text-gray-700 hover:dark:text-gray-200 transition-all duration-500"
                          >
                            {{ email.to[0].email }}
                          </a>
                        </div>
                      </template>
                      <div v-else>
                        {{ t('email.notAvailable') }}
                      </div>
                    </td>
                    <td
                      class="w-full max-w-0 whitespace-nowrap truncate px-6 py-4 text-sm text-gray-900 dark:text-gray-100"
                    >
                      {{ email.subject }}
                    </td>
                    <td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500 dark:text-gray-400">
                      <time
                        v-if="moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').isValid()"
                        :datetime="email.received_at"
                      >{{ moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').locale(locale).fromNow() }}
                      </time>
                    </td>
                    <td
                      class="whitespace-nowrap px-6 py-4 text-sm text-right flex justify-end space-x-1"
                    >
                      <a
                        class="
                        cursor-pointer
                        transition-all duration-500
                        text-gray-500 dark:text-gray-400
                        hover:text-gray-700 hover:dark:text-gray-300
                      "
                        @click.prevent="showEmail(email.id)"
                      >
                        <EyeIcon
                          class="w-5 h-5"
                          aria-hidden="true"
                        />
                      </a>
                      <a
                        class="
                        cursor-pointer
                        transition-all duration-500
                        text-gray-500 dark:text-gray-400
                        hover:text-gray-700 hover:dark:text-gray-300
                      "
                        @click.prevent="deleteEmail(email.id)"
                      >
                        <TrashIcon
                          class="w-5 h-5"
                          aria-hidden="true"
                        />
                      </a>
                    </td>
                  </tr>
                </tbody>
              </table>
              <!-- Pagination -->
              <nav
                class="flex items-center justify-between border-t border-gray-200 bg-gray-50 dark:bg-gray-800 px-4 py-3 sm:px-6"
                aria-label="Pagination"
              >
                <div class="hidden sm:block">
                  <p
                    class="text-sm text-gray-700 dark:text-gray-200"
                    v-html="t('pagination.text', {from: pagination.getFrom(), to: pagination.getTo(), of: pagination.getTotal()})"
                  />
                </div>
                <div class="flex flex-1 justify-between sm:justify-end">
                  <button
                    :disabled="pagination.isOnFirst()"
                    class="btn btn--default"
                    @click.prevent="goToDirection('prev')"
                  >
                    {{ t('pagination.prev') }}
                  </button>
                  <button
                    :disabled="pagination.isOnLast()"
                    class="ml-3 btn btn--default"
                    @click.prevent="goToDirection('next')"
                  >
                    {{ t('pagination.next') }}
                  </button>
                </div>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
  <Teleport to="#header-search">
    <form
      class="flex w-full md:ml-0"
      method="GET"
      @submit.prevent="getEmails(1)"
    >
      <label
        for="search-field"
        class="sr-only"
      >
        {{ t('inbox.search') }}
      </label>
      <div class="relative w-full text-gray-400 focus-within:text-gray-600">
        <div
          class="pointer-events-none absolute inset-y-0 left-0 flex items-center"
          aria-hidden="true"
        >
          <MagnifyingGlassIcon
            class="h-5 w-5"
            aria-hidden="true"
          />
        </div>
        <input
          id="search-field"
          v-model="queryParams.search"
          name="search-field"
          class="block h-full w-full  py-2 pl-8 pr-3  sm:text-sm
          focus:outline-none focus:ring-0
          border-transparent focus:border-transparent
          placeholder-gray-500
          bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100
          "
          :placeholder="t('inbox.search')"
          type="search"
        >
      </div>
    </form>
  </Teleport>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import moment from 'moment';
import 'moment/min/locales';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import {
  EyeIcon,
  TrashIcon,
  MagnifyingGlassIcon,
} from '@heroicons/vue/24/outline';
import Pagination from '../utils/pagination';

const { t, locale } = useI18n();
const router = useRouter();

const queryParams = ref({
  page: 1,
  per_page: 2,
  search: '',
});
const isRequesting = ref(false);
const emails = ref([]);
const pagination = ref(new Pagination());

let searchTimeout = null;
watch(() => queryParams.value.search, () => {
  if (searchTimeout) {
    clearTimeout(searchTimeout);
    searchTimeout = null;
  }
  searchTimeout = setTimeout(() => {
    getEmails(1);
  }, 500);
}, { deep: true });

const getEmails = (page = null) => {
  isRequesting.value = true;
  if (page) {
    queryParams.value.page = page;
  }
  window.MailHedgehog.request()
    .get('emails', {
      params: queryParams.value,
    })
    .then((response) => {
      if (response.data?.data) {
        emails.value = response.data?.data;
      } else {
        emails.value = [];
      }
      if (response.data?.meta?.pagination) {
        pagination.value = new Pagination(
          response.data?.meta?.pagination.current_page,
          response.data?.meta?.pagination.per_page,
          response.data?.meta?.pagination.last_page,
          response.data?.meta?.pagination.from,
          response.data?.meta?.pagination.to,
          response.data?.meta?.pagination.total,
        );
      } else {
        pagination.value = new Pagination();
      }
    })
    .finally(() => {
      isRequesting.value = false;
    });
};

onMounted(() => {
  getEmails(1);
});

const goToPage = (page) => {
  if (pagination.value.hasPage(page)) {
    getEmails(page);
  }
};

const goToDirection = (direction = 'next') => {
  getEmails(pagination.value.getPageFromDirection(direction));
};

const clearInbox = () => {
  isRequesting.value = true;
  window.MailHedgehog.request()
    .delete('emails')
    .then((response) => {
      getEmails(1);
      window.MailHedgehog.success(t('inbox.cleared'));
    })
    .catch(() => {
      isRequesting.value = false;
    });
};

const showEmail = (emailId) => {
  router.push({ name: 'email', params: { id: emailId } });
};

const deleteEmail = (emailId) => {
  isRequesting.value = true;
  window.MailHedgehog.request()
    .delete(`emails/${emailId}`)
    .then((response) => {
      if (pagination.value.count() == 0) {
        goToDirection('prev');
      } else {
        getEmails();
      }
      window.MailHedgehog.success(t('email.deleted'));
    })
    .catch(() => {
      isRequesting.value = false;
    });
};

</script>
