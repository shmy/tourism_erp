### 一个基于Golang的旅游业ERP系统

### 状态
> 开发中

### 功能
* - [x] 基于`iris`，地球上最快的Golang Web框架
* - [x] 数据存储采用 `mysql`，orm 采用 `gorm`
* - [x] 采用 `jsonwebtoken` 进行用户会话持久
* - [ ] 采用 `casbin`，更加灵活的权限验证

* 安装依赖
> 使用`glide`做依赖管理工具，操作系统请安装`upx`工具以压缩

```bash
$ glide update
```

* 构建

```bash
$ ./bash.sh
```

* 部署与启动

```bash
$ ./manage.sh {start|stop|restart|status}
```