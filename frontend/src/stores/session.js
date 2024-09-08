import { defineStore } from 'pinia'
import {
    GetTabSession,
    SetTabSessionPosition,
    SetTabSessionTabs,
} from 'wailsjs/go/service/sessionSvc.js'
const useSessionStore = defineStore('session', {
    state: () => ({
        tabs: [], // all group name set
        currentTab: {},
        tabScrollLeft:0,

    }),
    getters:{
        getTabs: (state) => state.tabs,
        getCurrentTab: (state) => state.currentTab,
        getTabScrollLeft: (state) => state.tabScrollLeft,
    },
    actions:{
        async loadTabSession() {
            const { data : {
                tab_scroll_left,
                tabs,
                current_tab
            }, success, msg }  = await GetTabSession()
            if (!success){
                return {success,msg}
            }
            if (tab_scroll_left===""){
                this.tabScrollLeft=0
            }else {
                this.tabScrollLeft=parseInt(tab_scroll_left)
            }
            if (!tabs){
                this.tabs=[]
            }else {
                this.tabs=tabs
            }
            this.currentTab=current_tab
            console.log( this.currentTab)
            console.log( this.tabs)
            return {success,msg}
        },
        async saveTabSessionTabs(tabs) {
            const { data = {}, success, msg } = await SetTabSessionTabs(tabs)
            if (!success){
                return { success, msg}
            }
            this.tabs=tabs
            return { success, msg}
        },

        async saveTabSessionPositions(currentTab,scrollLeft) {
            let scrollLeftStr=scrollLeft+""
            const { data = {}, success, msg } = await SetTabSessionPosition(currentTab,scrollLeftStr)
            if (!success){
                return { success, msg}
            }
            console.log(currentTab,scrollLeft)
            this.currentTab=currentTab
            this.tabScrollLeft=scrollLeft
            return { success, msg}
        },
        async addRequestTab() {
           this.tabs.push({
               type:"request",
               name:"测试",
               id:1,
           })
        }
    }
})

export default useSessionStore
