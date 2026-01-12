# 个人金融工具箱

## 项目目的

个人金融工具箱是一个综合应用程序，旨在提供高级金融分析功能。它利用 Go 语言进行后端 API 开发，使用 Python 进行复杂数据分析，使用户能够对股票数据执行技术分析，包括 MACD、RSI 和回测策略等指标。

## 架构设计

- **后端**: Go 1.21+ (控制器和服务层)
- **数据分析**: Python 3.10+ (用于复杂计算的 Worker 脚本)
- **通信方式**: Go 和 Python 进程之间通过 JSON 数据交换
- **API 框架**: Gin Web 框架

## 安装说明

### 前置条件

- Go 1.21 或更高版本
- Python 3.10 或更高版本
- Git

### 安装步骤

1. **克隆仓库**
   ```bash
   git clone <repository-url>
   cd financial_toolbox
   ```

2. **安装 Go 依赖**
   ```bash
   go mod download
   ```

3. **安装 Python 依赖**
   ```bash
   pip install -r scripts/requirements.txt
   ```

## 使用指南

### 运行应用

1. **启动 Go 服务器**
   ```bash
   go run main.go
   ```

2. **访问 API**
   服务器将在 `http://localhost:8080` 启动

### API 端点

#### GET /api/indicator

计算给定股票代码的技术指标。

**参数**:
- `symbol`: 股票代码 (例如: "AAPL", "GOOGL")
- `indicator`: 指标类型 (例如: "macd")

**示例请求**:
```bash
curl "http://localhost:8080/api/indicator?symbol=AAPL&indicator=macd"
```

**示例响应**:
```json
{
  "success": true,
  "result": "[{\"date\":\"2023-01-01\",\"MACD_12_26_9\":null,\"MACDh_12_26_9\":null,\"MACDs_12_26_9\":null},{\"date\":\"2023-01-02\",\"MACD_12_26_9\":null,\"MACDh_12_26_9\":null,\"MACDs_12_26_9\":null},...]"
}
```

## 功能特性

- **MACD 指标**: 移动平均收敛发散指标计算
- **RESTful API**: 易于使用的 API，用于访问金融指标
- **模块化架构**: 后端和数据分析层分离
- **错误处理**: 全面的错误处理，确保稳健运行
- **超时保护**: 防止 Python 脚本无响应导致的挂起

## 计划功能

- **更多指标**: RSI、布林带、简单移动平均线、指数移动平均线等
- **回测框架**: 针对历史数据测试交易策略
- **实时数据集成**: 连接到金融数据 API 获取实时数据
- **可视化**: 生成指标结果的图表和图形
- **用户认证**: API 的安全访问

## 贡献指南

### 分支策略

本项目采用简化的 Git Flow 分支策略，结合基于主干的开发：

1. **主分支**: 始终代表生产就绪代码
2. **功能分支**: 为新开发工作创建
   - 命名约定: `feature/feature-name`
   - 短期分支 (通常 1-3 天)
3. **代码审查**: 所有功能分支在合并到主分支之前需要审查
4. **测试**: 确保所有更改在集成前经过测试

### 贡献工作流

1. **创建功能分支**
   ```bash
   git checkout main
   git pull origin main
   git checkout -b feature/your-feature-name
   ```

2. **进行更改**并**提交**
   ```bash
   git add .
   git commit -m "描述你的更改"
   ```

3. **推送**分支到仓库
   ```bash
   git push origin feature/your-feature-name
   ```

4. **创建 Pull Request** 进行代码审查

5. **合并**经过批准和测试成功的分支

### 代码规范

- **Go 代码**: 遵循 `gofmt` 标准，包含中文注释
- **Python 代码**: 遵循 PEP 8 风格指南
- **文档**: 为新功能更新文档
- **测试**: 为新功能添加测试

## 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

## 联系方式

如有问题或建议，请在仓库中打开 Issue。