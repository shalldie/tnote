# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

`tnote` 是一个运行在终端里的云笔记本（TUI），基于 Gist 构建。所有笔记以文件形式存储在同一个专用的私有 Gist 中（通过特殊的 description 标识），可在多设备间同步。后端同时支持 GitHub 与 Gitee。

## 常用命令

```bash
# 本地运行（需在环境变量中设置 TNOTE_GIST_TOKEN 或 TNOTE_GIST_TOKEN_GITEE）
go run main.go

# 交叉编译发布产物到 ./output（linux/darwin/windows，amd64/arm64），并用 upx 压缩
bash ./scripts/build.sh

# 单目标构建
go build -o output/tnote main.go

# 静态检查 / 格式化
go vet ./...
gofmt -l .
```

仓库中没有测试代码。需要 `go@1.26+`（见 `go.mod`）。

### 必需的环境变量
- `TNOTE_GIST_TOKEN`（GitHub token）或 `TNOTE_GIST_TOKEN_GITEE`（Gitee token）：至少设置一个，否则启动时直接退出（`internal/conf/conf.go`）。若两者都设置，启动时会提示选择平台。
- `TNOTE_LANG` / `LANG`：`en_US`（默认）或 `zh_CN`。

## 架构

### Bubble Tea（Elm 架构）
UI 基于 charmbracelet 的 **v2** 系列。注意导入路径：bubbletea/bubbles/lipgloss/glamour 均为 vanity 域名 `charm.land/*/v2`，bubblezone 例外，是 `github.com/lrstanley/bubblezone/v2`。每个组件都遵循 `Init` / `Update(msg) (model, cmd)` / `View()` 三元组，且都是**值类型**——`Update` 返回的是新副本，因此修改后必须重新赋值回去（参见 `internal/app/app.go` 的 `AppModel.propagate`）。根模型 `AppModel` 组合了四个子组件并向它们转发消息：`FileList`、`FilePanel`、`StatusBar`、`Dialog`。

> v2 关键差异（迁移易踩坑）：只有根模型 `AppModel.View()` 返回 `tea.View`，且程序级开关（AltScreen、鼠标模式、窗口标题）是在这个 `tea.View` 的字段上设置的（不再走 `NewProgram` options）；**所有子组件的 `View()` 仍返回 `string`**。鼠标消息按类型分裂为 `tea.MouseClickMsg`/`tea.MouseWheelMsg` 等（不再是单结构体 + Action）；键盘消息用 `tea.KeyPressMsg`（`tea.KeyMsg` 是接口）。`lipgloss.Style.Width(n)` 在 v2 是**含边框在内**的目标宽度（`BoxModel.Render` 据此用 `Width(m.Width)`）。

### store：全局状态 + 消息总线（`internal/app/store`）
这是串联各组件的核心，是理解本项目的关键：

- `store.State`：全局单例，保存跨组件的 UI 状态：`Editing`、`InputFocus`、`DialogMode`，以及当前选中的文件 `file`（通过 `GetFile`/`SetFile` 加互斥锁保护）。
- `store.Gist`：全局 Gist 客户端。
- `store.Send(cmd)`：中央分发器。它先对特定命令类型更新 `State`（例如 `CMD_INVOKE_EDIT` 会设置 `State.Editing`），再通过 `store.SendImpl` 把命令转发给 Bubble Tea 程序；`app.Run()` 会把 `SendImpl` 接到 `app.Send`。这也是非 UI 协程把消息推入事件循环的方式。
- 命令（`commands.go`）被定义为**各自独立的 Go 类型**（如 `type CMD_REFRESH_FILES string`、`type CMD_APP_FOCUS int`）。组件在其 `Update` 中通过类型 switch 做出响应。新增跨组件动作的做法：在 `commands.go` 定义类型 → `store.Send` 它 → 在相关组件的 `Update` 中处理。

### 并发模式
阻塞性工作（对 Gist 的网络 I/O）通过 `go store.Send(...)` / `go m.someAction()` 异步派发，以保持 UI 线程响应。回调再把结果命令（如 `CMD_REFRESH_FILES`、`StatusPayload`）`store.Send` 回事件循环。注意这个模式——用 `go` 启动的组件方法运行在 Bubble Tea 线程之外，只能通过 `store.Send` 通信。

