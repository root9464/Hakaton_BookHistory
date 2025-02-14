import svgr from '@svgr/rollup';
import tailwindcss from '@tailwindcss/vite';
import { TanStackRouterVite } from '@tanstack/router-plugin/vite';
import react from '@vitejs/plugin-react-swc';
import { defineConfig } from 'vite';

export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
    TanStackRouterVite(),
    svgr({
      include: '**/*.svg',

      svgoConfig: {
        floatPrecision: 2,
      },

      typescript: true,
      ref: true,
      memo: true,
      svgProps: {
        ref: 'ref',
      },
      prettierConfig: {
        parser: 'typescript',
      },
    }),
  ],
});
