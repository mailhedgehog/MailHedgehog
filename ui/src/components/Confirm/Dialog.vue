<script setup lang="ts">
import { computed } from 'vue';
import {
  Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot,
} from '@headlessui/vue';
import { ExclamationTriangleIcon } from '@heroicons/vue/24/outline';
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const store = useStore();
const open = computed(() => store.getters['confirmDialog/show']);
</script>

<template>
  <TransitionRoot
    as="template"
    :show="open"
  >
    <Dialog
      as="div"
      class="relative z-10"
      @close="store.dispatch('confirmDialog/decline')"
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
        <div class="fixed inset-0 bg-context-500 dark:bg-context-900 bg-opacity-75 dark:bg-opacity-75 transition-opacity" />
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
              relative transform overflow-hidden rounded-lg px-4 pt-5 pb-4
              text-left shadow-xl transition-all
              sm:my-8 w-full sm:max-w-lg sm:p-6
              bg-context-50 dark:bg-context-900
            "
            >
              <div class="sm:flex sm:items-start">
                <div class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-primary-100 sm:mx-0 sm:h-10 sm:w-10">
                  <ExclamationTriangleIcon
                    class="h-6 w-6 text-primary-600"
                    aria-hidden="true"
                  />
                </div>
                <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                  <DialogTitle
                    as="h3"
                    class="text-lg font-medium leading-6 text-context-900 dark:text-context-50"
                  >
                    {{ store.getters['confirmDialog/modalTitle'] || t('confirmDialog.modalTitle') }}
                  </DialogTitle>
                  <div
                    class="mt-2"
                  >
                    <p class="text-sm text-context-500 dark:text-context-400">
                      {{ store.getters['confirmDialog/modalText'] || t('confirmDialog.modalText') }}
                    </p>
                  </div>
                </div>
              </div>
              <div class="mt-5 sm:mt-4 sm:flex justify-end">
                <button
                  ref="cancelButtonRef"
                  type="button"
                  class="btn btn--default w-full sm:w-36 justify-center mb-4 sm:mb-0 sm:mr-4"
                  @click="store.dispatch('confirmDialog/decline')"
                >
                  {{ store.getters['confirmDialog/btnCancel'] || t('confirmDialog.btnCancel') }}
                </button>
                <button
                  type="button"
                  class="btn btn--primary w-full sm:w-36 justify-center"
                  @click="store.dispatch('confirmDialog/accept')"
                >
                  {{ store.getters['confirmDialog/btnAccept'] || t('confirmDialog.btnAccept') }}
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
