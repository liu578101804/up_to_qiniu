package config

import (
	"encoding/json"
	"fmt"
	"github.com/liu578101804/up_to_qiniu/utils"
	"io/ioutil"
	"os"
	"regexp"
)

var jsonData map[string]interface{}

func initJSON()  {
	var(
		bytes []byte
		err error
		configStr string
	)
	bytes, err = ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr = string(bytes[:])

	//清空注释
	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err = json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type dbConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}
//数据库配置
var DBConfig dbConfig

func initDB()  {
	utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}


type serverConfig struct {
	Port				int
}
//服务配置
var ServerConfig serverConfig

func initServerConfig()  {
	utils.SetStructByJSON(&ServerConfig, jsonData["server"].(map[string]interface{}))
}



type qiniuConfig struct {
	AccessKey 	string
	SecretKey 	string
	Bucket		string
	PipeLine	string
	Domain		string
	MaxSize   	int
}
//七牛云配置
var QiniuConfig qiniuConfig

func initQiniuConfig()  {
	utils.SetStructByJSON(&QiniuConfig, jsonData["qiniu"].(map[string]interface{}))
}

func init()  {
	initJSON()
	initDB()
	initServerConfig()
	initQiniuConfig()
}
