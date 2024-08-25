import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import Components from 'unplugin-vue-components/vite'

const rootPath = new URL('.', import.meta.url).pathname
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
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
