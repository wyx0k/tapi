import {createApp} from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { i18n } from '@/i18n/i18n.js'
import { router } from '@/routers/index.js'
import './style.css';

const app =createApp(App)
app.use(i18n)
app.use(createPinia())
app.use(router)
app.mount('#app')
