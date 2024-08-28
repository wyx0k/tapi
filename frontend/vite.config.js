import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import Components from 'unplugin-vue-components/vite'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

const rootPath = new URL('.', import.meta.url).pathname
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
      createSvgIconsPlugin({
          iconDirs: [path.resolve(process.cwd(), 'src/icons')],
          symbolId: 'icon-[name]',
      }),
      vue(),
      Components({
        resolvers: [NaiveUiResolver()],
      })
  ],

  resolve: {
    alias: {
      '@': rootPath + 'src',
      stores: rootPath + 'src/stores',
      wailsjs: rootPath + 'wailsjs',
    },
  },
})
