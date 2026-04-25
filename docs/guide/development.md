# 开发指南

## 常用命令

```bash
mise run tidy
mise run test
mise run vet
mise run build
```

## 添加命令

在 `cmd` 目录里添加新的 Cobra 命令文件，并在 `init` 函数里通过 `rootCmd.AddCommand` 挂载即可。

## 文档站

```bash
mise run docs:install
mise run docs:dev
mise run docs:build
```

## 更新 CI / CD

```bash
# 使用 Pinact 更新 CI / CD 版本
mise run action:pin
```
