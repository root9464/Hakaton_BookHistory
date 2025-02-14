import { default as js, default as pluginJs } from '@eslint/js';
import tanstack from '@tanstack/eslint-plugin-query';
import pluginReact from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import globals from 'globals';
import tseslint from 'typescript-eslint';

export default tseslint.config(
  { ignores: ['dist'] },
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ['**/*.{ts,tsx}'],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node,
      },
      ecmaVersion: 2023,
      parserOptions: {
        sourceType: 'module',
        ecmaFeatures: {
          jsx: true,
        },
      },
      parser: '@typescript-eslint/parser',
    },
  },

  // expoted rules
  {
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      '@tanstack/query': tanstack,
      '@typescript-eslint': tseslint.plugin,
      react: pluginReact,
    },

    rules: {
      ...reactHooks.configs.recommended.rules,
      ...tanstack.configs.recommended.rules,
      ...reactRefresh.configs.recommended.rules,
    },
  },
  // all rules
  {
    rules: {
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }],
      'react/react-in-jsx-scope': 'off',
      '@typescript-eslint/naming-convention': 'off',
      'no-unneeded-ternary': ['error', { defaultAssignment: false }],
     
    },
  },

  pluginJs.configs.recommended,
  pluginReact.configs.flat['jsx-runtime'],
  ...tseslint.configs.recommended,
);