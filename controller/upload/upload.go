package upload

import (
	"fmt"
	"github.com/liu578101804/up_to_qiniu/config"
	"github.com/liu578101804/up_to_qiniu/controller/common"
	"github.com/liu578101804/up_to_qiniu/model"
	"github.com/liu578101804/up_to_qiniu/utils"
	"github.com/plimble/ace"
	"io/ioutil"
)

func Upload(c *ace.C)  {

	var(
		fileName string
	)
	SendErrorJSON := common.SendErrJSON

	random := c.MustPostString("random", "T")
	file, fileHander, err := c.Request.FormFile("file")
	if err != nil {
		SendErrorJSON("没找到文件 file", c)
		return
	}
	defer file.Close()

	//判断文件大小
	type Sizer interface {
		Size() int64
	}
	fSize := file.(Sizer).Size()
	maxSize := int64(config.QiniuConfig.MaxSize*1024*1024)
	if fSize > maxSize {
		SendErrorJSON(fmt.Sprintf("超过文件大小限制 %vM", config.QiniuConfig.MaxSize), c)
		return
	}

	//是否随机生成文件名
	fileName = fileHander.Filename
	if random == "T" {
		fileName = utils.RandomFileName()
	}

	//读取文件
	fileBytes, err := ioutil.ReadAll(file)
	//上传文件
	res, err := model.UpResumeData(fileName, fileBytes)
	if err != nil {
		SendErrorJSON(err.Error(), c)
		return
	}

	c.JSON(200, map[string]interface{}{
		"code": 200,
		"msg": "操作成功",
		"data": map[string]interface{}{
			"url": config.QiniuConfig.Domain+"/"+res.Key,
		},
	})

}