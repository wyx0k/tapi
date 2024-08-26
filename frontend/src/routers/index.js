import { createMemoryHistory, createRouter } from 'vue-router'

import Api from '@/components/main/api/Api.vue'

const routes = [
    {path: '/',redirect:'/api'},
    { path: '/api', component: Api },
]

export const router = createRouter({
    history: createMemoryHistory(),
    routes,
})
