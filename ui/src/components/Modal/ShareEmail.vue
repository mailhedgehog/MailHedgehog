<script setup lang="ts">
import {
  Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot,
} from '@headlessui/vue';
import {
  ChevronUpDownIcon,
  ClipboardDocumentListIcon,
} from '@heroicons/vue/24/outline';
import { useI18n } from 'vue-i18n';
import { inject } from 'vue';
import { Email } from '@/interfaces/Email';
import { useEmailShare } from '@/composables/email-share';
import { MailHedgehog } from '@/main';

const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
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
  shareForm,
  shareLink,
  shareEmail,
} = useEmailShare();

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
                    {{ t('sharedEmail.modalTitle') }}
                  </DialogTitle>
                  <div
                    class="
                    mt-6
                    text-left
                    text-context-900 dark:text-context-100
                    "
                  >
                    <div v-if="shareLink">
                      <label
                        for="shareLink"
                        class="form-label"
                      >
                        {{ t('sharedEmail.form.shareLink.label') }}
                      </label>
                      <div class="mt-1 flex">
                        <input
                          id="shareLink"
                          v-model="shareLink"
                          readonly
                          name="shareLink"
                          type="text"
                          autocomplete="off"
                          required
                          class="form-input rounded-r-none"
                        >
                        <a
                          class="cursor-pointer form-input border-l-0 rounded-l-none w-auto
                        text-context-400 dark:text-context-600 hover:text-primary-600 transition-all duration-500"
                          @click.prevent="mailHedgehog?.copyToClipboard(shareLink as string)"
                        >
                          <ClipboardDocumentListIcon
                            class="h-6 flex-shrink-0  pointer-events-none"
                          />
                        </a>
                      </div>
                    </div>
                    <div v-else>
                      <label
                        for="expiration_in"
                        class="form-label"
                      >
                        {{ t('sharedEmail.form.expiration.label') }}
                      </label>
                      <div class="mt-1 relative">
                        <select
                          id="expiration_in"
                          v-model="shareForm.expiration_in"
                          name="expiration_in"
                          class="form-input cursor-pointer"
                        >
                          <option :value="1">
                            {{ t('sharedEmail.form.expiration.1h') }}
                          </option>
                          <option :value="3">
                            {{ t('sharedEmail.form.expiration.3h') }}
                          </option>
                          <option :value="24">
                            {{ t('sharedEmail.form.expiration.1d') }}
                          </option>
                          <option :value="168">
                            {{ t('sharedEmail.form.expiration.7d') }}
                          </option>
                          <option :value="720">
                            {{ t('sharedEmail.form.expiration.30d') }}
                          </option>
                        </select>
                        <ChevronUpDownIcon
                          class="h-6 flex-shrink-0 text-context-400 dark:text-context-600 pointer-events-none
                              absolute top-1/2 right-2 -translate-y-1/2
                            "
                          aria-hidden="true"
                        />
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
                <div class="px-2" />
                <button
                  type="button"
                  class="btn btn--primary"
                  :disabled="!!shareLink"
                  @click="shareEmail(email)"
                >
                  {{ t('sharedEmail.form.submitButton.label') }}
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
