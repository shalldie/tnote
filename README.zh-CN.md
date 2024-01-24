<!-- 中英文切换 -->
<div align="right">

[English](./README.md) | 中文

</div>
<!-- 中英文切换 end -->

<!-- 封面区域 -->
<div align="center">

<img src="https://user-images.githubusercontent.com/9987486/229472271-62a5d923-f7b7-416c-913e-c842ecc2de4d.png" width="320" />

### Note in terminal, based on Github Gist. 🧑‍💻

`终端运行的记事本，基于 Github Gist 构建。`

[![Release Version](https://img.shields.io/github/v/release/shalldie/tnote?display_name=tag&logo=github&style=flat-square)](https://github.com/shalldie/tnote)
[![Docker Image Version](https://img.shields.io/docker/v/shalldie/tnote/latest?style=flat-square&logo=docker)](https://hub.docker.com/r/shalldie/tnote/tags)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/tnote?label=go&logo=go&style=flat-square)](https://github.com/shalldie/tnote)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/tnote.svg)](https://pkg.go.dev/github.com/shalldie/tnote)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shalldie/tnote/ci.yml?logo=github&style=flat-square)](https://github.com/shalldie/tnote/actions)
[![License](https://img.shields.io/github/license/shalldie/tnote?logo=github&style=flat-square)](https://github.com/shalldie/tnote)

<img src="https://github.com/shalldie/tnote/assets/9987486/1964cbc5-3a10-47a1-9d13-66f8debb8ad0" width="900">

</div>

<!-- 封面区域 end -->

`tnote` 是一个运行在 `Terminal` 的记事本应用程序，让你可以在不同设备快速访问、同步内容，记录自己的生活。

- [x] 应用 🎯
  - [x] 快捷键、鼠标操作
  - [x] 存储云端持久化
  - [x] 国际化
- [x] 文件
  - [x] 增删查改
- [x] 详情 📝
  - [x] Markdown 高亮
  - [x] 编辑、保存
- [x] 安装
  - [x] binary
  - [x] go install
  - [x] docker

## 准备&配置

应用基于 github gist 构建，需要去 [申请 access token](https://github.com/settings/tokens/new)，然后添加到环境变量 `TNOTE_GIST_TOKEN`。

```bash
# ~/.bashrc
export TNOTE_GIST_TOKEN="<your_access_token>"
```

| 环境变量              | 默认值  | 描述                                 |
| :-------------------- | :-----: | :----------------------------------- |
| `TNOTE_GIST_TOKEN`    |         | 申请到的 access token                |
| `TNOTE_LANG` / `LANG` | `en_US` | 使用的语言，可选值：`en_US`、`zh_CN` |

## 安装&运行

### 1. binary

[Download](https://github.com/shalldie/tnote/releases)，下载后直接执行即可，加入 `PATH` 更佳。

| 文件                 | 适用系统                 |
| :------------------- | :----------------------- |
| `tnote.darwin-amd64` | `Mac amd64`、`Mac arm64` |
| `tnote.linux-amd64`  | `Linux amd64`            |
| `tnote.linux-arm64`  | `Linux arm64`            |

example:

```bash
# install
wget -O tnote [url]
sudo chmod a+x tnote
sudo mv tnote /usr/local/bin/tnote
# run
tnote
```

### 2. go install

需要 `go@1.20+` 环境

```bash
# install
go install github.com/shalldie/tnote@latest
# run
tnote
```

### 3. docker

```bash
docker run -it -e TNOTE_GIST_TOKEN=$TNOTE_GIST_TOKEN shalldie/tnote:latest
```

## LICENSE

MIT
