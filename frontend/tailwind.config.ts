import { heroui } from '@heroui/theme';

import { Config } from 'tailwindcss';

export default {
  content: [
    './index.html',
    './src/**/*.{js,ts,jsx,tsx}',
    './node_modules/@nextui-org/theme/dist/components/(button|ripple|spinner).js',
    './node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}',
    './node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      borderRadius: {},

      colors: {
        uiLightGray: '#DADBE0',
        uiDeepGray: '#373B40',
        uiLightWhite: '#F4F4F5',
        uiGhostGray: '#71717A',
      },
    },
  },
  plugins: [heroui()],
} satisfies Config;
