# 第一阶段以golang:alpine镜像为基础，提供安装了 Go 编程语言的最小 Alpine Linux 环境。
FROM golang:alpine as builder

# 该WORKDIR指令为其余指令设置工作目录。
WORKDIR /go/src/gin-project/server

# 该COPY指令将应用程序源代码从本地目录复制到容器中的工作目录中。
COPY . .

# 该RUN指令运行一系列命令来构建应用程序。
# 这些命令设置了一些 Go 环境变量，例如模块模式和代理设置，启用 GoCGO标志，
# 并运行go mod tidy以下载所需的依赖项。
# 最后，该go build命令用于将应用程序编译成可执行二进制文件。
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

# 第二阶段以alpine:latest镜像为基础，提供了一个没有安装 Go 的最小 Alpine Linux 环境。
# 注：Alpine是一个轻量级的操作系统，它特别适合用于嵌入式系统和Docker容器。它的设计目标是最小化系统的大小，以便于最大化资源的利用。
FROM alpine:latest

WORKDIR /go/src/gin-project/server

# 此阶段将编译后的二进制文件从第一阶段复制到映像中，并使用指令公开应用程序的端口 8888 EXPOSE。
COPY --from=0 /go/src/gin-project/server ./

EXPOSE 8888

# 该ENTRYPOINT指令指定容器启动时要运行的命令。
ENTRYPOINT ./server -c config.docker.yaml