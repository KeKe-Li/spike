#### spike
用的是Golang 1.112版本的go module,配合Golang的web框架gin实现的.

<p align="center">
<img width="100%" align="center" src="public/images/logo.png" />
</p>

<p align='center'>
<img src="https://img.shields.io/badge/build-passing-brightgreen.svg">
<a href="https://twitter.com/perfactsen"><img src="https://img.shields.io/badge/twitter-keke-green.svg?style=flat&colorA=009df2"></a>
<a href="https://www.zhihu.com/people/sencoed.com/activities"><img src="https://img.shields.io/badge/%E7%9F%A5%E4%B9%8E-keke-green.svg?style=flat&colorA=009df2"></a>
</p>

#### docker

```docker
#母镜像
FROM golang
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
```

然后运行下面的命令把当前编译过项目打进docker镜像:

```bash
> docker build -t main.go .

 ---> 797daa9977c6
Successfully built 797daa9977c6
Successfully tagged main.go:latest
```

```bash
> docker images 
main.go                               latest              797daa9977c6        8 minutes ago       801MB
```

表示把项目成功打进docker镜像了.

```bash
> docker run -p 8080:8080 -d main.go
```

docker run -p 80:8080 oracle:latest　可以直接docker运行本地镜像启动go项目


#### golang编程

觉得此文章不错，支持我的话可以给我star ，:star:！如果有问题可以加我的微信,也可以加入我们的交流群一起交流goalng技术！

#### License
This is free software distributed under the terms of the MIT license
