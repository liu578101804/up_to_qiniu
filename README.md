# 七牛云文件上传工具

这个工具是基于Go语言开发，使用七牛云的Go SDK开发，你只需要配置好相关秘钥，指定好去区块，就能有一个http网关。

**此项目仅适合探讨学习，用于生产还需要二次开发。**

## 功能说明：

- [x] 上传文件到七牛


## 未来

- [ ] 上传缓存到mysql
- [ ] 文件删除
- [ ] 访问日志记录
- [ ] 支持限定上传文件类型

## API列表
- /upload 文件上传


## docker支持

喜欢docker部署的同学，此项目我已经打包到hub.docker.com开源镜像上了，可以直接去 [https://hub.docker.com/r/liu578101804/up_to_qiniu](https://hub.docker.com/r/liu578101804/up_to_qiniu)
 上查看镜像相关信息。

```
docker pull liu578101804/up_to_qiniu
```

## 本地编译docker镜像


```
make build_docker
```

## 镜像使用方法

程序的允许需要依赖一个配置文件(config.json)，这个配置文件，主要是配置七牛云的相关参数以及程序监听的端口，这里给一个例子：

```json
{
  "server":{
    "Port": 8023                 /*监听端口*/
  },
  "qiniu":{
    "AccessKey": "",      /*从七牛云后台能拿到*/
    "SecretKey": "",      /*从七牛云后台能拿到*/
    "Bucket": "",         /*区块名字*/
    "PipeLine": "",       /*七牛的多媒体作业*/
    "Domain": "",         /*域名*/
    "MaxSize": 4          /*上传图片最大允许大小，单位M*/
  }
}
```

这个配置文件里面支持写这样的注释 `/*注释*/`，在代码里面做了过滤。

把这个配置文件挂在到 `/app/config/config.json` 下面就能正常启动了。

## docker-compose的使用 

```yaml

version: '2'

services:

  upQiniu:
    container_name: up_qiniu
    image: liu578101804/up_to_qiniu:1.0
    ports:
      - "8023:8023"
    environment:
      - TZ=Asia/Shanghai
    volumes:
      - ./config:/app/config


```

感兴趣可以关注我的个人微信公众号，微信公众号名称【GoLang全栈】，或者微信号 【GolangStackDev】，了解更多学习Go语言的资料。
