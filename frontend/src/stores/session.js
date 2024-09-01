import { defineStore } from 'pinia'
import {
    GetTabSession,
    SetTabSession,

} from 'wailsjs/go/services/sessionSvc.js'
const useSessionStore = defineStore('session', {
    state: () => ({
        tabs: [], // all group name set

    }),
    getters:{

    },
    actions:{

    }
})

export default useSessionStore
