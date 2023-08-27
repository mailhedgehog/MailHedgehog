<script setup lang="ts">
import {
  Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot,
} from '@headlessui/vue';
import { ArchiveBoxArrowDownIcon, ArchiveBoxXMarkIcon } from '@heroicons/vue/24/outline';
import { useI18n } from 'vue-i18n';
import { onMounted } from 'vue';
import { useEmailRelease } from '@/composables/email-release';
import { Email } from '@/interfaces/Email';

const { t } = useI18n();

interface Props {
  email: Email|null
  show: boolean
  requesting: boolean
}

const props = withDefaults(defineProps<Props>(), {
  email: null,
  show: false,
  requesting: false,
});

const emit = defineEmits(['close']);

const {
  releaseForm,
  clearFromLocalStorage,
  fromLocalStorage,
  saveToLocalStorage,
  releaseEmail,
} = useEmailRelease();

onMounted(() => {
  fromLocalStorage();
});

</script>

<template>
  <TransitionRoot
    as="template"
    :show="show"
  >
    <Dialog
      as="div"
      class="relative z-10"
      :class="{'pointer-events-none': requesting}"
      @close="emit('close')"
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
              <div>
                <div class="text-center sm:mt-5">
                  <DialogTitle
                    as="h3"
                    class="text-lg font-medium leading-6 text-context-900 dark:text-context-100"
                  >
                    {{ t('email.releaseHint') }}
                  </DialogTitle>
                  <div
                    class="
                    mt-6
                    text-left
                    text-context-900 dark:text-context-100
                    "
                  >
                    <div
                      class="space-y-6"
                    >
                      <div>
                        <label
                          for="host"
                          class="form-label"
                        >
                          {{ t('release.host') }}
                        </label>
                        <div class="mt-1 flex">
                          <input
                            id="host"
                            v-model="releaseForm.host"
                            name="host"
                            type="text"
                            autocomplete="host"
                            required
                            class="form-input rounded-r-none"
                            placeholder="smtp.provider.com"
                          >
                          <input
                            id="port"
                            v-model="releaseForm.port"
                            v-tooltip="t('release.port')"
                            name="port"
                            type="number"
                            min="1"
                            max="9999"
                            step="1"
                            autocomplete="port"
                            required
                            class="form-input rounded-l-none w-20"
                            placeholder="25"
                          >
                        </div>
                      </div>
                      <div>
                        <label
                          for="smtp-username"
                          class="form-label"
                        >
                          {{ t('release.auth') }}
                        </label>
                        <div class="mt-1">
                          <input
                            id="smtp-username"
                            v-model="releaseForm.username"
                            name="smtp-username"
                            type="text"
                            autocomplete="smtp-username"
                            class="form-input rounded-b-none"
                            :placeholder="t('release.username')"
                          >
                          <input
                            id="smtp-password"
                            v-model="releaseForm.password"
                            name="smtp-password"
                            type="password"
                            autocomplete="smtp-password"
                            class="form-input rounded-t-none"
                            :placeholder="t('release.password')"
                          >
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
                  @click="emit('close')"
                >
                  {{ t('email.cancel') }}
                </button>
                <div class="flex space-x-2 mx-2">
                  <a
                    v-tooltip="t('release.saveToStorage')"
                    href="#"
                    class="font-medium text-primary-600 hover:text-primary-500"
                    @click.prevent="saveToLocalStorage"
                  >
                    <ArchiveBoxArrowDownIcon class="h-6 w-6" />
                  </a>
                  <a
                    v-tooltip="t('release.deleteSavedToStorage')"
                    href="#"
                    class="font-medium text-primary-600 hover:text-primary-500"
                    @click.prevent="clearFromLocalStorage"
                  >
                    <ArchiveBoxXMarkIcon class="h-6 w-6" />
                  </a>
                </div>
                <button
                  type="button"
                  class="btn btn--primary"
                  @click="releaseEmail(email)"
                >
                  {{ t('email.release') }}
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
