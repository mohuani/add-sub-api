
> 这里是client，通过grpc client 调用 服务方

本项目设计的三个项目地址，将这三个项目放到同一层级目录就行
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
- 准备 Rpc Request 参数
- 调用 grpc service 的具体 方法
- 接收 Rpc Response
  

---

### TODO

- [ ]  后面有时间的话，计划录一节视频来讲解一下自己的理解

### 参考项目
- https://www.bilibili.com/video/BV1Np4y1x7JB