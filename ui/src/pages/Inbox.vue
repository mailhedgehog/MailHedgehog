<template>
  <!-- Page header -->
  <div class="lg:-mt-px bg-white dark:bg-gray-900 shadow dark:shadow-gray-500">
    <div class="px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8">
      <div class="py-6 md:flex md:items-center md:justify-between">
        <div class="min-w-0 flex-1">
          <!-- Profile -->
          <div class="flex items-center">
            <div>
              <div class="flex items-center">
                <h1 class="ml-3 text-2xl font-bold leading-7 text-gray-900 dark:text-gray-100 sm:truncate sm:leading-9">
                  Hi, "vtp"
                </h1>
              </div>
            </div>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3 md:mt-0 md:ml-4">
          <button type="button"
                  class="btn btn--default">
            Add money
          </button>
          <button type="button"
                  class="btn btn--primary">
            Send money
          </button>
        </div>
      </div>
    </div>
  </div>

  <div class="mt-8">
    <h2 class="
          mx-auto mt-8 max-w-6xl px-4
          text-lg font-medium leading-6
          text-gray-900 dark:text-gray-100
          sm:px-6 lg:px-8">
      Emails
    </h2>

    <!-- Activity list (smallest breakpoint only) -->
    <div class="shadow dark:shadow-gray-500 sm:hidden">
      <ul role="list" class="mt-2 divide-y divide-gray-200 overflow-hidden shadow dark:shadow-gray-500 sm:hidden">
        <li v-for="transaction in transactions" :key="transaction.id">
          <a :href="transaction.href" class="block bg-white dark:bg-gray-800 px-4 py-4 hover:bg-gray-50">
                  <span class="flex items-center space-x-4">
                    <span class="flex flex-1 space-x-2 truncate">
                      <BanknotesIcon class="h-5 w-5 flex-shrink-0 text-gray-400 dark:text-gray-300" aria-hidden="true"/>
                      <span class="flex flex-col truncate text-sm text-gray-500 dark:text-gray-400">
                        <span class="truncate">{{ transaction.name }}</span>
                        <span
                        ><span class="font-medium text-gray-900 dark:text-gray-100">{{ transaction.amount }}</span>
                          {{ transaction.currency }}
                        </span
                        >
                        <time :datetime="transaction.datetime">{{ transaction.date }}</time>
                      </span>
                    </span>
                    <ChevronRightIcon class="h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true"/>
                  </span>
          </a>
        </li>
      </ul>

      <nav class="flex items-center justify-between border-t border-gray-200 bg-white dark:bg-gray-800 px-4 py-3"
           aria-label="Pagination">
        <div class="flex flex-1 justify-between">
          <a href="#"
             class="btn btn--default">
            Previous
          </a>
          <a href="#"
             class="btn btn--default">
            Next
          </a>
        </div>
      </nav>
    </div>

    <!-- Activity table (small breakpoint and up) -->
    <div class="hidden sm:block">
      <div class="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
        <div class="mt-2 flex flex-col">
          <div class="min-w-full overflow-hidden overflow-x-auto align-middle shadow sm:rounded-lg">
            <table class="min-w-full divide-y divide-gray-200">
              <thead>
              <tr>
                <th class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100" scope="col">
                  Transaction
                </th>
                <th class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-right text-sm font-semibold text-gray-900 dark:text-gray-100" scope="col">Amount
                </th>
                <th class="hidden bg-gray-50 dark:bg-gray-700 px-6 py-3 text-left text-sm font-semibold text-gray-900 dark:text-gray-100 md:block"
                    scope="col">Status
                </th>
                <th class="bg-gray-50 dark:bg-gray-700 px-6 py-3 text-right text-sm font-semibold text-gray-900 dark:text-gray-100" scope="col">Date
                </th>
              </tr>
              </thead>
              <tbody class="divide-y divide-gray-200 bg-white dark:bg-gray-800">
              <tr v-for="transaction in transactions" :key="transaction.id" class="bg-white dark:bg-gray-800">
                <td class="w-full max-w-0 whitespace-nowrap px-6 py-4 text-sm text-gray-900 dark:text-gray-100">
                  <div class="flex">
                    <a :href="transaction.href" class="group inline-flex space-x-2 truncate text-sm">
                      <BanknotesIcon class="h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500 dark:group-hover:text-gray-100"
                                     aria-hidden="true"/>
                      <p class="truncate text-gray-500 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-gray-100">{{ transaction.name }}</p>
                    </a>
                  </div>
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-right text-sm text-gray-500 dark:text-gray-400">
                  <span class="font-medium text-gray-900 dark:text-gray-100">{{ transaction.amount }}</span>
                  {{ transaction.currency }}
                </td>
                <td class="hidden whitespace-nowrap px-6 py-4 text-sm text-gray-500 md:block">
                      <span
                          :class="[statusStyles[transaction.status], 'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium capitalize']">{{
                          transaction.status
                        }}</span>
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-right text-sm text-gray-500 dark:text-gray-400">
                  <time :datetime="transaction.datetime">{{ transaction.date }}</time>
                </td>
              </tr>
              </tbody>
            </table>
            <!-- Pagination -->
            <nav class="flex items-center justify-between border-t border-gray-200 bg-white dark:bg-gray-800 px-4 py-3 sm:px-6"
                 aria-label="Pagination">
              <div class="hidden sm:block">
                <p class="text-sm text-gray-700 dark:text-gray-200">
                  Showing
                  {{ ' ' }}
                  <span class="font-medium">1</span>
                  {{ ' ' }}
                  to
                  {{ ' ' }}
                  <span class="font-medium">10</span>
                  {{ ' ' }}
                  of
                  {{ ' ' }}
                  <span class="font-medium">20</span>
                  {{ ' ' }}
                  results
                </p>
              </div>
              <div class="flex flex-1 justify-between sm:justify-end">
                <a href="#"
                   class="btn btn--default">Previous</a>
                <a href="#"
                   class="ml-3 btn btn--default">Next</a>
              </div>
            </nav>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Teleport to="#header-search">
    <form class="flex w-full md:ml-0" action="#" method="GET">
      <label for="search-field" class="sr-only">Search</label>
      <div class="relative w-full text-gray-400 focus-within:text-gray-600">
        <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center" aria-hidden="true">
          <MagnifyingGlassIcon class="h-5 w-5" aria-hidden="true"/>
        </div>
        <input id="search-field" name="search-field"
               class="block h-full w-full border-transparent py-2 pl-8 pr-3 bg-white dark:bg-gray-900 text-gray-900 dark:text-gray-100 placeholder-gray-500 focus:border-transparent focus:outline-none focus:ring-0 sm:text-sm"
               placeholder="Search transactions" type="search"/>
      </div>
    </form>
  </Teleport>
