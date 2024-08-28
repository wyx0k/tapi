<template>
  <svg :style="{width, height}">
    <use :xlink:href="id" :fill="fill" />
  </svg>
</template>

<script setup>
import {computed} from "vue"
const props = defineProps({
  prefix: {
    type: String,
    default: "icon"
  },
  name: {
    type: String,
    required: true
  },
  width: {
    type: String,
    default: "16px"
  },
  height: {
    type: String,
    default: "16px"
  },
  fill: {
    type: String,
    default: '#333'
  }
})
const id = computed(() => {
  return `#${props.prefix}-${props.name}`
})
</script>