**触网的后台任务必须用 `store.SafeGo(fn)` 启动**（`internal/app/store/safego.go`），它在 `go` 外面包了一层 `recover` 兜底：任何 panic（如网络超时——`internal/gist/fetch.go` 的 HTTP 请求失败会 `panic`）都会被转成状态栏错误提示并关闭 loading，而不是让整个 TUI 崩溃。反例：不要对触网任务直接写 `go func(){...}()`，否则会绕过兜底。非触网的轻量消息（如 `go store.Send(CMD_APP_FOCUS(1))`）仍可直接 `go`。

### 焦点模型
`AppModel.focusPanel(index)` 先 blur 全部再 focus 单个面板（1 = 文件列表，2 = 文件面板）；但当有对话框激活或输入框获得焦点（`store.State.InputFocus`）时它是空操作。`left`/`right` 方向键切换面板。`FilePanel` 根据 `store.State.Editing` 显示 `Markdown`（阅读）或 `Editor`（编辑）。

### Gist 后端抽象（`internal/gist`）
`Gist` 用一个类型封装了 GitHub 与 Gitee 两套 REST API。`conf.IsGitee()` 用来切换 API 前缀、鉴权头，以及 token 的传递方式（Gitee 用 `access_token` 查询/请求体参数，GitHub 用 bearer 头）。`Setup()` 通过匹配 `SPECIAL_DESCRIPTION` 定位应用专用的 Gist，最多扫描两页，找不到则创建。每条笔记都是这个 Gist 内的一个文件。

### i18n（`internal/i18n`）
不是全局词表。每个组件在自己的 `i18n_tpl.go` 中定义各自的 `i18nTpl` 字符串（一段 YAML）；`i18n.Get(tpl, key)` 解析它并返回当前 `LANG` 对应的文案。新增面向用户的文案时，往对应组件的 `i18nTpl` 里加一个同时含 `en_US` 和 `zh_CN` 的 key。

### 共享模型基类（`internal/app/pkgs/model`）
- `BaseModel`：ID/Width/Height/Active + `Resize`/`Focus`/`Blur`，各组件内嵌它。
- `BoxModel`：内嵌 `BaseModel`，增加带 header/footer 标题的边框盒渲染。

### 鼠标区域（mouse zones）
用 `lrstanley/bubblezone/v2` 标记可点击区域。`zone.Scan` 包裹根视图；组件用 `zone.Mark(m.ID, ...)` 标记自己的输出，并在鼠标事件里用 `zone.Get(id).InBounds(msg)` 判断命中。ID 前缀由 `NewBaseModel` 中的 `zone.NewPrefix()` 生成。

bubblezone 靠往字符串里注入**零宽的私有 ANSI 转义序列**来标记区域，最后由 `zone.Scan` 解析出坐标。这决定了叠加渲染的实现方式——见下。

### 模态弹框叠加（`internal/app/pkgs/dialog/position.go`）
`PlaceOverlay(x, y, fg, bg)` 把对话框 `fg` 叠加到主界面 `bg` 之上（`app.go` 的 `AppModel.View` 调用），做法是**字符串级逐行裁剪拼接**（依赖 `charmbracelet/x/ansi` 的 `Truncate`/`TruncateLeft`/`StringWidth`）。

**为什么不用 lipgloss v2 原生的 `Layer`/`Compositor`**：Compositor 基于 cell-buffer 重绘，合成时只保留可见字符，会**剥离 bubblezone 注入的 ANSI 区域标记**，导致对话框内按钮/输入框的鼠标点击全部失效。字符串拼接则原样保留 `fg` 里的标记，交由外层 `zone.Scan` 正确解析。二者是「视觉合成」与「鼠标命中」两个正交问题，不能互相替代，混用还会互相破坏。

**未来计划**：待 bubblezone 支持把区域标记写入 cell-buffer（而非零宽 ANSI 流）后，即可迁移到原生 `lipgloss.Compositor` 并删除 `position.go`。在此之前，`position.go` 是一段只依赖底层 `x/ansi` 的自有工具代码，跟随 `x/ansi` 升级即可，无需等待任何上游 PR 合并。


## CI / 发布
`.github/workflows/ci.yml` 在每次 push 时通过 `scripts/build.sh` 构建；在打 tag（`v*`）时通过 GitHub Release 发布二进制。`.github/workflows/docker.yml` 在打 tag 时构建并推送多架构 Docker 镜像。应用版本号硬编码在 `internal/conf/conf.go`（`VERSION`）。
