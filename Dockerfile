#母镜像
FROM golang:latest
#维护者信息
MAINTAINER keke
#工作目录
WORKDIR $GOPATH/src/go
#将文件复制到镜像中
ADD . $GOPATH/src/go
#执行操作
RUN go build  main.go
#暴露端口
EXPOSE 8080
#程序入口
ENTRYPOINT ["./main"]