export default {
  app: {
    title: 'MailHedgehog',
  },
  menu: {
    inbox: 'Boîte de réception',
    logout: 'Se déconnecter',
    users: 'Utilisateurs',
  },
  sidebar: {
    open: 'Ouvrir la barre latérale',
    close: 'Fermer la barre latérale',
  },
  response: {
    error: 'Erreur de réponse du serveur',
  },
  colorMode: {
    title: 'Mode de couleur',
    light: 'Clair',
    dark: 'Foncé',
    system: 'Système',
  },
  locales: {
    title: 'Langue',
    en: 'Anglais',
    uk: 'Ukrainien',
    fr: 'Français',
  },
  clipboard: {
    success: 'Copié dans le presse-papier',
    error: 'Erreur lors du copié dans le presse-papiers',
  },
  pagination: {
    requesting: 'Chargement...',
    prev: 'Précédent',
    next: 'Prochain',
    text: 'Affichage <span class="font-medium">{from}</span> à <span class="font-medium">{to}</span> sur <span class="font-medium">{of}</span> résultats',
    actions: 'Actions',
  },
  email: {
    notAvailable: 'n/a',
    headersTitle: 'En-têtes d\'e-mails',
    headersSubtitle: 'Un en-tête d\'e-mail est une métadonnée qui accompagne chaque e-mail et contient des informations détaillées',
    hintHideHeaders: 'Afficher uniquement les en-têtes importants',
    hintShowHeaders: 'Afficher tous les en-têtes',
    from: 'De',
    to: 'À',
    subject: 'Objet',
    received_at: 'Reçu à',
    delete: 'Effacer',
    deleteHint: 'Effacer e-mail',
    deleted: 'E-mail supprimé',
    download: 'Télécharger',
    downloadHint: 'Télécharger l\'e-mail en tant que fichier',
    share: 'Partager',
    shareHint: 'Autoriser l\'accès aux e-mails par lien de partage',
    release: 'Transfér',
    releaseHint: 'Transfér l\'e-mail',
    released: 'E-mail transféré',
    linkCreated: 'Lien créé',
    cancel: 'Annuler',
    back: 'Retour',
    tab: {
      html: 'HTML',
      plain: 'Texte brut',
      source: 'Source',
      attachments: 'Pièces jointes',
    },
    htmlEmpty: 'L\'e-mail ne contient pas de contenu HTML',
    plainEmpty: 'L\'e-mail ne contient pas de contenu en texte brut',
    noAttachments: 'L\'e-mail n\'a pas de pièces jointes',
  },
  sharedEmail: {
    notFound: 'E-mail partagé introuvable ou lien de partage expiré.',
    modalTitle: 'Partager',
    form: {
      expiration: {
        label: 'Expiration dans',
        '10m': '10 minutes',
        '30m': '30 minutes',
        '1h': '1 heure',
        '3h': '3 heures',
        '1d': 'Un jour',
        '7d': '7 jours',
      },
      shareLink: {
        label: 'Lien',
      },
      submitButton: {
        label: 'Partager',
      },
    },
  },
  inbox: {
    hello: 'Bonjour, "{msg}"',
    pageTitle: 'Boîte de réception',
    clear: 'Vider la boîte de réception',
    cleared: 'Boîte de réception effacée',
    search: 'Rechercher des e-mails',
    empty: 'Aucun e-mail trouvé',
  },
  release: {
    host: 'Host',
    port: 'Port',
    auth: 'Identifiants',
    username: 'Nom d\'utilisateur',
    password: 'Mot de passe',
    saveToStorage: 'Enregistrer les informations d\'identification sur le stockage local (non sécurisé)',
    saved: 'Enregistré',
    deleteSavedToStorage: 'Supprimer les informations d\'identification enregistrées',
    deleted: 'Supprimé',
  },
  confirmDialog: {
    modalTitle: 'Êtes-vous sûr?',
    modalText: 'Voulez-vous exécuter cette action ?',
    btnCancel: 'No',
    btnAccept: 'Oui',
  },
  login: {
    title: 'Se connecter à votre compte',
    form: {
      username: {
        label: 'Username',
      },
      password: {
        label: 'Mot de passe',
      },
      login: {
        label: 'S\'identifier',
      },
    },
  },
  users: {
    pageTitle: 'Utilisateurs',
    search: 'Rechercher des utilisateurs',
    empty: 'Aucun utilisateur trouvé',
    create: 'Ajouter un utilisateur',
    deleted: 'Utilisateur supprimé',
    updated: 'Utilisateur mis à jour',
    created: 'Utilisateur créé',
    username: 'Nom d\'utilisateur',
    hubPassword: 'UI mot de passe',
    smtpPassword: 'SMTP mot de passe',
    emptySmtpPasswordHint: 'Laissé vide pour définir le même mot de passe que UI',
    emptyPasswordHint: 'Laissé vide pour conserver l\'ancien mot de passe',
    modal: {
      createTitle: 'Créer un utilisateur',
      editTitle: 'Modifier l\'utilisateur \'{user}\'',
      cancel: 'Annuler',
      submit: 'Soumettre',
    },
  },
};
