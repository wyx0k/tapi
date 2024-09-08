<script setup>
import { useOsTheme, darkTheme , zhCN, dateZhCN} from "naive-ui";
import {computed} from "vue";
import { } from "naive-ui";
import './style.css'
import Layout from "./components/Layout.vue";
const osTheme = useOsTheme();
const theme = computed(() => osTheme.value === "dark" ? darkTheme : null)
const locale = zhCN
const dateLocale = dateZhCN
</script>

<template>
  <n-config-provider :theme="theme" :locale="locale" :date-locale="dateLocale">
    <n-notification-provider>
    <n-modal-provider>
    <n-message-provider>
      <!--    <n-card> 当前操作系统的主题是 {{ osTheme }}。 </n-card>-->
      <layout v-no-ctx-menu></layout>
    </n-message-provider>
    </n-modal-provider>
  </n-notification-provider>

  </n-config-provider>
</template>

<style>

</style>
