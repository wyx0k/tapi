<script setup>
import { useOsTheme, darkTheme , zhCN, dateZhCN} from "naive-ui";
import {computed} from "vue";
import { } from "naive-ui";

const osTheme = useOsTheme();
const theme = computed(() => osTheme.value === "dark" ? darkTheme : null)
const local = zhCN
const dateLocale = dateZhCN
</script>

<template>
  <n-config-provider :theme="theme" :locale="locale" :date-locale="dateLocale">
    <n-card> 当前操作系统的主题是 {{ osTheme }}。 </n-card>
  </n-config-provider>
</template>

<style>

</style>
