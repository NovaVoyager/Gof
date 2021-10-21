# GOF

![](https://img.shields.io/badge/golang-1.15-brightgreen) ![](https://img.shields.io/badge/gin-1.7.4-red) ![](https://img.shields.io/badge/redis-8.11.4-yellow) ![](https://img.shields.io/badge/gorm-1.21.16-orange) ![](https://img.shields.io/badge/mongo-1.7.3-blue)

### 介绍

`GOF` 是基于 `gin` 框架开发，用于快速构建 *api* 接口，框架已实现 `redis` `mysql` `mongo` 的基本配置和连接，支持配置文件热加载，日志记录使用了  `uber/zap `
日志库，已完成请求日志的记录，记录格式为 `json` 方便保存至 `mongo` 或 `ES` ，可根据需要快速且方便的进行自定义调整。

### 使用说明

> golang 版本 >= 1.15
> 设置 GO111MODULE=on
> 设置代理 GOPROXY=https://goproxy.cn,direct

**安装**

> go get github.com/miaogu-go/Gof

**运行**

> go run main.go

**注意：**在运行前请先设置好 config.toml 里对应的值

### 目录结构

```
├── Gof
    ├── api                   (控制器层)
    │   └── v1                (v1版本接口)
    ├── config                (配置结构体)
    ├── core                  (核心文件)
    │	├── initialize        (初始化文件)
    │	└── bootstrap.go      (启动入口文件)
    ├── global                (全局对象)                                                
    ├── middleware            (中间件层)                        
    ├── model                 (模型层)                    
    │   ├── common            (公共结构体)
    │   ├── dao               (数据库结构体层)
    │   ├── data              (公共数据组装层——函数复用)
    │   └── service           (服务层)                           
    ├── resource              (资源层)
    │   └── errno             (错误码注册)                     
    ├── router                (路由层)    
    │   └── bootstrap.go      (路由启动文件)               
    └── utils                 (工具包)       
```