
## 环境安装
开局第一步，先解决环境问题，本人之前也通过一些标准的文档来安装，但是总会出现消小小的问题

### 标准安装

- 1 安装protoc
```shell
https://github.com/protocolbuffers/protobuf/releases/tag/v3.14.0
```
- 2 安装protoc-gen-go
```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

- 3 安装grpc

```shell
go get google.golang.org/grpc
```


### 通过 go-zero 提供的goctl一键安装
```shell
https://go-zero.dev/cn/docs/prepare/protoc-install
```

### 参考项目
- https://www.bilibili.com/video/BV1Np4y1x7JB


## 编码实现

本项目设计的三个项目地址，将这三个项目放到同一层级目录就行，`grpc-go-add`和`grpc-go-sub`是server，`add-sub-api`是client
- https://github.com/mohuani/grpc-go-add
- https://github.com/mohuani/grpc-go-sub
- https://github.com/mohuani/add-sub-api


### 编写 proto 文件
- client端的 proto 文件应该直接使用服务方的 proto 文件`add.proto`，`sub.proto`
    

### 生成pb文件
```shell
protoc --go_out=plugins=grpc:. add.proto
protoc --go_out=plugins=grpc:. sub.proto
```

### 编写 grpc service 服务
- 实现具体的逻辑

### 编写 grpc client 调用 grpc server
- 连接 grpc server xx端口
   
- 实例化 grpc service client
- 准备 rpc Request 参数
- 调用 grpc service 的具体 方法
- 接收 rpc Response
  

---

### TODO

- [ ]  后面有时间的话，计划录一节视频来讲解一下自己的理解
- [ ]  支持 http访问
- [ ]  支持 swagger 文档

