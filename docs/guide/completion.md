# Shell 补全

模板使用 Cobra 原生补全能力，不需要额外引入 `goreleaser`、`usage` 或其他生成工具。

## 生成补全脚本

构建二进制后，可以通过 `completion` 子命令生成 Bash、Zsh、Fish 和 PowerShell 的补全脚本。

```bash
mise run build
./bin/your-cli completion zsh > _your-cli
./bin/your-cli completion bash > your-cli.bash
./bin/your-cli completion fish > your-cli.fish
./bin/your-cli completion powershell > your-cli.ps1
```

## 安装示例

Zsh 可以把生成的 `_your-cli` 放到 `$fpath` 中已有的目录，或者放到自定义目录后在 `~/.zshrc` 里加入该目录。

```bash
mkdir -p ~/.zsh/completions
./bin/your-cli completion zsh > ~/.zsh/completions/_your-cli
```

```zsh
fpath=(~/.zsh/completions $fpath)
autoload -Uz compinit
compinit
```

Bash 可以把补全脚本放到本地目录后手动 `source`，也可以交给系统的 bash-completion 目录管理。

```bash
mkdir -p ~/.bash_completion.d
./bin/your-cli completion bash > ~/.bash_completion.d/your-cli.bash
source ~/.bash_completion.d/your-cli.bash
```

Fish 可以直接写入用户补全目录。

```bash
mkdir -p ~/.config/fish/completions
./bin/your-cli completion fish > ~/.config/fish/completions/your-cli.fish
```
