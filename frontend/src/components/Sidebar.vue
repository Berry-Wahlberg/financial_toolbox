<template>
  <div class="sidebar-container" :class="{ 'is-collapsed': isCollapsed }">
    <div class="sidebar-header">
      <el-icon v-if="!isCollapsed" class="logo-icon" :size="24">
        <component :is="MenuIcon" />
      </el-icon>
      <span v-if="!isCollapsed" class="app-title">Financial Toolbox</span>
      <el-button
        class="collapse-button"
        :icon="isCollapsed ? Expand : Fold"
        @click="toggleCollapse"
        circle
        size="small"
      />
    </div>

    <div class="sidebar-menu">
      <SidebarItem
        v-for="item in visibleMenuItems"
        :key="item.id"
        :item="item"
        :is-active="isActiveRoute(item.route)"
        @click="handleItemClick"
      />
    </div>

    <div class="sidebar-footer">
      <el-tooltip v-if="isCollapsed" content="Expand" placement="right">
        <el-button :icon="Expand" @click="toggleCollapse" circle size="small" />
      </el-tooltip>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { Expand, Fold, Menu as MenuIcon } from '@element-plus/icons-vue'
import SidebarItem from './SidebarItem.vue'
import { defaultSidebarConfig, getVisibleMenuItems } from './SidebarConfig'
import type { MenuItem } from './SidebarConfig'

const router = useRouter()
const route = useRoute()

const isCollapsed = ref(defaultSidebarConfig.defaultCollapsed || false)
const sidebarConfig = ref(defaultSidebarConfig)

const visibleMenuItems = computed(() => {
  return getVisibleMenuItems(sidebarConfig.value)
})

const isActiveRoute = (itemRoute: string): boolean => {
  return route.path === itemRoute
}

const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value
}

const handleItemClick = (item: MenuItem) => {
  console.log('Menu item clicked:', item.name, item.route)
}

const toggleMenuItemVisibility = (id: string) => {
  sidebarConfig.value = {
    ...sidebarConfig.value,
    items: sidebarConfig.value.items.map((item) =>
      item.id === id ? { ...item, visible: !item.visible } : item,
    ),
  }
}

const reorderMenuItems = (newOrder: string[]) => {
  const itemMap = new Map(sidebarConfig.value.items.map((item) => [item.id, item]))
  const reorderedItems = newOrder
    .map((id) => itemMap.get(id))
    .filter((item): item is MenuItem => item !== undefined) as MenuItem[]

  sidebarConfig.value = {
    ...sidebarConfig.value,
    items: reorderedItems,
  }
}

defineExpose({
  toggleCollapse,
  toggleMenuItemVisibility,
  reorderMenuItems,
})
</script>

<style scoped>
.sidebar-container {
  display: flex;
  flex-direction: column;
  width: 240px;
  height: 100vh;
  background: linear-gradient(180deg, #f5f7fa 0%, #ffffff 100%);
  border-right: 1px solid #e4e7ed;
  transition: width 0.3s ease;
}

.sidebar-container.is-collapsed {
  width: 64px;
}

.sidebar-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px;
  border-bottom: 1px solid #e4e7ed;
  background-color: #409eff;
  color: white;
}

.logo-icon {
  flex-shrink: 0;
}

.app-title {
  flex: 1;
  font-size: 16px;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.collapse-button {
  flex-shrink: 0;
  background-color: rgba(255, 255, 255, 0.2);
  border: none;
  color: white;
}

.collapse-button:hover {
  background-color: rgba(255, 255, 255, 0.3);
}

.sidebar-menu {
  flex: 1;
  overflow-y: auto;
  padding: 16px 8px;
}

.sidebar-menu::-webkit-scrollbar {
  width: 6px;
}

.sidebar-menu::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.sidebar-menu::-webkit-scrollbar-thumb {
  background: #c0c4cc;
  border-radius: 3px;
}

.sidebar-menu::-webkit-scrollbar-thumb:hover {
  background: #a0c4e8;
}

.sidebar-footer {
  padding: 16px;
  border-top: 1px solid #e4e7ed;
  display: flex;
  justify-content: center;
}

@media (max-width: 768px) {
  .sidebar-container {
    position: fixed;
    left: 0;
    top: 0;
    z-index: 1000;
    box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  }

  .sidebar-container.is-collapsed {
    transform: translateX(-100%);
  }
}
</style>
