# 使用官方的Golang镜像作为构建环境
FROM golang:1.22-alpine as builder
# 安装 CA 证书、sshpass 和其他必要的工具
RUN apk --no-cache add ca-certificates sshpass tzdata
# 设置工作目录
WORKDIR /app
COPY . .
# 下载所有依赖项
RUN go mod download
# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/SEP/main.go

# 使用Alpine作为最小运行环境（不使用scratch以支持sshpass）
FROM alpine:latest
WORKDIR /app
# 从Alpine安装sshpass、bash和openssh-client
RUN apk --no-cache add sshpass bash openssh-client
# 从builder镜像中复制构建的二进制文件
COPY --from=builder /app/main ./main
# 从builder镜像中复制时区数据
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# 从builder镜像中复制 CA 证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# 从builder镜像中复制配置和日志目录
COPY --from=builder /app/configs/ ./configs/
COPY --from=builder /app/logs/ ./logs/
COPY --from=builder /app/uploads/ ./uploads/
# 设置时区环境变量
ENV TZ=Asia/Shanghai
# 暴露端口
EXPOSE 714
# 运行应用程序
CMD ["./main"]
