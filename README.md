<!-- 封面区域 -->
<div align="center">

<img src="https://user-images.githubusercontent.com/9987486/229472271-62a5d923-f7b7-416c-913e-c842ecc2de4d.png" width="320" />

### Note in terminal, based on github gist.

`终端运行的记事本，基于 github gist 构建。`

[![Release Version](https://img.shields.io/github/v/release/shalldie/tnote?display_name=tag&logo=github&style=flat-square)](https://github.com/shalldie/tnote)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/tnote?label=go&logo=go&style=flat-square)](https://github.com/shalldie/tnote)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/tnote.svg)](https://pkg.go.dev/github.com/shalldie/tnote)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shalldie/tnote/ci.yml?branch=master&label=build&logo=github&style=flat-square)](https://github.com/shalldie/tnote/actions)
[![License](https://img.shields.io/github/license/shalldie/tnote?logo=github&style=flat-square)](https://github.com/shalldie/tnote)

<img src="https://user-images.githubusercontent.com/9987486/229672987-6cc48582-fed0-4537-8192-aa2780cf1316.png" width="800">

</div>

<!-- 封面区域 end -->

> 目前 github api v3，删除包含中文名文件时候报错 <br>
> 已工单沟通，待修复

## 功能

- [x] 面板
  - [x] 切换
  - [x] 存储持久化
  - [x] 快捷键、鼠标
  - [x] Confirm
- [ ] 文件
  - [x] 新建
  - [ ] 重命名
  - [ ] 删除
    - [x] 英文
    - [ ] 中文
- [x] 详情
  - [x] Markdown 高亮
  - [x] 编辑、保存
- [ ] 国际化
  - [ ] 多语言切换
  - [x] 中文文档
  - [ ] 英文文档
- [ ] 安装
  - [x] go install
  - [x] binary
  - [x] docker

## 准备

应用基于 github gist 构建，需要去 [申请 access token](https://github.com/settings/tokens/new)，然后把值加入环境变量 `$TNOTE_GIST_TOKEN`

## 安装

### 1. install 方式

需要 `go@1.19+` 环境

```bash
go install github.com/shalldie/tnote@tag
```

### 2. binary 方式

下载地址：[download](https://github.com/shalldie/tnote/releases)

| 环境           | 适用系统                     |
| :------------- | :--------------------------- |
| `darwin-amd64` | `Mac amd64`、`Mac arm64(M1)` |
| `linux-amd64`  | `Linux amd64`                |
| `linux-arm64`  | `Linux arm64`                |

下载后直接执行即可，加入 `PATH` 更佳。

example:

```bash
wget -O tnote [url]
sudo chmod a+x tnote
sudo mv tnote /usr/local/bin/tnote
```

### 3. docker

```bash
docker run --rm -it -e TNOTE_GIST_TOKEN=$TNOTE_GIST_TOKEN shalldie/tnote
```

## LICENSE

MIT
