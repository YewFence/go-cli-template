# Go CLI Template

This is a ready-to-use Go CLI project template. It uses Cobra for command organization, mise for toolchain and task management, and includes a VitePress documentation site, GitHub Actions continuous integration, documentation publishing, GitHub Actions update checks, and automated releases.

## Requirements

Install [mise](https://github.com/jdx/mise) first.

Other development tools are declared in [mise.toml](mise.toml). Run `mise install` to install them into the current project environment. If you do not use mise, install the tools manually using the links and versions in `mise.toml`.

This template commits [mise.lock](mise.lock) to pin the resolved tools declared in `mise.toml`. `mise.toml` can declare major versions, minor versions, exact versions, or `latest` based on project needs. To update the toolchain, run `mise lock`, commit the refreshed lockfile, and the CI and Release workflows will install tools from the lockfile for reproducible builds.

## Included Tooling

| Capability | Description |
| --- | --- |
| CLI framework | Cobra is wired in with a root command, a `version` subcommand, and shell completion commands |
| Toolchain | Go, Node, pnpm, actions-up, golangci-lint, and related tools are managed through `mise` |
| Development tasks | Built-in mise tasks include `tidy`, `update`, `test`, `fmt`, `vet`, `lint`, `check`, `build`, and `run` |
| Documentation site | The `docs` directory includes a VitePress documentation site and a GitHub Pages workflow |
| CI checks | GitHub Actions update actions, run Go tests, build the project, audit dependencies, and build the documentation |
| Release workflow | Pushing to `main` creates semantic releases from Conventional Commits. Manual tags and pushed `v*` tags are also supported |

## Quick Start

1. Click `Use this template` to create a new repository, or clone this repository directly.

```bash
git clone https://github.com/YewFence/go-cli-template.git your-cli
cd your-cli
```

2. Trust the mise configuration and install development tools.

```bash
mise trust
mise install
```
3. Run the initialization task.

```bash
mise run init
```

This task asks for the Go module path, command name, GitHub owner or organization, repository name, and project description. It then replaces the template content, removes the template repository origin, and reinitializes the Git history in the current directory so the cloned project starts from a clean `main` branch.

You can also pass all values at once, which is useful for scripted project creation.

```bash
mise run init -- \
  --module github.com/you/your-cli \
  --name your-cli \
  --owner you \
  --repo your-cli \
  --description "Your CLI description"
```

Initialization replaces these template defaults.

| Template default | Replaced with |
| --- | --- |
| `github.com/example/your-cli` | Your Go module path |
| `your-cli` | Your command name |
| `example` | Your GitHub owner or organization |
| `Your CLI description` | Your project description |

4. Clean up template-only files.

After initialization is complete and the replacement result looks correct, remove the template-only initialization tool.

```bash
rm -rf tools/init-template
```

Then remove the `[tasks.init]` section from `mise.toml`.

## Apply To An Existing Project

If you already have a Go project and only want to reuse this template's project configuration, run the existing-project apply tool. It does not modify application code, `go.mod`, README, Git origin, or Git history. It only overwrites `mise.toml`, `.gitignore`, and `.github/workflows`, and downloads the template documentation site only when the current project does not already have a `docs` directory.

Before running it, make sure the Git working tree is clean. The tool also asks you to type `yes` before continuing. After it finishes, use `git diff` to review the changes and keep or adjust them as needed.

```bash
go run github.com/YewFence/go-cli-template/tools/apply-existing@latest
git diff
```

Pass `--ref` to download template files from a specific branch or commit.

```bash
go run github.com/YewFence/go-cli-template/tools/apply-existing@latest --ref main
```

You can also download the single file first and run it locally, which is useful when you want to review the source before executing it.

```bash
tmp="$(mktemp -d)"
curl -fsSL https://raw.githubusercontent.com/YewFence/go-cli-template/main/tools/apply-existing/main.go -o "$tmp/apply-existing.go"
go run "$tmp/apply-existing.go"
```

## Development

See [README.template.md](README.template.md) for details.

## License

[MIT License](LICENSE)
