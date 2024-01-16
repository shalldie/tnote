# --- 构建 ---
FROM golang:1.20-alpine AS builder

COPY ./ /app/

WORKDIR /app

# 编译
RUN CGO_ENABLED=0 \
    go build \
    -gcflags "all=-trimpath=/app" \
    -ldflags="-s -w" \
    -o /output/tnote /app/main.go

# 压缩
RUN apk add upx && upx /output/*

# --- 产出 ---
FROM alpine:latest

WORKDIR /

COPY --from=builder /output/tnote /tnote

# https://github.com/charmbracelet/lipgloss/issues/31
ENV COLORTERM="truecolor"

CMD [ "/tnote" ]
