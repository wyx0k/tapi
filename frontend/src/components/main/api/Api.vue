<script setup>
import {NSplit,NInput,useMessage} from "naive-ui";
import { useI18n } from 'vue-i18n'
import {onMounted, ref, watch, watchEffect} from "vue";
import TTabs from "../../common/TTabs.vue";
import useSessionStore from "stores/session.js";
const i18n=useI18n()
const sessionStore=useSessionStore()
const data = [
  {
    label: '0',
    key: '0',
    children: [
      {
        label: '0-0',
        key: '0-0',
        children: [
          { label: '0-0-0', key: '0-0-0' },
          { label: '0-0-1', key: '0-0-1' }
        ]
      },
      {
        label: '0-1',
        key: '0-1',
        children: [
          { label: '0-1-0', key: '0-1-0' },
          { label: '0-1-1', key: '0-1-1' }
        ]
      }
    ]
  },
  {
    label: '1',
    key: '1',
    children: [
      {
        label: '1-0',
        key: '1-0',
        children: [
          { label: '1-0-0', key: '1-0-0' },
          { label: '1-0-1', key: '1-0-1' }
        ]
      },
      {
        label: '1-1',
        key: '1-1',
        children: [
          { label: '1-1-0', key: '1-1-0' },
          { label: '1-1-1', key: '1-1-1' }
        ]
      }
    ]
  }
]
const pattern = ref('')
const message = useMessage()
// init data


let currentTabScrollLeft = sessionStore.getTabScrollLeft

// handle tab events
const moveTab = (e)=>{

}

const  chooseTab = async (item)=>{
  console.log("1111")
  await sessionStore.saveTabSessionPositions(item,currentTabScrollLeft)
  console.log("2222")
}
const closeTab = (e)=>{

}
const closeAllTabs = (e)=>{

}
const closeOtherTabs = (e)=>{

}
const saveAllTabs = (e)=>{

}
const addRequest = async ()=>{
  console.log("add request")
  await sessionStore.addRequestTab()
  console.log(sessionStore.getTabs)

}
// mounted
onMounted(async ()=>{
 const {success,msg} = await sessionStore.loadTabSession()
  if (!success){
    message.warning(msg)
  }

})
</script>

<template>
  <n-split direction="horizontal" default-size="200px" :max="0.35" min="200px">
    <template #1>
      <div class="sider">
        <n-input size="small" v-model:value="pattern" placeholder="搜索" class="sider-search" />
        <n-tree
            :show-irrelevant-nodes="true"
            :pattern="pattern"
            :data="data"
            block-line
            class="sider-tree"
        />
      </div>

    </template>
    <template #2>
      <t-tabs
          :items="sessionStore.getTabs"
          :active-tab="sessionStore.getCurrentTab"
          v-model:tab-scroll-left="currentTabScrollLeft"
          @choose-tab="chooseTab"
          @close-tab=""
          @move-tab=""
          @close-all-tabs=""
          @close-other-tabs=""
          @save-all-tabs=""
          @add-request="addRequest"
      ></t-tabs>
      <router-view></router-view>
    </template>
  </n-split>
</template>

<style scoped>
.sider{
  padding: 4px;
}
.sider-search{
  margin-bottom: 10px;
  text-align: left;

}
.sider-tree{
  text-align: left;
}
</style>