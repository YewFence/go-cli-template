# 快速开始

## 初始化模板

```bash
mise trust
mise install
mise run init
mise run tidy
```

如果要在脚本里初始化，可以直接传入参数跳过交互。

```bash
mise run init -- --module github.com/you/your-cli --name your-cli --owner you --repo your-cli --description "Your CLI description"
mise run tidy
```

## 运行 CLI

```bash
mise run run
mise run run -- version
```

## 构建二进制

```bash
mise run build
./bin/your-cli version
```
