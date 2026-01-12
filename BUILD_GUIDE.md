# Financial Toolbox - 桌面应用程序构建指南

## 项目概述

Financial Toolbox 是一个基于 Wails v2 框架的桌面应用程序，集成了 Go 后端、Vue 3 前端和 Python 量化分析模块。

## 技术栈

- **后端**: Go 1.21+ + Wails v2
- **前端**: Vue 3 + TypeScript + Element Plus + ECharts
- **量化**: Python 3.10+（可选）
- **构建**: Wails v2 自动打包

## 项目结构

```
financial_toolbox/
├── wails.json              # Wails 配置文件
├── main.go                 # Wails 入口文件
├── app.go                  # 应用程序逻辑和绑定方法
├── internal/                # 内部包
│   ├── api/             # API 相关
│   │   └── handlers/    # HTTP 处理器（可选）
│   ├── service/          # 业务逻辑
│   ├── model/            # 数据模型
│   ├── middleware/       # 中间件
│   └── python/           # Python 集成
├── frontend/               # Vue 前端
│   ├── src/
│   │   ├── main.ts       # 前端入口
│   │   ├── App.vue       # 主应用组件
│   │   ├── components/    # Vue 组件
│   │   ├── assets/       # 静态资源
│   │   └── wailsjs/      # Wails 生成的绑定
│   ├── index.html        # HTML 入口
│   ├── package.json      # NPM 依赖
│   └── vite.config.ts    # Vite 配置
├── python/                 # Python 量化模块
│   ├── analyzer.py       # 量化分析器
│   └── requirements.txt  # Python 依赖
├── build.bat               # Windows 构建脚本
└── build/                 # 构建输出目录
    └── bin/
        └── financial-toolbox.exe  # 最终可执行文件
```

## 开发环境设置

### 前置要求

1. **Go 1.21+**
   - 下载: https://golang.org/dl/
   - 安装后确保 `go` 命令可用

2. **Node.js 18+**
   - 下载: https://nodejs.org/
   - 安装后确保 `npm` 命令可用

3. **Python 3.10+**（可选）
   - 下载: https://www.python.org/downloads/
   - 安装后确保 `python` 命令可用

4. **Wails CLI**
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

### 安装依赖

#### 后端依赖
```bash
go mod download
go mod tidy
```

#### 前端依赖
```bash
cd frontend
npm install
```

#### Python 依赖（可选）
```bash
cd python
pip install -r requirements.txt
```

## 开发模式

### 启动开发服务器

```bash
wails dev
```

这将：
1. 启动 Go 后端
2. 启动 Vite 开发服务器
3. 生成 Wails 绑定代码
4. 自动重新加载前端更改

### 开发模式端口

- **前端**: http://localhost:34115（自动分配）
- **后端**: 通过 Wails IPC 自动通信

## 构建应用程序

### Windows 构建

#### 方法1：使用构建脚本（推荐）

```bash
build.bat
```

#### 方法2：使用 Wails CLI

```bash
wails build -platform windows/amd64
```

### 构建输出

构建完成后，可执行文件位于：
```
build/bin/financial-toolbox.exe
```

### 构建选项

```bash
# 清理构建
wails build -clean

# 指定平台
wails build -platform windows/amd64
wails build -platform windows/arm64
wails build -platform darwin/amd64    # macOS
wails build -platform darwin/arm64     # macOS Apple Silicon
wails build -platform linux/amd64      # Linux

# 跳过前端构建
wails build -skipfrontend

# 调试模式
wails build -debug
```

## 运行应用程序

### 开发模式

```bash
wails dev
```

### 生产模式

双击 `build/bin/financial-toolbox.exe` 文件启动应用程序。

## 功能特性

### 已实现功能

- ✅ Go 后端核心逻辑
- ✅ Vue 3 前端界面
- ✅ ECharts 图表集成
- ✅ Element Plus UI 组件
- ✅ 技术指标计算（MA、RSI、MACD 等）
- ✅ 策略编辑器
- ✅ Python 子进程管理
- ✅ Wails IPC 通信
- ✅ 跨平台构建支持

### 待实现功能

- ⏳ 实时股票数据
- ⏳ 数据持久化
- ⏳ 用户认证
- ⏳ 策略回测
- ⏳ 更多技术指标
- ⏳ 性能优化

## API 接口

### Go 绑定方法

前端可以通过 `window.$wails` 调用以下方法：

```typescript
// 获取股票数据
const data = await window.$wails.GetStockData(symbol: string)

// 计算指标
const result = await window.$wails.CalculateIndicator(reqJSON: string)

// 获取可用指标
const indicators = await window.$wails.GetAvailableIndicators()

// 执行 Python 脚本
const output = await window.$wails.ExecutePythonScript(scriptPath: string, inputData: string)

// 获取应用版本
const version = await window.$wails.GetAppVersion()

// 获取系统信息
const info = await window.$wails.GetSystemInfo()
```

## 故障排除

### 常见问题

#### 1. Wails dev 启动失败

**问题**: `no go.mod file found`

**解决**: 确保 `go.mod` 文件在项目根目录

#### 2. 前端构建失败

**问题**: `npm install` 失败

**解决**: 
```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
```

#### 3. Python 脚本执行失败

**问题**: `python: command not found`

**解决**: 确保 Python 已安装并添加到 PATH

#### 4. 构建的 exe 文件无法运行

**问题**: 缺少依赖或运行时错误

**解决**: 
- 使用 `-clean` 选项重新构建
- 检查 Windows 防火墙设置
- 在开发模式下测试

## 性能优化

### 减小应用体积

1. **前端优化**
   ```bash
   cd frontend
   npm run build
   ```
   Vite 会自动进行代码分割和压缩

2. **Go 优化**
   ```bash
   wails build -ldflags "-s -w"
   ```
   使用链接器标志减小二进制大小

3. **资源压缩**
   - 压缩图片和静态资源
   - 使用 gzip 压缩

### 加速启动时间

1. **延迟加载**
   - 按需加载 ECharts 组件
   - 懒加载 Python 模块

2. **缓存优化**
   - 使用浏览器缓存
   - 缓存计算结果

## 部署

### Windows 安装程序

使用 Inno Setup 或 NSIS 创建安装程序：

1. 创建安装程序脚本
2. 包含 `financial-toolbox.exe`
3. 添加快捷方式
4. 配置卸载程序

### 便携版本

直接分发 `financial-toolbox.exe`，无需安装。

## 许可证

MIT License

## 贡献指南

欢迎提交 Issue 和 Pull Request！

## 联系方式

- Email: financial-toolbox@example.com
- GitHub: https://github.com/Berry-Wahlberg/financial_toolbox
