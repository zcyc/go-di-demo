# Go依赖注入演示

本项目演示了Go中三个流行的依赖注入框架：

1. **Dig**（由Uber开发）- 一个基于反射的依赖注入框架
2. **Fx**（由Uber开发）- 一个基于依赖注入的应用程序框架，使用Dig
3. **Wire**（由Google开发）- 一个编译时依赖注入框架，生成代码

## 项目结构

```
├── cmd/                # 命令可执行文件
│   ├── dig/            # Dig示例
│   ├── fx/             # Fx示例
│   └── wire/           # Wire示例
├── pkg/                # 库代码
│   ├── common/         # 共享服务接口和实现
│   ├── dig/            # Dig特定代码
│   ├── fx/             # Fx特定代码
│   └── wire/           # Wire特定代码
├── go.mod              # Go模块定义
├── go.sum              # Go模块校验和
├── main.go             # 运行所有示例的主入口点
└── README.md           # 本文件
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
- 适用于简单应用或需要运行时灵活性的场景

### Fx
- 基于Dig构建，提供应用程序生命周期管理
- 为完整应用提供更多结构和功能
- 包含启动/关闭钩子
- 适合复杂应用

### Wire
- 使用代码生成的编译时依赖注入
- 没有运行时反射成本
- 类型安全，具有编译时错误检查
- 适用于性能关键型应用

## 依赖

- Go 1.16+
- github.com/google/wire
- go.uber.org/dig
- go.uber.org/fx