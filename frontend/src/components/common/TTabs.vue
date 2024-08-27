<script setup>
import {VueDraggable} from "vue-draggable-plus";
import {ref, watch} from "vue";

const props = defineProps({
  items: {
    type: Array,
    default: () => [],
    validator(value) {
      return value.every(item => {
        return typeof item === 'object' && 'id' in item && 'name' in item && 'type' in item;
      });
    },
    required: true
  },
  activeItem:{
    type: Object,
    default: () => {},
    validator(value) {
      return typeof value === 'object' && 'id' in value && 'type' in value;
    }
  }

});
const tabsArr = ref(props.items)
const tab = ref(props.activeItem)
watch(()=>props.activeItem,
    (newValue, oldValue) => {
        chooseTab(tab.value)
      },
    { deep: false })


const emits = defineEmits(['closeTab','moveItem','switchTab']);
const chooseTab = (item)=>{
  tab.value=item
}
const switchTab =  (item)=>{
  chooseTab(item)
  emits('switchTab',item)
}
const closeTab =  (item)=>{
  emits('closeTab',item)
}
const onEnd = (e)=>{
  emits('moveItem',tabsArr.value)
}
const isSelected=(t)=>{
  let cur= tab.value
  if (cur){
    if (t.id===cur.id && t.type===cur.type){
      return true
    }
  }

  return false
}
// scroll
const scrollableContainer = ref(null);
const handleWheel = (event) => {
  event.preventDefault();
  const container = scrollableContainer.value;
  if (container) {
    let newScrollLeft = container.scrollLeft + event.deltaY;
    if (!isNaN(newScrollLeft)) {
      container.scrollLeft = newScrollLeft;
    }
  }
};

</script>

<template>
  <div class="t-tab-wrapper">
      <div class="t-tabs-scroll"  @wheel="handleWheel"  ref="scrollableContainer">
        <VueDraggable v-model="tabsArr" :animation="150" class="t-tabs"  @end="onEnd">
          <div v-for="item in tabsArr" :key="item" :class="['t-tab', { 't-tab-active': isSelected(item) }]" @click="switchTab(item)">{{ item.name }}</div>
        </VueDraggable>
      </div>
    <div class="t-tab-overflow-select">V</div>
  </div>
</template>

<style scoped>
.t-tab-wrapper{
  display: flex;
  border-bottom: rgb(239, 239, 245) solid 1px;
}

.t-tabs-scroll{
  flex-grow: 1;
  overflow-x: auto;
  white-space: nowrap;
}
.t-tabs-scroll::-webkit-scrollbar{
  display: none;
}
.t-tabs{
  flex-grow: 1;
  display: flex;

}
.t-tab{
  flex-shrink: 0;
  height: 36px;
  width: 180px;
  margin-left:  10px;
  border-right: rgb(239, 239, 245) solid 1px;
  border-left:  rgb(239, 239, 245) solid 1px;
  cursor: pointer;
}
.t-tab-active{
  background-color: #fc6;
}
.t-tab-overflow-select{
  flex-shrink: 0;
  padding-right: 10px;
}
</style>