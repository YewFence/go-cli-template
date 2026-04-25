# 发布流程

## GitHub Release

发布流程使用 GitHub Actions 直接调用 `go build`，推送 `v*` 标签后自动构建 Linux、macOS 和 Windows 可执行文件并上传到 GitHub Release。

```bash
git tag v0.1.0
git push origin v0.1.0
```

## 本地构建

```bash
mise run build
```