</template>

<script lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import HelloWorld from './components/HelloWorld.vue'

import {ref} from 'vue'
import {
  Dialog,
  DialogPanel,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue'
import {
  Bars3CenterLeftIcon,
  BellIcon,
  ClockIcon,
  CogIcon,
  CreditCardIcon,
  DocumentChartBarIcon,
  HomeIcon,
  QuestionMarkCircleIcon,
  ScaleIcon,
  ShieldCheckIcon,
  UserGroupIcon,
  XMarkIcon,
} from '@heroicons/vue/24/outline'
import {
  BanknotesIcon,
  BuildingOfficeIcon,
  CheckCircleIcon,
  ChevronDownIcon,
  ChevronRightIcon,
  MagnifyingGlassIcon,
} from '@heroicons/vue/20/solid'


export default {
  name: "Inbox",
  components: {
    Dialog,
    DialogPanel,
    Menu,
    MenuButton,
    MenuItem,
    MenuItems,
    TransitionChild,
    TransitionRoot,
    Bars3CenterLeftIcon,
    BellIcon,
    ClockIcon,
    CogIcon,
    CreditCardIcon,
    DocumentChartBarIcon,
    HomeIcon,
    QuestionMarkCircleIcon,
    ScaleIcon,
    ShieldCheckIcon,
    UserGroupIcon,
    XMarkIcon,
    BanknotesIcon,
    BuildingOfficeIcon,
    CheckCircleIcon,
    ChevronDownIcon,
    ChevronRightIcon,
    MagnifyingGlassIcon,
  },
  setup() {
    const navigation = [
      {name: 'Inbox', href: '#', icon: HomeIcon, current: true},
    ]
    const secondaryNavigation = [
      {name: 'Settings', href: '#', icon: CogIcon},
    ]
    const transactions = [
      {
        id: 1,
        name: 'Payment to Molly Sanders',
        href: '#',
        amount: '$20,000',
        currency: 'USD',
        status: 'success',
        date: 'July 11, 2020',
        datetime: '2020-07-11',
      },
      // More transactions...
    ]
    const statusStyles = {
      success: 'bg-green-100 text-green-800',
      processing: 'bg-yellow-100 text-yellow-800',
      failed: 'bg-gray-100 text-gray-800',
    }

    const sidebarOpen = ref(false)

    return {
      navigation,
      secondaryNavigation,
      transactions,
      statusStyles,
      sidebarOpen,
    }
  }
}

</script>
