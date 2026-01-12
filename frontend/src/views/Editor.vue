<template>
  <div class="editor">
    <el-container>
      <el-header>
        <h1>策略编辑器</h1>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>组件库</span>
                </div>
              </template>
              <div class="component-list">
                <el-button
                  v-for="component in components"
                  :key="component.id"
                  @click="addComponent(component)"
                  style="width: 100%; margin-bottom: 10px"
                >
                  {{ component.name }}
                </el-button>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>编辑区域</span>
                  <el-button type="primary" size="small" @click="saveStrategy">
                    保存策略
                  </el-button>
                </div>
              </template>
              <div class="editor-area">
                <div
                  v-for="(item, index) in strategyComponents"
                  :key="index"
                  class="strategy-item"
                >
                  <span>{{ item.name }}</span>
                  <el-button
                    type="danger"
                    size="small"
                    @click="removeComponent(index)"
                  >
                    删除
                  </el-button>
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>属性面板</span>
                </div>
              </template>
              <div class="property-panel">
                <el-form v-if="selectedComponent" label-width="80px">
                  <el-form-item label="名称">
                    <el-input v-model="selectedComponent.name" />
                  </el-form-item>
                  <el-form-item label="参数">
                    <el-input
                      v-model="selectedComponent.params"
                      type="textarea"
                      :rows="4"
                    />
                  </el-form-item>
                </el-form>
                <div v-else>请选择一个组件</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

interface Component {
  id: string
  name: string
  type: string
  params?: string
}

const components = ref<Component[]>([
  { id: '1', name: '买入条件', type: 'buy' },
  { id: '2', name: '卖出条件', type: 'sell' },
  { id: '3', name: '止损设置', type: 'stop-loss' },
  { id: '4', name: '止盈设置', type: 'take-profit' },
  { id: '5', name: '仓位管理', type: 'position' },
])

const strategyComponents = ref<Component[]>([])
const selectedComponent = ref<Component | null>(null)

const addComponent = (component: Component) => {
  strategyComponents.value.push({
    ...component,
    params: '',
  })
  ElMessage.success(`已添加 ${component.name}`)
}

const removeComponent = (index: number) => {
  strategyComponents.value.splice(index, 1)
  ElMessage.success('已删除组件')
}

const saveStrategy = () => {
  const strategy = {
    name: '新策略',
    components: strategyComponents.value,
    createdAt: new Date().toISOString(),
  }
  console.log('保存策略:', strategy)
  ElMessage.success('策略保存成功')
}
</script>

<style scoped>
.editor {
  height: 100%;
}

.el-header {
  background-color: #409eff;
  color: white;
  text-align: center;
  line-height: 60px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.component-list {
  max-height: 600px;
  overflow-y: auto;
}

.editor-area {
  min-height: 500px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  padding: 10px;
}

.strategy-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  margin-bottom: 10px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.property-panel {
  max-height: 600px;
  overflow-y: auto;
}
</style>
