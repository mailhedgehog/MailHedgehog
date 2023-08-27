import {
  computed, inject, Ref, ref,
} from 'vue';
import { useI18n } from 'vue-i18n';
import { useRouter, useRoute } from 'vue-router';
import { SMTPReleaseCredentials } from '@/interfaces/SMTPReleaseCredentials';
import { MailHedgehog } from '@/main';
import { Email } from '@/interfaces/Email';

export function useEmailShare() {
  const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
  const { t } = useI18n();
  const router = useRouter();
  const route = useRoute();

  const isSharingEnabled = computed(() => mailHedgehog?.configValue('sharing.enabled', false));
  const isRequestingShareEmail = ref(false);
  const isOpenShareModal = ref(false);
  const shareLink = ref<string|null>(null);

  const shareForm = ref({
    expiration_in: 1440,
  });

  const shareEmail = (email: Email|null) => {
    if (!email) {
      return;
    }

    isRequestingShareEmail.value = true;
    mailHedgehog?.request()
      .post(`emails/${email?.id}/share`, shareForm.value)
      .then((response) => {
        mailHedgehog?.success(t('email.linkCreated'));

        const shareLinkId = response.data?.data?.id;
        if (!shareLinkId) {
          mailHedgehog?.error(t('response.error'));
          return;
        }

        shareLink.value = `${window.location.origin}/${router.resolve({ name: 'sharedMail', params: { id: shareLinkId } }).href}`;
      })
      .catch(() => {
        mailHedgehog?.error(t('response.error'));
      })
      .finally(() => {
        isRequestingShareEmail.value = false;
      });
  };

  return {
    isSharingEnabled,
    isOpenShareModal,
    isRequestingShareEmail,
    shareForm,
    shareLink,
    shareEmail,
  };
}
