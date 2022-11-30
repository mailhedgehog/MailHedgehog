<template>
  <!-- {{ $route.params.id }}  -->
  <!-- Mail actions -->
  <div class="lg:-mt-px bg-white dark:bg-gray-900 shadow dark:shadow-gray-500 mb-1">
    <div class="px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8">
      <div class="py-6 md:flex md:items-center md:justify-between">
        <div class="min-w-0 flex-1">
          <!-- Profile -->
          <div class="flex items-center">
            <div>
              <div class="flex items-center space-x-3">
                <button
                  type="button"
                  class="btn btn--default"
                >
                  Add money
                </button>
                <button
                  type="button"
                  class="btn btn--primary"
                >
                  Send money
                </button>
              </div>
            </div>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3 md:mt-0 md:ml-4">
          <div
            class="rounded-full bg-white dark:bg-gray-900 p-1 text-gray-400 dark:text-gray-600 hover:text-gray-500 hover:dark:text-gray-500 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
          >
            <span class="sr-only">View notifications</span>
            <BellIcon
              class="h-6 w-6"
              aria-hidden="true"
            />
          </div>
          <Menu
            as="div"
            class="relative ml-3"
          >
            <div>
              <MenuButton
                class="rounded-full bg-white dark:bg-gray-900 p-1 text-gray-400 dark:text-gray-600 hover:text-gray-500 hover:dark:text-gray-500 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
              >
                <span class="sr-only">View notifications</span>
                <BellIcon
                  class="h-6 w-6"
                  aria-hidden="true"
                />
              </MenuButton>
            </div>
            <transition
              enter-active-class="transition ease-out duration-100"
              enter-from-class="transform opacity-0 scale-95"
              enter-to-class="transform opacity-100 scale-100"
              leave-active-class="transition ease-in duration-75"
              leave-from-class="transform opacity-100 scale-100"
              leave-to-class="transform opacity-0 scale-95"
            >
              <MenuItems
                class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white dark:bg-gray-700 py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
              >
                <MenuItem v-slot="{ active }">
                  <a
                    href="#"
                    :class="[active ? 'bg-gray-100 dark:bg-gray-800' : '', 'block px-4 py-2 text-sm text-gray-700 dark:text-gray-400']"
                  >Your
                    Profile</a>
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <a
                    href="#"
                    :class="[active ? 'bg-gray-100 dark:bg-gray-800' : '', 'block px-4 py-2 text-sm text-gray-700 dark:text-gray-400']"
                  >Settings</a>
                </MenuItem>
                <MenuItem v-slot="{ active }">
                  <a
                    href="#"
                    :class="[active ? 'bg-gray-100 dark:bg-gray-800' : '', 'block px-4 py-2 text-sm text-gray-700 dark:text-gray-400']"
                  >Logout</a>
                </MenuItem>
              </MenuItems>
            </transition>
          </Menu>
        </div>
      </div>
    </div>
  </div>
  <!-- Mail headers -->
  <div>
    <h3 class="font-medium text-gray-900">
      Information
    </h3>
    <dl class="mt-2 divide-y divide-gray-200 border-t border-b border-gray-200">
      <div
        v-for="key in Object.keys(currentFile.information)"
        :key="key"
        class="flex justify-between py-3 text-sm font-medium"
      >
        <dt class="text-gray-500">
          {{ key }}
        </dt>
        <dd class="whitespace-nowrap text-gray-900">
          {{ currentFile.information[key] || '' }}
        </dd>
      </div>
    </dl>
  </div>
  <!-- Mail Tabs -->
</template>

<script>
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
  BellIcon,
} from '@heroicons/vue/24/outline';

export default {
  name: 'Message',
  components: {
    Dialog,
    DialogPanel,
    Menu,
    MenuButton,
    MenuItem,
    MenuItems,
    TransitionChild,
    TransitionRoot,
    BellIcon,
  },
  setup() {
    const currentFile = {
      name: 'IMG_4985.HEIC',
      size: '3.9 MB',
      source:
          'https://images.unsplash.com/photo-1582053433976-25c00369fc93?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=512&q=80',
      information: {
        'Uploaded by': 'Marie Culver',
        Created: 'June 8, 2020',
        'Last modified': 'June 8, 2020',
        Dimensions: '4032 x 3024',
        Resolution: '72 x 72',
      },
      sharedWith: [
        {
          id: 1,
          name: 'Aimee Douglas',
          imageUrl:
              'https://images.unsplash.com/photo-1502685104226-ee32379fefbe?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=3&w=1024&h=1024&q=80',
        },
        {
          id: 2,
          name: 'Andrea McMillan',
          imageUrl:
              'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixqx=oilqXxSqey&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80',
        },
      ],
    };

    return {
      currentFile,
    };
  },
};
</script>
