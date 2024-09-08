<script setup>
import {VueDraggable} from "vue-draggable-plus";
import {computed, nextTick, onMounted, reactive, ref, watch, watchEffect} from "vue";
import { useI18n } from 'vue-i18n'


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
  activeTab:{
    type: Object,
    default: () => {},
    validator(value) {
      return typeof value === 'object' && 'id' in value && 'type' in value;
    }
  },
  tabScrollLeft:{
    type: Number,
    default: () => 0
  }

});
const tabsArr = ref(props.items)
const tab = ref(props.activeTab)
const scrollLeft = ref(props.tabScrollLeft)
watch(()=>props.activeTab,
    (newValue, oldValue) => {
  console.log(newValue)
        changeCurrentTab(tab.value)
      },
    { deep: false }
)
watch(()=>props.items,(n,o)=>{
  Object.assign(tabsArr.value,n)
},{deep:true})

const emits = defineEmits(['moveTab','chooseTab','closeTab','closeOtherTabs',
  'closeAllTabs','saveAllTabs','update:tabScrollLeft','addRequest']);

// i18n
const i18n = useI18n()
// refs
const scrollableContainer = ref(null);
const tabs = ref([]);

// tab btn
const changeCurrentTab = (item)=>{
  tab.value=item
}

const chooseTab = (item)=>{
  changeCurrentTab(item)
  emits('chooseTab',item)
}
const closeTab =  (e,item)=>{
  e.stopPropagation()
  emits('closeTab',item)
}
const onChange = (e)=>{
  Object.assign(tabs.value,scrollableContainer.value.querySelectorAll('.t-tab'))
  emits('moveTab',tabsArr.value)
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
      emits('update:tabScrollLeft',newScrollLeft)
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
const IsShowHidingTabsBtn = computed(() => hidingTabs.length > 0);
const showHidingTabs=()=>{
  if (!scrollableContainer.value){
    return
  }
  const rect = scrollableContainer.value.getBoundingClientRect()
  let options = []
  tabs.value.forEach((tab,index) => {
    const tabRect = tab.getBoundingClientRect()
    if (tabRect.right>(rect.right+2)||tabRect.left<rect.left){
      let item = tabsArr.value[index]
      options.push({
        label: item.name,
        key: item
      })
    }
  });
  hidingTabs.length = 0;
  hidingTabs.push(...options);
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

const operationList = [
  {
    label: i18n.t("tabs.operation.close-all"),
    key: "tabs.operation.close-all",
  },
  {
    label: i18n.t("tabs.operation.save-all"),
    key: "tabs.operation.save-all",
  },  {
    label: i18n.t("tabs.operation.add"),
    key: "tabs.operation.add",
  },
]
const handleSelectOperation=(key)=>{
  switch (key){
    case "tabs.operation.close-all":
      emits('closeAllTabs')
          break
    case "tabs.operation.save-all":
      emits('saveAllTabs')
    case "tabs.operation.add":
      emits('addRequest')
          break
  }
}

// on mounted
const handleResize = () => {
  showHidingTabs()
};
onMounted(()=>{
  window.addEventListener('resize', handleResize);
  const container = scrollableContainer.value;
  if (!isNaN(scrollLeft.value)) {
    container.scrollLeft = scrollLeft.value;
  }
  showHidingTabs()
})

// ctx-menu
const showTabMenuRef =  ref(false)
const showTabMenuXRef = ref(0)
const showTabMenuYRef = ref(0)
let tabRightClickChoose = ref(null)
const onTabRightClick=(e,item)=>{
  e.preventDefault()
  tabRightClickChoose.value=item
  showTabMenuRef.value = false
  nextTick().then(() => {
    showTabMenuRef.value = true
    showTabMenuXRef.value = e.clientX
    showTabMenuYRef.value = e.clientY
  })
}
const onTabRightMenuClickOutside=()=>{
  showTabMenuRef.value = false
  tabRightClickChoose.value=null
}

const tabCtxMenu = [
  {
    label: i18n.t("tabs.tab-ctx-menu.close"),
    key: "tabs.tab-ctx-menu.close",
  },
  {
    label: i18n.t("tabs.tab-ctx-menu.close-all"),
    key: "tabs.tab-ctx-menu.close-all",
  },
  {
    label: i18n.t("tabs.tab-ctx-menu.close-others"),
    key: "tabs.tab-ctx-menu.close-others",
  }
]

const onTabRightMenuSelect=(key)=>{
  showTabMenuRef.value = false
  switch (key){
    case "tabs.tab-ctx-menu.close":
      emits('closeTab',[tabRightClickChoose.value])
      break
    case "tabs.tab-ctx-menu.close-all":
      emits('closeAllTabs')
      break
    case "tabs.tab-ctx-menu.close-others":
      emits('closeOtherTabs',[tabRightClickChoose.value])
      break
  }
}

// method icon
const methodIcon = (m)=>{
  switch (m){
    case "GET":
      return "method-get"
    case "POST":
      return "method-post"
    case "PUT":
      return "method-put"
    case "DELETE":
      return "method-delete"
    default:
      return "method-other"
  }
}

</script>

<template>
  <div class="t-tab-wrapper">
      <div class="t-tabs-scroll"  @wheel="handleWheel"  ref="scrollableContainer">
        <VueDraggable v-model="tabsArr" :animation="150" class="t-tabs"  @change="onChange">
          <div v-for="item in tabsArr" :key="item" :class="['t-tab', { 't-tab-active': isSelected(item) }]"
               @click="chooseTab(item)" @mouseenter="showClose(item)" @mouseleave="hideClose"
               ref="tabs"
               @contextmenu.prevent="onTabRightClick($event,item)"
               @mousedown="onTabRightMenuClickOutside"
          >

            <n-ellipsis style="max-width: 176px" :tooltip="{placement:'bottom'}">
              <t-icon name="collection" width="1em" height="1em" v-if="item.type==='collection'" style="margin-right: 6px;margin-left: 10px"/>
              <span v-else style="margin-right: 6px;margin-left: 10px">
              <span v-if="item.method&&item.method.length>0" :class="['method',methodIcon(item.method)]" >{{item.method}}</span>
               <t-icon name="request" width="1em" height="1em" v-else />
            </span>
              {{ item.name }}
            </n-ellipsis>
              <n-button quaternary circle class="t-tab-close" size="tiny" v-show="isHovered(item)" @click="closeTab($event,item)">
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
          v-if="IsShowHidingTabsBtn"
          :options="hidingTabs"
          @select="handleSelectHidingTab"
          style="text-align: left"
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
<!--  ctx menu  -->
    <n-dropdown
        placement="bottom-start"
        trigger="manual"
        style="text-align: left"
        :x="showTabMenuXRef"
        :y="showTabMenuYRef"
        :options="tabCtxMenu"
        :show="showTabMenuRef"
        :on-clickoutside="onTabRightMenuClickOutside"
        @select="onTabRightMenuSelect"
    />
  </div>
</template>

<style scoped>
.t-tab-wrapper{
  height: 46px;
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
  text-align: left;
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
.method{
  font-size: 0.8em;
  font-weight: bold;
}

.method-get{
  color: #7eb65e;
}
.method-post{
  color: #db9b4b;
}
.method-put{
  color: #6274e1;
}
.method-delete{
  color: rgba(216, 55, 55, 0.82);
}

.method-other{
  color: #909090;
}

</style>