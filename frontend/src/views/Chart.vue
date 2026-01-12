<template>
  <div class="chart">
    <el-container>
      <el-header>
        <h1>图表分析</h1>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card>
              <el-input v-model="symbol" placeholder="股票代码" />
              <el-select v-model="indicator" placeholder="指标" style="margin-top: 10px">
                <el-option label="MA" value="ma" />
                <el-option label="RSI" value="rsi" />
                <el-option label="MACD" value="macd" />
              </el-select>
              <el-button type="primary" @click="loadData" style="margin-top: 10px">
                加载数据
              </el-button>
            </el-card>
          </el-col>
          <el-col :span="18">
            <el-card>
              <v-chart :option="chartOption" style="height: 500px" />
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart, BarChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent,
} from 'echarts/components'
import type { EChartsOption } from 'echarts'

use([
  CanvasRenderer,
  LineChart,
  BarChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent,
])

const symbol = ref('AAPL')
const indicator = ref('ma')
const chartOption = ref<EChartsOption>({})

const loadData = async () => {
  try {
    const [stockRes, indicatorRes] = await Promise.all([
      fetch(`/api/stock/${symbol.value}`),
      fetch('/api/indicator/calculate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          symbol: symbol.value,
          indicator: indicator.value,
        }),
      }),
    ])

    const stockData = await stockRes.json()
    const indicatorData = await indicatorRes.json()

    const dates = stockData.data.map((item: any) => item.date)
    const closes = stockData.data.map((item: any) => item.close)
    const indicatorValues = indicatorData.data.map((item: any) => item.value)

    chartOption.value = {
      title: {
        text: `${symbol.value} - ${indicator.value.toUpperCase()}`,
      },
      tooltip: {
        trigger: 'axis',
      },
      legend: {
        data: ['价格', indicator.value.toUpperCase()],
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
      yAxis: [
        {
          type: 'value',
          name: '价格',
        },
        {
          type: 'value',
          name: indicator.value.toUpperCase(),
        },
      ],
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
          name: '价格',
          type: 'line',
          data: closes,
        },
        {
          name: indicator.value.toUpperCase(),
          type: 'line',
          yAxisIndex: 1,
          data: indicatorValues,
        },
      ],
    }
  } catch (error) {
    console.error('加载数据失败:', error)
  }
}
</script>

<style scoped>
.chart {
  height: 100%;
}

.el-header {
  background-color: #409eff;
  color: white;
  text-align: center;
  line-height: 60px;
}
</style>
