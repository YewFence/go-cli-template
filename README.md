# Go CLI Template

这是一个开箱即用的 Go 命令行项目模板，默认使用 Cobra 组织命令，使用 mise 管理工具链和常用任务，并内置 VitePress 文档站、GitHub Actions 持续集成、文档发布、Actions 更新检查和自动发布流程

## 依赖

只需要先安装 [mise](https://github.com/jdx/mise)。

其他开发工具由 [mise.toml](mise.toml) 声明，执行 `mise install` 即可安装到当前项目环境。不使用 mise 时，请参考 `mise.toml` 中的工具链接和版本自行安装。

本模板提交 [mise.lock](mise.lock) 来固定 `mise.toml` 中声明的工具解析结果。`mise.toml` 可以按项目需要声明主版本、次版本、精确版本或 `latest`，开发者想更新工具链时可以运行 `mise lock` 刷新锁文件并提交变更，CI 和 Release 工作流会使用锁文件安装工具，保证构建可复现。

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

## 应用到已有项目

如果已有 Go 项目只想复用本模板的工程配置，可以运行已有项目应用工具。该工具不会修改业务代码、`go.mod`、README、Git origin 或 Git 历史，只会覆盖 `mise.toml`、`.gitignore` 和 `.github/workflows`，并且仅在当前项目没有 `docs` 目录时下载模板文档站。

运行前必须确保 Git 工作区是干净的，工具也会要求输入 `yes` 才继续。运行完成后请使用 `git diff` 查看变更，并按项目需要保留或调整。

```bash
go run github.com/YewFence/go-cli-template/tools/apply-existing@latest
git diff
```

如果需要从指定分支或提交下载模板文件，可以传入 `--ref`。

```bash
go run github.com/YewFence/go-cli-template/tools/apply-existing@latest --ref main
```

也可以先下载单文件再运行，适合想先审查源码的场景。

```bash
tmp="$(mktemp -d)"
curl -fsSL https://raw.githubusercontent.com/YewFence/go-cli-template/main/tools/apply-existing/main.go -o "$tmp/apply-existing.go"
go run "$tmp/apply-existing.go"
```

## 开发

详情请参考 [README.template.md](README.template.md)

## 许可证

[MIT License](LICENSE)
