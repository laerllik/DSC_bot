import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';
import prettier from 'eslint-config-prettier';
import prettierPlugin from 'eslint-plugin-prettier';

export default tseslint.config(
  // Игнорируем файлы и папки
  {
    ignores: [
      'node_modules',
      'dist',
      'build',
      'coverage',
      '*.log',
      '.env',
      '.env.local',
      '.DS_Store'
    ]
  },

  eslint.configs.recommended,

  ...tseslint.configs.recommended.map(config => ({
    ...config,
    files: ['**/*.ts', '**/*.tsx'],
  })),

  ...tseslint.configs.strict.map(config => ({
    ...config,
    files: ['**/*.ts', '**/*.tsx'],
  })),

  {
    plugins: {
      prettier: prettierPlugin
    },
    rules: {
      'prettier/prettier': 'error',
      'no-console': 'warn',
      'no-debugger': 'warn',
      'eqeqeq': ['error', 'always'],
      'curly': 'error'
    }
  },

  // Отключаем конфликтующие с Prettier правила для всех файлов
  prettier
);
