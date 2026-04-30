# Go CLI Template

这是一个开箱即用的 Go 命令行项目模板，默认使用 Cobra 组织命令，使用 mise 管理工具链和常用任务，并内置 VitePress 文档站、GitHub Actions 持续集成、文档发布、Actions 更新检查和自动发布流程

## 依赖

[mise](https://github.com/jdx/mise)

## 模板内置工具

| 能力 | 说明 |
| --- | --- |
| CLI 框架 | 已接入 Cobra，包含根命令、`version` 子命令和 Shell 补全命令 |
| 工具链 | 通过 `mise` 引入 Go、Node、pnpm、actions-up 和 golangci-lint 等工具 |
| 开发任务 | 内置 `tidy`、`update`、`test`、`fmt`、`vet`、`lint`、`check`、`build`、`run` 等 mise 任务 |
| 文档站 | `docs` 目录内置 VitePress 文档站和 GitHub Pages 工作流 |
| CI 检查 | GitHub Actions 会自动更新 Actions、运行 Go 测试、构建、依赖审计和文档构建 |
| 发布流程 | 推送到 `main` 后按约定式提交自动发布语义化版本，也支持手动指定标签或推送 `v*` 标签发布 |

## 快速开始

1. 点击 `Use this template` 创建新仓库，也可以直接克隆。

```bash
git clone https://github.com/YewFence/go-cli-template.git your-cli
cd your-cli
```

2. 信任 mise 配置文件并安装开发工具

```bash
mise trust
mise install
```
3. 运行初始化任务

```bash
mise run init
```

该任务会询问 Go 模块路径、命令名、GitHub 所属账号或组织、仓库名和项目描述。然后替换模板内容，移除模板仓库的 origin，并重新初始化当前目录的 Git 历史，让克隆下来的项目从干净的 `main` 分支开始。

也可以一次性传入参数，适合脚本化创建项目。

```bash
mise run init -- \
  --module github.com/you/your-cli \
  --name your-cli \
  --owner you \
  --repo your-cli \
  --description "Your CLI description"
```

初始化会替换项目里的这些模板默认值。

| 模板默认值 | 替换为 |
| --- | --- |
| `github.com/example/your-cli` | 你的 Go 模块路径 |
| `your-cli` | 你的命令名 |
| `example` | 你的 GitHub 所属账号或组织 |
| `Your CLI description` | 你的项目描述 |

4. 清理模板

初始化完成并确认替换结果没有问题后，可以删除模板专用初始化工具。

```bash
rm -rf tools/init-template
```

然后从 `mise.toml` 删除 `[tasks.init]` 一段即可

## 开发

运行命令行程序。

```bash
mise run run
```

整理或更新 Go 依赖。

```bash
mise run tidy
mise run update
```

运行测试、格式检查、静态检查和 lint。

```bash
mise run test
mise run fmt
mise run vet
mise run lint
```

也可以使用聚合任务运行格式检查、静态检查、构建和 lint。

```bash
mise run check
```

构建本地二进制文件。

```bash
mise run build
```

构建产物会输出到 `bin/` 目录。

## 使用命令

初始化后的项目默认包含根命令和版本命令。

```bash
your-cli
your-cli version
```

生成 Shell 补全脚本。

```bash
your-cli completion zsh > _your-cli
your-cli completion bash > your-cli.bash
your-cli completion fish > your-cli.fish
your-cli completion powershell > your-cli.ps1
```

## 文档站

安装文档依赖。

```bash
mise run docs:install
```

启动本地文档站。

```bash
mise run docs:dev
```

构建文档站。

```bash
mise run docs:build
```

## GitHub Actions 维护

CI 会检查工作流里的 Action 是否有可更新版本，发现更新时会失败并在摘要中列出结果。

```bash
mise run action:check
```

交互式更新 GitHub Action 版本，或自动更新全部 Action 版本。

```bash
mise run action:update
mise run action:update:all
```

## 发布

推送到 `main` 后，Release 工作流会根据 Conventional Commits 解析版本；当存在需要发布的变更时，会自动创建 `v*` 标签、构建多平台二进制文件并发布到 GitHub Release。

```bash
git push origin main
```

也可以推送指定的 `v*` 标签，或在 GitHub Actions 页面手动触发 Release 工作流并输入要发布的标签。

```bash
git tag v0.1.0
git push origin v0.1.0
```

## 许可证

[MIT License](LICENSE)
