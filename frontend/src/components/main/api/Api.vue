<script setup>
import {NSplit,NInput,NTabs,NTab} from "naive-ui";
import { useI18n } from 'vue-i18n'
import {ref} from "vue";
import {VueDraggable} from "vue-draggable-plus";
const i18n=useI18n()

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
const list = ref([
  {
    name: 'Joao',
    id: 1
  },
  {
    name: 'Jean',
    id: 2
  },
  {
    name: 'Johanna',
    id: 3
  },
  {
    name: 'Juan',
    id: 4
  }
])
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
      <VueDraggable v-model="list" :animation="150" target=".n-tabs-wrapper">
        <n-tabs type="card" closable>
            <n-tab v-for="item in list" :key="item.id" >
              {{ item.name }}
            </n-tab>
        </n-tabs>
      </VueDraggable>
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