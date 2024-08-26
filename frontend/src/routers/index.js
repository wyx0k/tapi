import { createMemoryHistory, createRouter } from 'vue-router'

import Api from '../components/main/api/Api.vue'
import Collection from '../components/main/api/Collection.vue'
import Request from '../components/main/api/Request.vue'
import NoChoose from "../components/main/api/NoChoose.vue";

const routes = [
    {path: '/',redirect:'/api/no-choose'},
    {   path: '/api',
        component: Api,
        children: [
            {
                path: 'collection',
                component: Collection,
            },
            {
                path: 'request',
                component: Request,
            },
            {
                path: 'no-choose',
                component: NoChoose,
            }
        ]
    },
]

export const router = createRouter({
    history: createMemoryHistory(),
    routes,
})
