<!-- 封面区域 -->
<div align="center">

<img src="images/ttm.png" width="320" height="136" />

### Task manager in terminal.

`终端运行的任务管理工具`

[![Release Version](https://img.shields.io/github/v/release/shalldie/ttm?display_name=tag&logo=github&style=flat-square)](https://github.com/shalldie/ttm)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/ttm?label=go&logo=go&style=flat-square)](https://github.com/shalldie/ttm)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/ttm.svg)](https://pkg.go.dev/github.com/shalldie/ttm)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/ttm/ci?label=build&logo=github&style=flat-square)](https://github.com/shalldie/ttm/actions)
[![License](https://img.shields.io/github/license/shalldie/ttm?logo=github&style=flat-square)](https://github.com/shalldie/ttm)

<img src="https://user-images.githubusercontent.com/9987486/206672150-24f34157-72e0-4c81-929c-ee07eb668ac8.png" width="1000">

</div>

<!-- 封面区域 end -->

## 功能

- [x] 面板
  - [x] 切换
  - [x] 存储持久化
  - [x] 快捷键
- [ ] 项目
  - [x] 新建
  - [ ] 重命名
  - [x] 删除
- [ ] 任务
  - [x] 新建
  - [ ] 重命名
  - [x] 删除
  - [ ] 状态、筛选
- [ ] 详情
  - [x] 新建
  - [x] Markdown 高亮、编辑、保存

## 操作

你可以用快捷键或者鼠标来操作。

## 安装

### 1. install 方式

需要 `go@1.18+` 环境

```bash
go install github.com/shalldie/ttm@latest
```

### 2. binary 方式

下载地址：[download](https://github.com/shalldie/ttm/releases)

| 环境           | 适用系统                     |
| :------------- | :--------------------------- |
| `darwin-amd64` | `Mac amd64`、`Mac arm64(M1)` |
| `linux-amd64`  | `Linux amd64`                |
| `linux-arm64`  | `Linux arm64`                |

下载后直接执行即可，加入 `PATH` 更佳。

example:

```bash
wget -O ttm [url]
sudo chmod a+x ttm
sudo mv ttm /usr/local/bin/ttm
```

## LICENSE

MIT
