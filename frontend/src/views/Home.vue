<template>
  <div class="home">
    <el-container>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-card class="card">
              <template #header>
                <div class="card-header">
                  <span>股票数据</span>
                </div>
              </template>
              <div class="card-content">
                <el-input v-model="symbol" placeholder="输入股票代码" />
                <el-button type="primary" @click="fetchStockData" style="margin-top: 10px">
                  查询
                </el-button>
              </div>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card class="card">
              <template #header>
                <div class="card-header">
                  <span>技术指标</span>
                </div>
              </template>
              <div class="card-content">
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
              </div>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card class="card">
              <template #header>
                <div class="card-header">
                  <span>快速导航</span>
                </div>
              </template>
              <div class="card-content">
                <el-button @click="$router.push('/chart')">图表分析</el-button>
                <el-button @click="$router.push('/editor')">编辑器</el-button>
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

const symbol = ref('AAPL')
const indicator = ref('ma')
const indicators = ref<string[]>([])

onMounted(async () => {
  try {
    const response = await fetch('/api/indicators')
    const data = await response.json()
    indicators.value = data.data
  } catch (error) {
    ElMessage.error('加载指标失败')
  }
})

const fetchStockData = async () => {
  try {
    const response = await fetch(`/api/stock/${symbol.value}`)
    const data = await response.json()
    ElMessage.success('数据加载成功')
  } catch (error) {
    ElMessage.error('加载失败')
  }
}

const calculateIndicator = async () => {
  try {
    const response = await fetch('/api/indicator/calculate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        symbol: symbol.value,
        indicator: indicator.value,
      }),
    })
    const data = await response.json()
    ElMessage.success('计算成功')
  } catch (error) {
    ElMessage.error('计算失败')
  }
}
</script>

<style scoped>
.home {
  height: 100%;
}

.el-header {
  background-color: #409eff;
  color: white;
  text-align: center;
  line-height: 60px;
}

.card {
  margin-bottom: 20px;
}

.card-header {
  font-weight: bold;
}

.card-content {
  text-align: center;
}
</style>
