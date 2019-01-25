# spike
用的是Golang 1.112版本的go module实现的.

#### docker

```markdown
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
