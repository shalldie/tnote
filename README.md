<!-- 中英文切换 -->
<div align="right">

English | [中文](./README.zh-CN.md)

</div>
<!-- 中英文切换 end -->

<!-- 封面区域 -->
<div align="center">

<img src="https://user-images.githubusercontent.com/9987486/229472271-62a5d923-f7b7-416c-913e-c842ecc2de4d.png" width="320" />

### Cloud Notebook in terminal, based on Gist. 🦋

`终端中运行的云笔记本，基于 Gist 构建。`

[![Release Version](https://img.shields.io/github/v/release/shalldie/tnote?display_name=tag&logo=github&style=flat-square)](https://github.com/shalldie/tnote)
[![Docker Image Version](https://img.shields.io/docker/v/shalldie/tnote/latest?style=flat-square&logo=docker)](https://hub.docker.com/r/shalldie/tnote/tags)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/tnote?label=go&logo=go&style=flat-square)](https://github.com/shalldie/tnote)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/tnote.svg)](https://pkg.go.dev/github.com/shalldie/tnote)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shalldie/tnote/ci.yml?logo=github&style=flat-square)](https://github.com/shalldie/tnote/actions)
[![License](https://img.shields.io/github/license/shalldie/tnote?logo=github&style=flat-square)](https://github.com/shalldie/tnote)

<img src="https://github.com/shalldie/tnote/assets/9987486/57902934-4a58-482a-adff-fafba92dbae9" width="900">

</div>

<!-- 封面区域 end -->

`tnote` is a notebook application running on `Terminal`, allowing you to quickly access, synchronize content, and record your idea on different devices.

- [x] Application 🎯
  - [x] Shortcut keys, mouse operations
  - [x] Cloud storage persistence, support Github/Gitee
  - [x] i18n
- [x] Files
  - [x] Add, delete, check, and modify
- [x] Detail 📝
  - [x] Markdown highlighted
  - [x] Edit and save
- [x] Installation
  - [x] binary
  - [x] go install
  - [x] docker

## Prepare & Config

The application is built on gist, support Github/Gitee.

- [Applying for Github access token](https://github.com/settings/tokens/new), and then add it to the environment variable `TNOTE_GIST_TOKEN`。
- [Applying for Gitee access token](https://gitee.com/profile/personal_access_tokens/new), and then add it to the environment variable `TNOTE_GIST_TOKEN_GITEE`。

```bash
# ~/.bashrc, github
export TNOTE_GIST_TOKEN="<your_access_token>"
```

| Environment Variable                          | Default | Description                                           |
| :-------------------------------------------- | :-----: | :---------------------------------------------------- |
| `TNOTE_GIST_TOKEN` / `TNOTE_GIST_TOKEN_GITEE` |         | `access token` applied for                            |
| `TNOTE_LANG` / `LANG`                         | `en_US` | Language preferred, optional values: `en_US`、`zh_CN` |

## Installation

### 1. binary

[Download](https://github.com/shalldie/tnote/releases), download and execute it, adding to `PATH` would be even better.

| File                    | OS            |
| :---------------------- | :------------ |
| `tnote.darwin-amd64`    | `Mac amd64`   |
| `tnote.darwin-arm64`    | `Mac arm64`   |
| `tnote.linux-amd64`     | `Linux amd64` |
| `tnote.linux-arm64`     | `Linux arm64` |
| `tnote.windows-amd64.exe` | `Windows amd64` |

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

Need `go@1.26+` environment.

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
