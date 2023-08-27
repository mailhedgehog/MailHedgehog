import {
  computed, inject, Ref, ref,
} from 'vue';
import { useI18n } from 'vue-i18n';
import { SMTPReleaseCredentials } from '@/interfaces/SMTPReleaseCredentials';
import { MailHedgehog } from '@/main';
import { Email } from '@/interfaces/Email';

export function useEmailRelease() {
  const mailHedgehog: MailHedgehog|undefined = inject('MailHedgehog');
  const { t } = useI18n();

  const isRequestingReleaseEmail = ref(false);
  const isOpenReleaseModal = ref(false);
  const releaseForm = ref({
    host: '',
    port: 25,
    username: '',
    password: '',
  });

  const fromLocalStorage = () => {
    const savedSmtpReleaseCredentials = localStorage.getItem('smtpReleaseCredentials');
    if (savedSmtpReleaseCredentials) {
      try {
        const smtpReleaseCredentials: SMTPReleaseCredentials = JSON.parse(atob(savedSmtpReleaseCredentials));
        if (
          typeof smtpReleaseCredentials === 'object'
          && !Array.isArray(smtpReleaseCredentials)
          && smtpReleaseCredentials !== null
        ) {
          if (smtpReleaseCredentials.host) {
            releaseForm.value.host = smtpReleaseCredentials.host;
          }
          if (smtpReleaseCredentials.port) {
            releaseForm.value.port = smtpReleaseCredentials.port;
          }
          if (smtpReleaseCredentials.username) {
            releaseForm.value.username = smtpReleaseCredentials.username;
          }
          if (smtpReleaseCredentials.password) {
            releaseForm.value.password = smtpReleaseCredentials.password;
          }
        }
      } catch (e: Error|any) {
        mailHedgehog?.error(e?.message);
      }
      console.error(releaseForm.value);
    }
  };
  const clearFromLocalStorage = () => {
    localStorage.removeItem('smtpReleaseCredentials');
    mailHedgehog?.success(t('release.deleted'));
  };
  const saveToLocalStorage = () => {
    const smtpReleaseCredentials: SMTPReleaseCredentials = {};
    if (releaseForm.value.host) {
      smtpReleaseCredentials.host = releaseForm.value.host;
    }
    if (releaseForm.value.port) {
      smtpReleaseCredentials.port = releaseForm.value.port;
    }
    if (releaseForm.value.username) {
      smtpReleaseCredentials.username = releaseForm.value.username;
    }
    if (releaseForm.value.password) {
      smtpReleaseCredentials.password = releaseForm.value.password;
    }
    localStorage.setItem('smtpReleaseCredentials', btoa(JSON.stringify(smtpReleaseCredentials)));
    fromLocalStorage();
    mailHedgehog?.success(t('release.saved'));
  };

  const releaseEmail = (email: Email|null) => {
    if (!email) {
      return;
    }

    isRequestingReleaseEmail.value = true;
    mailHedgehog?.request()
      .post(`emails/${email?.id}/release`, releaseForm.value)
      .then(() => {
        isOpenReleaseModal.value = false;
        mailHedgehog?.success(t('email.released'));
      })
      .catch(() => {
        mailHedgehog?.error(t('response.error'));
      })
      .finally(() => {
        isRequestingReleaseEmail.value = false;
      });
  };

  return {
    isOpenReleaseModal,
    releaseForm,
    fromLocalStorage,
    clearFromLocalStorage,
    saveToLocalStorage,
    releaseEmail,
    isRequestingReleaseEmail,
  };
}
