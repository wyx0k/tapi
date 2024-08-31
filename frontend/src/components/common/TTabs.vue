<script setup>
import {VueDraggable} from "vue-draggable-plus";
import {computed, nextTick, reactive, ref, watch, watchEffect} from "vue";

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
        changeCurrentTab(tab.value)
      },
    { deep: false })


const emits = defineEmits(['changeCurrentTab','moveItem','chooseTab']);

// refs
const scrollableContainer = ref(null);
const tabs = ref([]);

// tab btn
const changeCurrentTab = (item)=>{
  console.log("change tab")
  tab.value=item
}
const chooseTab =  (item)=>{
  console.log("choose tab")
  changeCurrentTab(item)
  emits('chooseTab',item)
}
const closeTab =  (item)=>{
  emits('closeTab',item)
}
const onChange = (e)=>{
  Object.assign(tabs.value,scrollableContainer.value.querySelectorAll('.t-tab'))
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
// close
const hoveredTab = ref(null)
const showClose = (item)=>{
  hoveredTab.value=item
}
const hideClose = ()=>{
  hoveredTab.value=null
}
const isHovered = (item)=>{
  if (hoveredTab.value){
    let a=hoveredTab.value
    if (a.id===item.id&&a.type ===item.type){
      return true
    }
  }
  return false
}
// hiding tabs
const hidingTabs = reactive([])

const showHidingTabs=()=>{

  const rect = scrollableContainer.value.getBoundingClientRect();
  let options = []
  tabs.value.forEach((tab,index) => {
    const tabRect = tab.getBoundingClientRect();
    if (tabRect.right>(rect.right+2)||tabRect.left<rect.left){
      let item = tabsArr.value[index]
      options.push({
        label: item.name,
        key: item
      })
    }
  });
  Object.assign(hidingTabs,options)

}

const handleSelectHidingTab = (item)=>{
  tabsArr.value.forEach((tab,index) => {
    if (tab.type ===item.type&&tab.id===item.id){
      const targetRect = tabs.value[index].getBoundingClientRect()
      scrollableContainer.value.scrollLeft=targetRect.width*index
    }
  })
  chooseTab(item)
}

// tabs operatioin

const operationList = []
const handleSelectOperation=(item)=>{
  console.log(item)
}
</script>

<template>
  <div class="t-tab-wrapper">
      <div class="t-tabs-scroll"  @wheel="handleWheel"  ref="scrollableContainer">
        <VueDraggable v-model="tabsArr" :animation="150" class="t-tabs"  @change="onChange">
          <div v-for="item in tabsArr" :key="item" :class="['t-tab', { 't-tab-active': isSelected(item) }]"
               @click="chooseTab(item)" @mouseenter="showClose(item)" @mouseleave="hideClose"
               ref="tabs"
          >
            <n-ellipsis style="max-width: 165px" :tooltip="{placement:'bottom'}">
              {{ item.name }}
            </n-ellipsis>
              <n-button quaternary circle class="t-tab-close" size="tiny" v-show="isHovered(item)">
                <t-icon name="close" width="16px" height="16px"/>
              </n-button>
            <div class="t-tab-selected-sign"></div>
          </div>
        </VueDraggable>
      </div>
    <div class="t-tab-overflow-select">
      <n-dropdown
          placement="bottom-start"
          trigger="click"
          size="small"
          :options="hidingTabs"
          @select="handleSelectHidingTab"
      >
        <n-button quaternary  size="tiny" class="t-tab-overflow-select-item" @click="showHidingTabs">
          <t-icon name="dropdown" width="20px" height="20px"/>
        </n-button>
      </n-dropdown>
      <n-dropdown
          placement="bottom-start"
          trigger="click"
          size="small"
          :options="operationList"
          @select="handleSelectOperation"
      >
      <n-button quaternary  size="tiny" class="t-tab-overflow-select-item">
        <t-icon name="more-operation" width="20px" height="20px"/>
      </n-button>
      </n-dropdown>
    </div>
  </div>
</template>

<style scoped>
.t-tab-wrapper{
  display: flex;
  border-bottom: rgb(239, 239, 245) solid 1px;
  /*box-shadow: inset 10px 0 10px -10px rgba(0, 0, 0, 0.3);*/
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
  padding-top: 10px;
  padding-right: 25px;
  position: relative;
  flex-shrink: 0;
  height: 36px;
  width: 180px;
  margin-left:  10px;
  border-right: rgb(239, 239, 245) solid 1px;
  border-left:  rgb(239, 239, 245) solid 1px;
  cursor: pointer;
}
.t-tab-active{
}

.t-tab-overflow-select{
  border-left: rgba(43,35,151,0.46) solid 3px;
  flex-shrink: 0;
  padding: 0 12px 0 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  /*box-shadow: -10px 0 10px rgba(0, 0, 0, 0.1);*/
}
.t-tab-overflow-select-item{
  background-color: #fff;
  margin-left: 2px;
}
.t-tab-selected-sign{
  position: absolute;
  width: 100%;
  height: 8px;
  display: none;
}
.t-tab-active .t-tab-selected-sign{
  display: block;
  background-color: #6cf;
  bottom: -2px;
}
.t-tab-close{
  position: absolute;
  right: 6px;
  top:6px
}


</style>