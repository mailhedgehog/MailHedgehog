module.exports = {
    env: {
        browser: true,
        es2021: true,
    },
    extends: [
        'airbnb-base',
        'plugin:vue/vue3-recommended',
    ],
    parserOptions: {
        ecmaVersion: 12,
        sourceType: 'module',
        parser: '@typescript-eslint/parser',
    },
    plugins: [
        'vue',
        '@typescript-eslint',
    ],
    rules: {
        'no-param-reassign': [2, {props: false}],
        'max-len': 'off',
        'vue/no-v-html': 0,
        'vue/max-len': [
            'error',
            {
                code: 120,
                template: 160,
                ignoreTemplateLiterals: true,
                ignoreUrls: true,
                ignoreStrings: true,
            },
        ],
        'import/extensions': [
            'error',
            'ignorePackages',
            {
                js: 'never',
                jsx: 'never',
                ts: 'never',
                tsx: 'never',
            },
        ],
    },
    globals: {
        MailHedgehog: true,
    },
    settings: {},
};
