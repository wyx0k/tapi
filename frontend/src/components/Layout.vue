<script setup>
import { NSpace,NLayout,NLayoutHeader,NLayoutContent,NLayoutSider,NLayoutFooter,NSplit } from 'naive-ui'

import { useI18n } from 'vue-i18n'

const i18n = useI18n()
</script>

<template>
  <n-space vertical size="large" >
    <n-layout class="n-layout">
      <n-layout-header bordered class="n-layout-header">{{i18n.t("name")}}</n-layout-header>
      <n-layout  has-sider class="n-layout-inner">
        <n-layout-sider bordered
                        :width="60"
                        >
          collection
        </n-layout-sider>
        <n-layout position="absolute">
          <component ></component>
                  <n-layout-content  class="n-layout-content">
                    <n-split direction="horizontal" default-size="200px" :max="0.35" min="200px">
                      <template #1>
                        {{i18n.t("name")}}

                      </template>
                      <template #2>
                        main
                      </template>
                    </n-split>
                  </n-layout-content>
        </n-layout>

      </n-layout>
      <n-layout-footer bordered class="n-layout-footer">footer</n-layout-footer>
    </n-layout>
  </n-space>
</template>

<style scoped>
.n-layout {
  height: 100vh;
}
.n-layout-inner {
  height: calc(100vh - 83px);
}
.n-layout-header {
  height: 60px;
}
.n-layout-footer {
  height: 20px;
}
.n-layout-content {
  flex:1;
}
</style>