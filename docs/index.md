---
layout: home

hero:
  name: 'your-cli'
  text: Go CLI 模板
  tagline: 'Your CLI description'
  actions:
    - theme: brand
      text: 快速开始
      link: /guide/getting-started
    - theme: alt
      text: 开发指南
      link: /guide/development
    - theme: alt
      text: GitHub Repo
      link: https://github.com/example/your-cli

features:
  - title: Cobra 命令行
    details: 模板默认提供根命令、version 子命令和 Cobra 原生 Shell 补全，适合继续扩展业务命令。
  - title: mise 工作流
    details: Go、Node 和 pnpm 版本统一写在 mise.toml，初始化、测试、构建和文档命令都可以一键运行。
  - title: 文档站内置
    details: 使用 VitePress 提供轻量文档站，并带 GitHub Pages 部署工作流。
  - title: 发布足够简单
    details: 推送 v* 标签即可触发 GitHub Actions，使用 go build 构建多平台二进制并上传到 GitHub Release。
---
