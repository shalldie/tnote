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

<img src="https://github.com/shalldie/tnote/assets/9987486/05fcde85-eb08-4caf-9bd8-ef955c93cabc" width="800">

</div>

<!-- 封面区域 end -->

## 功能 🎯

- [x] 面板
  - [x] 切换
  - [x] 存储持久化
  - [x] 快捷键、鼠标
  - [x] Confirm
  - [x] Prompt
- [x] 文件
  - [x] 新建
  - [x] 重命名
  - [x] 删除
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

下载地址：[download](https://github.com/shalldie/tnote/releases)

| 环境           | 适用系统                     |
| :------------- | :--------------------------- |
| `darwin-amd64` | `Mac amd64`、`Mac arm64(M1)` |
| `linux-amd64`  | `Linux amd64`                |
| `linux-arm64`  | `Linux arm64`                |

下载后直接执行即可，加入 `PATH` 更佳。

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
docker run --rm -it -e TNOTE_GIST_TOKEN=$TNOTE_GIST_TOKEN shalldie/tnote
```

## LICENSE

MIT
