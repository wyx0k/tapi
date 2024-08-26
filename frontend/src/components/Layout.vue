<script setup>
import { NSpace,NLayout,NLayoutHeader,NLayoutContent,NLayoutSider,NLayoutFooter,NSplit } from 'naive-ui'

import { useI18n } from 'vue-i18n'

const i18n = useI18n()
</script>

<template>
  <n-space vertical size="large" >
    <n-layout class="n-layout" >
      <n-layout-header bordered class="n-layout-header">{{i18n.t("name")}}</n-layout-header>
      <n-layout has-sider class="n-layout-inner">
        <n-layout-sider bordered
                        :width="40"
                        >
          menu
        </n-layout-sider>
        <n-layout-content class="content-wrapper">
          <router-view></router-view>
        </n-layout-content>
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
  height: calc(100vh - 70px);
}
.n-layout-header {
  height: 50px;
}
.n-layout-footer {
  position: absolute;
  width: 100vw;
  height: 20px;
  bottom: 0;
}
.content-wrapper{
  position: relative;
  height: calc(100vh - 70px);
  flex: 1;
}
</style>