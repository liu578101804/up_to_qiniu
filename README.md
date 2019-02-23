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


