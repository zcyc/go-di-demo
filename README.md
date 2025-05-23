# Go 依赖注入示例

本项目演示了 Go 中四个流行的依赖注入框架：

1. **Dig**（由 Uber 开发）- 一个基于反射的依赖注入框架
2. **Fx**（由 Uber 开发）- 一个基于依赖注入的应用程序框架，使用Dig
3. **Wire**（由 Google 开发）- 一个编译时依赖注入框架，生成代码
4. **Do**（由 Samber 开发）- 一个轻量级的依赖注入框架，使用泛型提供类型安全

## 项目结构

```
├── cmd/                # 可执行文件
│   ├── dig/            # Dig 示例
│   ├── do/             # Do示例
│   ├── fx/             # Fx 示例
│   └── wire/           # Wire 示例
├── pkg/                # 库代码
│   ├── common/         # 共享服务接口和实现
│   ├── dig/            # Dig 特定代码
│   ├── do/             # Do 特定代码
│   ├── fx/             # Fx 特定代码
│   └── wire/           # Wire 特定代码
├── go.mod
├── go.sum
├── main.go             # 运行所有示例的主入口
└── README.md
```

## 运行示例

运行所有示例：

```bash
go run main.go
```

运行单个示例：

```bash
# Dig示例
go run cmd/dig/main.go

# Fx示例
go run cmd/fx/main.go

# Wire示例（首先需要Wire代码生成）
go run cmd/wire/main.go

# samber/do示例
go run cmd/do/main.go
```

## Wire代码生成

Wire使用前需要生成代码。你可以安装Wire工具并生成wire_gen.go文件：

```bash
go install github.com/google/wire/cmd/wire@latest
cd pkg/wire
wire
```

## 框架之间的主要区别

### Dig
- 使用反射的运行时依赖注入
- 简单的API，专注于依赖解析
- 不需要代码生成

### Fx
- 基于Dig构建，提供应用程序生命周期管理
- 为完整应用提供更多结构和功能
- 包含启动/关闭钩子

### Wire
- 使用代码生成的编译时依赖注入
- 没有运行时反射成本
- 类型安全，具有编译时错误检查

### Do
- 轻量级，API简单直观
- 使用泛型提供类型安全的依赖注入
- 无需反射，性能更好
- 显式依赖管理

## 依赖注入框架对比表

| 特性 | Do | Dig | Wire | Fx |
|------|-----------|-----|------|----|
| 类型安全 | ✅ | ❌ | ✅ | ❌ |
| 编译时检查 | ✅ | ❌ | ✅ | ❌ |
| 运行时注入 | ✅ | ✅ | ❌ | ✅ |
| 生命周期钩子 | ✅ | ❌ | ❌ | ✅ |
| 反射 | ❌ | ✅ | ❌ | ✅ |
| 代码生成 | ❌ | ❌ | ✅ | ❌ |

## 依赖

- Go 1.18+ (Do 需要泛型支持)
- github.com/google/wire
- go.uber.org/dig
- go.uber.org/fx
- github.com/samber/do
