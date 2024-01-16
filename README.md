<!-- 封面区域 -->
<div align="center">

<img src="https://user-images.githubusercontent.com/9987486/229472271-62a5d923-f7b7-416c-913e-c842ecc2de4d.png" width="320" />

### Note in terminal, based on github gist. 🧑‍💻

`终端运行的记事本，基于 github gist 构建。`

[![Release Version](https://img.shields.io/github/v/release/shalldie/tnote?display_name=tag&logo=github&style=flat-square)](https://github.com/shalldie/tnote)
[![Docker Image Version](https://img.shields.io/docker/v/shalldie/tnote?label=docker&logo=docker&style=flat-square)](https://github.com/shalldie/tnote)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/tnote?label=go&logo=go&style=flat-square)](https://github.com/shalldie/tnote)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/tnote.svg)](https://pkg.go.dev/github.com/shalldie/tnote)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shalldie/tnote/ci.yml?logo=github&style=flat-square)](https://github.com/shalldie/tnote/actions)
[![License](https://img.shields.io/github/license/shalldie/tnote?logo=github&style=flat-square)](https://github.com/shalldie/tnote)

<img src="https://github.com/shalldie/tnote/assets/9987486/4f7f7b51-766f-49a9-b388-8e40c0692fd2" width="900">

</div>

<!-- 封面区域 end -->

## 功能 🎯

- [x] 面板
  - [x] 快捷键操作
  - [x] 存储持久化
  - [x] Confirm、Prompt、Alert
- [x] 文件
  - [x] 增删查改
- [x] 详情 📝
  - [x] Markdown 高亮
  - [x] 编辑、保存
- [x] 安装
  - [x] binary
  - [x] go install
  - [x] docker

## 准备&前提

应用基于 github gist 构建，需要去 [申请 access token](https://github.com/settings/tokens/new)，然后添加到环境变量 `TNOTE_GIST_TOKEN`。

```bash
# ~/.bashrc
export TNOTE_GIST_TOKEN="<your_access_token>"
```

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
go install github.com/shalldie/tnote
# run
tnote
```

### 3. docker

```bash
docker run -it -e TNOTE_GIST_TOKEN=$TNOTE_GIST_TOKEN shalldie/tnote
```

## LICENSE

MIT
