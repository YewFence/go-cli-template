# {{PROJECT_NAME}}

[![Release](https://img.shields.io/github/v/release/{{GITHUB_OWNER}}/{{REPO_NAME}}?sort=semver)](https://github.com/{{GITHUB_OWNER}}/{{REPO_NAME}}/releases)
[![Docs](https://img.shields.io/badge/docs-online-blue)](https://{{GITHUB_OWNER}}.github.io/{{REPO_NAME}}/)
[![License](https://img.shields.io/github/license/{{GITHUB_OWNER}}/{{REPO_NAME}})](LICENSE)

{{PROJECT_DESCRIPTION}}

> [!NOTE]
> This project is in an early development stage. Core features may be missing, and backward compatibility is not guaranteed.

## Quick Start

### Installation

#### Mise

```bash
# Applies only to the current directory. Add -g to install globally.
mise use github:{{GITHUB_OWNER}}/{{REPO_NAME}}
```

#### Build From Source

```bash
go install github.com/{{GITHUB_OWNER}}/{{REPO_NAME}}/cmd/{{PROJECT_NAME}}@latest
```

#### Build Locally

```bash
git clone https://github.com/{{GITHUB_OWNER}}/{{REPO_NAME}}.git
cd {{REPO_NAME}}
mise trust
mise install
mise run build
```

### Usage

```bash
{{PROJECT_NAME}}
{{PROJECT_NAME}} version
```

Generate shell completion scripts.

```bash
{{PROJECT_NAME}} completion zsh > _{{PROJECT_NAME}}
{{PROJECT_NAME}} completion bash > {{PROJECT_NAME}}.bash
{{PROJECT_NAME}} completion fish > {{PROJECT_NAME}}.fish
{{PROJECT_NAME}} completion powershell > {{PROJECT_NAME}}.ps1
```

## Documentation

See the [documentation site](https://{{GITHUB_OWNER}}.github.io/{{REPO_NAME}}) for more information.

## Development

### Requirements

[mise](https://github.com/jdx/mise) is recommended for managing development tools.

The development tools required by this project are declared in [mise.toml](mise.toml). Run `mise install` to install them into the current project environment. If you do not use mise, install the tools manually using the links and versions in `mise.toml`.

This project commits [mise.lock](mise.lock) to pin the resolved tools declared in `mise.toml`. To update the toolchain, run `mise lock`, commit the refreshed lockfile, and the CI and Release workflows will install tools from the lockfile for reproducible builds.

This project is intentionally shaped as a single-command CLI. The executable entrypoint lives under `cmd/{{PROJECT_NAME}}`, and the mise build, run, install, and release tasks use `./cmd/...` so Go derives the binary name from that entrypoint directory. If you add more executable entrypoints under `cmd/`, adjust those tasks and the release packaging logic at the same time.

### Common Commands

Run `mise tasks` to see the full task list.

#### CLI

```bash
# Run the CLI
mise run cli
# Install the CLI into your local Go bin directory
mise run cli:install
# Tidy Go module dependencies
mise run tidy
# Run formatting checks, static checks, build, and lint
mise run check
# Build a local executable. Build artifacts are written to the `bin/` directory.
mise run build
```

#### Documentation Site

```bash
# Install dependencies
mise run docs:install
# Start the local documentation development server
mise run docs:dev
```

#### GitHub Actions Maintenance

GitHub Actions are automatically updated by [this workflow](.github/workflows/actions-up.yml) when a pull request is opened in this repository. You can also update GitHub Actions versions interactively with the following command.

> Pull requests opened from other repositories only check Action versions and do not update them automatically.

```bash
mise run action:update
```

#### Release

After pushing to `main`, the Release workflow derives the version from Conventional Commits. When there are releasable changes, it automatically creates a `v*` tag, builds multi-platform binaries, and publishes them to GitHub Releases.

```bash
git push origin main
```

You can also push a specific `v*` tag, or trigger the Release workflow manually from the GitHub Actions page and enter the tag to publish.

```bash
git tag v0.1.0
git push origin v0.1.0
```

## License

[MIT License](LICENSE)
