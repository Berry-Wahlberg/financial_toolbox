<template>
  <div id="app">
    <el-container>
      <el-header>
        <h1>Financial Toolbox</h1>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>股票数据</span>
                </div>
              </template>
              <el-input v-model="symbol" placeholder="输入股票代码" />
              <el-button type="primary" @click="fetchStockData" style="margin-top: 10px">
                查询
              </el-button>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>技术指标</span>
                </div>
              </template>
              <el-select v-model="indicator" placeholder="选择指标">
                <el-option
                  v-for="item in indicators"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
              <el-button type="primary" @click="calculateIndicator" style="margin-top: 10px">
                计算
              </el-button>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>快速导航</span>
                </div>
              </template>
              <el-button @click="currentView = 'chart'">图表分析</el-button>
              <el-button @click="currentView = 'editor'">编辑器</el-button>
            </el-card>
          </el-col>
        </el-row>

        <el-row v-if="currentView === 'chart'" :gutter="20" style="margin-top: 20px">
          <el-col :span="24">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>图表分析</span>
                </div>
              </template>
              <v-chart :option="chartOption" style="height: 500px" />
            </el-card>
          </el-col>
        </el-row>

        <el-row v-if="currentView === 'editor'" :gutter="20" style="margin-top: 20px">
          <el-col :span="24">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>策略编辑器</span>
                </div>
              </template>
              <div class="editor-area">
                <el-button
                  v-for="component in components"
                  :key="component.id"
                  @click="addComponent(component)"
                  style="margin-right: 10px; margin-bottom: 10px"
                >
                  {{ component.name }}
                </el-button>
                <div class="strategy-list">
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
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent,
} from 'echarts/components'

use([
  CanvasRenderer,
  LineChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent,
])

const symbol = ref('AAPL')
const indicator = ref('ma')
const indicators = ref<string[]>([])
const currentView = ref('home')
const chartOption = ref<any>({})

const components = ref([
  { id: '1', name: '买入条件' },
  { id: '2', name: '卖出条件' },
  { id: '3', name: '止损设置' },
  { id: '4', name: '止盈设置' },
  { id: '5', name: '仓位管理' },
])

const strategyComponents = ref<any[]>([])

onMounted(async () => {
  try {
    const result = await window.$wails.GetAvailableIndicators()
    indicators.value = result
  } catch (error) {
    ElMessage.error('加载指标失败')
  }
})

const fetchStockData = async () => {
  try {
    const result = await window.$wails.GetStockData(symbol.value)
    ElMessage.success('数据加载成功')
    console.log(result)
  } catch (error) {
    ElMessage.error('加载失败')
  }
}

const calculateIndicator = async () => {
  try {
    const req = {
      symbol: symbol.value,
      indicator: indicator.value,
      params: {},
    }
    const reqJSON = JSON.stringify(req)
    const result = await window.$wails.CalculateIndicator(reqJSON)
    const data = JSON.parse(result)
    
    if (currentView.value === 'chart') {
      updateChart(data)
    }
    
    ElMessage.success('计算成功')
  } catch (error) {
    ElMessage.error('计算失败')
  }
}

const updateChart = (data: any) => {
  const dates = data.data.map((item: any) => item.date)
  const values = data.data.map((item: any) => item.value)

  chartOption.value = {
    title: {
      text: `${symbol.value} - ${indicator.value.toUpperCase()}`,
    },
    tooltip: {
      trigger: 'axis',
    },
    legend: {
      data: [indicator.value.toUpperCase()],
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    xAxis: {
      type: 'category',
      data: dates,
    },
    yAxis: {
      type: 'value',
    },
    dataZoom: [
      {
        type: 'inside',
        start: 0,
        end: 100,
      },
      {
        start: 0,
        end: 100,
      },
    ],
    series: [
      {
        name: indicator.value.toUpperCase(),
        type: 'line',
        data: values,
      },
    ],
  }
}

const addComponent = (component: any) => {
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
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100vh;
  margin: 0;
  padding: 0;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.el-header {
  background-color: #409eff;
  color: white;
  text-align: center;
  line-height: 60px;
}

.card-header {
  font-weight: bold;
}

.editor-area {
  min-height: 500px;
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
</style>
