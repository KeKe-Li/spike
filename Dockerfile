#母镜像
FROM golang:latest as build
#维护者信息
MAINTAINER keke

ENV GOPROXY https://go.likeli.top
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .

# 构建缓存包含了该项所有的依赖起到加速构建的作用
RUN go mod download

#工作目录
WORKDIR /go/release

#将文件复制到镜像中
ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o spike main.go

# scratch空的基础镜像，最小的基础镜像
FROM scratch as prod

COPY --from=build /go/release/spike /

EXPOSE 8080

CMD ["/spike"]
