# Go CLI Template

这是一个开箱即用的 Go 命令行项目模板，默认使用 Cobra 组织命令，使用 mise 管理工具链和常用任务，并内置 VitePress 文档站、GitHub Actions 持续集成、文档发布和标签发布流程。

> [!NOTE]
> 这个仓库是模板项目，创建真实项目后建议先执行初始化，再清理模板专用脚本。

## 模板里有什么

| 能力 | 说明 |
| --- | --- |
| CLI 框架 | 已接入 Cobra，包含根命令、`version` 子命令和 Shell 补全命令 |
| 工具链 | 通过 `mise.toml` 固定 Go、Node、pnpm 和 pinact |
| 开发任务 | 内置 `tidy`、`test`、`vet`、`build`、`run` 等 mise 任务 |
| 文档站 | `docs` 目录内置 VitePress 文档站和 GitHub Pages 工作流 |
| 发布流程 | 推送 `v*` 标签后自动构建多平台二进制并发布到 GitHub Release |

## 创建项目

可以在 GitHub 页面点击 `Use this template` 创建新仓库，也可以直接克隆后改成自己的仓库地址。

```bash
git clone https://github.com/example/go-cli-template.git your-cli
cd your-cli
```

准备本地工具链。

```bash
mise trust
mise install
```

## 初始化模板

交互式初始化会提示输入 Go 模块路径、命令名、GitHub 所属账号或组织、仓库名和项目描述。

```bash
mise run init
```

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

## 清理模板

初始化完成并确认替换结果没有问题后，可以删除模板专用初始化工具。

```bash
rm -rf tools/init-template
```

然后从 `mise.toml` 删除 `[tasks.init]` 这一段，避免真实项目里继续保留模板初始化任务。

最后检查是否还有占位内容残留。

```bash
git grep -n "github.com/example/your-cli\\|your-cli\\|example\\|Your CLI description"
```

## 开发

运行命令行程序。

```bash
mise run run
```

运行测试和静态检查。

```bash
mise run tidy
mise run test
mise run vet
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

## 发布

推送 `v*` 标签后，GitHub Actions 会构建多平台二进制文件并发布到 GitHub Release。

```bash
git tag v0.1.0
git push origin v0.1.0
```

## 许可证

[MIT License](LICENSE)
