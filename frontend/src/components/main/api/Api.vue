<script setup>
import {NSplit,NInput,NSwitch} from "naive-ui";
import { useI18n } from 'vue-i18n'
const i18n=useI18n()

</script>

<template>
  <n-split direction="horizontal" default-size="200px" :max="0.35" min="200px">
    <template #1>
      <n-input v-model:value="pattern" placeholder="搜索" />
      <n-switch v-model:value="showIrrelevantNodes">
        <template #checked>
          展示搜索无关的节点
        </template>
        <template #unchecked>
          隐藏搜索无关的节点
        </template>
      </n-switch>
      <n-tree
          :show-irrelevant-nodes="showIrrelevantNodes"
          :pattern="pattern"
          :data="data"
          block-line
      />
    </template>
    <template #2>
      main
    </template>
  </n-split>
</template>

<style scoped>

</style>