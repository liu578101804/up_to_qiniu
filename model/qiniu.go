package model

import (
	"context"
	"github.com/liu578101804/up_to_qiniu/config"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"strings"
	"bytes"
)

// 自定义返回值结构体
type PutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func UpResumeData(key string, data []byte) (PutRet, error) {

	config := config.QiniuConfig

	dataLen := int64(len(data))
	scope := config.Bucket + ":" + key
	persistentOps := strings.Join([]string{}, ";")
	putPolicy := storage.PutPolicy{
		Scope: scope,
		PersistentOps: persistentOps,
		//多媒体队列
		PersistentPipeline: config.PipeLine,
		//数据返回格式
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	//过期时间7小时
	putPolicy.Expires = 60*60*7
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	//cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := PutRet{}
	putExtra := storage.RputExtra{}
	err := resumeUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)

	return ret, err
}