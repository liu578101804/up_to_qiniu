package router

import (
	"github.com/liu578101804/up_to_qiniu/controller/upload"
	"github.com/liu578101804/up_to_qiniu/middleware"
	"github.com/plimble/ace"
)

func Router(app *ace.Ace)  {

	app.Use(middleware.APIStatsD())

	app.POST("/upload", upload.Upload)
}