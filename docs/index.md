---
layout: home

hero:
  name: '{{PROJECT_NAME}}'
  text: '{{PROJECT_DESCRIPTION}}'
  tagline: '轻量、可发布的 Go 命令行工具。'
  actions:
    - theme: brand
      text: GitHub Repo
      link: https://github.com/{{GITHUB_OWNER}}/{{REPO_NAME}}
    - theme: alt
      text: Shell 补全
      link: /guide/completion

features:
  - title: 安装
    details: 请参考仓库 README 中的安装方式，按需使用 mise 或从源码构建。
  - title: 使用
    details: README 会覆盖常用命令和基础示例，文档站只补充适合线上查阅的内容。
  - title: Shell 补全
    details: 支持生成 Bash、Zsh、Fish 和 PowerShell 补全脚本。
---
