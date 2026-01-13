<template>
  <div
    class="sidebar-item"
    :class="{ 'is-active': isActive }"
    @click="handleClick"
  >
    <el-icon v-if="item.icon" class="item-icon">
      <component :is="getIconComponent(item.icon)" />
    </el-icon>
    <span class="item-label">{{ item.name }}</span>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import type { MenuItem } from './SidebarConfig'
import {
  House,
  DataAnalysis,
  Document,
} from '@element-plus/icons-vue'

interface Props {
  item: MenuItem
  isActive?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  isActive: false,
})

const emit = defineEmits<{
  click: [item: MenuItem]
}>()

const router = useRouter()

const getIconComponent = (iconName: string) => {
  const iconMap: Record<string, any> = {
    House,
    DataAnalysis,
    Document,
  }
  return iconMap[iconName] || House
}

const handleClick = () => {
  emit('click', props.item)
  router.push(props.item.route)
}
</script>

<style scoped>
.sidebar-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  margin: 4px 8px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #606266;
}

.sidebar-item:hover {
  background-color: #ecf5ff;
  color: #409eff;
}

.sidebar-item.is-active {
  background-color: #409eff;
  color: white;
  font-weight: 500;
}

.item-icon {
  margin-right: 12px;
  font-size: 18px;
}

.item-label {
  font-size: 14px;
  white-space: nowrap;
}
</style>
