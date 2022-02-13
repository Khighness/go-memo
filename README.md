## go-memo

基于`gin`和`gorm`实现的备忘录后端。



### 项目结构

``` 
go-memo
    ├─api        声明接口
    ├─cache      缓存操作
    ├─conf       配置文件
    ├─model      数据库模型
    ├─pkg
    │  ├─e      定义错误码
    │  └─util   工具函数
    ├─router     路由逻辑
    ├─serializer 序列化
    ├─service    业务实现
    └─web        服务配置
```



### 快速开始

1. 下载依赖

```
go mod tidy
```

2. 运行项目

```
go run main.go
```

   
