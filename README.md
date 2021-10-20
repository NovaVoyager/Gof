# GOF

![](https://img.shields.io/badge/golang-1.15-brightgreen) ![](https://img.shields.io/badge/gin-1.7.4-red) ![](https://img.shields.io/badge/redis-8.11.4-yellow) ![](https://img.shields.io/badge/gorm-1.21.16-orange) ![](https://img.shields.io/badge/mongo-1.7.3-blue)

### 介绍

`GOF` 是基于 `gin` 框架开发，用于快速构建 *api* 接口，框架已实现 `redis` `mysql` `mongo` 的基本配置和连接，日志记录使用了  `uber/zap `
日志库，已完成请求日志的记录，记录格式为 `json` 方便保存至 `mongo` 或 `ES` ，可根据需要快速且方便的进行自定义调整。

### 使用说明

> golang 版本 >= 1.15

**安装**

> go get github.com/miaogu-go/Gof

**运行**

> go run main.go

### 目录结构
