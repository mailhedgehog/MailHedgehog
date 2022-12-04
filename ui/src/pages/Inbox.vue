<template>
  <!-- Page header -->
  <div class="lg:-mt-px bg-white dark:bg-gray-900 shadow dark:shadow-gray-500">
    <div class="px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8">
      <div class="py-6 flex items-center justify-between">
        <h1
          class="ml-3 text-xl md:text-2xl font-bold leading-7 text-gray-900 dark:text-gray-100 truncate sm:leading-9"
        >
          {{ $t('inbox.hello', {msg: 'vtp'}) }}
          {{ $t('inbox.pageTitle') }}
        </h1>
        <div class="flex justify-end ml-4">
          <button
            type="button"
            class="btn btn--default whitespace-nowrap"
            :title="$t('inbox.clear')"
          >
            <TrashIcon class="w-4 h-4 md:mr-2"/>
            <span
              class="hidden md:inline"
            >
              {{ $t('inbox.clear') }}
            </span>
          </button>
        </div>
      </div>
    </div>
  </div>

  <div class="mt-8">
    <!-- List (smallest breakpoint only) -->
    <div class="shadow dark:shadow-gray-500 md:hidden">
      <ul
        role="list"
        class="mt-2 divide-y divide-gray-200 overflow-hidden shadow dark:shadow-gray-500"
      >
        <li
          v-for="email in emails"
          :key="email.id"
        >
          <div
            class="block bg-white dark:bg-gray-800 px-4 py-4 hover:bg-gray-50"
          >
            <span class="flex items-center space-x-4">
              <span class="flex flex-1 space-x-2 truncate">
                <span class="flex flex-col truncate text-sm text-gray-500 dark:text-gray-400">
                  <span
                    v-if="email.from"
                    class="truncate"
                  >
                    <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">From</span>
                    {{ email.from.name }}({{ email.from.email }})
                  </span>
                  <span
                    v-if="email.to[0]"
                    class="truncate"
                  >
                    <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">To</span>
                    {{ email.to[0].name }}({{ email.to[0].email }})
                  </span>
                  <span>
                    <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">Subject</span>
                    {{ email.subject }}
                  </span>
                  <span>
                    <span class="font-medium text-gray-900 dark:text-gray-100 mr-2">Received at</span>
                    <time
                      v-if="moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').isValid()"
                      :datetime="email.received_at"
                    >{{ moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').fromNow() }}
                    </time>
                  </span>
                </span>
              </span>
              <div class="space-y-2">
                <EyeIcon
                  class="h-5 w-5 flex-shrink-0 text-gray-400"
                  aria-hidden="true"
                />
                <TrashIcon
                  class="h-5 w-5 flex-shrink-0 text-gray-400"
                  aria-hidden="true"
                />
              </div>
            </span>
          </div>
        </li>
      </ul>

      <nav
        class="flex items-center justify-between border-t border-gray-200 bg-white dark:bg-gray-800 px-4 py-3"
        aria-label="Pagination"
      >
        <div class="flex flex-1 justify-between">
          <a
            href="#"
            class="btn btn--default"
          >
            Previous
          </a>
          <a
            href="#"
            class="btn btn--default"
          >
            Next
          </a>
        </div>
      </nav>
    </div>

    <!-- List (small breakpoint and up) -->
    <div
      v-if="pagination"
      class="hidden md:block"
    >
      <div
        v-if="pagination.isEmpty()"
        class=""
      >
        <h3 class="ml-3 text-lg md:text-xl text-center text-gray-900 dark:text-gray-100 select-none">
          <template v-if="isRequesting">
            {{ $t('pagination.requesting') }}
          </template>
          <template v-else>
            {{ $t('inbox.empty') }}
          </template>
        </h3>
      </div>
      <div
        v-else
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
                  From
                </th>
                <th
                  class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                  scope="col"
                >
                  To
                </th>
                <th
                  class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                  scope="col"
                >
                  Subject
                </th>
                <th
                  class="whitespace-nowrap bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                  scope="col"
                >
                  Received at
                </th>
                <th
                  class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-sm font-semibold text-gray-900 dark:text-gray-100 text-right"
                  scope="col"
                >
                  Actions
                </th>
              </tr>
              </thead>
              <tbody class="divide-y divide-gray-200 bg-white dark:bg-gray-800">
              <tr
                v-for="email in emails"
                :key="email.id"
                class="bg-white dark:bg-gray-800"
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
                    n/a
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
                    n/a
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
                  >{{ moment(email.received_at, 'YYYY-MM-DD HH:mm:ss').fromNow() }}
                  </time>
                </td>
                <td
                  class="whitespace-nowrap px-6 py-4 text-sm text-right text-gray-500 dark:text-gray-400 flex justify-end space-x-1"
                >
                  <EyeIcon class="w-5 h-5"/>
                  <TrashIcon class="w-5 h-5"/>
                </td>
              </tr>
              </tbody>
            </table>
            <!-- Pagination -->
            <nav
              class="flex items-center justify-between border-t border-gray-200 bg-white dark:bg-gray-800 px-4 py-3 sm:px-6"
              aria-label="Pagination"
            >
              <div class="hidden sm:block">
                <p
                  class="text-sm text-gray-700 dark:text-gray-200"
                  v-html="$t('pagination.text', {from: pagination.getFrom(), to: pagination.getTo(), of: pagination.getTotal()})"
                />
              </div>
              <div class="flex flex-1 justify-between sm:justify-end">
                <button
                  :disabled="pagination.isOnFirst()"
                  class="btn btn--default"
                  @click.prevent="goToDirection('prev')"
                >
                  {{ $t('pagination.prev') }}
                </button>
                <button
                  :disabled="pagination.isOnLast()"
                  class="ml-3 btn btn--default"
                  @click.prevent="goToDirection('next')"
                >
                  {{ $t('pagination.next') }}
                </button>
              </div>
            </nav>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Teleport to="#header-search">
    <form
      class="flex w-full md:ml-0"
      method="GET"
    >
      <label
        for="search-field"
        class="sr-only"
      >
        {{ $t('inbox.search') }}
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
          bg-white dark:bg-gray-900 text-gray-900 dark:text-gray-100
          "
          :placeholder="$t('inbox.search')"
          type="search"
        >
      </div>
    </form>
  </Teleport>
</template>

<script lang="ts">
import {ref, onMounted, watch} from 'vue';
import moment from 'moment';
import {
  EyeIcon,
  TrashIcon,
  MagnifyingGlassIcon,
} from '@heroicons/vue/24/outline';
import Pagination from '@/utils/pagination';

export default {
  name: 'Inbox',
  components: {
    TrashIcon,
    EyeIcon,
    MagnifyingGlassIcon,
  },
  setup() {
    const queryParams = ref({
      page: 1,
      per_page: 1,
      search: '',
    });
    const isRequesting = ref(false);
    const emails = ref([]);
    const pagination = ref(new Pagination());

    let searchTimeout: NodeJS.Timeout | null = null;
    watch(() => queryParams.value.search, () => {
      if (searchTimeout) {
        clearTimeout(searchTimeout);
        searchTimeout = null;
      }
      searchTimeout = setTimeout(() => {
        getEmails(1);
      }, 500);
    }, {deep: true});

    const getEmails = (page: number) => {
      isRequesting.value = true;
      queryParams.value.page = page;
      MailHedgehog.request()
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

    const goToPage = (page: number) => {
      if (pagination.value.hasPage(page)) {
        getEmails(page);
      }
    };

    const goToDirection = (direction = 'next') => {
      getEmails(pagination.value.getPageFromDirection(direction));
    };

    return {
      isRequesting,
      emails,
      pagination,
      queryParams,
      goToDirection,
      moment,
    };
  },
};

</script>
