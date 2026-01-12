# Financial Toolbox

一个现代化的金融分析工具箱，使用最新的技术栈构建。

## 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: Gin Web Framework
- **特性**: RESTful API、CORS支持

### 前端
- **框架**: Vue 3 + TypeScript
- **构建工具**: Vite
- **UI库**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router
- **图表**: ECharts + vue-echarts

### 编辑器
- **技术**: React + TypeScript
- **功能**: 可视化策略编辑器

### 量化分析
- **语言**: Python 3.10+
- **库**: pandas, pandas_ta, numpy, scikit-learn
- **功能**: 技术指标计算、回测分析

## 项目结构

```
financial_toolbox/
├── backend/              # Go后端
│   ├── internal/
│   │   ├── api/         # API路由和处理器
│   │   ├── service/     # 业务逻辑
│   │   ├── model/       # 数据模型
│   │   └── middleware/  # 中间件
│   ├── main.go          # 入口文件
│   └── go.mod           # Go模块配置
├── frontend/            # Vue前端
│   ├── src/
│   │   ├── components/  # 组件
│   │   ├── views/       # 页面
│   │   ├── router/      # 路由配置
│   │   ├── store/       # 状态管理
│   │   ├── api/         # API调用
│   │   └── types/       # TypeScript类型
│   ├── package.json     # 依赖配置
│   └── vite.config.ts   # Vite配置
├── editor/              # React编辑器（待开发）
├── python/              # Python量化模块
│   ├── analyzer.py      # 分析器
│   └── requirements.txt # Python依赖
└── docs/               # 文档
```

## 快速开始

### 后端

```bash
cd backend
go mod download
go run main.go
```

后端服务将在 `http://localhost:8080` 启动

### 前端

```bash
cd frontend
npm install
npm run dev
```

前端服务将在 `http://localhost:3000` 启动

### Python模块

```bash
cd python
pip install -r requirements.txt
python analyzer.py
```

## API接口

### 健康检查
```
GET /api/health
```

### 获取股票数据
```
GET /api/stock/:symbol
```

### 计算指标
```
POST /api/indicator/calculate
Content-Type: application/json

{
  "symbol": "AAPL",
  "indicator": "ma",
  "params": {}
}
```

### 获取可用指标
```
GET /api/indicators
```

## 功能特性

- [x] RESTful API
- [x] 技术指标计算（MA、RSI、MACD等）
- [x] 交互式图表（ECharts）
- [x] 策略编辑器
- ] 回测功能
- ] 实时数据
- ] 用户认证

## 开发计划

- [ ] 完善回测功能
- [ ] 添加更多技术指标
- [ ] 实现实时数据推送
- [ ] 添加用户认证系统
- [ ] 优化性能和用户体验

## 贡献指南

欢迎提交Issue和Pull Request！

## 许可证

MIT License
