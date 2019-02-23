package main

import (
	"fmt"
	"github.com/liu578101804/up_to_qiniu/config"
	"github.com/liu578101804/up_to_qiniu/router"
	"github.com/plimble/ace"
	"github.com/plimble/ace-contrib/cors"
)

func main() {

	app := ace.New()

	//跨域
	app.Use(cors.Cors(cors.Options{}))

	//初始化路由
	router.Router(app)

	fmt.Println("server is run listing at", config.ServerConfig.Port)
	app.Run(fmt.Sprintf(":%d", config.ServerConfig.Port))

}
