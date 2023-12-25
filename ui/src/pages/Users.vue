<script setup lang="ts">
import {
  ref, onMounted, watch, computed, inject,
} from 'vue';
// FIXME: locales not works
import moment from 'moment';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import {
  PencilIcon,
  UserPlusIcon,
  TrashIcon,
  PlusIcon,
  MagnifyingGlassIcon,
} from '@heroicons/vue/24/outline';
import {
  Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot,
} from '@headlessui/vue';
import Pagination from '@/utils/pagination';
import { MailHedgehog } from '@/main';
import { User } from '@/interfaces/User';

const { t } = useI18n();
const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
const router = useRouter();
const store = useStore();

const queryParams = ref({
  page: 1,
  per_page: 25,
  search: '',
});
const mounted = ref(false);
const isRequesting = ref(false);
const users = ref<Array<User>>([]);
const pagination = ref(new Pagination());

const getUsers = (page: number|null = null) => {
  isRequesting.value = true;
  if (page) {
    queryParams.value.page = page;
  }
  mailHedgehog?.request()
    .get('users', {
      params: queryParams.value,
    })
    .then((response) => {
      if (response.data?.data) {
        users.value = response.data?.data;
      } else {
        users.value = [];
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

let searchTimeout: NodeJS.Timeout|null = null;
watch(() => queryParams.value.search, () => {
  if (searchTimeout) {
    clearTimeout(searchTimeout);
    searchTimeout = null;
  }
  searchTimeout = setTimeout(() => {
    getUsers(1);
  }, 500);
}, { deep: true });

onMounted(() => {
  getUsers(1);
  mounted.value = true;
});

const goToDirection = (direction = 'next') => {
  getUsers(pagination.value.getPageFromDirection(direction));
};

const editingUser = ref<User|null>(null);

const userForm = ref({
  new_username: null as null|string,
  dashboard_password: null as null|string,
  smtp_password: null as null|string,
  smtp_auth_ips: [] as string[],
  smtp_allow_listed_ips: [] as string[],
  dashboard_auth_emails: [] as string[],
});

const editUser = (user:User) => {
  userForm.value.smtp_auth_ips = user.smtp_auth_ips;
  userForm.value.smtp_allow_listed_ips = user.smtp_allow_listed_ips;
  userForm.value.dashboard_auth_emails = user.dashboard_auth_emails;

  editingUser.value = user;
};

const openCreateUserModal = () => {
  editingUser.value = {
    username: '',
    smtp_auth_ips: [],
    smtp_allow_listed_ips: [],
    dashboard_auth_emails: [],
  };
};

const updateOrCreateUser = () => {
  isRequesting.value = true;
  const request = mailHedgehog?.request();
  let promise;
  if (editingUser.value?.username) {
    promise = request?.put(`users/${editingUser.value.username}`, {
      dashboard_password: userForm.value.dashboard_password,
      smtp_password: userForm.value.smtp_password,
      smtp_auth_ips: userForm.value.smtp_auth_ips,
      smtp_allow_listed_ips: userForm.value.smtp_allow_listed_ips,
      dashboard_auth_emails: userForm.value.dashboard_auth_emails,
    });
  } else {
    promise = request?.post('users', {
      username: userForm.value.new_username,
      dashboard_password: userForm.value.dashboard_password,
      smtp_password: userForm.value.smtp_password,
      smtp_auth_ips: userForm.value.smtp_auth_ips,
      smtp_allow_listed_ips: userForm.value.smtp_allow_listed_ips,
      dashboard_auth_emails: userForm.value.dashboard_auth_emails,
    });
  }
  promise?.then((response) => {
    getUsers();
    if (editingUser.value?.username) {
      mailHedgehog?.success(t('users.updated'));
    } else {
      mailHedgehog?.success(t('users.created'));
    }
    closeModal();
  })
    .catch((error) => {
      mailHedgehog?.onResponseError(error, 'Response Error');
    })
    .finally(() => {
      isRequesting.value = false;
    });
};

const closeModal = () => {
  editingUser.value = null;
  userForm.value = {
    new_username: null,
    dashboard_password: null,
    smtp_password: null,
    smtp_auth_ips: [],
    smtp_allow_listed_ips: [],
    dashboard_auth_emails: [],
  };
};

const deleteUser = (user:User) => {
  store.dispatch('confirmDialog/confirm')
    .then(() => {
      isRequesting.value = true;
      mailHedgehog?.request()
        .delete(`users/${user.username}`)
        .then(() => {
          if (pagination.value.count() === 0) {
            goToDirection('prev');
          } else {
            getUsers();
          }
          mailHedgehog?.success(t('users.deleted'));
        })
        .catch((error) => {
          mailHedgehog.onResponseError(error, 'Response Error');
        })
        .catch(() => {
          isRequesting.value = false;
        });
    });
};

</script>

<template>
  <!-- Page header -->
  <div class="lg:-mt-px bg-context-50 dark:bg-context-900 shadow dark:shadow-context-500">
    <div class="px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8">
      <div class="py-6 flex items-center justify-between">
        <h1
          class="ml-3 text-xl md:text-2xl font-bold leading-7 text-context-900 dark:text-context-100 truncate sm:leading-9"
        >
          {{ t('users.pageTitle') }}
        </h1>
        <div
          class="flex justify-end ml-4 transition-all duration-200"
          :class="{
            'pointer-events-none opacity-75': isRequesting
          }"
        >
          <button
            v-if="pagination && !pagination.isEmpty()"
            v-tooltip="t('users.create')"
            type="button"
            class="btn btn--default whitespace-nowrap"
            @click.prevent="openCreateUserModal"
          >
            <UserPlusIcon class="w-4 h-4 md:mr-2" />
            <span
              class="hidden md:inline"
            >
              {{ t('users.create') }}
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
      <h3 class="ml-3 text-lg md:text-xl text-center text-context-900 dark:text-context-100 select-none">
        <template v-if="isRequesting">
          {{ t('pagination.requesting') }}
        </template>
        <template v-else>
          {{ t('users.empty') }}
        </template>
      </h3>
    </div>
    <template v-else>
      <!-- List (smallest breakpoint only) -->
      <div
        v-if="pagination"
        class="shadow dark:shadow-context-500 md:hidden"
        :class="{
          'pointer-events-none opacity-75': isRequesting
        }"
      >
        <ul
          role="list"
          class="mt-2 divide-y divide-context-200 overflow-hidden shadow dark:shadow-context-500"
        >
          <li
            v-for="user in users"
            :key="user.username"
          >
            <div
              class="block bg-context-50 dark:bg-context-800 px-4 py-4 hover:bg-context-50"
            >
              <span class="flex items-center space-x-4">
                <span class="flex flex-1 space-x-2 truncate">
                  <span class="flex flex-col truncate text-sm text-context-500 dark:text-context-400">
                    <span
                      v-if="user.username"
                      class="truncate"
                    >
                      <span class="font-medium text-context-900 dark:text-context-100 mr-2">
                        {{ t('users.username') }}
                      </span>
                      {{ user.username }}
                    </span>
                  </span>
                </span>
                <div class="space-y-2">
                  <a
                    class="
                        cursor-pointer block
                        transition-all duration-500
                        text-context-500 dark:text-context-400
                        hover:text-context-700 hover:dark:text-context-300
                      "
                    @click.prevent="editUser(user)"
                  >
                    <PencilIcon
                      class="h-5 w-5 flex-shrink-0"
                      aria-hidden="true"
                    />
                  </a>
                  <a
                    class="
                        cursor-pointer block
                        transition-all duration-500
                        text-context-500 dark:text-context-400
                        hover:text-context-700 hover:dark:text-context-300
                      "
                    @click.prevent="deleteUser(user)"
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
          class="flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3"
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
              <table class="min-w-full divide-y divide-context-200">
                <thead>
                  <tr>
                    <th
                      class="bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100"
                      scope="col"
                    >
                      {{ t('users.username') }}
                    </th>
                    <th
                      class="bg-context-50 dark:bg-context-700 px-6 py-3 text-sm font-semibold text-context-900 dark:text-context-100 text-right"
                      scope="col"
                    >
                      {{ t('pagination.actions') }}
                    </th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-context-200 bg-context-50 dark:bg-context-800">
                  <tr
                    v-for="user in users"
                    :key="user.username"
                    class="bg-context-50 dark:bg-context-800"
                  >
                    <td class="max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-context-900 dark:text-context-100">
                      <div class="truncate">
                        {{ user.username }}
                      </div>
                    </td>
                    <td
                      class="whitespace-nowrap px-6 py-4 text-sm text-right flex justify-end space-x-1"
                    >
                      <a
                        class="
                        cursor-pointer
                        transition-all duration-500
                        text-context-500 dark:text-context-400
                        hover:text-context-700 hover:dark:text-context-300
                      "
                        @click.prevent="editUser(user)"
                      >
                        <PencilIcon
                          class="w-5 h-5"
                          aria-hidden="true"
                        />
                      </a>
                      <a
                        class="
                        cursor-pointer
                        transition-all duration-500
                        text-context-500 dark:text-context-400
                        hover:text-context-700 hover:dark:text-context-300
                      "
                        @click.prevent="deleteUser(user)"
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
                class="flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3 sm:px-6"
                aria-label="Pagination"
              >
                <div class="hidden sm:block">
                  <p
                    class="text-sm text-context-700 dark:text-context-200"
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
  <TransitionRoot
    as="template"
    :show="editingUser !== null"
  >
    <Dialog
      as="div"
      class="relative z-10"
      :class="{'pointer-events-none': isRequesting}"
      @close="closeModal"
    >
      <TransitionChild
        as="template"
        enter="ease-out duration-300"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="ease-in duration-200"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
      </TransitionChild>

      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
          <TransitionChild
            as="template"
            enter="ease-out duration-300"
            enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            enter-to="opacity-100 translate-y-0 sm:scale-100"
            leave="ease-in duration-200"
            leave-from="opacity-100 translate-y-0 sm:scale-100"
            leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
          >
            <DialogPanel
              class="
            bg-context-50 dark:bg-context-900
            relative transform overflow-hidden
            rounded-lg px-4 pt-5 pb-4
            text-left shadow-xl transition-all
            sm:my-8 sm:w-full sm:max-w-lg sm:p-6"
            >
              <form @submit.prevent="updateOrCreateUser">
                <div>
                  <div class="text-center sm:mt-5">
                    <DialogTitle
                      as="h3"
                      class="text-lg font-medium leading-6 text-context-900 dark:text-context-100"
                    >
                      <template v-if="editingUser?.username === ''">
                        {{ t('users.modal.createTitle') }}
                      </template>
                      <template v-else>
                        {{ t('users.modal.editTitle', {user: editingUser?.username}) }}
                      </template>
                    </DialogTitle>
                    <div
                      class="
                    block
                    mt-6
                    text-left
                    text-context-900 dark:text-context-100
                    "
                    >
                      <div
                        class="space-y-6"
                      >
                        <div v-if="editingUser?.username === ''">
                          <label
                            for="new_username"
                            class="form-label"
                          >
                            {{ t('users.username') }}
                          </label>
                          <div class="mt-1 flex">
                            <input
                              id="new_username"
                              v-model="userForm.new_username"
                              name="new_username"
                              type="text"
                              autocomplete="off"
                              :required="editingUser.username === ''"
                              class="form-input"
                              :placeholder="t('users.username')"
                            >
                          </div>
                        </div>
                        <div>
                          <label
                            for="hub_password"
                            class="form-label"
                          >
                            {{ t('users.hubPassword') }}
                          </label>
                          <div class="mt-1 flex">
                            <input
                              id="hub_password"
                              v-model="userForm.dashboard_password"
                              name="hub_password"
                              type="password"
                              autocomplete="hub_password"
                              :required="editingUser?.username === ''"
                              class="form-input"
                              :placeholder="t('users.hubPassword')"
                            >
                          </div>
                          <div class="form-hint">
                            <template v-if="editingUser?.username === ''" />
                            <template v-else>
                              {{ t('users.emptyPasswordHint') }}
                            </template>
                          </div>
                        </div>
                        <div>
                          <label
                            for="smtp_password"
                            class="form-label"
                          >
                            {{ t('users.smtpPassword') }}
                          </label>
                          <div class="mt-1 flex">
                            <input
                              id="smtp_password"
                              v-model="userForm.smtp_password"
                              name="smtp_password"
                              type="password"
                              autocomplete="smtp_password"
                              class="form-input"
                              :placeholder="t('users.smtpPassword')"
                            >
                          </div>
                          <div class="form-hint">
                            <template v-if="editingUser?.username === ''">
                              {{ t('users.emptySmtpPasswordHint') }}
                            </template>
                            <template v-else>
                              {{ t('users.emptyPasswordHint') }}
                            </template>
                          </div>
                        </div>
                        <div>
                          <div class="flex justify-between">
                            <label
                              for="smtp_auth_ips"
                              class="form-label"
                            >
                              {{ t('users.smtpNoPassIps') }}
                            </label>
                            <a
                              v-tooltip="t('users.addIPHint')"
                              class="
                                  cursor-pointer
                                  transition-all duration-500
                                  text-context-500 dark:text-context-400
                                  hover:text-primary-500 hover:dark:text-primary-500
                                  flex justify-center items-center w-auto
                                  px-3
                                "
                              @click.prevent="userForm?.smtp_auth_ips?.push('127.0.0.1')"
                            >
                              <PlusIcon
                                class="h-5 w-5 flex-shrink-0"
                                aria-hidden="true"
                              />
                            </a>
                          </div>
                          <div class="mt-1 flex flex-col -space-y-px">
                            <div
                              v-for="(noPassIP, i) in userForm.smtp_auth_ips"
                              :key="i"
                              class="flex form-inputs-list-item"
                            >
                              <input
                                :id="`smtp_auth_ips_${i}`"
                                v-model="userForm.smtp_auth_ips[i]"
                                name="smtp_auth_ips[]"
                                type="text"
                                autocomplete="off"
                                class="form-input"
                                :placeholder="t('users.ipPlaceholder')"
                                maxlength="39"
                                required
                              >
                              <a
                                class="
                                  cursor-pointer
                                  transition-all duration-500
                                  text-context-500 dark:text-context-400
                                  hover:text-context-700 hover:dark:text-context-300
                                  flex justify-center items-center form-input w-auto
                                  border-l-0
                                "
                                @click.prevent="userForm.smtp_auth_ips.splice(i, 1)"
                              >
                                <TrashIcon
                                  class="h-5 w-5 flex-shrink-0"
                                  aria-hidden="true"
                                />
                              </a>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="mt-6 sm:mt-8 flex items-center justify-between">
                  <button
                    type="button"
                    class="btn btn--default"
                    @click.prevent="closeModal"
                  >
                    {{ t('users.modal.cancel') }}
                  </button>
                  <button
                    type="submit"
                    class="btn btn--primary"
                  >
                    {{ t('users.modal.submit') }}
                  </button>
                </div>
              </form>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
  <Teleport
    v-if="mounted"
    to="#header-search"
  >
    <form
      class="flex w-full md:ml-0"
      method="GET"
      @submit.prevent="getUsers(1)"
    >
      <label
        for="search-field"
        class="sr-only"
      >
        {{ t('inbox.search') }}
      </label>
      <div class="relative w-full text-context-400 focus-within:text-context-600">
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
          placeholder-context-500
          bg-context-50 dark:bg-context-900 text-context-900 dark:text-context-100
          "
          :placeholder="t('users.search')"
          type="search"
        >
      </div>
    </form>
  </Teleport>
</template>
